package rigo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tiloso/rigo"
)

type testObj struct {
	key, vclock []byte
	indexes     []rigo.KVPair
	Name        string
}

func (o *testObj) Key() []byte {
	return o.key
}

func (o *testObj) SetKey(v []byte) {
	o.key = v
}

func (o *testObj) Vclock() []byte {
	return o.vclock
}

func (o *testObj) SetVclock(v []byte) {
	o.vclock = v
}

func (o *testObj) Indexes() []rigo.KVPair {
	return o.indexes
}

var (
	client *rigo.Client

	typeName   = []byte("rigo_test")
	bucketName = []byte("rigo_test")
)

// func TestGet(t *testing.T) {
// 	var err error
// 	ms, err = rigo.Dial("127.0.0.1:8087", 1)
// 	if err != nil {
// 		t.Fatalf("err dialing: %v", err)
// 	}
//
// 	s := ms.Type(testType).Bucket(testBucket).Key([]byte("Rj0zGrn3oa2kgTqX99bSTNSXWho"))
//
// 	objs := []testObj{}
// 	if err := s.GetI(&objs); err != nil {
// 		t.Errorf("err c.Get: %v", err)
// 	}
// }

// func clearDB() error {
// 	keys := [][]byte{}
// 	cs := ms.Type(testType).Bucket(testBucket)
// 	dc, ec := cs.Keys()
//
// WaitResponse:
// 	for {
// 		select {
// 		case key, ok := <-dc:
// 			if !ok {
// 				break WaitResponse
// 			}
// 			keys = append(keys, key)
// 		case err := <-ec:
// 			return err
// 		}
// 	}
//
// 	for _, v := range keys {
// 		objs := []testObj{}
// 		if err := cs.Key(v).GetI(&objs); err != nil {
// 			return err
// 		}
//
// 		if len(objs) < 1 {
// 			return errors.New("no objs")
// 		}
//
// 		if len(objs) > 1 {
// 			var err error
// 			objs, err = resolveSiblings(objs)
// 			if err != nil {
// 				return err
// 			}
// 		}
//
// 		if err := cs.Key(objs[0].Key()).Vclock(objs[0].Vclock()).Delete(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func resolveSiblings(objs []testObj) ([]testObj, error) {
// 	cs := ms.Type(testType).Bucket(testBucket)
//
// 	for len(objs) > 1 {
// 		objs[0].Name = "resolved"
//
// 		if err := cs.StoreI(&objs[0]); err != nil {
// 			return nil, err
// 		}
//
// 		objs2 := []testObj{}
// 		if err := cs.Key(objs[0].Key()).GetI(&objs2); err != nil {
// 			return nil, err
// 		}
// 		objs = objs2
// 	}
//
// 	return objs, nil
// }

func TestDial2(t *testing.T) {
	c, err := rigo.Dial("127.0.0.1:8087", 1)
	if err != nil {
		t.Fatalf("err dialing: %v", err)
	}

	client = c
	// TODO
	// s.Close()

	// if err := clearDB(); err != nil {
	// 	t.Fatalf("err clearing db: %v", err)
	// }

	time.Sleep(3 * time.Second)
}

// requires bucket type setting `notfound_ok: true` to pass
func TestGetNotFound(t *testing.T) {
	key := []byte("not-existent")
	var objs []*testObj
	if err := client.T(typeName).B(bucketName).K(key).GetKVI(&objs); err != nil {
		t.Skipf("err get: %v", err)
	}

	if len(objs) > 0 {
		t.Errorf("got %v objs, expected 0", len(objs))
	}
}

// func TestStoreWithKey(t *testing.T) {
// 	s := ms.Type(testType).Bucket(testBucket)
// 	key := []byte("teststorewithkey")
//
// 	oldObjs := []testObj{}
// 	if err := s.Key(key).GetI(&oldObjs); err != nil {
// 		t.Errorf("err c.Get: %v", err)
// 	}
//
// 	if len(oldObjs) > 0 {
// 		t.Errorf("old objects present: %v", len(oldObjs))
// 	}
//
// 	o1 := &testObj{
// 		key:  key,
// 		Name: "TestStoreWithKey",
// 	}
//
// 	if err := s.StoreI(o1); err != nil {
// 		t.Errorf("err c.Store: %v", err)
// 	}
//
// 	objs := []testObj{}
// 	if err := s.Key(key).GetI(&objs); err != nil {
// 		t.Errorf("err second c.Get: %v", err)
// 	}
//
// 	if len(objs) != 1 {
// 		t.Errorf("got %v objs, expected 1", len(objs))
// 	}
//
// 	if objs[0].Vclock() == nil || len(objs[0].Vclock()) == 0 {
// 		t.Error("vclock not present")
// 	}
// }

