package rigo_test

// var c2 *rigo.Client
//
// func TestSetID(t *testing.T) {
// 	c2 = rigo.NewClient("127.0.0.1:8087", "rigo", 1)
//
// 	if err := c2.Dial(); err != nil {
// 		t.Fatalf("error dialing c: %v", err)
// 	}
//
// 	clientID := []byte("testclientid")
//
// 	if err := c2.SetID(clientID); err != nil {
// 		t.Errorf("error SetID(): %v", err)
// 	}
//
// 	cID, err := c2.ID()
// 	if err != nil {
// 		t.Errorf("error ID(): %v", err)
// 	}
//
// 	if !bytes.Equal(cID, clientID) {
// 		t.Errorf("client ids not equal: got %v, expected %v", cID, clientID)
// 	}
// }
//
// func TestPing(t *testing.T) {
// 	if err := c2.Ping(); err != nil {
// 		t.Errorf("err: %v", err)
// 	}
// }
