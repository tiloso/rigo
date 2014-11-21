package rigo

import "github.com/tiloso/rigo/rpb"

type Bucket struct {
	*Type
	bucket []byte
}

func (b *Bucket) K(v []byte) *Object {
	return &Object{
		Bucket: b,
		key:    v,
	}
}

func (b *Bucket) I(v []byte) *Index {
	return &Index{
		Bucket: b,
		index:  v,
	}
}

func (b *Bucket) Keys() (<-chan []byte, <-chan error) {
	dc := make(chan []byte)
	ec := make(chan error)

	go func() {
		rpbReq := &rpb.RpbListKeysReq{
			Bucket: b.bucket,
			Type:   b.typ,
		}

		req, err := marshalRPB(rpbListKeysReqCode, rpbReq)
		if err != nil {
			ec <- err
			return
		}

		s := b.session()
		defer s.release()
		if err := s.writeRequest(req); err != nil {
			ec <- err
			return
		}

		rpbRes := &rpb.RpbListKeysResp{}

		for rpbRes.Done == nil || !*rpbRes.Done {
			res, err := s.readResponse()
			if err != nil {
				ec <- err
				return
			}

			if err := unmarshalRPB(res, rpbListKeysResCode, rpbRes); err != nil {
				ec <- err
				return
			}

			for _, v := range rpbRes.Keys {
				dc <- v
			}
		}
		close(dc)
	}()
	return dc, ec
}