// func TestStoreWithoutKey(t *testing.T) {
// 	t.Parallel()
// 	s := ms.Type(testType).Bucket(testBucket)
//
// 	o := &testObj{Name: "TestStoreWithoutKey"}
// 	if err := s.StoreI(o); err != nil {
// 		t.Errorf("err c.Store: %v", err)
// 	}
//
// 	if o.Key() == nil {
// 		t.Error("key empty, expect it to be not empty")
// 	}
//
// 	objs := []testObj{}
// 	if err := s.Key(o.Key()).GetI(&objs); err != nil {
// 		t.Errorf("err c.Get: %v", err)
// 	}
//
// 	if len(objs) != 1 {
// 		t.Errorf("got %v objs, expected 1", len(objs))
// 	}
//
// 	if objs[0].Name != o.Name {
// 		t.Errorf("names don't match: got %v, expected %v", objs[0].Name, o.Name)
// 	}
// }

// requires bucket type setting `allow_mult: true` to pass
// func TestStoreExistingWithoutVclock(t *testing.T) {
// 	t.Parallel()
// 	s := ms.Type(testType).Bucket(testBucket)
//
// 	names := []string{"a", "b", "c"}
// 	key := []byte("testupdatewithoutvclock")
// 	// should create siblings
// 	for i, n := range names {
// 		if err := s.StoreI(&testObj{key: key, Name: n}); err != nil {
// 			t.Errorf("err %v. c.Store: %v", i, err)
// 		}
// 	}
//
// 	objs := []testObj{}
// 	if err := s.Key(key).GetI(&objs); err != nil {
// 		t.Errorf("err c.Get: %v", err)
// 	}
//
// 	if len(objs) != 3 {
// 		t.Errorf("got %v objs, expected 3", len(objs))
// 	}
//
// 	// should resolve siblings
// 	if err := s.StoreI(&objs[0]); err != nil {
// 		t.Errorf("err c.Store: %v", err)
// 	}
//
// 	objs = []testObj{}
// 	if err := s.Key(key).GetI(&objs); err != nil {
// 		t.Errorf("err second c.Get: %v", err)
// 	}
//
// 	if len(objs) != 1 {
// 		t.Errorf("got %v objs, expected 1", len(objs))
// 	}
// }

// func TestStoreWithVclock(t *testing.T) {
// 	t.Parallel()
// 	s := ms.Type(testType).Bucket(testBucket)
// 	names := []string{"a", "b", "c"}
// 	key := []byte("TestUpdateWithVclock")
//
// 	o := &testObj{key: key}
// 	for i, v := range names {
// 		o.Name = v
// 		if err := s.StoreI(o); err != nil {
// 			t.Errorf("err %v. c.Store: %v", i, err)
// 		}
//
// 		objs := []testObj{}
// 		if err := s.Key(key).GetI(&objs); err != nil {
// 			t.Errorf("err %v c.Get: %v", i, err)
// 		}
//
// 		if len(objs) != 1 {
// 			t.Errorf("got %v objs, expected 1", len(objs))
// 		}
//
// 		o = &objs[0]
// 	}
// }

// func TestDelete(t *testing.T) {
// 	t.Parallel()
// 	s := ms.Type(testType).Bucket(testBucket)
// 	o := &testObj{
// 		Name: "test delete",
// 	}
//
// 	if err := s.StoreI(o); err != nil {
// 		t.Errorf("err c.Store: %v", err)
// 	}
//
// 	objs := []testObj{}
// 	if err := s.Key(o.Key()).GetI(&objs); err != nil {
// 		t.Errorf("err c.Get: %v", err)
// 	}
//
// 	if err := s.DeleteI(&objs[0]); err != nil {
// 		t.Errorf("err c.Delete: %v", err)
// 	}
// }

