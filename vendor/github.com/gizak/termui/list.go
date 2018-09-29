// Copyright 2018 The MATRIX Authors as well as Copyright 2014-2017 The go-ethereum Authors
// This file is consisted of the MATRIX library and part of the go-ethereum library.
//
// The MATRIX-ethereum library is free software: you can redistribute it and/or modify it under the terms of the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, 
//and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject tothe following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, 
//WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISINGFROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
//OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import "strings"

// List displays []string as its items,
// it has a Overflow option (default is "hidden"), when set to "hidden",
// the item exceeding List's width is truncated, but when set to "wrap",
// the overflowed text breaks into next line.
/*
  strs := []string{
		"[0] github.com/gizak/termui",
		"[1] editbox.go",
		"[2] iterrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go"}

  ls := termui.NewList()
  ls.Items = strs
  ls.ItemFgColor = termui.ColorYellow
  ls.BorderLabel = "List"
  ls.Height = 7
  ls.Width = 25
  ls.Y = 0
*/
type List struct {
	Block
	Items       []string
	Overflow    string
	ItemFgColor Attribute
	ItemBgColor Attribute
}

// NewList returns a new *List with current theme.
func NewList() *List {
	l := &List{Block: *NewBlock()}
	l.Overflow = "hidden"
	l.ItemFgColor = ThemeAttr("list.item.fg")
	l.ItemBgColor = ThemeAttr("list.item.bg")
	return l
}

// Buffer implements Bufferer interface.
func (l *List) Buffer() Buffer {
	buf := l.Block.Buffer()

	switch l.Overflow {
	case "wrap":
		cs := DefaultTxBuilder.Build(strings.Join(l.Items, "\n"), l.ItemFgColor, l.ItemBgColor)
		i, j, k := 0, 0, 0
		for i < l.innerArea.Dy() && k < len(cs) {
			w := cs[k].Width()
			if cs[k].Ch == '\n' || j+w > l.innerArea.Dx() {
				i++
				j = 0
				if cs[k].Ch == '\n' {
					k++
				}
				continue
			}
			buf.Set(l.innerArea.Min.X+j, l.innerArea.Min.Y+i, cs[k])

			k++
			j++
		}

	case "hidden":
		trimItems := l.Items
		if len(trimItems) > l.innerArea.Dy() {
			trimItems = trimItems[:l.innerArea.Dy()]
		}
		for i, v := range trimItems {
			cs := DTrimTxCls(DefaultTxBuilder.Build(v, l.ItemFgColor, l.ItemBgColor), l.innerArea.Dx())
			j := 0
			for _, vv := range cs {
				w := vv.Width()
				buf.Set(l.innerArea.Min.X+j, l.innerArea.Min.Y+i, vv)
				j += w
			}
		}
	}
	return buf
}
