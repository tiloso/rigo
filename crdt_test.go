package rigo_test

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/tiloso/rigo"
)

var c *rigo.Client

// How to proceed
// - create helper which allows to unmarshal fetched data into structs (or whatever)
// - proceed with UpdateDT method => more datatypes, more tests, more examples

// main difference compared to riak KV =>
// only store delta / update operations vs store whole object again

// proceed with  more examples
// create new crdt, store, update, get, update, get, delete,...

// how would it work with existing stuff, e.g. User of trendxp?

// implement second datatype, e.g. counter (which may be integrated into map
// and stored/ fetched directly)

// is it possible to create an `or` interface{}? e.g. either Keyer or Vclocker

func TestCreateDT(t *testing.T) {
	var err error
	c, err = rigo.Dial("127.0.0.1:8087", 1)
	if err != nil {
		t.Fatalf("err dialing: %v", err)
	}

	// 	m := rigo.NewMap([]byte("n9"))
	// 	m.Update("phone", rigo.NewRegister([]byte("nexus 8")))
	// 	m.Update("brand", rigo.NewRegister([]byte("google")))
	// 	m.Update("views", rigo.NewCounter(1024))
	//
	// 	res, err := c.T([]byte("map")).B([]byte("crdt")).UpdateDT(m)
	// 	if err != nil {
	// 		t.Errorf("err: %s", err)
	// 	}
	//
	// 	fmt.Printf("res: %+v\n", res)
}

// func TestUpdateDT(t *testing.T) {
// 	rm, err := c.T([]byte("map")).B([]byte("crdt")).GetDT([]byte("n9"))
// 	if err != nil {
// 		t.Errorf("err: %s", err)
// 	}
//
// 	m := rm.(rigo.Map)
// 	m.Update("phone", rigo.NewRegister([]byte("nexus 6")))
// 	m.Update("release", rigo.NewRegister([]byte("2014")))
//
// 	res, err := c.T([]byte("map")).B([]byte("crdt")).UpdateDT(m)
// 	if err != nil {
// 		t.Errorf("err: %s", err)
// 	}
//
// 	fmt.Printf("res: %+v\n", res)
//
// 	// if err := c.T([]byte("map")).B([]byte("crdt")).K([]byte("a")).Delete(); err != nil {
// 	// 	t.Errorf("err: %s", err)
// 	// }
// }

// TODO
// rewrite mapstruct for my purpose => converts datatypes to go types
//   (directly without aggregating map)

// mapping riak datatypes => go types
// counter => int / int64?
// register => []byte
// flag => bool
// set => [][]byte
// map => map[string]interface{} || struct with

// how to set key / vclock on top level datatypes (map, counter, set)?
// how to limit types of map/interface{} at compiletime? => not; according to
//   mapstruct
// write simple interface which return simple datatypes + ctx (key not necessary)
//
// may be map[string]inter

type Counter int64

func (c Counter) SetContext(v []byte) {
}

func (c Counter) Context() []byte {
	return nil
}

// => not really possile => take interfaces as input which!
// (e.g. counter => Diff)

// PROCEED HERE
// updating props of a map type => how to? (e.g. for a user?)
// test with counters etc. => how to make sure we get the ∆? => interface!
// how do other drivers handle this?

type User struct {
	key, ctx []byte
	Phone    []byte
	Brand    []byte
	Views    *int64
}

func TestGetDT(t *testing.T) {
	rm, err := c.T([]byte("map")).B([]byte("crdt")).GetDT([]byte("n9"))
	if err != nil {
		t.Errorf("err: %s", err)
	}

	m := rm.(rigo.Map)

	u := User{}
	if err := mapstructure.Decode(m.Data(), &u); err != nil {
		t.Errorf("err: %s", err)
	}

	fmt.Printf("User: %+v\n", u)
	// TODO
	// assign Key() + Context()
}

// type User struct {
// 	k, c []byte
// 	up   map[string]interface{}
//
// 	Email  string
// 	Logins int
// }
//
// func (u *User) Updates() map[string]interface{} {
// 	return u.up
// }
//
// type Updater func() map[string]interface{}
//
// func (u Updater) Updates() map[string]interface{} {
// 	return u()
// }
//
// func SetUpdates(k string, v []byte) Updater {
// 	return func() map[string]interface{} {
// 		return map[string]interface{}{
// 			k: v,
// 		}
// 	}
// }
//
// func (u *User) UpdateEmail(v string) {
// 	updater := SetUpdates("email", v)
//
// 	// u.up = map[string]interface{}{
// 	// 	"Email": []byte(v),
// 	// }
//
// 	if err := c.T([]byte("map")).B([]byte("crdt")).UpdateDT(updater); err != nil {
// 		return err
// 	}
//
// 	u.Email = v
// }
//
// // kind of annoying... => simpler
// // update one prop only?
