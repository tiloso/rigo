// Code generated by protoc-gen-go.
// source: riak_kv.proto
// DO NOT EDIT!

/*
Package riak_kv is a generated protocol buffer package.

It is generated from these files:
	riak_kv.proto

It has these top-level messages:
	RpbGetClientIdResp
	RpbSetClientIdReq
	RpbGetReq
	RpbGetResp
	RpbPutReq
	RpbPutResp
	RpbDelReq
	RpbListBucketsReq
	RpbListBucketsResp
	RpbListKeysReq
	RpbListKeysResp
	RpbMapRedReq
	RpbMapRedResp
	RpbIndexReq
	RpbIndexResp
	RpbCSBucketReq
	RpbCSBucketResp
	RpbIndexObject
	RpbContent
	RpbLink
	RpbCounterUpdateReq
	RpbCounterUpdateResp
	RpbCounterGetReq
	RpbCounterGetResp
*/
package rpb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RpbIndexReq_IndexQueryType int32

const (
	RpbIndexReq_eq    RpbIndexReq_IndexQueryType = 0
	RpbIndexReq_range RpbIndexReq_IndexQueryType = 1
)

var RpbIndexReq_IndexQueryType_name = map[int32]string{
	0: "eq",
	1: "range",
}
var RpbIndexReq_IndexQueryType_value = map[string]int32{
	"eq":    0,
	"range": 1,
}

func (x RpbIndexReq_IndexQueryType) Enum() *RpbIndexReq_IndexQueryType {
	p := new(RpbIndexReq_IndexQueryType)
	*p = x
	return p
}
func (x RpbIndexReq_IndexQueryType) String() string {
	return proto.EnumName(RpbIndexReq_IndexQueryType_name, int32(x))
}
func (x *RpbIndexReq_IndexQueryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpbIndexReq_IndexQueryType_value, data, "RpbIndexReq_IndexQueryType")
	if err != nil {
		return err
	}
	*x = RpbIndexReq_IndexQueryType(value)
	return nil
}

// Get ClientId Request - no message defined, just send RpbGetClientIdReq message code
type RpbGetClientIdResp struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetClientIdResp) Reset()         { *m = RpbGetClientIdResp{} }
func (m *RpbGetClientIdResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetClientIdResp) ProtoMessage()    {}

func (m *RpbGetClientIdResp) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

type RpbSetClientIdReq struct {
	ClientId         []byte `protobuf:"bytes,1,req,name=client_id" json:"client_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbSetClientIdReq) Reset()         { *m = RpbSetClientIdReq{} }
func (m *RpbSetClientIdReq) String() string { return proto.CompactTextString(m) }
func (*RpbSetClientIdReq) ProtoMessage()    {}

func (m *RpbSetClientIdReq) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

// Get Request - retrieve bucket/key
type RpbGetReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	R                *uint32 `protobuf:"varint,3,opt,name=r" json:"r,omitempty"`
	Pr               *uint32 `protobuf:"varint,4,opt,name=pr" json:"pr,omitempty"`
	BasicQuorum      *bool   `protobuf:"varint,5,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk       *bool   `protobuf:"varint,6,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	IfModified       []byte  `protobuf:"bytes,7,opt,name=if_modified" json:"if_modified,omitempty"`
	Head             *bool   `protobuf:"varint,8,opt,name=head" json:"head,omitempty"`
	Deletedvclock    *bool   `protobuf:"varint,9,opt,name=deletedvclock" json:"deletedvclock,omitempty"`
	Timeout          *uint32 `protobuf:"varint,10,opt,name=timeout" json:"timeout,omitempty"`
	SloppyQuorum     *bool   `protobuf:"varint,11,opt,name=sloppy_quorum" json:"sloppy_quorum,omitempty"`
	NVal             *uint32 `protobuf:"varint,12,opt,name=n_val" json:"n_val,omitempty"`
	Type             []byte  `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbGetReq) Reset()         { *m = RpbGetReq{} }
func (m *RpbGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbGetReq) ProtoMessage()    {}

func (m *RpbGetReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbGetReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbGetReq) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbGetReq) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbGetReq) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbGetReq) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

func (m *RpbGetReq) GetIfModified() []byte {
	if m != nil {
		return m.IfModified
	}
	return nil
}

func (m *RpbGetReq) GetHead() bool {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return false
}

func (m *RpbGetReq) GetDeletedvclock() bool {
	if m != nil && m.Deletedvclock != nil {
		return *m.Deletedvclock
	}
	return false
}

func (m *RpbGetReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbGetReq) GetSloppyQuorum() bool {
	if m != nil && m.SloppyQuorum != nil {
		return *m.SloppyQuorum
	}
	return false
}

func (m *RpbGetReq) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbGetReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get Response - if the record was not found there will be no content/vclock
type RpbGetResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Unchanged        *bool         `protobuf:"varint,3,opt,name=unchanged" json:"unchanged,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RpbGetResp) Reset()         { *m = RpbGetResp{} }
func (m *RpbGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbGetResp) ProtoMessage()    {}

func (m *RpbGetResp) GetContent() []*RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbGetResp) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbGetResp) GetUnchanged() bool {
	if m != nil && m.Unchanged != nil {
		return *m.Unchanged
	}
	return false
}

