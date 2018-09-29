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
// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

// Package threefish implements the Threefish tweakable block cipher.
// Threefish is designed to be the core function of the Skein hash function
// family.
// There are three versions of Threefish
// 		- Threefish256  processes 256  bit blocks
//		- Threefish512  processes 512  bit blocks
//		- Threefish1024 processes 1024 bit blocks
package threefish

import (
	"crypto/cipher"
	"errors"
)

const (
	// The size of the tweak in bytes.
	TweakSize = 16
	// C240 is the key schedule constant
	C240 = 0x1bd11bdaa9fc1a22
	// The block size of Threefish-256 in bytes.
	BlockSize256 = 32
	// The block size of Threefish-512 in bytes.
	BlockSize512 = 64
	// The block size of Threefish-1024 in bytes.
	BlockSize1024 = 128
)

var errKeySize = errors.New("invalid key size")

// NewCipher returns a cipher.Block implementing the Threefish cipher.
// The length of the key must be 32, 64 or 128 byte.
// The length of the tweak must be TweakSize.
// The returned cipher implements:
//		- Threefish-256  - if len(key) = 32
//		- Threefish-512  - if len(key) = 64
// 		- Threefish-1024 - if len(key) = 128
func NewCipher(tweak *[TweakSize]byte, key []byte) (cipher.Block, error) {
	switch k := len(key); k {
	default:
		return nil, errKeySize
	case BlockSize256:
		return newCipher256(tweak, key), nil
	case BlockSize512:
		return newCipher512(tweak, key), nil
	case BlockSize1024:
		return newCipher1024(tweak, key), nil
	}
}

// Increment the tweak by the ctr argument.
// Skein can consume messages up to 2^96 -1 bytes.
func IncrementTweak(tweak *[3]uint64, ctr uint64) {
	t0 := tweak[0]
	tweak[0] += ctr
	if tweak[0] < t0 {
		t1 := tweak[1]
		tweak[1] = (t1 + 1) & 0x00000000FFFFFFFF
	}
}

// The threefish-256 tweakable blockcipher
type threefish256 struct {
	keys  [5]uint64
	tweak [3]uint64
}

func (t *threefish256) BlockSize() int { return BlockSize256 }

// The threefish-512 tweakable blockcipher
type threefish512 struct {
	keys  [9]uint64
	tweak [3]uint64
}

func (t *threefish512) BlockSize() int { return BlockSize512 }

// The threefish-1024 tweakable blockcipher
type threefish1024 struct {
	keys  [17]uint64
	tweak [3]uint64
}

func (t *threefish1024) BlockSize() int { return BlockSize1024 }
