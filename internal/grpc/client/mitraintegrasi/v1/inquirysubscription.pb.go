// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.0
// source: inquirysubscription.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InqSubsParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefNo string `protobuf:"bytes,1,opt,name=ref_no,json=refNo,proto3" json:"ref_no,omitempty"`
}

func (x *InqSubsParam) Reset() {
	*x = InqSubsParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inquirysubscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InqSubsParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InqSubsParam) ProtoMessage() {}

func (x *InqSubsParam) ProtoReflect() protoreflect.Message {
	mi := &file_inquirysubscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InqSubsParam.ProtoReflect.Descriptor instead.
func (*InqSubsParam) Descriptor() ([]byte, []int) {
	return file_inquirysubscription_proto_rawDescGZIP(), []int{0}
}

func (x *InqSubsParam) GetRefNo() string {
	if x != nil {
		return x.RefNo
	}
	return ""
}

type InqResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status              string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ResponseCode        string `protobuf:"bytes,2,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	ResponseDescription string `protobuf:"bytes,3,opt,name=response_description,json=responseDescription,proto3" json:"response_description,omitempty"`
	Data                *Data  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InqResponse) Reset() {
	*x = InqResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inquirysubscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InqResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InqResponse) ProtoMessage() {}

func (x *InqResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inquirysubscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InqResponse.ProtoReflect.Descriptor instead.
func (*InqResponse) Descriptor() ([]byte, []int) {
	return file_inquirysubscription_proto_rawDescGZIP(), []int{1}
}

func (x *InqResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InqResponse) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *InqResponse) GetResponseDescription() string {
	if x != nil {
		return x.ResponseDescription
	}
	return ""
}

func (x *InqResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefNo                  string `protobuf:"bytes,1,opt,name=ref_no,json=refNo,proto3" json:"ref_no,omitempty"`
	Paymentpoolid          string `protobuf:"bytes,2,opt,name=paymentpoolid,proto3" json:"paymentpoolid,omitempty"`
	PaymentBankBicCode     string `protobuf:"bytes,3,opt,name=payment_bank_bic_code,json=paymentBankBicCode,proto3" json:"payment_bank_bic_code,omitempty"`
	NavDate                string `protobuf:"bytes,4,opt,name=nav_date,json=navDate,proto3" json:"nav_date,omitempty"`
	InvestorFundUnitAcNo   string `protobuf:"bytes,5,opt,name=investor_fund_unit_ac_no,json=investorFundUnitAcNo,proto3" json:"investor_fund_unit_ac_no,omitempty"`
	InvestorFundUnitAcName string `protobuf:"bytes,6,opt,name=investor_fund_unit_ac_name,json=investorFundUnitAcName,proto3" json:"investor_fund_unit_ac_name,omitempty"`
	Sid                    string `protobuf:"bytes,7,opt,name=sid,proto3" json:"sid,omitempty"`
	Status                 string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	FundCode               string `protobuf:"bytes,9,opt,name=fund_code,json=fundCode,proto3" json:"fund_code,omitempty"`
	AmountNominal          string `protobuf:"bytes,10,opt,name=amount_nominal,json=amountNominal,proto3" json:"amount_nominal,omitempty"`
	Unit                   string `protobuf:"bytes,11,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                  string `protobuf:"bytes,12,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inquirysubscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_inquirysubscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_inquirysubscription_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetRefNo() string {
	if x != nil {
		return x.RefNo
	}
	return ""
}

func (x *Data) GetPaymentpoolid() string {
	if x != nil {
		return x.Paymentpoolid
	}
	return ""
}

func (x *Data) GetPaymentBankBicCode() string {
	if x != nil {
		return x.PaymentBankBicCode
	}
	return ""
}

func (x *Data) GetNavDate() string {
	if x != nil {
		return x.NavDate
	}
	return ""
}

func (x *Data) GetInvestorFundUnitAcNo() string {
	if x != nil {
		return x.InvestorFundUnitAcNo
	}
	return ""
}

func (x *Data) GetInvestorFundUnitAcName() string {
	if x != nil {
		return x.InvestorFundUnitAcName
	}
	return ""
}

func (x *Data) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Data) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Data) GetFundCode() string {
	if x != nil {
		return x.FundCode
	}
	return ""
}

func (x *Data) GetAmountNominal() string {
	if x != nil {
		return x.AmountNominal
	}
	return ""
}

func (x *Data) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Data) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

var File_inquirysubscription_proto protoreflect.FileDescriptor