// Put request - if options.return_body is set then the updated metadata/data for
//               the key will be returned.
type RpbPutReq struct {
	Bucket           []byte      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Vclock           []byte      `protobuf:"bytes,3,opt,name=vclock" json:"vclock,omitempty"`
	Content          *RpbContent `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	W                *uint32     `protobuf:"varint,5,opt,name=w" json:"w,omitempty"`
	Dw               *uint32     `protobuf:"varint,6,opt,name=dw" json:"dw,omitempty"`
	ReturnBody       *bool       `protobuf:"varint,7,opt,name=return_body" json:"return_body,omitempty"`
	Pw               *uint32     `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	IfNotModified    *bool       `protobuf:"varint,9,opt,name=if_not_modified" json:"if_not_modified,omitempty"`
	IfNoneMatch      *bool       `protobuf:"varint,10,opt,name=if_none_match" json:"if_none_match,omitempty"`
	ReturnHead       *bool       `protobuf:"varint,11,opt,name=return_head" json:"return_head,omitempty"`
	Timeout          *uint32     `protobuf:"varint,12,opt,name=timeout" json:"timeout,omitempty"`
	Asis             *bool       `protobuf:"varint,13,opt,name=asis" json:"asis,omitempty"`
	SloppyQuorum     *bool       `protobuf:"varint,14,opt,name=sloppy_quorum" json:"sloppy_quorum,omitempty"`
	NVal             *uint32     `protobuf:"varint,15,opt,name=n_val" json:"n_val,omitempty"`
	Type             []byte      `protobuf:"bytes,16,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RpbPutReq) Reset()         { *m = RpbPutReq{} }
func (m *RpbPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbPutReq) ProtoMessage()    {}

func (m *RpbPutReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbPutReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbPutReq) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbPutReq) GetContent() *RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbPutReq) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbPutReq) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbPutReq) GetReturnBody() bool {
	if m != nil && m.ReturnBody != nil {
		return *m.ReturnBody
	}
	return false
}

func (m *RpbPutReq) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbPutReq) GetIfNotModified() bool {
	if m != nil && m.IfNotModified != nil {
		return *m.IfNotModified
	}
	return false
}

func (m *RpbPutReq) GetIfNoneMatch() bool {
	if m != nil && m.IfNoneMatch != nil {
		return *m.IfNoneMatch
	}
	return false
}

func (m *RpbPutReq) GetReturnHead() bool {
	if m != nil && m.ReturnHead != nil {
		return *m.ReturnHead
	}
	return false
}

func (m *RpbPutReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbPutReq) GetAsis() bool {
	if m != nil && m.Asis != nil {
		return *m.Asis
	}
	return false
}

func (m *RpbPutReq) GetSloppyQuorum() bool {
	if m != nil && m.SloppyQuorum != nil {
		return *m.SloppyQuorum
	}
	return false
}

func (m *RpbPutReq) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbPutReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Put response - same as get response with optional key if one was generated
type RpbPutResp struct {
	Content          []*RpbContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Vclock           []byte        `protobuf:"bytes,2,opt,name=vclock" json:"vclock,omitempty"`
	Key              []byte        `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *RpbPutResp) Reset()         { *m = RpbPutResp{} }
