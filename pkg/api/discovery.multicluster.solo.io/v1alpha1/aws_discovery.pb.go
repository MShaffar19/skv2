// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/multicluster/discovery/v1alpha1/aws_discovery.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Discovery configuration for AWS resources.
type AwsDiscoverySpec struct {
	// If unspecified, by default discovery will run for EKS clusters in all regions.
	EksSelector          []*AwsDiscoverySpec_AwsResourceSelector `protobuf:"bytes,2,rep,name=eks_selector,json=eksSelector,proto3" json:"eks_selector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AwsDiscoverySpec) Reset()         { *m = AwsDiscoverySpec{} }
func (m *AwsDiscoverySpec) String() string { return proto.CompactTextString(m) }
func (*AwsDiscoverySpec) ProtoMessage()    {}
func (*AwsDiscoverySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8f6d4627dc227, []int{0}
}
func (m *AwsDiscoverySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsDiscoverySpec.Unmarshal(m, b)
}
func (m *AwsDiscoverySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsDiscoverySpec.Marshal(b, m, deterministic)
}
func (m *AwsDiscoverySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsDiscoverySpec.Merge(m, src)
}
func (m *AwsDiscoverySpec) XXX_Size() int {
	return xxx_messageInfo_AwsDiscoverySpec.Size(m)
}
func (m *AwsDiscoverySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsDiscoverySpec.DiscardUnknown(m)
}

var xxx_messageInfo_AwsDiscoverySpec proto.InternalMessageInfo

func (m *AwsDiscoverySpec) GetEksSelector() []*AwsDiscoverySpec_AwsResourceSelector {
	if m != nil {
		return m.EksSelector
	}
	return nil
}

// For a given resource_selector to apply to a resource, the resource must match all of the resource_selector's match criteria.
type AwsDiscoverySpec_AwsResourceSelector struct {
	// Types that are valid to be assigned to SelectorType:
	//	*AwsDiscoverySpec_AwsResourceSelector_Arn
	//	*AwsDiscoverySpec_AwsResourceSelector_Matcher_
	SelectorType         isAwsDiscoverySpec_AwsResourceSelector_SelectorType `protobuf_oneof:"selector_type"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AwsDiscoverySpec_AwsResourceSelector) Reset()         { *m = AwsDiscoverySpec_AwsResourceSelector{} }
func (m *AwsDiscoverySpec_AwsResourceSelector) String() string { return proto.CompactTextString(m) }
func (*AwsDiscoverySpec_AwsResourceSelector) ProtoMessage()    {}
func (*AwsDiscoverySpec_AwsResourceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8f6d4627dc227, []int{0, 0}
}
func (m *AwsDiscoverySpec_AwsResourceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector.Unmarshal(m, b)
}
func (m *AwsDiscoverySpec_AwsResourceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector.Marshal(b, m, deterministic)
}
func (m *AwsDiscoverySpec_AwsResourceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector.Merge(m, src)
}
func (m *AwsDiscoverySpec_AwsResourceSelector) XXX_Size() int {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector.Size(m)
}
func (m *AwsDiscoverySpec_AwsResourceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector proto.InternalMessageInfo

type isAwsDiscoverySpec_AwsResourceSelector_SelectorType interface {
	isAwsDiscoverySpec_AwsResourceSelector_SelectorType()
	Equal(interface{}) bool
}

type AwsDiscoverySpec_AwsResourceSelector_Arn struct {
	Arn string `protobuf:"bytes,1,opt,name=arn,proto3,oneof" json:"arn,omitempty"`
}
type AwsDiscoverySpec_AwsResourceSelector_Matcher_ struct {
	Matcher *AwsDiscoverySpec_AwsResourceSelector_Matcher `protobuf:"bytes,2,opt,name=matcher,proto3,oneof" json:"matcher,omitempty"`
}

func (*AwsDiscoverySpec_AwsResourceSelector_Arn) isAwsDiscoverySpec_AwsResourceSelector_SelectorType() {
}
func (*AwsDiscoverySpec_AwsResourceSelector_Matcher_) isAwsDiscoverySpec_AwsResourceSelector_SelectorType() {
}

func (m *AwsDiscoverySpec_AwsResourceSelector) GetSelectorType() isAwsDiscoverySpec_AwsResourceSelector_SelectorType {
	if m != nil {
		return m.SelectorType
	}
	return nil
}

func (m *AwsDiscoverySpec_AwsResourceSelector) GetArn() string {
	if x, ok := m.GetSelectorType().(*AwsDiscoverySpec_AwsResourceSelector_Arn); ok {
		return x.Arn
	}
	return ""
}

func (m *AwsDiscoverySpec_AwsResourceSelector) GetMatcher() *AwsDiscoverySpec_AwsResourceSelector_Matcher {
	if x, ok := m.GetSelectorType().(*AwsDiscoverySpec_AwsResourceSelector_Matcher_); ok {
		return x.Matcher
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AwsDiscoverySpec_AwsResourceSelector) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AwsDiscoverySpec_AwsResourceSelector_Arn)(nil),
		(*AwsDiscoverySpec_AwsResourceSelector_Matcher_)(nil),
	}
}