// func TestListBuckets(t *testing.T) {
// 	t.Parallel()
// 	buckets := [][]byte{}
// 	dc, ec := ms.Buckets()
//
// WaitResponse:
// 	for {
// 		select {
// 		case bucket, ok := <-dc:
// 			if !ok {
// 				break WaitResponse
// 			}
// 			buckets = append(buckets, bucket)
// 		case err := <-ec:
// 			t.Fatalf("err c.ListBuckets: %v", err)
// 		}
//
// 	}
//
// 	if len(buckets) < 1 {
// 		t.Errorf("got %v buckets, want > 0", len(buckets))
// 	}
// }

// func TestListKeys(t *testing.T) {
// 	t.Parallel()
// 	keys := [][]byte{}
// 	dc, ec := ms.Bucket(testBucket).Keys()
//
// WaitResponse:
// 	for {
// 		select {
// 		case key, ok := <-dc:
// 			if !ok {
// 				break WaitResponse
// 			}
// 			keys = append(keys, key)
// 		case err := <-ec:
// 			t.Fatalf("err c.ListKeys: %v", err)
// 		}
// 	}
//
// 	if len(keys) < 1 {
// 		t.Errorf("got %v keys, want > 0", len(keys))
// 	}
// }

// func storeObjsWithIndex() error {
// 	for i := 0; i < 2; i++ {
// 		obj := testObj{
// 			Name: "test index",
// 			indexes: []rigo.Index{
// 				rigo.Index{
// 					Key:   []byte("test_bin"),
// 					Value: []byte("b"),
// 				},
// 			},
// 		}
//
// 		if err := ms.Type(testType).Bucket(testBucket).StoreI(&obj); err != nil {
// 			return fmt.Errorf("err c.Store: %v", err)
// 		}
//
// 	}
// 	return nil
// }

// func TestIndexKey(t *testing.T) {
// 	s := ms.Type(testType).Bucket(testBucket)
// 	if err := storeObjsWithIndex(); err != nil {
// 		t.Fatalf("err storeObjsWithIndex: %v", err)
// 	}
//
// 	testIndex := []byte("test_bin")
// 	keys := [][]byte{}
//
// 	dc, ec := s.Index(testIndex).Key([]byte("b"))
//
// WaitResponse:
// 	for {
// 		select {
// 		case key, ok := <-dc:
// 			if !ok {
// 				break WaitResponse
// 			}
// 			keys = append(keys, key)
// 		case err := <-ec:
// 			t.Fatalf("err c.Index: %v", err)
// 		}
// 	}
//
// 	if len(keys) != 2 {
// 		t.Errorf("got %v keys, want 2", len(keys))
// 	}
// }

// func TestIndexRange(t *testing.T) {
// 	testIndex := []byte("test_bin")
// 	keys := [][]byte{}
//
// 	dc, ec := c.Index(testBucket, testIndex, []byte("a"), []byte("b"))
// WaitResponse:
// 	for {
// 		select {
// 		case key, ok := <-dc:
// 			if !ok {
// 				break WaitResponse
// 			}
// 			keys = append(keys, key)
// 		case err := <-ec:
// 			t.Fatalf("err c.Index: %v", err)
// 		}
// 	}
//
// 	if len(keys) != 2 {
// 		t.Errorf("got %v keys, want 2", len(keys))
// 	}
// }

func Example() {
	c, err := rigo.Dial("127.0.0.1:8087", 1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	bType := []byte("rigo-example-type")
	bucket := []byte("rigo-example-bucket")
	key := []byte("example")

	rgBucket := c.T([]byte("rigo-type")).B([]byte("rigo-bucket"))

	o := &testObj{
		key:  key,
		Name: "nexus",
	}
	if err := c.T(bType).B(bucket).StoreKVI(o); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if err := rgBucket.StoreKVI(o); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	v := []*testObj{&testObj{}}
	if err := rgBucket.K(key).GetKVI(v); err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// TODO
	// err := c.T("foo").B("bar").I("foo_bin").K("foobar").Get()
	// err := c.T("foo").B("bar").I("foo_bin").R("foobar").Get()

	// cType := rigo.ContentType("application/json")
	// enc := rigo.ContentEncoding("utf-8")

	// c.T("foo").B("bar").K("abc").O(enc, typ).Get()
}