func (m *RpbPutResp) String() string { return proto.CompactTextString(m) }
func (*RpbPutResp) ProtoMessage()    {}

func (m *RpbPutResp) GetContent() []*RpbContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RpbPutResp) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbPutResp) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// Delete request
type RpbDelReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Rw               *uint32 `protobuf:"varint,3,opt,name=rw" json:"rw,omitempty"`
	Vclock           []byte  `protobuf:"bytes,4,opt,name=vclock" json:"vclock,omitempty"`
	R                *uint32 `protobuf:"varint,5,opt,name=r" json:"r,omitempty"`
	W                *uint32 `protobuf:"varint,6,opt,name=w" json:"w,omitempty"`
	Pr               *uint32 `protobuf:"varint,7,opt,name=pr" json:"pr,omitempty"`
	Pw               *uint32 `protobuf:"varint,8,opt,name=pw" json:"pw,omitempty"`
	Dw               *uint32 `protobuf:"varint,9,opt,name=dw" json:"dw,omitempty"`
	Timeout          *uint32 `protobuf:"varint,10,opt,name=timeout" json:"timeout,omitempty"`
	SloppyQuorum     *bool   `protobuf:"varint,11,opt,name=sloppy_quorum" json:"sloppy_quorum,omitempty"`
	NVal             *uint32 `protobuf:"varint,12,opt,name=n_val" json:"n_val,omitempty"`
	Type             []byte  `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbDelReq) Reset()         { *m = RpbDelReq{} }
func (m *RpbDelReq) String() string { return proto.CompactTextString(m) }
func (*RpbDelReq) ProtoMessage()    {}

func (m *RpbDelReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbDelReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbDelReq) GetRw() uint32 {
	if m != nil && m.Rw != nil {
		return *m.Rw
	}
	return 0
}

func (m *RpbDelReq) GetVclock() []byte {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *RpbDelReq) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbDelReq) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbDelReq) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbDelReq) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbDelReq) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbDelReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbDelReq) GetSloppyQuorum() bool {
	if m != nil && m.SloppyQuorum != nil {
		return *m.SloppyQuorum
	}
	return false
}

func (m *RpbDelReq) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbDelReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// List buckets request
type RpbListBucketsReq struct {
	Timeout          *uint32 `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
	Stream           *bool   `protobuf:"varint,2,opt,name=stream" json:"stream,omitempty"`
	Type             []byte  `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbListBucketsReq) Reset()         { *m = RpbListBucketsReq{} }
func (m *RpbListBucketsReq) String() string { return proto.CompactTextString(m) }
func (*RpbListBucketsReq) ProtoMessage()    {}

func (m *RpbListBucketsReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbListBucketsReq) GetStream() bool {
	if m != nil && m.Stream != nil {
		return *m.Stream
	}
	return false
}

func (m *RpbListBucketsReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// List buckets response - one or more of these packets will be sent
// the last one will have done set true (and may not have any buckets in it)
type RpbListBucketsResp struct {
	Buckets          [][]byte `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	Done             *bool    `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbListBucketsResp) Reset()         { *m = RpbListBucketsResp{} }
func (m *RpbListBucketsResp) String() string { return proto.CompactTextString(m) }
func (*RpbListBucketsResp) ProtoMessage()    {}

func (m *RpbListBucketsResp) GetBuckets() [][]byte {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *RpbListBucketsResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

// List keys in bucket request
type RpbListKeysReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Timeout          *uint32 `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	Type             []byte  `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbListKeysReq) Reset()         { *m = RpbListKeysReq{} }
func (m *RpbListKeysReq) String() string { return proto.CompactTextString(m) }
func (*RpbListKeysReq) ProtoMessage()    {}

func (m *RpbListKeysReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbListKeysReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbListKeysReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// List keys in bucket response - one or more of these packets will be sent
// the last one will have done set true (and may not have any keys in it)
type RpbListKeysResp struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Done             *bool    `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RpbListKeysResp) Reset()         { *m = RpbListKeysResp{} }
func (m *RpbListKeysResp) String() string { return proto.CompactTextString(m) }
func (*RpbListKeysResp) ProtoMessage()    {}

