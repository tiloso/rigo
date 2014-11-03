package rigo

import (
	"encoding/json"
	"reflect"

	"github.com/tiloso/rigo/rpb"
)

// v := []KVI{&testObjs{}}
// B("abc").K("nexus").GetI(v)
// B("abc").StoreI(v)
// B"abc").DeleteI(v)

// B("abc").K("nexus").GetKVI(v)
// B("abc").StoreKVI(v)
// B"abc").DeleteKVI(v)

// type KeyVclockIndexer
type KVIer interface {
	Key() []byte
	SetKey([]byte)
	Vclock() []byte
	SetVclock([]byte)
	Indexes() []*rpb.RpbPair
}

func (o *Object) GetKVI(v []KVIer) error {
	resultv := reflect.ValueOf(v)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("obj argument must be a slice address")
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
				return err
			}
		}

		kv, ok := elemp.Interface().(KVIer)
		if ok {
			kv.SetKey(o.key)
			kv.SetVclock(rpbRes.GetVclock())
		}

		slicev = reflect.Append(slicev, elemp.Elem())
	}
	resultv.Elem().Set(slicev)
	return nil
}

func (b *Bucket) DeleteKVI(v KVIer) error {
	return b.K(v.Key()).Vclock(v.Vclock()).Delete()
}

func (b *Bucket) StoreKVI(v KVIer) error {
	om, err := json.Marshal(v)
	if err != nil {
		return err
	}

	rpbRes, err := b.K(v.Key()).
		Vclock(v.Vclock()).
		Indexes(v.Indexes()).
		ContentType([]byte("application/json")).
		Store(om)

	if err != nil {
		return err
	}

	if v.Key() == nil {
		v.SetKey(rpbRes.GetKey())
	}

	return nil
}
