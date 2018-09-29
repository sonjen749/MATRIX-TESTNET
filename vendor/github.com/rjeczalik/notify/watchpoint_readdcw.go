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
// Copyright (c) 2014-2015 The Notify Authors. All rights reserved.
// Use of this source code is governed by the MIT license that can be
// found in the LICENSE file.

// +build windows

package notify

// eventmask uses ei to create a new event which contains internal flags used by
// notify package logic. If one of FileAction* masks is detected, this function
// adds corresponding FileNotifyChange* values. This allows non registered
// FileAction* events to be passed on.
func eventmask(ei EventInfo, extra Event) (e Event) {
	if e = ei.Event() | extra; e&fileActionAll != 0 {
		if ev, ok := ei.(*event); ok {
			switch ev.ftype {
			case fTypeFile:
				e |= FileNotifyChangeFileName
			case fTypeDirectory:
				e |= FileNotifyChangeDirName
			case fTypeUnknown:
				e |= fileNotifyChangeModified
			}
			return e &^ fileActionAll
		}
	}
	return
}

// matches reports a match only when:
//
//   - for user events, when event is present in the given set
//   - for internal events, when additionally both event and set have omit bit set
//
// Internal events must not be sent to user channels and vice versa.
func matches(set, event Event) bool {
	return (set&omit)^(event&omit) == 0 && (set&event == event || set&fileNotifyChangeModified&event != 0)
}