func (m *RpbListKeysResp) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RpbListKeysResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

// Map/Reduce request
type RpbMapRedReq struct {
	Request          []byte `protobuf:"bytes,1,req,name=request" json:"request,omitempty"`
	ContentType      []byte `protobuf:"bytes,2,req,name=content_type" json:"content_type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbMapRedReq) Reset()         { *m = RpbMapRedReq{} }
func (m *RpbMapRedReq) String() string { return proto.CompactTextString(m) }
func (*RpbMapRedReq) ProtoMessage()    {}

func (m *RpbMapRedReq) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RpbMapRedReq) GetContentType() []byte {
	if m != nil {
		return m.ContentType
	}
	return nil
}

// Map/Reduce response
// one or more of these packets will be sent the last one will have done set
// true (and may not have phase/data in it)
type RpbMapRedResp struct {
	Phase            *uint32 `protobuf:"varint,1,opt,name=phase" json:"phase,omitempty"`
	Response         []byte  `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Done             *bool   `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbMapRedResp) Reset()         { *m = RpbMapRedResp{} }
func (m *RpbMapRedResp) String() string { return proto.CompactTextString(m) }
func (*RpbMapRedResp) ProtoMessage()    {}

func (m *RpbMapRedResp) GetPhase() uint32 {
	if m != nil && m.Phase != nil {
		return *m.Phase
	}
	return 0
}

func (m *RpbMapRedResp) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RpbMapRedResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

// Secondary Index query request
type RpbIndexReq struct {
	Bucket       []byte                      `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Index        []byte                      `protobuf:"bytes,2,req,name=index" json:"index,omitempty"`
	Qtype        *RpbIndexReq_IndexQueryType `protobuf:"varint,3,req,name=qtype,enum=RpbIndexReq_IndexQueryType" json:"qtype,omitempty"`
	Key          []byte                      `protobuf:"bytes,4,opt,name=key" json:"key,omitempty"`
	RangeMin     []byte                      `protobuf:"bytes,5,opt,name=range_min" json:"range_min,omitempty"`
	RangeMax     []byte                      `protobuf:"bytes,6,opt,name=range_max" json:"range_max,omitempty"`
	ReturnTerms  *bool                       `protobuf:"varint,7,opt,name=return_terms" json:"return_terms,omitempty"`
	Stream       *bool                       `protobuf:"varint,8,opt,name=stream" json:"stream,omitempty"`
	MaxResults   *uint32                     `protobuf:"varint,9,opt,name=max_results" json:"max_results,omitempty"`
	Continuation []byte                      `protobuf:"bytes,10,opt,name=continuation" json:"continuation,omitempty"`
	Timeout      *uint32                     `protobuf:"varint,11,opt,name=timeout" json:"timeout,omitempty"`
	Type         []byte                      `protobuf:"bytes,12,opt,name=type" json:"type,omitempty"`
	TermRegex    []byte                      `protobuf:"bytes,13,opt,name=term_regex" json:"term_regex,omitempty"`
	// Whether to use pagination sort for non-paginated queries
	PaginationSort   *bool  `protobuf:"varint,14,opt,name=pagination_sort" json:"pagination_sort,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbIndexReq) Reset()         { *m = RpbIndexReq{} }
func (m *RpbIndexReq) String() string { return proto.CompactTextString(m) }
func (*RpbIndexReq) ProtoMessage()    {}

func (m *RpbIndexReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbIndexReq) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *RpbIndexReq) GetQtype() RpbIndexReq_IndexQueryType {
	if m != nil && m.Qtype != nil {
		return *m.Qtype
	}
	return RpbIndexReq_eq
}

func (m *RpbIndexReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbIndexReq) GetRangeMin() []byte {
	if m != nil {
		return m.RangeMin
	}
	return nil
}

func (m *RpbIndexReq) GetRangeMax() []byte {
	if m != nil {
		return m.RangeMax
	}
	return nil
}

func (m *RpbIndexReq) GetReturnTerms() bool {
	if m != nil && m.ReturnTerms != nil {
		return *m.ReturnTerms
	}
	return false
}

func (m *RpbIndexReq) GetStream() bool {
	if m != nil && m.Stream != nil {
		return *m.Stream
	}
	return false
}

func (m *RpbIndexReq) GetMaxResults() uint32 {
	if m != nil && m.MaxResults != nil {
		return *m.MaxResults
	}
	return 0
}

func (m *RpbIndexReq) GetContinuation() []byte {
	if m != nil {
		return m.Continuation
	}
	return nil
}

func (m *RpbIndexReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbIndexReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RpbIndexReq) GetTermRegex() []byte {
	if m != nil {
		return m.TermRegex
	}
	return nil
}

func (m *RpbIndexReq) GetPaginationSort() bool {
	if m != nil && m.PaginationSort != nil {
		return *m.PaginationSort
	}
	return false
}

// Secondary Index query response
type RpbIndexResp struct {
	Keys             [][]byte   `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Results          []*RpbPair `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	Continuation     []byte     `protobuf:"bytes,3,opt,name=continuation" json:"continuation,omitempty"`
	Done             *bool      `protobuf:"varint,4,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbIndexResp) Reset()         { *m = RpbIndexResp{} }
func (m *RpbIndexResp) String() string { return proto.CompactTextString(m) }
func (*RpbIndexResp) ProtoMessage()    {}

func (m *RpbIndexResp) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RpbIndexResp) GetResults() []*RpbPair {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *RpbIndexResp) GetContinuation() []byte {
	if m != nil {
		return m.Continuation
	}
	return nil
}

