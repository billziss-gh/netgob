// Copyright 2009 Bill Zissimopoulos. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

// NetgobEncoder is the interface used to encode channels.
// An object that implements NetgobEncoder will be called
// when a channel needs to be encoded.
//
// In order for a NetgobEncoder to be used it must be set
// on an Encoder using SetNetgobEncoder.
type NetgobEncoder interface {
	// NetgobEncode returns a byte slice representing the encoding of the
	// passed channel.
	NetgobEncode(interface{}) ([]byte, error)
}

// NetgobDecoder is the interface used to decode channels.
// An object that implements NetgobDecoder will be called
// when a channel needs to be decoded.
//
// In order for a NetgobDecoder to be used it must be set
// on an Decoder using SetNetgobDecoder.
type NetgobDecoder interface {
	// NetgobDecode overwrites the passed channel, which must be a pointer,
	// with the value represented by the byte slice.
	NetgobDecode(interface{}, []byte) error
}
