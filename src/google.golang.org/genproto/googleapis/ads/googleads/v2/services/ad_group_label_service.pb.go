// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [AdGroupLabelService.GetAdGroupLabel][google.ads.googleads.v2.services.AdGroupLabelService.GetAdGroupLabel].
type GetAdGroupLabelRequest struct {
	// Required. The resource name of the ad group label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupLabelRequest) Reset()         { *m = GetAdGroupLabelRequest{} }
func (m *GetAdGroupLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupLabelRequest) ProtoMessage()    {}
func (*GetAdGroupLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8693fa91b2cc8640, []int{0}
}

func (m *GetAdGroupLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupLabelRequest.Unmarshal(m, b)
}
func (m *GetAdGroupLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupLabelRequest.Merge(m, src)
}
func (m *GetAdGroupLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupLabelRequest.Size(m)
}
func (m *GetAdGroupLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupLabelRequest proto.InternalMessageInfo

func (m *GetAdGroupLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupLabelService.MutateAdGroupLabels][google.ads.googleads.v2.services.AdGroupLabelService.MutateAdGroupLabels].
type MutateAdGroupLabelsRequest struct {
	// Required. ID of the customer whose ad group labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group labels.
	Operations []*AdGroupLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupLabelsRequest) Reset()         { *m = MutateAdGroupLabelsRequest{} }
func (m *MutateAdGroupLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelsRequest) ProtoMessage()    {}
func (*MutateAdGroupLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8693fa91b2cc8640, []int{1}
}

func (m *MutateAdGroupLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelsRequest.Merge(m, src)
}
func (m *MutateAdGroupLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelsRequest.Size(m)
}
func (m *MutateAdGroupLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelsRequest proto.InternalMessageInfo

func (m *MutateAdGroupLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupLabelsRequest) GetOperations() []*AdGroupLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group label.
type AdGroupLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupLabelOperation_Create
	//	*AdGroupLabelOperation_Remove
	Operation            isAdGroupLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AdGroupLabelOperation) Reset()         { *m = AdGroupLabelOperation{} }
func (m *AdGroupLabelOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupLabelOperation) ProtoMessage()    {}
func (*AdGroupLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8693fa91b2cc8640, []int{2}
}

func (m *AdGroupLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupLabelOperation.Unmarshal(m, b)
}
func (m *AdGroupLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupLabelOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupLabelOperation.Merge(m, src)
}
func (m *AdGroupLabelOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupLabelOperation.Size(m)
}
func (m *AdGroupLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupLabelOperation proto.InternalMessageInfo

type isAdGroupLabelOperation_Operation interface {
	isAdGroupLabelOperation_Operation()
}

type AdGroupLabelOperation_Create struct {
	Create *resources.AdGroupLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupLabelOperation_Create) isAdGroupLabelOperation_Operation() {}

func (*AdGroupLabelOperation_Remove) isAdGroupLabelOperation_Operation() {}

func (m *AdGroupLabelOperation) GetOperation() isAdGroupLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupLabelOperation) GetCreate() *resources.AdGroupLabel {
	if x, ok := m.GetOperation().(*AdGroupLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupLabelOperation_Create)(nil),
		(*AdGroupLabelOperation_Remove)(nil),
	}
}

// Response message for an ad group labels mutate.
type MutateAdGroupLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateAdGroupLabelsResponse) Reset()         { *m = MutateAdGroupLabelsResponse{} }
func (m *MutateAdGroupLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelsResponse) ProtoMessage()    {}
func (*MutateAdGroupLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8693fa91b2cc8640, []int{3}
}

func (m *MutateAdGroupLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelsResponse.Merge(m, src)
}
func (m *MutateAdGroupLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelsResponse.Size(m)
}
func (m *MutateAdGroupLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelsResponse proto.InternalMessageInfo

func (m *MutateAdGroupLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupLabelsResponse) GetResults() []*MutateAdGroupLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for an ad group label mutate.
type MutateAdGroupLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupLabelResult) Reset()         { *m = MutateAdGroupLabelResult{} }
func (m *MutateAdGroupLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupLabelResult) ProtoMessage()    {}
func (*MutateAdGroupLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8693fa91b2cc8640, []int{4}
}

