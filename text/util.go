// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package text

// Min returns the minimum of the arguments
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of the arguments
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Clamp clamps v to be in the range of _min and _max
func Clamp(_min, _max, v int) int {
	return Max(_min, Min(_max, v))
}

// Abs returns the absolute value of a.
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
