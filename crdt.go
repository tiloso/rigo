package rigo

import (
	"fmt"

	"github.com/tiloso/rigo/rpb"
)

type DT interface {
	Counter
	Map
}

type Contexter interface {
	Context() []byte
}

type Map interface {
	Update(string, interface{})
	Remove(string)

	Key() []byte
	Context() []byte
	Data() map[string]interface{}
	Updates() map[string]interface{}
	Removes() []string
}

func NewMap(k []byte) Map {
	return &mp{
		key: k,
	}
}

// Really needed? could check for byte in type switches
type Register func() []byte

func NewRegister(v []byte) Register {
	return func() []byte {
		return v
	}
}

type Counter interface {
	Value() int64
	Diff() *int64
	Increase() int64
	Decrease() int64
}

type counter struct {
	v    int64
	diff int64
}

func (c *counter) Value() int64 {
	return c.v
}

func (c *counter) Increase() int64 {
	c.diff++
	return c.v + c.diff
}

func (c *counter) Decrease() int64 {
	c.diff--
	return c.v + c.diff
}

func (c *counter) Diff() *int64 {
	return &c.diff
}

func NewCounter(v int64) Counter {
	return &counter{
		diff: v,
	}
}

type mp struct {
	key, context []byte

	data    map[string]interface{}
	updates map[string]interface{}
	removes []string
}

func (m *mp) Update(k string, v interface{}) {
	if m.updates == nil {
		m.updates = make(map[string]interface{})
	}
	m.updates[k] = v
}

func (m *mp) Remove(k string) {
	delete(m.updates, k)
}

func (m *mp) Key() []byte {
	return m.key
}

func (m *mp) Context() []byte {
	return m.context
}

func (m *mp) Updates() map[string]interface{} {
	return m.updates
}

func (m *mp) Removes() []string {
	return m.removes
}

func (m *mp) Data() map[string]interface{} {
	return m.data
}

func (b *Bucket) UpdateDT(dt interface{}) (*rpb.DtUpdateResp, error) {
	op := &rpb.DtOp{}

	switch t := dt.(type) {
	case Map:
		mUpdates := []*rpb.MapUpdate{}
		mRemoves := []*rpb.MapField{}

		for k, v := range t.Updates() {
			var tp rpb.MapField_MapFieldType
			u := &rpb.MapUpdate{
				Field: &rpb.MapField{
					Name: []byte(k),
				},
			}

			switch mt := v.(type) {
			case Register:
				tp = rpb.MapField_REGISTER
				u.RegisterOp = mt()
			case Counter:
				tp = rpb.MapField_COUNTER
				u.CounterOp = &rpb.CounterOp{
					Increment: mt.Diff(),
				}
			default:
				return nil, fmt.Errorf("neither Register nor Counter")
			}

			u.Field.Type = &tp
			mUpdates = append(mUpdates, u)
		}

		op.MapOp = &rpb.MapOp{
			Updates: mUpdates,
			Removes: mRemoves,
		}
	default:
		return nil, fmt.Errorf("not a map operation")
	}

	ker := dt.(Keyer)
	ctxer := dt.(Contexter)

	rpbReq := &rpb.DtUpdateReq{
		Bucket:  b.bucket,
		Type:    b.typ,
		Key:     ker.Key(),
		Context: ctxer.Context(),
		Op:      op,
	}

	req, err := marshalRPB(DtUpdateReq, rpbReq)
	if err != nil {
		return nil, err
	}

	s := b.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return nil, err
	}

	rpbRes := &rpb.DtUpdateResp{}
	if err := unmarshalRPB(res, DtUpdateRes, rpbRes); err != nil {
		return nil, err
	}

	return rpbRes, nil
}

func (b *Bucket) GetDT(key []byte) (interface{}, error) {
	rpbReq := &rpb.DtFetchReq{
		Type:   b.typ,
		Bucket: b.bucket,
		Key:    key,
	}

	req, err := marshalRPB(DtFetchReq, rpbReq)
	if err != nil {
		return nil, err
	}

	s := b.session()
	res, err := s.roundTrip(req)
	s.release()
	if err != nil {
		return nil, err
	}

	rpbRes := &rpb.DtFetchResp{}
	if err := unmarshalRPB(res, DtFetchRes, rpbRes); err != nil {
		return nil, err
	}

	if *rpbRes.Type == rpb.DtFetchResp_MAP {
		dt := make(map[string]interface{})
		for _, v := range rpbRes.Value.MapValue {
			switch *v.Field.Type {
			case rpb.MapField_COUNTER:
				dt[string(v.Field.Name)] = v.CounterValue
			case rpb.MapField_SET:
				dt[string(v.Field.Name)] = v.SetValue
			case rpb.MapField_REGISTER:
				dt[string(v.Field.Name)] = v.RegisterValue
			case rpb.MapField_FLAG:
				dt[string(v.Field.Name)] = v.FlagValue
			case rpb.MapField_MAP:
				dt[string(v.Field.Name)] = v.MapValue
			}
		}
		return &mp{
			key:     key,
			context: rpbRes.Context,
			data:    dt,
		}, nil
	}

	return rpbRes, nil
}
