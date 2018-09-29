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
// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package whisperv5

import (
	"encoding/json"

	"github.com/matrix/go-matrix/common/hexutil"
)

var _ = (*newMessageOverride)(nil)

func (n NewMessage) MarshalJSON() ([]byte, error) {
	type NewMessage struct {
		SymKeyID   string        `json:"symKeyID"`
		PublicKey  hexutil.Bytes `json:"pubKey"`
		Sig        string        `json:"sig"`
		TTL        uint32        `json:"ttl"`
		Topic      TopicType     `json:"topic"`
		Payload    hexutil.Bytes `json:"payload"`
		Padding    hexutil.Bytes `json:"padding"`
		PowTime    uint32        `json:"powTime"`
		PowTarget  float64       `json:"powTarget"`
		TargetPeer string        `json:"targetPeer"`
	}
	var enc NewMessage
	enc.SymKeyID = n.SymKeyID
	enc.PublicKey = n.PublicKey
	enc.Sig = n.Sig
	enc.TTL = n.TTL
	enc.Topic = n.Topic
	enc.Payload = n.Payload
	enc.Padding = n.Padding
	enc.PowTime = n.PowTime
	enc.PowTarget = n.PowTarget
	enc.TargetPeer = n.TargetPeer
	return json.Marshal(&enc)
}

func (n *NewMessage) UnmarshalJSON(input []byte) error {
	type NewMessage struct {
		SymKeyID   *string        `json:"symKeyID"`
		PublicKey  *hexutil.Bytes `json:"pubKey"`
		Sig        *string        `json:"sig"`
		TTL        *uint32        `json:"ttl"`
		Topic      *TopicType     `json:"topic"`
		Payload    *hexutil.Bytes `json:"payload"`
		Padding    *hexutil.Bytes `json:"padding"`
		PowTime    *uint32        `json:"powTime"`
		PowTarget  *float64       `json:"powTarget"`
		TargetPeer *string        `json:"targetPeer"`
	}
	var dec NewMessage
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.SymKeyID != nil {
		n.SymKeyID = *dec.SymKeyID
	}
	if dec.PublicKey != nil {
		n.PublicKey = *dec.PublicKey
	}
	if dec.Sig != nil {
		n.Sig = *dec.Sig
	}
	if dec.TTL != nil {
		n.TTL = *dec.TTL
	}
	if dec.Topic != nil {
		n.Topic = *dec.Topic
	}
	if dec.Payload != nil {
		n.Payload = *dec.Payload
	}
	if dec.Padding != nil {
		n.Padding = *dec.Padding
	}
	if dec.PowTime != nil {
		n.PowTime = *dec.PowTime
	}
	if dec.PowTarget != nil {
		n.PowTarget = *dec.PowTarget
	}
	if dec.TargetPeer != nil {
		n.TargetPeer = *dec.TargetPeer
	}
	return nil
}
