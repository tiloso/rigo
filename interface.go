package rigo

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Interface interface {
	Keyer
	KeySetter
	Vclocker
	VclockSetter
	Indexer
}

type Keyer interface {
	Key() []byte
}

type KeySetter interface {
	SetKey([]byte)
}

type Vclocker interface {
	Vclock() []byte
}

type VclockSetter interface {
	SetVclock([]byte)
}

type Indexer interface {
	Indexes() []KVPair
}

type KVPair struct {
	Key, Value []byte
}

// GetI fetches object and siblings from riak, unmarshals data into provided
// interface and set the Key and Vector Clock on every element if they are type
// of KeySetter / VclockSetter
// It expects a pointer to a slice of the expected result type as argument
func (o *Object) GetI(v interface{}) error {
	resultv := reflect.ValueOf(v)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("rigo: obj argument must be a slice address")
	}

	slicev := resultv.Elem()
	elemt := slicev.Type().Elem()

	rpbRes, err := o.Get()
	if err != nil {
		return err
	}

	// method used by rgo might improve performance
	// (set elemv at index of slicev, append and set len == cap if index > len)
	// https://github.com/go-mgo/mgo/blob/v2/session.go#L2817
	// func (iter *Iter) All(v interface{}) error {...}
	for _, c := range rpbRes.GetContent() {
		elemp := reflect.New(elemt)

		if len(c.Value) > 0 {
			if err := json.Unmarshal(c.Value, elemp.Interface()); err != nil {
				return fmt.Errorf("rigo: %v", err)
			}
		}

		if i, ok := elemp.Elem().Interface().(KeySetter); ok {
			i.SetKey(o.key)
		}

		if i, ok := elemp.Elem().Interface().(VclockSetter); ok {
			i.SetVclock(rpbRes.GetVclock())
		}

		slicev = reflect.Append(slicev, elemp.Elem())
	}
	resultv.Elem().Set(slicev)
	return nil
}

// The argument to DeleteI needs to satisfy Keyer interface. If it satisfies
// the Vclocker interfaces aswell, rigo provides riak with a vector clock aswell
func (b *Bucket) DeleteI(v Keyer) error {
	o := b.K(v.Key())

	if vc, ok := v.(Vclocker); ok {
		o.SetVclock(vc.Vclock())
	}

	return o.Delete()
}

func (b *Bucket) StoreI(v interface{}) error {
	om, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("rigo: %v", err)
	}

	var k []byte
	kr, isKr := v.(Keyer)
	if isKr {
		k = kr.Key()
	}

	o := b.K(k)

	if i, ok := v.(Vclocker); ok {
		o.SetVclock(i.Vclock())
	}

	if i, ok := v.(Indexer); ok {
		o.SetIndexes(i.Indexes())
	}

	rpbRes, err := o.SetContentType([]byte("application/json")).Store(om)
	if err != nil {
		return err
	}

	// set key if a key hasn't been provided and key is settable
	if isKr && kr.Key() != nil {
		return nil
	}

	if ks, ok := v.(KeySetter); ok {
		ks.SetKey(rpbRes.GetKey())
	}
	return nil
}
