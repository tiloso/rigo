package rigo

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"code.google.com/p/goprotobuf/proto"
	"github.com/tiloso/rigo/rpb"
)

var ErrZeroLength = errors.New("response was only 0 bytes long")

func unmarshalRPB(data []byte, code byte, pm proto.Message) error {
	if len(data) == 0 {
		return ErrZeroLength
	}

	if data[0] == rpbErrResCode {
		out := &rpb.RpbErrorResp{}
		if err := proto.Unmarshal(data[1:], out); err != nil {
			return err
		}
		return fmt.Errorf("riak error [%d]: %s", out.GetErrcode(), out.GetErrmsg())
	}

	if data[0] != code {
		return fmt.Errorf("invalid rpb code: expected %v, got %v", code, data[0])
	}

	return proto.Unmarshal(data[1:], pm)
}

func marshalRPB(code int, pm proto.Message) ([]byte, error) {
	in, err := proto.Marshal(pm)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	if err := binary.Write(b, binary.BigEndian, int32(len(in)+1)); err != nil {
		return nil, err
	}
	if err := binary.Write(b, binary.BigEndian, int8(code)); err != nil {
		return nil, err
	}
	if err := binary.Write(b, binary.BigEndian, in); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func marshalRPBCode(code int) ([]byte, error) {
	b := &bytes.Buffer{}
	if err := binary.Write(b, binary.BigEndian, int32(1)); err != nil {
		return nil, err
	}
	if err := binary.Write(b, binary.BigEndian, int8(code)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
