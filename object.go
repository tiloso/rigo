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
	indexes         []*rpb.RpbPair
}

// type ObjectSetter func(*Object) *Object

func (o *Object) Vclock(v []byte) *Object {
	o.vclock = v
	return o
}

func (o *Object) Indexes(v []*rpb.RpbPair) *Object {
	o.indexes = v
	return o
}

func (o *Object) ContentType(v []byte) *Object {
	o.contentType = v
	return o
}

func (o *Object) ContentEncoding(v []byte) *Object {
	o.contentEncoding = v
	return o
}

func (o *Object) Get() (*rpb.RpbGetResp, error) {
	rpbReq := &rpb.RpbGetReq{
		Type:   o.typ,
		Bucket: o.bucket,
		Key:    o.key,
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
	rpbReq := &rpb.RpbPutReq{
		Bucket: o.bucket,
		Type:   o.typ,
		Key:    o.key,
		Vclock: o.vclock,
		Content: &rpb.RpbContent{
			Value:           dt,
			ContentType:     o.contentType,
			ContentEncoding: o.contentEncoding,
			Indexes:         o.indexes,
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
