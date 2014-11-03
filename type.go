package rigo

import "github.com/tiloso/rigo/rpb"

type Type struct {
	*Client
	typ []byte
}

func (t *Type) B(v []byte) *Bucket {
	return &Bucket{
		Type:   t,
		bucket: v,
	}
}

func (t *Type) Buckets() (<-chan []byte, <-chan error) {
	dc := make(chan []byte)
	ec := make(chan error)

	go func() {
		rpbReq := &rpb.RpbListBucketsReq{
			Type:   t.typ,
			Stream: &tval,
		}

		req, err := marshalRPB(rpbListBucketsReqCode, rpbReq)
		if err != nil {
			ec <- err
			return
		}

		s := t.session()
		defer s.release()

		if err := s.writeRequest(req); err != nil {
			ec <- err
			return
		}

		rpbRes := &rpb.RpbListBucketsResp{}
		for rpbRes.Done == nil || !*rpbRes.Done {
			res, err := s.readResponse()
			if err != nil {
				ec <- err
				return
			}

			if err := unmarshalRPB(res, rpbListBucketsResCode, rpbRes); err != nil {
				ec <- err
				return
			}

			for _, v := range rpbRes.Buckets {
				dc <- v
			}
		}
		close(dc)
	}()
	return dc, ec
}