func (m *MutateAdGroupLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupLabelResult.Unmarshal(m, b)
}
func (m *MutateAdGroupLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupLabelResult.Merge(m, src)
}
func (m *MutateAdGroupLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupLabelResult.Size(m)
}
func (m *MutateAdGroupLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupLabelResult proto.InternalMessageInfo

func (m *MutateAdGroupLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupLabelRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupLabelRequest")
	proto.RegisterType((*MutateAdGroupLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupLabelsRequest")
	proto.RegisterType((*AdGroupLabelOperation)(nil), "google.ads.googleads.v2.services.AdGroupLabelOperation")
	proto.RegisterType((*MutateAdGroupLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupLabelsResponse")
	proto.RegisterType((*MutateAdGroupLabelResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_label_service.proto", fileDescriptor_8693fa91b2cc8640)
}

var fileDescriptor_8693fa91b2cc8640 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6b, 0xd4, 0x4e,
	0x14, 0xff, 0x26, 0xfb, 0xa5, 0xda, 0xd9, 0xd6, 0xc2, 0x94, 0xd6, 0xb8, 0x15, 0x5c, 0x62, 0xa5,
	0x65, 0xd1, 0x04, 0x23, 0x58, 0x89, 0x2e, 0x92, 0x05, 0xbb, 0x2d, 0xa8, 0x2d, 0xa9, 0xec, 0x41,
	0x16, 0xc2, 0x34, 0x99, 0xae, 0x81, 0x6c, 0x26, 0xce, 0x4c, 0x16, 0x4a, 0xe9, 0xc5, 0x9b, 0x67,
	0xff, 0x03, 0x8f, 0xfe, 0x0b, 0x9e, 0xbd, 0xf4, 0xea, 0xad, 0xa7, 0x1e, 0x04, 0xa1, 0x07, 0xff,
	0x00, 0x2f, 0x4a, 0x7e, 0xcc, 0x6e, 0xd2, 0xee, 0xb2, 0xd8, 0xdb, 0xcb, 0x7b, 0x9f, 0xf7, 0x3e,
	0xef, 0xe7, 0x04, 0x34, 0x7b, 0x84, 0xf4, 0x02, 0xac, 0x23, 0x8f, 0xe9, 0x99, 0x98, 0x48, 0x03,
	0x43, 0x67, 0x98, 0x0e, 0x7c, 0x17, 0x33, 0x1d, 0x79, 0x4e, 0x8f, 0x92, 0x38, 0x72, 0x02, 0xb4,
	0x8f, 0x03, 0x27, 0xd7, 0x6b, 0x11, 0x25, 0x9c, 0xc0, 0x7a, 0xe6, 0xa3, 0x21, 0x8f, 0x69, 0x43,
	0x77, 0x6d, 0x60, 0x68, 0xc2, 0xbd, 0xf6, 0x78, 0x12, 0x01, 0xc5, 0x8c, 0xc4, 0xf4, 0x32, 0x43,
	0x16, 0xb9, 0x76, 0x5b, 0xf8, 0x45, 0xbe, 0x8e, 0xc2, 0x90, 0x70, 0xc4, 0x7d, 0x12, 0xb2, 0xdc,
	0x7a, 0xb3, 0x60, 0x75, 0x03, 0x1f, 0x87, 0x3c, 0x37, 0xdc, 0x29, 0x18, 0x0e, 0x7c, 0x1c, 0x78,
	0xce, 0x3e, 0x7e, 0x87, 0x06, 0x3e, 0xa1, 0x39, 0xe0, 0x56, 0x01, 0x20, 0x52, 0xb8, 0x10, 0x94,
	0x46, 0xae, 0xce, 0x38, 0xe2, 0x71, 0xce, 0xa6, 0x06, 0x60, 0xb9, 0x8d, 0xb9, 0xe5, 0xb5, 0x93,
	0x2c, 0x5f, 0x26, 0x49, 0xda, 0xf8, 0x7d, 0x8c, 0x19, 0x87, 0x36, 0x98, 0x17, 0x41, 0x9c, 0x10,
	0xf5, 0xb1, 0x22, 0xd5, 0xa5, 0xf5, 0xd9, 0xd6, 0x83, 0x33, 0x4b, 0xfe, 0x6d, 0xad, 0x81, 0x7b,
	0xa3, 0x9e, 0xe4, 0x52, 0xe4, 0x33, 0xcd, 0x25, 0x7d, 0xbd, 0x14, 0x6c, 0x4e, 0xc4, 0x78, 0x8d,
	0xfa, 0x58, 0xfd, 0x25, 0x81, 0xda, 0xab, 0x98, 0x23, 0x8e, 0x8b, 0x20, 0x26, 0x28, 0x57, 0x41,
	0xd5, 0x8d, 0x19, 0x27, 0x7d, 0x4c, 0x1d, 0xdf, 0xcb, 0x09, 0x2b, 0x67, 0x96, 0x6c, 0x03, 0xa1,
	0xdf, 0xf6, 0x60, 0x17, 0x00, 0x12, 0x61, 0x9a, 0x35, 0x4d, 0x91, 0xeb, 0x95, 0xf5, 0xaa, 0xb1,
	0xa1, 0x4d, 0x9b, 0x96, 0x56, 0x64, 0xdc, 0x11, 0xfe, 0x79, 0xf4, 0x51, 0x3c, 0xb8, 0x06, 0x16,
	0x22, 0x44, 0xb9, 0x8f, 0x02, 0xe7, 0x00, 0xf9, 0x41, 0x4c, 0xb1, 0x52, 0xa9, 0x4b, 0xeb, 0xd7,
	0xed, 0x1b, 0xb9, 0x7a, 0x33, 0xd3, 0xc2, 0xbb, 0x60, 0x7e, 0x80, 0x02, 0xdf, 0x43, 0x1c, 0x3b,
	0x24, 0x0c, 0x0e, 0x95, 0xff, 0x53, 0xd8, 0x9c, 0x50, 0xee, 0x84, 0xc1, 0xa1, 0xfa, 0x51, 0x02,
	0x4b, 0x63, 0x89, 0xe1, 0x36, 0x98, 0x71, 0x29, 0x46, 0x3c, 0xeb, 0x6b, 0xd5, 0xd0, 0x27, 0x56,
	0x30, 0xdc, 0xa6, 0x52, 0x09, 0x5b, 0xff, 0xd9, 0x79, 0x00, 0xa8, 0x80, 0x19, 0x8a, 0xfb, 0x64,
	0x80, 0x15, 0x39, 0xe9, 0x58, 0x62, 0xc9, 0xbe, 0x5b, 0x55, 0x30, 0x3b, 0x2c, 0x4d, 0xfd, 0x2a,
	0x81, 0x95, 0xb1, 0xcd, 0x67, 0x11, 0x09, 0x19, 0x86, 0x9b, 0x60, 0xe9, 0x42, 0xe5, 0x0e, 0xa6,
	0x94, 0xd0, 0xb4, 0xfe, 0xaa, 0x01, 0x45, 0x82, 0x34, 0x72, 0xb5, 0xbd, 0x74, 0x87, 0xec, 0xc5,
	0x72, 0x4f, 0x5e, 0x24, 0x70, 0xf8, 0x06, 0x5c, 0xa3, 0x98, 0xc5, 0x01, 0x17, 0xc3, 0x31, 0xa7,
	0x0f, 0xe7, 0x72, 0x5e, 0x76, 0x1a, 0xc2, 0x16, 0xa1, 0xd4, 0xe7, 0x40, 0x99, 0x04, 0x4a, 0x46,
	0x31, 0x66, 0x55, 0xcb, 0xbb, 0x67, 0xfc, 0xac, 0x80, 0xc5, 0xa2, 0xef, 0x5e, 0xc6, 0x0d, 0xbf,
	0x49, 0x60, 0xe1, 0xc2, 0x09, 0xc0, 0x27, 0xd3, 0x33, 0x1e, 0x7f, 0x35, 0xb5, 0x7f, 0x1d, 0xa3,
	0xda, 0x3e, 0xb5, 0xca, 0xc9, 0x7f, 0xf8, 0xfe, 0xe3, 0x93, 0xfc, 0x10, 0xea, 0xc9, 0x43, 0x72,
	0x54, 0xb2, 0x34, 0xc5, 0x11, 0x30, 0xbd, 0xa1, 0xa3, 0xe2, 0x0c, 0xf5, 0xc6, 0x31, 0x3c, 0x97,
	0xc0, 0xe2, 0x98, 0xf1, 0xc2, 0x67, 0x57, 0xe9, 0xbe, 0x38, 0xc9, 0x5a, 0xf3, 0x8a, 0xde, 0xd9,
	0x4e, 0xa9, 0x9d, 0x53, 0x6b, 0xb9, 0x70, 0xd2, 0xf7, 0x47, 0x87, 0x96, 0x96, 0xb9, 0xa1, 0x1a,
	0x49, 0x99, 0xa3, 0xba, 0x8e, 0x0a, 0xe0, 0x66, 0xe3, 0xb8, 0x5c, 0xa5, 0xd9, 0x4f, 0x99, 0x4c,
	0xa9, 0x51, 0x5b, 0x39, 0xb1, 0x94, 0x49, 0x0f, 0x50, 0xeb, 0x8f, 0x04, 0x56, 0x5d, 0xd2, 0x9f,
	0x9a, 0x79, 0x4b, 0x19, 0xb3, 0x0f, 0xbb, 0xc9, 0xb3, 0xb8, 0x2b, 0xbd, 0xdd, 0xca, 0xbd, 0x7b,
	0x24, 0x40, 0x61, 0x4f, 0x23, 0xb4, 0xa7, 0xf7, 0x70, 0x98, 0x3e, 0x9a, 0xfa, 0x88, 0x6f, 0xf2,
	0xcf, 0xe5, 0xa9, 0x10, 0x3e, 0xcb, 0x95, 0xb6, 0x65, 0x7d, 0x91, 0xeb, 0xed, 0x2c, 0xa0, 0xe5,
	0x31, 0x2d, 0x13, 0x13, 0xa9, 0x63, 0x68, 0x39, 0x31, 0x3b, 0x11, 0x90, 0xae, 0xe5, 0xb1, 0xee,
	0x10, 0xd2, 0xed, 0x18, 0x5d, 0x01, 0x39, 0x97, 0x57, 0x33, 0xbd, 0x69, 0x5a, 0x1e, 0x33, 0xcd,
	0x21, 0xc8, 0x34, 0x3b, 0x86, 0x69, 0x0a, 0xd8, 0xfe, 0x4c, 0x9a, 0xe7, 0xa3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xd1, 0x0d, 0xc8, 0x03, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupLabelServiceClient is the client API for AdGroupLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupLabelServiceClient interface {
	// Returns the requested ad group label in full detail.
	GetAdGroupLabel(ctx context.Context, in *GetAdGroupLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupLabel, error)
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error)
}

type adGroupLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupLabelServiceClient(cc grpc.ClientConnInterface) AdGroupLabelServiceClient {
	return &adGroupLabelServiceClient{cc}
}

func (c *adGroupLabelServiceClient) GetAdGroupLabel(ctx context.Context, in *GetAdGroupLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupLabel, error) {
	out := new(resources.AdGroupLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupLabelService/GetAdGroupLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupLabelServiceClient) MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error) {
	out := new(MutateAdGroupLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupLabelService/MutateAdGroupLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupLabelServiceServer is the server API for AdGroupLabelService service.
type AdGroupLabelServiceServer interface {
	// Returns the requested ad group label in full detail.
	GetAdGroupLabel(context.Context, *GetAdGroupLabelRequest) (*resources.AdGroupLabel, error)
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	MutateAdGroupLabels(context.Context, *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error)
}

// UnimplementedAdGroupLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupLabelServiceServer struct {
}

func (*UnimplementedAdGroupLabelServiceServer) GetAdGroupLabel(ctx context.Context, req *GetAdGroupLabelRequest) (*resources.AdGroupLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupLabel not implemented")
}
func (*UnimplementedAdGroupLabelServiceServer) MutateAdGroupLabels(ctx context.Context, req *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupLabels not implemented")
}

func RegisterAdGroupLabelServiceServer(s *grpc.Server, srv AdGroupLabelServiceServer) {
	s.RegisterService(&_AdGroupLabelService_serviceDesc, srv)
}

func _AdGroupLabelService_GetAdGroupLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupLabelServiceServer).GetAdGroupLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupLabelService/GetAdGroupLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupLabelServiceServer).GetAdGroupLabel(ctx, req.(*GetAdGroupLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupLabelService_MutateAdGroupLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupLabelService/MutateAdGroupLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, req.(*MutateAdGroupLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupLabelService",
	HandlerType: (*AdGroupLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupLabel",
			Handler:    _AdGroupLabelService_GetAdGroupLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupLabels",
			Handler:    _AdGroupLabelService_MutateAdGroupLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_label_service.proto",
}
