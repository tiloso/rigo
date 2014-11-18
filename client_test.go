package rigo_test

import (
	"bytes"
	"testing"

	"github.com/tiloso/rigo"
)

var gc *rigo.Client

func TestDialNotRunning(t *testing.T) {
	_, err := rigo.Dial("127.0.0.1:8888", 1)
	if err == nil {
		t.Error("got nil, expected error")
	}
}

func TestDial(t *testing.T) {
	_, err := rigo.Dial("127.0.0.1:8087", 1)
	if err != nil {
		t.Fatalf("err dialing: %v", err)
	}
}

func TestSetID(t *testing.T) {
	c, err := rigo.Dial("127.0.0.1:8087", 1)
	if err != nil {
		t.Fatalf("error dialing c: %v", err)
	}

	gc = c

	clientID := []byte("testclientid")

	if err := c.SetID(clientID); err != nil {
		t.Errorf("error SetID(): %v", err)
	}

	cID, err := c.ID()
	if err != nil {
		t.Errorf("error ID(): %v", err)
	}

	if !bytes.Equal(cID, clientID) {
		t.Errorf("client ids not equal: got %v, expected %v", cID, clientID)
	}
}

func TestPing(t *testing.T) {
	if err := gc.Ping(); err != nil {
		t.Errorf("err: %v", err)
	}
}
