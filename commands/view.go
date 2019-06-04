// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import "github.com/jxo/lime"

type (
	// Close command closes the currently opened view.
	Close struct {
		lime.DefaultCommand
	}

	// NextView command switches to the view which is
	// immediately to the next of the current view.
	NextView struct {
		lime.DefaultCommand
	}

	// PrevView command switches to the view
	// which is immediately before hte current view.
	PrevView struct {
		lime.DefaultCommand
	}

	// SetFileType command will let us set the file type
	// for the currently active view, eg: for Syntax highlighting.
	SetFileType struct {
		lime.DefaultCommand
		Syntax string
	}
)

// Run executes the Close command.
func (c *Close) Run(w *lime.Window) error {
	if v := w.ActiveView(); v != nil {
		v.Close()
	} else {
		w.Close()
	}
	return nil
}

// Run executes the NextView command.
func (c *NextView) Run(w *lime.Window) error {
	for i, v := range w.Views() {
		if v == w.ActiveView() {
			i++
			if i == len(w.Views()) {
				i = 0
			}
			w.SetActiveView(w.Views()[i])
			break
		}
	}

	return nil
}

// Run executes the PrevView command.
func (c *PrevView) Run(w *lime.Window) error {
	for i, v := range w.Views() {
		if v == w.ActiveView() {
			if i == 0 {
				i = len(w.Views())
			}
			i--
			w.SetActiveView(w.Views()[i])
			break
		}
	}

	return nil
}

// Run executes the SetFileType command.
func (c *SetFileType) Run(v *lime.View, e *lime.Edit) error {
	v.SetSyntaxFile(c.Syntax)
	return nil
}

func init() {
	register([]lime.Command{
		&Close{},
		&NextView{},
		&PrevView{},
		&SetFileType{},
	})
}
