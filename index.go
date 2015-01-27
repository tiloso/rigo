package rigo

import "github.com/tiloso/rigo/rpb"

type Index struct {
	*Bucket
	index          []byte
	paginationSort *bool
	returnTerms    *bool
	maxResults     *uint32
	timeout        *uint32
	rpbReq         *rpb.RpbIndexReq
}

func (i *Index) setTimeout(v uint32) {
	i.timeout = &v
}

func (i *Index) setMaxResults(v uint32) {
	i.maxResults = &v
}

func (i *Index) setPaginationSort(v bool) {
	i.paginationSort = &v
}

func (i *Index) Option(opts ...option) *Index {
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func (i *Index) Key(k []byte) *Index {
	i.rpbReq = &rpb.RpbIndexReq{
		Bucket:         i.bucket,
		Index:          i.index,
		Type:           i.typ,
		Qtype:          rpb.RpbIndexReq_eq.Enum(),
		Key:            k,
		PaginationSort: i.paginationSort,
	}
	return i
}

func (i *Index) Stream() (<-chan []byte, <-chan error) {
	i.rpbReq.Stream = &tval

	dc := make(chan []byte)
	ec := make(chan error)

	go func() {
		req, err := marshalRPB(rpbIndexReqCode, i.rpbReq)
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

func (i *Index) Range(start, end []byte) *Index {
	i.rpbReq = &rpb.RpbIndexReq{
		Bucket:         i.bucket,
		Index:          i.index,
		Type:           i.typ,
		Qtype:          rpb.RpbIndexReq_range.Enum(),
		RangeMin:       start,
		RangeMax:       end,
		PaginationSort: i.paginationSort,
	}
	return i
}
