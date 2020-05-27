// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ObjectRef
func (this *ObjectRef) MarshalJSON() ([]byte, error) {
	str, err := CoreMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ObjectRef
func (this *ObjectRef) UnmarshalJSON(b []byte) error {
	return CoreUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Status
func (this *Status) MarshalJSON() ([]byte, error) {
	str, err := CoreMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Status
func (this *Status) UnmarshalJSON(b []byte) error {
	return CoreUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CoreMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	CoreUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)