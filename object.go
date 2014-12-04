package rigo

import (
	"fmt"

	"github.com/tiloso/rigo/rpb"
)

type Object struct {
	*Bucket
	key             []byte
	vclock          []byte
	contentType     []byte
	contentEncoding []byte
	indexes         []KVPair
	timeout         *uint32
	ifModified      []byte
	head            *bool
}

func (o *Object) setHead(v bool) {
	o.head = &v
}

func (o *Object) setIfModified(v []byte) {
	o.ifModified = v
}

// type ObjectSetter func(*Object) *Object
func (o *Object) SetVclock(v []byte) *Object {
	o.vclock = v
	return o
}

func (o *Object) SetIndexes(v []KVPair) *Object {
	o.indexes = v
	return o
}

func (o *Object) SetContentType(v []byte) *Object {
	o.contentType = v
	return o
}

func (o *Object) SetContentEncoding(v []byte) *Object {
	o.contentEncoding = v
	return o
}

func (o *Object) Option(opts ...option) *Object {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *Object) Get() (*rpb.RpbGetResp, error) {
	rpbReq := &rpb.RpbGetReq{
		Type:       o.typ,
		Bucket:     o.bucket,
		Key:        o.key,
		Timeout:    o.timeout,
		IfModified: o.ifModified,
		Head:       o.head,
	}

	req, err := marshalRPB(rpbGetReqCode, rpbReq)
	if err != nil {
		return nil, err
	}

	s := o.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return nil, err
	}

	rpbRes := &rpb.RpbGetResp{}
	if err := unmarshalRPB(res, rpbGetResCode, rpbRes); err != nil {
		return nil, err
	}
	return rpbRes, nil
}

func (o *Object) Store(dt []byte) (*rpb.RpbPutResp, error) {
	rpbPairs := make([]*rpb.RpbPair, len(o.indexes))
	for i, v := range o.indexes {
		rpbPairs[i] = &rpb.RpbPair{
			Key:   v.Key,
			Value: v.Value,
		}
	}

	rpbReq := &rpb.RpbPutReq{
		Bucket: o.bucket,
		Type:   o.typ,
		Key:    o.key,
		Vclock: o.vclock,
		Content: &rpb.RpbContent{
			Value:           dt,
			ContentType:     o.contentType,
			ContentEncoding: o.contentEncoding,
			Indexes:         rpbPairs,
		},
	}

	req, err := marshalRPB(rpbPutReqCode, rpbReq)
	if err != nil {
		return nil, err
	}

	s := o.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return nil, err
	}

	rpbRes := &rpb.RpbPutResp{}
	if err := unmarshalRPB(res, rpbPutResCode, rpbRes); err != nil {
		return nil, err
	}
	return rpbRes, nil
}

func (o *Object) Delete() error {
	rpbReq := &rpb.RpbDelReq{
		Bucket: o.bucket,
		Key:    o.key,
		Vclock: o.vclock,
		Type:   o.typ,
	}

	req, err := marshalRPB(rpbDelReqCode, rpbReq)
	if err != nil {
		return err
	}

	s := o.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return err
	}

	if len(res) != 1 || res[0] != rpbDelResCode {
		return fmt.Errorf(unexpectedRPBResFormat, rpbDelResCode, len(res), res[0])
	}
	return nil
}