func (m *RpbIndexResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

// added solely for riak_cs currently
// for folding over a bucket and returning
// objects.
type RpbCSBucketReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	StartKey         []byte  `protobuf:"bytes,2,req,name=start_key" json:"start_key,omitempty"`
	EndKey           []byte  `protobuf:"bytes,3,opt,name=end_key" json:"end_key,omitempty"`
	StartIncl        *bool   `protobuf:"varint,4,opt,name=start_incl,def=1" json:"start_incl,omitempty"`
	EndIncl          *bool   `protobuf:"varint,5,opt,name=end_incl,def=0" json:"end_incl,omitempty"`
	Continuation     []byte  `protobuf:"bytes,6,opt,name=continuation" json:"continuation,omitempty"`
	MaxResults       *uint32 `protobuf:"varint,7,opt,name=max_results" json:"max_results,omitempty"`
	Timeout          *uint32 `protobuf:"varint,8,opt,name=timeout" json:"timeout,omitempty"`
	Type             []byte  `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbCSBucketReq) Reset()         { *m = RpbCSBucketReq{} }
func (m *RpbCSBucketReq) String() string { return proto.CompactTextString(m) }
func (*RpbCSBucketReq) ProtoMessage()    {}

const Default_RpbCSBucketReq_StartIncl bool = true
const Default_RpbCSBucketReq_EndIncl bool = false

func (m *RpbCSBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbCSBucketReq) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *RpbCSBucketReq) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

func (m *RpbCSBucketReq) GetStartIncl() bool {
	if m != nil && m.StartIncl != nil {
		return *m.StartIncl
	}
	return Default_RpbCSBucketReq_StartIncl
}

func (m *RpbCSBucketReq) GetEndIncl() bool {
	if m != nil && m.EndIncl != nil {
		return *m.EndIncl
	}
	return Default_RpbCSBucketReq_EndIncl
}

func (m *RpbCSBucketReq) GetContinuation() []byte {
	if m != nil {
		return m.Continuation
	}
	return nil
}

func (m *RpbCSBucketReq) GetMaxResults() uint32 {
	if m != nil && m.MaxResults != nil {
		return *m.MaxResults
	}
	return 0
}

func (m *RpbCSBucketReq) GetTimeout() uint32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *RpbCSBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// return for CS bucket fold
type RpbCSBucketResp struct {
	Objects          []*RpbIndexObject `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty"`
	Continuation     []byte            `protobuf:"bytes,2,opt,name=continuation" json:"continuation,omitempty"`
	Done             *bool             `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RpbCSBucketResp) Reset()         { *m = RpbCSBucketResp{} }
func (m *RpbCSBucketResp) String() string { return proto.CompactTextString(m) }
func (*RpbCSBucketResp) ProtoMessage()    {}

func (m *RpbCSBucketResp) GetObjects() []*RpbIndexObject {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *RpbCSBucketResp) GetContinuation() []byte {
	if m != nil {
		return m.Continuation
	}
	return nil
}

func (m *RpbCSBucketResp) GetDone() bool {
	if m != nil && m.Done != nil {
		return *m.Done
	}
	return false
}

type RpbIndexObject struct {
	Key              []byte      `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Object           *RpbGetResp `protobuf:"bytes,2,req,name=object" json:"object,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RpbIndexObject) Reset()         { *m = RpbIndexObject{} }
func (m *RpbIndexObject) String() string { return proto.CompactTextString(m) }
func (*RpbIndexObject) ProtoMessage()    {}

func (m *RpbIndexObject) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbIndexObject) GetObject() *RpbGetResp {
	if m != nil {
		return m.Object
	}
	return nil
}

// Content message included in get/put responses
// Holds the value and associated metadata
type RpbContent struct {
	Value            []byte     `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	ContentType      []byte     `protobuf:"bytes,2,opt,name=content_type" json:"content_type,omitempty"`
	Charset          []byte     `protobuf:"bytes,3,opt,name=charset" json:"charset,omitempty"`
	ContentEncoding  []byte     `protobuf:"bytes,4,opt,name=content_encoding" json:"content_encoding,omitempty"`
	Vtag             []byte     `protobuf:"bytes,5,opt,name=vtag" json:"vtag,omitempty"`
	Links            []*RpbLink `protobuf:"bytes,6,rep,name=links" json:"links,omitempty"`
	LastMod          *uint32    `protobuf:"varint,7,opt,name=last_mod" json:"last_mod,omitempty"`
	LastModUsecs     *uint32    `protobuf:"varint,8,opt,name=last_mod_usecs" json:"last_mod_usecs,omitempty"`
	Usermeta         []*RpbPair `protobuf:"bytes,9,rep,name=usermeta" json:"usermeta,omitempty"`
	Indexes          []*RpbPair `protobuf:"bytes,10,rep,name=indexes" json:"indexes,omitempty"`
	Deleted          *bool      `protobuf:"varint,11,opt,name=deleted" json:"deleted,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbContent) Reset()         { *m = RpbContent{} }
func (m *RpbContent) String() string { return proto.CompactTextString(m) }
func (*RpbContent) ProtoMessage()    {}

func (m *RpbContent) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RpbContent) GetContentType() []byte {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *RpbContent) GetCharset() []byte {
	if m != nil {
		return m.Charset
	}
	return nil
}

func (m *RpbContent) GetContentEncoding() []byte {
	if m != nil {
		return m.ContentEncoding
	}
	return nil
}

func (m *RpbContent) GetVtag() []byte {
	if m != nil {
		return m.Vtag
	}
	return nil
}

func (m *RpbContent) GetLinks() []*RpbLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *RpbContent) GetLastMod() uint32 {
	if m != nil && m.LastMod != nil {
		return *m.LastMod
	}
	return 0
}

func (m *RpbContent) GetLastModUsecs() uint32 {
	if m != nil && m.LastModUsecs != nil {
		return *m.LastModUsecs
	}
	return 0
}

func (m *RpbContent) GetUsermeta() []*RpbPair {
	if m != nil {
		return m.Usermeta
	}
	return nil
}

func (m *RpbContent) GetIndexes() []*RpbPair {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *RpbContent) GetDeleted() bool {
	if m != nil && m.Deleted != nil {
		return *m.Deleted
	}
	return false
}

// Link metadata
type RpbLink struct {
	Bucket           []byte `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Key              []byte `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Tag              []byte `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbLink) Reset()         { *m = RpbLink{} }
