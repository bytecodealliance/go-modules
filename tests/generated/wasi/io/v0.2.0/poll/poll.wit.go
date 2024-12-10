// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package poll represents the imported interface "wasi:io/poll@0.2.0".
//
// A poll API intended to let users wait for I/O events on multiple handles
// at once.
package poll

import (
	"go.bytecodealliance.org/cm"
)

// Pollable represents the imported resource "wasi:io/poll@0.2.0#pollable".
//
// `pollable` represents a single I/O event which may be ready, or not.
//
//	resource pollable
type Pollable cm.Resource

// ResourceDrop represents the imported resource-drop for resource "pollable".
//
// Drops a resource handle.
//
//go:nosplit
func (self Pollable) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_PollableResourceDrop((uint32)(self0))
	return
}

// Block represents the imported method "block".
//
// `block` returns immediately if the pollable is ready, and otherwise
// blocks until ready.
//
// This function is equivalent to calling `poll.poll` on a list
// containing only this pollable.
//
//	block: func()
//
//go:nosplit
func (self Pollable) Block() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_PollableBlock((uint32)(self0))
	return
}

// Ready represents the imported method "ready".
//
// Return the readiness of a pollable. This function never blocks.
//
// Returns `true` when the pollable is ready, and `false` otherwise.
//
//	ready: func() -> bool
//
//go:nosplit
func (self Pollable) Ready() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_PollableReady((uint32)(self0))
	result = (bool)(cm.U32ToBool((uint32)(result0)))
	return
}

// Poll represents the imported function "poll".
//
// Poll for completion on a set of pollables.
//
// This function takes a list of pollables, which identify I/O sources of
// interest, and waits until one or more of the events is ready for I/O.
//
// The result `list<u32>` contains one or more indices of handles in the
// argument list that is ready for I/O.
//
// If the list contains more elements than can be indexed with a `u32`
// value, this function traps.
//
// A timeout can be implemented by adding a pollable from the
// wasi-clocks API to the list.
//
// This function does not return a `result`; polling in itself does not
// do any I/O so it doesn't fail. If any of the I/O sources identified by
// the pollables has an error, it is indicated by marking the source as
// being reaedy for I/O.
//
//	poll: func(in: list<borrow<pollable>>) -> list<u32>
//
//go:nosplit
func Poll(in cm.List[Pollable]) (result cm.List[uint32]) {
	in0, in1 := cm.LowerList(in)
	wasmimport_Poll((*Pollable)(in0), (uint32)(in1), &result)
	return
}