var file_inquirysubscription_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0c, 0x49,
	0x6e, 0x71, 0x53, 0x75, 0x62, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x72,
	0x65, 0x66, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66,
	0x4e, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x71, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x31, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9d, 0x03,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x66, 0x4e, 0x6f, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6f, 0x6c, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6f,
	0x6c, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x42,
	0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x61, 0x76, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x61, 0x76, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x36, 0x0a, 0x18, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x75,
	0x6e, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x61, 0x63, 0x5f, 0x6e, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x46, 0x75, 0x6e,
	0x64, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x63, 0x4e, 0x6f, 0x12, 0x3a, 0x0a, 0x1a, 0x69, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x61, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x46, 0x75, 0x6e, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x41,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x47, 0x0a,
	0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x71, 0x53, 0x75, 0x62, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x34, 0x0a, 0x13, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x49, 0x6e, 0x71, 0x53, 0x75, 0x62,
	0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0c, 0x2e, 0x49, 0x6e, 0x71, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inquirysubscription_proto_rawDescOnce sync.Once
	file_inquirysubscription_proto_rawDescData = file_inquirysubscription_proto_rawDesc
)

func file_inquirysubscription_proto_rawDescGZIP() []byte {
	file_inquirysubscription_proto_rawDescOnce.Do(func() {
		file_inquirysubscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_inquirysubscription_proto_rawDescData)
	})
	return file_inquirysubscription_proto_rawDescData
}

var file_inquirysubscription_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inquirysubscription_proto_goTypes = []interface{}{
	(*InqSubsParam)(nil), // 0: InqSubsParam
	(*InqResponse)(nil),  // 1: InqResponse
	(*Data)(nil),         // 2: Data
}
var file_inquirysubscription_proto_depIdxs = []int32{
	2, // 0: InqResponse.data:type_name -> Data
	0, // 1: SendInqSubsData.InquirySubscription:input_type -> InqSubsParam
	1, // 2: SendInqSubsData.InquirySubscription:output_type -> InqResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inquirysubscription_proto_init() }
func file_inquirysubscription_proto_init() {
	if File_inquirysubscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inquirysubscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InqSubsParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inquirysubscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InqResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inquirysubscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inquirysubscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inquirysubscription_proto_goTypes,
		DependencyIndexes: file_inquirysubscription_proto_depIdxs,
		MessageInfos:      file_inquirysubscription_proto_msgTypes,
	}.Build()
	File_inquirysubscription_proto = out.File
	file_inquirysubscription_proto_rawDesc = nil
	file_inquirysubscription_proto_goTypes = nil
	file_inquirysubscription_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendInqSubsDataClient is the client API for SendInqSubsData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendInqSubsDataClient interface {
	InquirySubscription(ctx context.Context, in *InqSubsParam, opts ...grpc.CallOption) (*InqResponse, error)
}

type sendInqSubsDataClient struct {
	cc grpc.ClientConnInterface
}

func NewSendInqSubsDataClient(cc grpc.ClientConnInterface) SendInqSubsDataClient {
	return &sendInqSubsDataClient{cc}
}

func (c *sendInqSubsDataClient) InquirySubscription(ctx context.Context, in *InqSubsParam, opts ...grpc.CallOption) (*InqResponse, error) {
	out := new(InqResponse)
	err := c.cc.Invoke(ctx, "/SendInqSubsData/InquirySubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendInqSubsDataServer is the server API for SendInqSubsData service.
type SendInqSubsDataServer interface {
	InquirySubscription(context.Context, *InqSubsParam) (*InqResponse, error)
}

// UnimplementedSendInqSubsDataServer can be embedded to have forward compatible implementations.
type UnimplementedSendInqSubsDataServer struct {
}

func (*UnimplementedSendInqSubsDataServer) InquirySubscription(context.Context, *InqSubsParam) (*InqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquirySubscription not implemented")
}

func RegisterSendInqSubsDataServer(s *grpc.Server, srv SendInqSubsDataServer) {
	s.RegisterService(&_SendInqSubsData_serviceDesc, srv)
}

func _SendInqSubsData_InquirySubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InqSubsParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendInqSubsDataServer).InquirySubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SendInqSubsData/InquirySubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendInqSubsDataServer).InquirySubscription(ctx, req.(*InqSubsParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _SendInqSubsData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SendInqSubsData",
	HandlerType: (*SendInqSubsDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InquirySubscription",
			Handler:    _SendInqSubsData_InquirySubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inquirysubscription.proto",
}