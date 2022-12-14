// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: delivery.proto

package proto

import (
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

type InDeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerAddress     string                        `protobuf:"bytes,1,opt,name=customer_address,json=customerAddress,proto3" json:"customer_address,omitempty"`
	CustomerCoordinates *InDeliveryCoordinatesRequest `protobuf:"bytes,2,opt,name=customer_coordinates,json=customerCoordinates,proto3" json:"customer_coordinates,omitempty"`
}

func (x *InDeliveryRequest) Reset() {
	*x = InDeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InDeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InDeliveryRequest) ProtoMessage() {}

func (x *InDeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InDeliveryRequest.ProtoReflect.Descriptor instead.
func (*InDeliveryRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *InDeliveryRequest) GetCustomerAddress() string {
	if x != nil {
		return x.CustomerAddress
	}
	return ""
}

func (x *InDeliveryRequest) GetCustomerCoordinates() *InDeliveryCoordinatesRequest {
	if x != nil {
		return x.CustomerCoordinates
	}
	return nil
}

type InDeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeliveryFee float64 `protobuf:"fixed64,1,opt,name=delivery_fee,json=deliveryFee,proto3" json:"delivery_fee,omitempty"`
}

func (x *InDeliveryResponse) Reset() {
	*x = InDeliveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InDeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InDeliveryResponse) ProtoMessage() {}

func (x *InDeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InDeliveryResponse.ProtoReflect.Descriptor instead.
func (*InDeliveryResponse) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *InDeliveryResponse) GetDeliveryFee() float64 {
	if x != nil {
		return x.DeliveryFee
	}
	return 0
}

type InDeliveryCoordinatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *InDeliveryCoordinatesRequest) Reset() {
	*x = InDeliveryCoordinatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InDeliveryCoordinatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InDeliveryCoordinatesRequest) ProtoMessage() {}

func (x *InDeliveryCoordinatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InDeliveryCoordinatesRequest.ProtoReflect.Descriptor instead.
func (*InDeliveryCoordinatesRequest) Descriptor() ([]byte, []int) {
	return file_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *InDeliveryCoordinatesRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *InDeliveryCoordinatesRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_delivery_proto protoreflect.FileDescriptor

var file_delivery_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x64, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x12,
	0x49, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x46, 0x65, 0x65, 0x22, 0x58, 0x0a, 0x1c, 0x49, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x32,
	0x77, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_delivery_proto_rawDescOnce sync.Once
	file_delivery_proto_rawDescData = file_delivery_proto_rawDesc
)

func file_delivery_proto_rawDescGZIP() []byte {
	file_delivery_proto_rawDescOnce.Do(func() {
		file_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_proto_rawDescData)
	})
	return file_delivery_proto_rawDescData
}

var file_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_delivery_proto_goTypes = []interface{}{
	(*InDeliveryRequest)(nil),            // 0: delivery_service.v1.InDeliveryRequest
	(*InDeliveryResponse)(nil),           // 1: delivery_service.v1.InDeliveryResponse
	(*InDeliveryCoordinatesRequest)(nil), // 2: delivery_service.v1.InDeliveryCoordinatesRequest
}
var file_delivery_proto_depIdxs = []int32{
	2, // 0: delivery_service.v1.InDeliveryRequest.customer_coordinates:type_name -> delivery_service.v1.InDeliveryCoordinatesRequest
	0, // 1: delivery_service.v1.deliveryService.DeliveryService:input_type -> delivery_service.v1.InDeliveryRequest
	1, // 2: delivery_service.v1.deliveryService.DeliveryService:output_type -> delivery_service.v1.InDeliveryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_delivery_proto_init() }
func file_delivery_proto_init() {
	if File_delivery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InDeliveryRequest); i {
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
		file_delivery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InDeliveryResponse); i {
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
		file_delivery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InDeliveryCoordinatesRequest); i {
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
			RawDescriptor: file_delivery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_proto_goTypes,
		DependencyIndexes: file_delivery_proto_depIdxs,
		MessageInfos:      file_delivery_proto_msgTypes,
	}.Build()
	File_delivery_proto = out.File
	file_delivery_proto_rawDesc = nil
	file_delivery_proto_goTypes = nil
	file_delivery_proto_depIdxs = nil
}
