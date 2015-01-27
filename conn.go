package rigo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"syscall"
)

type session struct {
	tc     *net.TCPConn
	client *Client
	active bool
}

func (c *session) release() {
	c.client.pool <- c
}

func (c *session) Dial(addr *net.TCPAddr) error {
	tc, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}
	tc.SetKeepAlive(true)
	c.tc = tc
	c.active = true
	return nil
}

func (c *session) roundTrip(req []byte) ([]byte, error) {
	// TODO => implement timeout
	if err := c.writeRequest(req); err != nil {
		return nil, err
	}

	return c.readResponse()
}

func (c *session) writeRequest(data []byte) error {
	count, err := c.tc.Write(data)
	if err != nil {
		if err == syscall.EPIPE {
			c.tc.Close()
		}
		return err
	}

	if count != len(data) {
		return fmt.Errorf("data length: %d, only wrote: %d", len(data), count)
	}

	return nil
}

func (c *session) readResponse() ([]byte, error) {
	// first 4 bytes are size of message
	buf := bytes.NewBuffer(nil)
	count, err := io.CopyN(buf, c.tc, 4)
	if err != nil {
		return nil, err
	}

	if count != 4 {
		return nil, fmt.Errorf("read bytes (%v) and expected length (%v) don't match\n", count, 4)
	}

	var size int32
	if err := binary.Read(buf, binary.BigEndian, &size); err != nil {
		return nil, err
	}

	data := make([]byte, size)
	// read rest of message and return it if no errors
	n, err := io.ReadFull(c.tc, data)
	if err != nil {
		if err == syscall.EPIPE {
			c.tc.Close()
		}
		return nil, err
	}

	if n != int(size) {
		return nil, fmt.Errorf("data length: %d, only read: %d", len(data), count)
	}

	return data, nil
}