func (m *RpbLink) String() string { return proto.CompactTextString(m) }
func (*RpbLink) ProtoMessage()    {}

func (m *RpbLink) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbLink) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbLink) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

// Counter update request
type RpbCounterUpdateReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Amount           *int64  `protobuf:"zigzag64,3,req,name=amount" json:"amount,omitempty"`
	W                *uint32 `protobuf:"varint,4,opt,name=w" json:"w,omitempty"`
	Dw               *uint32 `protobuf:"varint,5,opt,name=dw" json:"dw,omitempty"`
	Pw               *uint32 `protobuf:"varint,6,opt,name=pw" json:"pw,omitempty"`
	Returnvalue      *bool   `protobuf:"varint,7,opt,name=returnvalue" json:"returnvalue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbCounterUpdateReq) Reset()         { *m = RpbCounterUpdateReq{} }
func (m *RpbCounterUpdateReq) String() string { return proto.CompactTextString(m) }
func (*RpbCounterUpdateReq) ProtoMessage()    {}

func (m *RpbCounterUpdateReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbCounterUpdateReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbCounterUpdateReq) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *RpbCounterUpdateReq) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbCounterUpdateReq) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbCounterUpdateReq) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbCounterUpdateReq) GetReturnvalue() bool {
	if m != nil && m.Returnvalue != nil {
		return *m.Returnvalue
	}
	return false
}

// Counter update response? No message | error response
type RpbCounterUpdateResp struct {
	Value            *int64 `protobuf:"zigzag64,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbCounterUpdateResp) Reset()         { *m = RpbCounterUpdateResp{} }
