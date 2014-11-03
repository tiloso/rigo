package rigo

import "github.com/tiloso/rigo/rpb"

type Index struct {
	*Bucket
	index []byte
	// 	client *Client
	// 	typ    []byte
	// 	bucket []byte
	// 	index  []byte
}

func (i *Index) Key(k []byte) (<-chan []byte, <-chan error) {
	dc := make(chan []byte)
	ec := make(chan error)

	go func() {
		rpbReq := &rpb.RpbIndexReq{
			Bucket: i.bucket,
			Index:  i.index,
			Stream: &tval,
			Type:   i.typ,
			Qtype:  rpb.RpbIndexReq_eq.Enum(),
			Key:    k,
		}

		req, err := marshalRPB(rpbIndexReqCode, rpbReq)
		if err != nil {
			ec <- err
			return
		}

		s := i.session()
		defer s.release()
		if err := s.writeRequest(req); err != nil {
			ec <- err
			return
		}

		rpbRes := &rpb.RpbIndexResp{}

		for rpbRes.Done == nil || !*rpbRes.Done {
			res, err := s.readResponse()
			if err != nil {
				ec <- err
				return
			}

			if err := unmarshalRPB(res, rpbIndexResCode, rpbRes); err != nil {
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

func (i *Index) Range(start, end []byte) (<-chan []byte, <-chan error) {
	dc := make(chan []byte)
	ec := make(chan error)

	go func() {
		rpbReq := &rpb.RpbIndexReq{
			Bucket:   i.bucket,
			Index:    i.index,
			Stream:   &tval,
			Type:     i.typ,
			Qtype:    rpb.RpbIndexReq_range.Enum(),
			RangeMin: start,
			RangeMax: end,
		}

		req, err := marshalRPB(rpbIndexReqCode, rpbReq)
		if err != nil {
			ec <- err
			return
		}

		s := i.session()
		defer s.release()
		if err := s.writeRequest(req); err != nil {
			ec <- err
			return
		}

		rpbRes := &rpb.RpbIndexResp{}

		for rpbRes.Done == nil || !*rpbRes.Done {
			res, err := s.readResponse()
			if err != nil {
				ec <- err
				return
			}

			if err := unmarshalRPB(res, rpbIndexResCode, rpbRes); err != nil {
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