// Selects all resources that exist in the specified AWS region and possess the specified tags.
type AwsDiscoverySpec_AwsResourceSelector_Matcher struct {
	// AWS account IDs. If unspecified, select across any accessible AWS account.
	AccountIds []string `protobuf:"bytes,1,rep,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
	// AWS regions, e.g. us-east-2. If unspecified, select across all regions.
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	// AWS resource tags. If unspecified, match any tags.
	Tags                 map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) Reset() {
	*m = AwsDiscoverySpec_AwsResourceSelector_Matcher{}
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) String() string {
	return proto.CompactTextString(m)
}
func (*AwsDiscoverySpec_AwsResourceSelector_Matcher) ProtoMessage() {}
func (*AwsDiscoverySpec_AwsResourceSelector_Matcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8f6d4627dc227, []int{0, 0, 0}
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher.Unmarshal(m, b)
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher.Marshal(b, m, deterministic)
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher.Merge(m, src)
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) XXX_Size() int {
	return xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher.Size(m)
}
func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher.DiscardUnknown(m)
}

var xxx_messageInfo_AwsDiscoverySpec_AwsResourceSelector_Matcher proto.InternalMessageInfo

func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) GetAccountIds() []string {
	if m != nil {
		return m.AccountIds
	}
	return nil
}

func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *AwsDiscoverySpec_AwsResourceSelector_Matcher) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type AwsDiscoveryStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsDiscoveryStatus) Reset()         { *m = AwsDiscoveryStatus{} }
func (m *AwsDiscoveryStatus) String() string { return proto.CompactTextString(m) }
func (*AwsDiscoveryStatus) ProtoMessage()    {}
func (*AwsDiscoveryStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8f6d4627dc227, []int{1}
}
func (m *AwsDiscoveryStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsDiscoveryStatus.Unmarshal(m, b)
}
func (m *AwsDiscoveryStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsDiscoveryStatus.Marshal(b, m, deterministic)
}
func (m *AwsDiscoveryStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsDiscoveryStatus.Merge(m, src)
}
func (m *AwsDiscoveryStatus) XXX_Size() int {
	return xxx_messageInfo_AwsDiscoveryStatus.Size(m)
}
func (m *AwsDiscoveryStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsDiscoveryStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AwsDiscoveryStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AwsDiscoverySpec)(nil), "discovery.multicluster.solo.io.AwsDiscoverySpec")
	proto.RegisterType((*AwsDiscoverySpec_AwsResourceSelector)(nil), "discovery.multicluster.solo.io.AwsDiscoverySpec.AwsResourceSelector")
	proto.RegisterType((*AwsDiscoverySpec_AwsResourceSelector_Matcher)(nil), "discovery.multicluster.solo.io.AwsDiscoverySpec.AwsResourceSelector.Matcher")
	proto.RegisterMapType((map[string]string)(nil), "discovery.multicluster.solo.io.AwsDiscoverySpec.AwsResourceSelector.Matcher.TagsEntry")
	proto.RegisterType((*AwsDiscoveryStatus)(nil), "discovery.multicluster.solo.io.AwsDiscoveryStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/skv2/api/multicluster/discovery/v1alpha1/aws_discovery.proto", fileDescriptor_c0b8f6d4627dc227)
}

var fileDescriptor_c0b8f6d4627dc227 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xeb, 0x3a, 0x10, 0xf9, 0x06, 0x44, 0x35, 0x64, 0x61, 0x79, 0x51, 0xa2, 0xae, 0xb2,
	0x61, 0x46, 0x0d, 0x0b, 0x10, 0x3b, 0xaa, 0x22, 0x8a, 0x04, 0x42, 0x4c, 0x11, 0x0b, 0x36, 0xd1,
	0x74, 0x72, 0x35, 0x19, 0xec, 0x64, 0xac, 0xb9, 0xe3, 0xb4, 0x7e, 0x09, 0xc4, 0x63, 0xf0, 0x08,
	0x3c, 0x0f, 0x3b, 0x1e, 0x80, 0x3d, 0xb2, 0x1d, 0x87, 0x1f, 0x41, 0x57, 0xdd, 0xcd, 0xfd, 0x34,
	0x73, 0xee, 0xf1, 0xf1, 0x81, 0x37, 0xc6, 0x86, 0x65, 0x75, 0xc1, 0xb5, 0x5b, 0x09, 0x72, 0x85,
	0x7b, 0x68, 0x9d, 0xa0, 0x7c, 0x33, 0x13, 0xaa, 0xb4, 0x62, 0x55, 0x15, 0xc1, 0xea, 0xa2, 0xa2,
	0x80, 0x5e, 0x2c, 0x2c, 0x69, 0xb7, 0x41, 0x5f, 0x8b, 0xcd, 0xb1, 0x2a, 0xca, 0xa5, 0x3a, 0x16,
	0xea, 0x92, 0xe6, 0x3b, 0xcc, 0x4b, 0xef, 0x82, 0x63, 0x87, 0xbf, 0xc0, 0xef, 0xcf, 0x79, 0x23,
	0xce, 0xad, 0xcb, 0xc6, 0xc6, 0x19, 0xd7, 0x5e, 0x15, 0xcd, 0xa9, 0x7b, 0x95, 0x31, 0xbc, 0x0a,
	0x1d, 0xc4, 0xab, 0xd0, 0xb1, 0xa3, 0x4f, 0x03, 0x38, 0x78, 0x76, 0x49, 0xa7, 0xbd, 0xde, 0x79,
	0x89, 0x9a, 0x19, 0xb8, 0x83, 0x39, 0xcd, 0x09, 0x0b, 0xd4, 0xc1, 0xf9, 0x74, 0x7f, 0x12, 0x4f,
	0x47, 0xb3, 0x53, 0x7e, 0xfd, 0x56, 0xfe, 0xb7, 0x4e, 0x03, 0x24, 0x92, 0xab, 0xbc, 0xc6, 0xf3,
	0xad, 0x96, 0x1c, 0x61, 0x4e, 0xfd, 0x90, 0x7d, 0x8e, 0xe1, 0xfe, 0x3f, 0x2e, 0x31, 0x06, 0xb1,
	0xf2, 0xeb, 0x34, 0x9a, 0x44, 0xd3, 0xe4, 0x6c, 0x4f, 0x36, 0x03, 0x5b, 0xc2, 0x70, 0xa5, 0x82,
	0x5e, 0x62, 0xe3, 0x27, 0x9a, 0x8e, 0x66, 0xaf, 0x6e, 0xc2, 0x0f, 0x7f, 0xdd, 0x69, 0x9e, 0xed,
	0xc9, 0x5e, 0x3e, 0xfb, 0x1e, 0xc1, 0x70, 0x8b, 0xd9, 0x03, 0x18, 0x29, 0xad, 0x5d, 0xb5, 0x0e,
	0x73, 0xbb, 0xa0, 0x34, 0x9a, 0xc4, 0xd3, 0x44, 0xc2, 0x16, 0xbd, 0x5c, 0x10, 0x4b, 0x61, 0xe8,
	0xd1, 0x58, 0xb7, 0xa6, 0x36, 0xa6, 0x44, 0xf6, 0x23, 0xfb, 0x08, 0x83, 0xa0, 0x0c, 0xa5, 0x71,
	0x9b, 0xde, 0xfb, 0x9b, 0x74, 0xcb, 0xdf, 0x29, 0x43, 0xcf, 0xd7, 0xc1, 0xd7, 0xb2, 0xdd, 0x91,
	0x3d, 0x86, 0x64, 0x87, 0xd8, 0x01, 0xc4, 0x39, 0xd6, 0x5d, 0x7a, 0xb2, 0x39, 0xb2, 0x31, 0xdc,
	0xda, 0xa8, 0xa2, 0xc2, 0x36, 0xb9, 0x44, 0x76, 0xc3, 0xd3, 0xfd, 0x27, 0xd1, 0xc9, 0x3d, 0xb8,
	0xdb, 0xff, 0xe6, 0x79, 0xa8, 0x4b, 0x3c, 0x1a, 0x03, 0xfb, 0xc3, 0x49, 0x50, 0xa1, 0xa2, 0x93,
	0xb7, 0x5f, 0x7f, 0x0c, 0xa2, 0x2f, 0xdf, 0x0e, 0xa3, 0x0f, 0x2f, 0xfe, 0xd7, 0xe5, 0x32, 0x37,
	0x6d, 0x9f, 0xaf, 0xff, 0xcc, 0x5d, 0xaf, 0x2f, 0x6e, 0xb7, 0x05, 0x7c, 0xf4, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x44, 0x0b, 0xee, 0x1d, 0x03, 0x00, 0x00,
}

func (this *AwsDiscoverySpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoverySpec)
	if !ok {
		that2, ok := that.(AwsDiscoverySpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.EksSelector) != len(that1.EksSelector) {
		return false
	}
	for i := range this.EksSelector {
		if !this.EksSelector[i].Equal(that1.EksSelector[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsDiscoverySpec_AwsResourceSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoverySpec_AwsResourceSelector)
	if !ok {
		that2, ok := that.(AwsDiscoverySpec_AwsResourceSelector)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.SelectorType == nil {
		if this.SelectorType != nil {
			return false
		}
	} else if this.SelectorType == nil {
		return false
	} else if !this.SelectorType.Equal(that1.SelectorType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsDiscoverySpec_AwsResourceSelector_Arn) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoverySpec_AwsResourceSelector_Arn)
	if !ok {
		that2, ok := that.(AwsDiscoverySpec_AwsResourceSelector_Arn)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Arn != that1.Arn {
		return false
	}
	return true
}
func (this *AwsDiscoverySpec_AwsResourceSelector_Matcher_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoverySpec_AwsResourceSelector_Matcher_)
	if !ok {
		that2, ok := that.(AwsDiscoverySpec_AwsResourceSelector_Matcher_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	return true
}
func (this *AwsDiscoverySpec_AwsResourceSelector_Matcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoverySpec_AwsResourceSelector_Matcher)
	if !ok {
		that2, ok := that.(AwsDiscoverySpec_AwsResourceSelector_Matcher)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.AccountIds) != len(that1.AccountIds) {
		return false
	}
	for i := range this.AccountIds {
		if this.AccountIds[i] != that1.AccountIds[i] {
			return false
		}
	}
	if len(this.Regions) != len(that1.Regions) {
		return false
	}
	for i := range this.Regions {
		if this.Regions[i] != that1.Regions[i] {
			return false
		}
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsDiscoveryStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsDiscoveryStatus)
	if !ok {
		that2, ok := that.(AwsDiscoveryStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
