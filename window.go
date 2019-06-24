// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package lime

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime/debug"
	"sync"

	"github.com/jxo/lime/log"
	"github.com/jxo/lime/text"
	"github.com/jxo/lime/util"
)

type Window struct {
	util.HasID
	util.HasSettings
	views       []*View
	active_view *View
	project     *Project
	lock        sync.Mutex
}

// implement the fmt.Stringer interface
func (w *Window) String() string {
	return fmt.Sprintf("Window{id:%d}", w.ID())
}

func (w *Window) NewFile() *View {
	w.lock.Lock()
	defer w.lock.Unlock()

	v := newView(w)
	w.views = append(w.views, v)

	v.setBuffer(text.NewBuffer())
	v.selection.Clear()
	v.selection.Add(text.Region{A: 0, B: 0})
	v.Settings().Set("lime.last_save_change_count", v.ChangeCount())

	w.SetActiveView(v)
	OnNew.Call(v)

	return v
}

func (w *Window) Views() []*View {
	w.lock.Lock()
	defer w.lock.Unlock()
	ret := make([]*View, len(w.views))
	copy(ret, w.views)
	return ret
}

func (w *Window) remove(v *View) {
	w.lock.Lock()
	defer w.lock.Unlock()
	for i, vv := range w.views {
		if v == vv {
			end := len(w.views) - 1
			if i != end {
				copy(w.views[i:], w.views[i+1:])
			}
			w.views = w.views[:end]
			return
		}
	}
	log.Error("Wanted to remove view %s, but it doesn't appear to be a child of this window", v)
}

func (w *Window) OpenFile(filename string, flags int) *View {
	v := w.NewFile()

	v.SetScratch(true)
	e := v.BeginEdit()
	if fn, err := filepath.Abs(filename); err != nil {
		v.SetFileName(filename)
	} else {
		v.SetFileName(fn)
	}
	if d, err := ioutil.ReadFile(filename); err != nil {
		log.Error("Couldn't load file %s: %s", filename, err)
	} else {
		v.Insert(e, 0, string(d))
	}
	v.EndEdit(e)
	v.selection.Clear()
	v.selection.Add(text.Region{A: 0, B: 0})
	v.Settings().Set("lime.last_save_change_count", v.ChangeCount())
	v.SetScratch(false)

	OnLoad.Call(v)

	return v
}

func (w *Window) SetActiveView(v *View) {
	if w.active_view != nil {
		OnDeactivated.Call(w.active_view)
	}
	w.active_view = v
	if w.active_view != nil {
		OnActivated.Call(w.active_view)
	}
}

func (w *Window) ActiveView() *View {
	return w.active_view
}

// Closes the Window and all its Views.
// Returns "true" if the Window closed successfully. Otherwise returns "false".
func (w *Window) Close() bool {
	if !w.CloseAllViews() {
		return false
	}
	GetEditor().remove(w)

	return true
}

// Closes all of the Window's Views.
// Returns "true" if all the Views closed successfully. Otherwise returns "false".
func (w *Window) CloseAllViews() bool {
	for len(w.views) > 0 {
		if !w.views[0].Close() {
			return false
		}
	}

	return true
}

func (w *Window) runCommand(c WindowCommand, name string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Paniced while running window command %s %v: %v\n%s", name, c, r, string(debug.Stack()))
		}
	}()
	return c.Run(w)
}

func (w *Window) OpenProject(name string) *Project {
	if err := w.Project().Load(name); err != nil {
		log.Error(err)
		return nil
	}
	if abs, err := filepath.Abs(name); err != nil {
		w.Project().SetName(name)
	} else {
		w.Project().SetName(abs)
	}

	GetEditor().Watch(w.Project().FileName(), w.Project())
	OnProjectChanged.Call(w)

	return w.Project()
}

func (w *Window) Project() *Project {
	if w.project == nil {
		w.project = newProject(w)
	}
	return w.project
}
