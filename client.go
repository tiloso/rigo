package rigo

import (
	"fmt"
	"net"

	"github.com/tiloso/rigo/rpb"
)

type Client struct {
	tcpaddr *net.TCPAddr
	pool    chan *session
}

func (c *Client) session() *session {
	return <-c.pool
}

func (c *Client) T(v []byte) *Type {
	return &Type{c, v}
}

func (c *Client) B(v []byte) *Bucket {
	return &Bucket{
		Type:   c.T(nil),
		bucket: v,
	}
}

func Dial(addr string, poolLimit int) (*Client, error) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	c := &Client{
		tcpaddr: tcpaddr,
		pool:    make(chan *session, poolLimit),
	}

	for i := 0; i < poolLimit; i++ {
		c.pool <- &session{
			client: c,
		}
	}

	inactiveConns := 0
	for i := 0; i < len(c.pool); i++ {
		s := c.session()
		if err = s.Dial(tcpaddr); err != nil {
			inactiveConns++
			s.active = false
		}
		s.release()
	}

	if inactiveConns == len(c.pool) {
		return nil, fmt.Errorf("couldn't establish conn: %v", err)
	}

	return c, nil
}

func (c *Client) ID() ([]byte, error) {
	req, err := marshalRPBCode(rpbGetClientIDReqCode)
	if err != nil {
		return nil, err
	}

	s := c.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return nil, err
	}

	rpbRes := &rpb.RpbGetClientIdResp{}
	if err := unmarshalRPB(res, rpbGetClientIDResCode, rpbRes); err != nil {
		return nil, err
	}

	return rpbRes.ClientId, nil
}

func (c *Client) SetID(id []byte) error {
	rpbReq := &rpb.RpbSetClientIdReq{
		ClientId: id,
	}

	req, err := marshalRPB(rpbSetClientIDReqCode, rpbReq)
	if err != nil {
		return err
	}

	s := c.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return err
	}

	if len(res) != 1 || res[0] != rpbSetClientIDResCode {
		return fmt.Errorf(unexpectedRPBResFormat, rpbSetClientIDResCode, len(res), res[0])
	}
	return nil
}

func (c *Client) Ping() error {
	req, err := marshalRPBCode(rpbPingReqCode)
	if err != nil {
		return err
	}

	s := c.session()
	defer s.release()

	res, err := s.roundTrip(req)
	if err != nil {
		return err
	}

	if len(res) != 1 || res[0] != rpbPingResCode {
		return fmt.Errorf("unexpected ping res: len %v, code %v", len(res), res[0])
	}
	return nil
}
