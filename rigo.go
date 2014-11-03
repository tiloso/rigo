// Package rigo implements a simple go interface for riak. Currently it
// supports basic object and secondary index operations only. Furthermore the
// it's not yet possible to specify request options
//
// How do we distinguish between options (quorom, d, dw,..) and non options
// (Content-Type, ContentEncoding, ...) => query option <=> meta data which
// will be stored with the object
package rigo

// Riak Protocol Buffers message codes
const (
	rpbErrResCode         = 0
	rpbPingReqCode        = 1
	rpbPingResCode        = 2
	rpbGetClientIDReqCode = 3
	rpbGetClientIDResCode = 4
	rpbSetClientIDReqCode = 5
	rpbSetClientIDResCode = 6
	rpbGetReqCode         = 9
	rpbGetResCode         = 10
	rpbPutReqCode         = 11
	rpbPutResCode         = 12
	rpbDelReqCode         = 13
	rpbDelResCode         = 14
	rpbListBucketsReqCode = 15
	rpbListBucketsResCode = 16
	rpbListKeysReqCode    = 17
	rpbListKeysResCode    = 18
	rpbIndexReqCode       = 25
	rpbIndexResCode       = 26

	// RpbCounterUpdateReq     = 50
	// RpbCounterUpdateResp    = 51
	// RpbCounterGetReq        = 52
	// RpbCounterGetResp       = 53

	// RpbYokozunaIndexGetReq  = 54
	// RpbYokozunaIndexGetResp = 55
	// RpbYokozunaIndexPutReq  = 56

	// DtFetchReq   = 80
	// DtFetchResp  = 81
	// DtUpdateReq  = 82
	// DtUpdateResp = 83
)

const unexpectedRPBResFormat = "unexpected rpb res len / code: expected 1 / %v, got %v / %v"

var tval = true

// difference => more structured, no siblings?
// e.g. Update => DtOp CounterOp, SetOp, MapOp => Updates []*MapUpdate

// What's the meaning of context??

type Set struct {
	context []byte
}

func (s *Set) Context(v []byte) *Set {
	s.context = v
	return s
}

// SetOp => Add / Remove
// CounterOp => Increment
// MapOp => Remove / Update

func (s *Set) Add() error {
	return nil
}

func (s *Set) Remove() error {
	return nil
}

func (b *Bucket) S(v []byte) *Set {
	return nil
}

type Map struct {
}

func (m *Map) Update() error {
	return nil
}

func (m *Map) Delete() error {
	return nil
}

func (b *Bucket) M(v []byte) *Map {
	return nil
}

type Counter struct {
}

func (c *Counter) Increment() error {
	return nil
}

func (b *Bucket) C(v []byte) *Counter {
	return nil
}
