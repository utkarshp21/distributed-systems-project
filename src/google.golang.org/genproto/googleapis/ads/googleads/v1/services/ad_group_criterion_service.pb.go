// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for [AdGroupCriterionService.GetAdGroupCriterion][google.ads.googleads.v1.services.AdGroupCriterionService.GetAdGroupCriterion].
type GetAdGroupCriterionRequest struct {
	// Required. The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupCriterionRequest) Reset()         { *m = GetAdGroupCriterionRequest{} }
func (m *GetAdGroupCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupCriterionRequest) ProtoMessage()    {}
func (*GetAdGroupCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_214fd2be829a47dc, []int{0}
}

func (m *GetAdGroupCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Unmarshal(m, b)
}
func (m *GetAdGroupCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupCriterionRequest.Merge(m, src)
}
func (m *GetAdGroupCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupCriterionRequest.Size(m)
}
func (m *GetAdGroupCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupCriterionRequest proto.InternalMessageInfo

func (m *GetAdGroupCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupCriterionService.MutateAdGroupCriteria][google.ads.googleads.v1.services.AdGroupCriterionService.MutateAdGroupCriteria].
type MutateAdGroupCriteriaRequest struct {
	// Required. ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual criteria.
	Operations []*AdGroupCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupCriteriaRequest) Reset()         { *m = MutateAdGroupCriteriaRequest{} }
func (m *MutateAdGroupCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriteriaRequest) ProtoMessage()    {}
func (*MutateAdGroupCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_214fd2be829a47dc, []int{1}
}

func (m *MutateAdGroupCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriteriaRequest.Merge(m, src)
}
func (m *MutateAdGroupCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriteriaRequest.Size(m)
}
func (m *MutateAdGroupCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriteriaRequest proto.InternalMessageInfo

func (m *MutateAdGroupCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupCriteriaRequest) GetOperations() []*AdGroupCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on an ad group criterion.
type AdGroupCriterionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The list of policy violation keys that should not cause a
	// PolicyViolationError to be reported. Not all policy violations are
	// exemptable, please refer to the is_exemptible field in the returned
	// PolicyViolationError.
	//
	// Resources violating these polices will be saved, but will not be eligible
	// to serve. They may begin serving at a later time due to a change in
	// policies, re-review of the resource, or a change in advertiser
	// certificates.
	ExemptPolicyViolationKeys []*common.PolicyViolationKey `protobuf:"bytes,5,rep,name=exempt_policy_violation_keys,json=exemptPolicyViolationKeys,proto3" json:"exempt_policy_violation_keys,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupCriterionOperation_Create
	//	*AdGroupCriterionOperation_Update
	//	*AdGroupCriterionOperation_Remove
	Operation            isAdGroupCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AdGroupCriterionOperation) Reset()         { *m = AdGroupCriterionOperation{} }
func (m *AdGroupCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionOperation) ProtoMessage()    {}
func (*AdGroupCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_214fd2be829a47dc, []int{2}
}

func (m *AdGroupCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionOperation.Unmarshal(m, b)
}
func (m *AdGroupCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionOperation.Merge(m, src)
}
func (m *AdGroupCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionOperation.Size(m)
}
func (m *AdGroupCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionOperation proto.InternalMessageInfo

func (m *AdGroupCriterionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetExemptPolicyViolationKeys() []*common.PolicyViolationKey {
	if m != nil {
		return m.ExemptPolicyViolationKeys
	}
	return nil
}

type isAdGroupCriterionOperation_Operation interface {
	isAdGroupCriterionOperation_Operation()
}

type AdGroupCriterionOperation_Create struct {
	Create *resources.AdGroupCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupCriterionOperation_Update struct {
	Update *resources.AdGroupCriterion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupCriterionOperation_Create) isAdGroupCriterionOperation_Operation() {}

func (*AdGroupCriterionOperation_Update) isAdGroupCriterionOperation_Operation() {}

func (*AdGroupCriterionOperation_Remove) isAdGroupCriterionOperation_Operation() {}

func (m *AdGroupCriterionOperation) GetOperation() isAdGroupCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetCreate() *resources.AdGroupCriterion {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetUpdate() *resources.AdGroupCriterion {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionOperation_Create)(nil),
		(*AdGroupCriterionOperation_Update)(nil),
		(*AdGroupCriterionOperation_Remove)(nil),
	}
}

// Response message for an ad group criterion mutate.
type MutateAdGroupCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateAdGroupCriteriaResponse) Reset()         { *m = MutateAdGroupCriteriaResponse{} }
func (m *MutateAdGroupCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriteriaResponse) ProtoMessage()    {}
func (*MutateAdGroupCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_214fd2be829a47dc, []int{3}
}

func (m *MutateAdGroupCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriteriaResponse.Merge(m, src)
}
func (m *MutateAdGroupCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriteriaResponse.Size(m)
}
func (m *MutateAdGroupCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriteriaResponse proto.InternalMessageInfo

func (m *MutateAdGroupCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupCriteriaResponse) GetResults() []*MutateAdGroupCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateAdGroupCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupCriterionResult) Reset()         { *m = MutateAdGroupCriterionResult{} }
func (m *MutateAdGroupCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupCriterionResult) ProtoMessage()    {}
func (*MutateAdGroupCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_214fd2be829a47dc, []int{4}
}

func (m *MutateAdGroupCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Unmarshal(m, b)
}
func (m *MutateAdGroupCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupCriterionResult.Merge(m, src)
}
func (m *MutateAdGroupCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupCriterionResult.Size(m)
}
func (m *MutateAdGroupCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupCriterionResult proto.InternalMessageInfo

func (m *MutateAdGroupCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupCriterionRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupCriterionRequest")
	proto.RegisterType((*MutateAdGroupCriteriaRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriteriaRequest")
	proto.RegisterType((*AdGroupCriterionOperation)(nil), "google.ads.googleads.v1.services.AdGroupCriterionOperation")
	proto.RegisterType((*MutateAdGroupCriteriaResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriteriaResponse")
	proto.RegisterType((*MutateAdGroupCriterionResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_criterion_service.proto", fileDescriptor_214fd2be829a47dc)
}

var fileDescriptor_214fd2be829a47dc = []byte{
	// 851 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0xce, 0x52, 0xe8, 0x64, 0x17, 0xa4, 0x59, 0x2d, 0xeb, 0x86, 0x20, 0x22, 0x53, 0x89,
	0x92, 0x45, 0xb6, 0x92, 0x72, 0xc1, 0x65, 0x41, 0xce, 0x8a, 0x66, 0x57, 0xa8, 0xb4, 0x72, 0xa5,
	0xa8, 0x42, 0x91, 0xcc, 0xd4, 0x9e, 0x06, 0xab, 0xb6, 0xc7, 0xcc, 0x8c, 0x23, 0xa2, 0xaa, 0x07,
	0xe0, 0x1b, 0xc0, 0x37, 0xe0, 0xc8, 0xb7, 0xe0, 0xda, 0x6b, 0x6f, 0x3d, 0xf5, 0xc0, 0x89, 0x23,
	0xe2, 0xc4, 0x09, 0xd9, 0x33, 0x93, 0x7f, 0x8d, 0x15, 0xa9, 0xdc, 0x9e, 0xe7, 0xfd, 0xde, 0xef,
	0xfd, 0x7f, 0x06, 0xee, 0x88, 0x90, 0x51, 0x8c, 0x6d, 0x14, 0x32, 0x5b, 0x88, 0x85, 0x34, 0xee,
	0xd8, 0x0c, 0xd3, 0x71, 0x14, 0x60, 0x66, 0xa3, 0xd0, 0x1f, 0x51, 0x92, 0x67, 0x7e, 0x40, 0x23,
	0x8e, 0x69, 0x44, 0x52, 0x5f, 0xea, 0xac, 0x8c, 0x12, 0x4e, 0x60, 0x4b, 0xd8, 0x59, 0x28, 0x64,
	0xd6, 0x94, 0xc2, 0x1a, 0x77, 0x2c, 0x45, 0xd1, 0x78, 0x56, 0xe5, 0x24, 0x20, 0x49, 0x42, 0x52,
	0x3b, 0x23, 0x71, 0x14, 0x4c, 0x04, 0x5d, 0xc3, 0xa9, 0x02, 0x53, 0xcc, 0x48, 0x4e, 0x57, 0x87,
	0x24, 0x6d, 0x9b, 0xca, 0x36, 0x8b, 0x6c, 0x94, 0xa6, 0x84, 0x23, 0x1e, 0x91, 0x94, 0x49, 0xed,
	0xd3, 0x39, 0x6d, 0x10, 0x47, 0x38, 0xe5, 0x52, 0xf1, 0xfe, 0x9c, 0xe2, 0x2c, 0xc2, 0x71, 0xe8,
	0x9f, 0xe2, 0xef, 0xd0, 0x38, 0x22, 0x54, 0x02, 0xb6, 0xe6, 0x00, 0x2a, 0x0c, 0xa9, 0x92, 0xd9,
	0xdb, 0xe5, 0xd7, 0x69, 0x7e, 0x26, 0x09, 0x12, 0xc4, 0xce, 0x97, 0xdc, 0xd2, 0x2c, 0xb0, 0x19,
	0x47, 0x3c, 0x97, 0xf1, 0x98, 0x1c, 0x34, 0xfa, 0x98, 0xbb, 0x61, 0xbf, 0xc8, 0xe5, 0x85, 0x4a,
	0xc5, 0xc3, 0xdf, 0xe7, 0x98, 0x71, 0x38, 0x00, 0x8f, 0x94, 0x2b, 0x3f, 0x45, 0x09, 0x36, 0xb4,
	0x96, 0xb6, 0xb3, 0xd9, 0xeb, 0xdc, 0xba, 0xfa, 0xbf, 0xee, 0x33, 0xf0, 0xd1, 0xac, 0xd4, 0x52,
	0xca, 0x22, 0x66, 0x05, 0x24, 0xb1, 0xef, 0x10, 0x3e, 0x54, 0x3c, 0x5f, 0xa3, 0x04, 0x9b, 0xff,
	0x68, 0xa0, 0x79, 0x90, 0x73, 0xc4, 0xf1, 0x22, 0x10, 0x29, 0xc7, 0xdb, 0xa0, 0x1e, 0xe4, 0x8c,
	0x93, 0x04, 0x53, 0x3f, 0x0a, 0xa5, 0xdb, 0xda, 0xad, 0xab, 0x7b, 0x40, 0xbd, 0xbf, 0x0a, 0xe1,
	0xb7, 0x00, 0x90, 0x0c, 0x53, 0x51, 0x60, 0x43, 0x6f, 0xd5, 0x76, 0xea, 0xdd, 0x3d, 0x6b, 0xdd,
	0x28, 0x58, 0xcb, 0xc1, 0x1d, 0x2a, 0x0e, 0xe9, 0x61, 0xc6, 0x09, 0x3f, 0x04, 0x6f, 0x67, 0x88,
	0xf2, 0x08, 0xc5, 0xfe, 0x19, 0x8a, 0xe2, 0x9c, 0x62, 0xa3, 0xd6, 0xd2, 0x76, 0xde, 0xf4, 0xde,
	0x92, 0xcf, 0xfb, 0xe2, 0x15, 0x7e, 0x00, 0x1e, 0x8d, 0x51, 0x1c, 0x85, 0x88, 0x63, 0x9f, 0xa4,
	0xf1, 0xc4, 0x78, 0x50, 0xc2, 0x1e, 0xaa, 0xc7, 0xc3, 0x34, 0x9e, 0x98, 0xbf, 0xd4, 0xc0, 0x56,
	0xa5, 0x73, 0xb8, 0x07, 0xea, 0x79, 0x56, 0x12, 0x14, 0x8d, 0x2b, 0x09, 0xea, 0xdd, 0x86, 0x4a,
	0x47, 0xf5, 0xd6, 0xda, 0x2f, 0x7a, 0x7b, 0x80, 0xd8, 0xb9, 0x07, 0x04, 0xbc, 0x90, 0x21, 0x03,
	0x4d, 0xfc, 0x03, 0x4e, 0x32, 0xee, 0x8b, 0x41, 0xf6, 0xc7, 0x11, 0x89, 0x4b, 0x5e, 0xff, 0x1c,
	0x4f, 0x98, 0xf1, 0x7a, 0x59, 0x9c, 0x6e, 0x65, 0x71, 0xc4, 0x16, 0x58, 0x47, 0xa5, 0xf1, 0x40,
	0xd9, 0x7e, 0x85, 0x27, 0xde, 0x96, 0xe0, 0xbd, 0xab, 0x61, 0xf0, 0x00, 0x6c, 0x04, 0x14, 0x23,
	0x2e, 0xe6, 0xa2, 0xde, 0xdd, 0xad, 0xa4, 0x9f, 0xee, 0xcd, 0x9d, 0xe2, 0xbf, 0x7c, 0xcd, 0x93,
	0x24, 0x05, 0x9d, 0xc8, 0xc8, 0xd0, 0xff, 0x17, 0x9d, 0x20, 0x81, 0x06, 0xd8, 0xa0, 0x38, 0x21,
	0x63, 0xd1, 0xb2, 0xcd, 0x42, 0x23, 0xbe, 0x7b, 0x75, 0xb0, 0x39, 0xed, 0xb1, 0xf9, 0x87, 0x06,
	0xde, 0xab, 0x98, 0x45, 0x96, 0x91, 0x94, 0x61, 0xb8, 0x0f, 0x9e, 0x2c, 0x0d, 0x81, 0x8f, 0x29,
	0x25, 0xb4, 0xe4, 0xad, 0x77, 0xa1, 0x0a, 0x93, 0x66, 0x81, 0x75, 0x5c, 0x2e, 0x97, 0xf7, 0x78,
	0x71, 0x3c, 0xbe, 0x2c, 0xe0, 0xf0, 0x04, 0xbc, 0x41, 0x31, 0xcb, 0x63, 0xae, 0x66, 0xf5, 0xf3,
	0xf5, 0xb3, 0xba, 0x2a, 0xb2, 0x62, 0x3f, 0x0b, 0x1a, 0x4f, 0xd1, 0x99, 0x2f, 0x56, 0xaf, 0x93,
	0x02, 0x16, 0xd3, 0xb9, 0x62, 0x8f, 0x17, 0x97, 0xb2, 0xfb, 0xe3, 0x03, 0xf0, 0x74, 0xd9, 0xfe,
	0x58, 0xc4, 0x01, 0xaf, 0x35, 0xf0, 0x78, 0xc5, 0x9d, 0x80, 0x9f, 0xad, 0xcf, 0xa0, 0xfa, 0xbc,
	0x34, 0xee, 0xd3, 0x60, 0xf3, 0xd5, 0x8d, 0xbb, 0x98, 0xcc, 0x4f, 0xd7, 0x7f, 0xfe, 0xaa, 0xef,
	0xc2, 0x4e, 0x71, 0x9f, 0x2f, 0x16, 0x34, 0xcf, 0xd5, 0xad, 0x60, 0x76, 0xdb, 0x46, 0x8b, 0xdd,
	0xb5, 0xdb, 0x97, 0xf0, 0x6f, 0x0d, 0x3c, 0x59, 0xd9, 0x7a, 0x78, 0xcf, 0xce, 0xa8, 0xfb, 0xd5,
	0xf8, 0xe2, 0xde, 0xf6, 0x62, 0xe6, 0xcc, 0x93, 0x1b, 0xf7, 0x9d, 0xb9, 0x0b, 0xf8, 0xf1, 0xec,
	0x26, 0x95, 0xe9, 0x7e, 0x6a, 0x7e, 0x52, 0xfe, 0xbb, 0xa6, 0xf9, 0x5d, 0xcc, 0x81, 0x9f, 0xb7,
	0x2f, 0x97, 0xb3, 0x75, 0x92, 0xd2, 0x9b, 0xa3, 0xb5, 0x1b, 0xef, 0x5e, 0xb9, 0x46, 0xd5, 0xdd,
	0xee, 0xfd, 0xac, 0x83, 0xed, 0x80, 0x24, 0x6b, 0xa3, 0xef, 0x35, 0x2b, 0x26, 0xe5, 0xa8, 0x38,
	0x53, 0x47, 0xda, 0x37, 0x2f, 0x25, 0xc3, 0x88, 0xc4, 0x28, 0x1d, 0x59, 0x84, 0x8e, 0xec, 0x11,
	0x4e, 0xcb, 0x23, 0x66, 0xcf, 0x7c, 0x56, 0xff, 0xf2, 0xf7, 0x94, 0xf0, 0x9b, 0x5e, 0xeb, 0xbb,
	0xee, 0xef, 0x7a, 0xab, 0x2f, 0x08, 0xdd, 0x90, 0x59, 0x42, 0x2c, 0xa4, 0x41, 0xc7, 0x92, 0x8e,
	0xd9, 0x95, 0x82, 0x0c, 0xdd, 0x90, 0x0d, 0xa7, 0x90, 0xe1, 0xa0, 0x33, 0x54, 0x90, 0xbf, 0xf4,
	0x6d, 0xf1, 0xee, 0x38, 0x6e, 0xc8, 0x1c, 0x67, 0x0a, 0x72, 0x9c, 0x41, 0xc7, 0x71, 0x14, 0xec,
	0x74, 0xa3, 0x8c, 0x73, 0xf7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x6a, 0x03, 0x1b, 0x99,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupCriterionServiceClient is the client API for AdGroupCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetAdGroupCriterion(ctx context.Context, in *GetAdGroupCriterionRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error)
}

type adGroupCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionServiceClient {
	return &adGroupCriterionServiceClient{cc}
}

func (c *adGroupCriterionServiceClient) GetAdGroupCriterion(ctx context.Context, in *GetAdGroupCriterionRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterion, error) {
	out := new(resources.AdGroupCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupCriterionService/GetAdGroupCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionServiceClient) MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error) {
	out := new(MutateAdGroupCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupCriterionService/MutateAdGroupCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionServiceServer is the server API for AdGroupCriterionService service.
type AdGroupCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetAdGroupCriterion(context.Context, *GetAdGroupCriterionRequest) (*resources.AdGroupCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateAdGroupCriteria(context.Context, *MutateAdGroupCriteriaRequest) (*MutateAdGroupCriteriaResponse, error)
}

// UnimplementedAdGroupCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionServiceServer struct {
}

func (*UnimplementedAdGroupCriterionServiceServer) GetAdGroupCriterion(ctx context.Context, req *GetAdGroupCriterionRequest) (*resources.AdGroupCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupCriterion not implemented")
}
func (*UnimplementedAdGroupCriterionServiceServer) MutateAdGroupCriteria(ctx context.Context, req *MutateAdGroupCriteriaRequest) (*MutateAdGroupCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupCriteria not implemented")
}

func RegisterAdGroupCriterionServiceServer(s *grpc.Server, srv AdGroupCriterionServiceServer) {
	s.RegisterService(&_AdGroupCriterionService_serviceDesc, srv)
}

func _AdGroupCriterionService_GetAdGroupCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionServiceServer).GetAdGroupCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupCriterionService/GetAdGroupCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionServiceServer).GetAdGroupCriterion(ctx, req.(*GetAdGroupCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionService_MutateAdGroupCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupCriterionService/MutateAdGroupCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, req.(*MutateAdGroupCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupCriterionService",
	HandlerType: (*AdGroupCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterion",
			Handler:    _AdGroupCriterionService_GetAdGroupCriterion_Handler,
		},
		{
			MethodName: "MutateAdGroupCriteria",
			Handler:    _AdGroupCriterionService_MutateAdGroupCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_criterion_service.proto",
}