func (m *RpbCounterUpdateResp) String() string { return proto.CompactTextString(m) }
func (*RpbCounterUpdateResp) ProtoMessage()    {}

func (m *RpbCounterUpdateResp) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

// counter value
type RpbCounterGetReq struct {
	Bucket           []byte  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	R                *uint32 `protobuf:"varint,3,opt,name=r" json:"r,omitempty"`
	Pr               *uint32 `protobuf:"varint,4,opt,name=pr" json:"pr,omitempty"`
	BasicQuorum      *bool   `protobuf:"varint,5,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk       *bool   `protobuf:"varint,6,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbCounterGetReq) Reset()         { *m = RpbCounterGetReq{} }
func (m *RpbCounterGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbCounterGetReq) ProtoMessage()    {}

func (m *RpbCounterGetReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbCounterGetReq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbCounterGetReq) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbCounterGetReq) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbCounterGetReq) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbCounterGetReq) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

// Counter value response
type RpbCounterGetResp struct {
	Value            *int64 `protobuf:"zigzag64,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbCounterGetResp) Reset()         { *m = RpbCounterGetResp{} }
func (m *RpbCounterGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbCounterGetResp) ProtoMessage()    {}

func (m *RpbCounterGetResp) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("RpbIndexReq_IndexQueryType", RpbIndexReq_IndexQueryType_name, RpbIndexReq_IndexQueryType_value)
}
