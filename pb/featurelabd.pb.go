// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: featurelabd.proto

package pb

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

type GetTreatmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App      string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Feature  string `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"`
	Criteria string `protobuf:"bytes,3,opt,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *GetTreatmentRequest) Reset() {
	*x = GetTreatmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreatmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreatmentRequest) ProtoMessage() {}

func (x *GetTreatmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreatmentRequest.ProtoReflect.Descriptor instead.
func (*GetTreatmentRequest) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{0}
}

func (x *GetTreatmentRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *GetTreatmentRequest) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

func (x *GetTreatmentRequest) GetCriteria() string {
	if x != nil {
		return x.Criteria
	}
	return ""
}

type GetTreatmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Treatment string `protobuf:"bytes,1,opt,name=treatment,proto3" json:"treatment,omitempty"`
}

func (x *GetTreatmentResponse) Reset() {
	*x = GetTreatmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTreatmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTreatmentResponse) ProtoMessage() {}

func (x *GetTreatmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTreatmentResponse.ProtoReflect.Descriptor instead.
func (*GetTreatmentResponse) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{1}
}

func (x *GetTreatmentResponse) GetTreatment() string {
	if x != nil {
		return x.Treatment
	}
	return ""
}

type FetchFeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App     string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Feature string `protobuf:"bytes,2,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *FetchFeatureRequest) Reset() {
	*x = FetchFeatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureRequest) ProtoMessage() {}

func (x *FetchFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureRequest.ProtoReflect.Descriptor instead.
func (*FetchFeatureRequest) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{2}
}

func (x *FetchFeatureRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *FetchFeatureRequest) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

type FetchFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *Feature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *FetchFeatureResponse) Reset() {
	*x = FetchFeatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeatureResponse) ProtoMessage() {}

func (x *FetchFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeatureResponse.ProtoReflect.Descriptor instead.
func (*FetchFeatureResponse) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{3}
}

func (x *FetchFeatureResponse) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type FetchFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *FetchFeaturesRequest) Reset() {
	*x = FetchFeaturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeaturesRequest) ProtoMessage() {}

func (x *FetchFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeaturesRequest.ProtoReflect.Descriptor instead.
func (*FetchFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{4}
}

func (x *FetchFeaturesRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

type FetchFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []*Feature `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *FetchFeaturesResponse) Reset() {
	*x = FetchFeaturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFeaturesResponse) ProtoMessage() {}

func (x *FetchFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFeaturesResponse.ProtoReflect.Descriptor instead.
func (*FetchFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{5}
}

func (x *FetchFeaturesResponse) GetFeatures() []*Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App         string                 `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Allocations []*TreatmentAllocation `protobuf:"bytes,3,rep,name=allocations,proto3" json:"allocations,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{6}
}

func (x *Feature) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetAllocations() []*TreatmentAllocation {
	if x != nil {
		return x.Allocations
	}
	return nil
}

type TreatmentAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Treatment string `protobuf:"bytes,1,opt,name=treatment,proto3" json:"treatment,omitempty"`
	Weight    uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *TreatmentAllocation) Reset() {
	*x = TreatmentAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_featurelabd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreatmentAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreatmentAllocation) ProtoMessage() {}

func (x *TreatmentAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_featurelabd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreatmentAllocation.ProtoReflect.Descriptor instead.
func (*TreatmentAllocation) Descriptor() ([]byte, []int) {
	return file_featurelabd_proto_rawDescGZIP(), []int{7}
}

func (x *TreatmentAllocation) GetTreatment() string {
	if x != nil {
		return x.Treatment
	}
	return ""
}

func (x *TreatmentAllocation) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_featurelabd_proto protoreflect.FileDescriptor

var file_featurelabd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x61, 0x62, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65,
	0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x13,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x3d, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x28,
	0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x40, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x07, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0xde, 0x01, 0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c,
	0x61, 0x62, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x73, 0x6a, 0x65, 0x66, 0x66, 0x2f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x61, 0x62, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_featurelabd_proto_rawDescOnce sync.Once
	file_featurelabd_proto_rawDescData = file_featurelabd_proto_rawDesc
)

func file_featurelabd_proto_rawDescGZIP() []byte {
	file_featurelabd_proto_rawDescOnce.Do(func() {
		file_featurelabd_proto_rawDescData = protoimpl.X.CompressGZIP(file_featurelabd_proto_rawDescData)
	})
	return file_featurelabd_proto_rawDescData
}

var file_featurelabd_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_featurelabd_proto_goTypes = []interface{}{
	(*GetTreatmentRequest)(nil),   // 0: pb.GetTreatmentRequest
	(*GetTreatmentResponse)(nil),  // 1: pb.GetTreatmentResponse
	(*FetchFeatureRequest)(nil),   // 2: pb.FetchFeatureRequest
	(*FetchFeatureResponse)(nil),  // 3: pb.FetchFeatureResponse
	(*FetchFeaturesRequest)(nil),  // 4: pb.FetchFeaturesRequest
	(*FetchFeaturesResponse)(nil), // 5: pb.FetchFeaturesResponse
	(*Feature)(nil),               // 6: pb.Feature
	(*TreatmentAllocation)(nil),   // 7: pb.TreatmentAllocation
}
var file_featurelabd_proto_depIdxs = []int32{
	6, // 0: pb.FetchFeatureResponse.feature:type_name -> pb.Feature
	6, // 1: pb.FetchFeaturesResponse.features:type_name -> pb.Feature
	7, // 2: pb.Feature.allocations:type_name -> pb.TreatmentAllocation
	0, // 3: pb.FeatureLab.GetTreatment:input_type -> pb.GetTreatmentRequest
	2, // 4: pb.FeatureLab.FetchFeature:input_type -> pb.FetchFeatureRequest
	4, // 5: pb.FeatureLab.FetchFeatures:input_type -> pb.FetchFeaturesRequest
	1, // 6: pb.FeatureLab.GetTreatment:output_type -> pb.GetTreatmentResponse
	3, // 7: pb.FeatureLab.FetchFeature:output_type -> pb.FetchFeatureResponse
	5, // 8: pb.FeatureLab.FetchFeatures:output_type -> pb.FetchFeaturesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_featurelabd_proto_init() }
func file_featurelabd_proto_init() {
	if File_featurelabd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_featurelabd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreatmentRequest); i {
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
		file_featurelabd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTreatmentResponse); i {
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
		file_featurelabd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureRequest); i {
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
		file_featurelabd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeatureResponse); i {
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
		file_featurelabd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeaturesRequest); i {
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
		file_featurelabd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFeaturesResponse); i {
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
		file_featurelabd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_featurelabd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreatmentAllocation); i {
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
			RawDescriptor: file_featurelabd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_featurelabd_proto_goTypes,
		DependencyIndexes: file_featurelabd_proto_depIdxs,
		MessageInfos:      file_featurelabd_proto_msgTypes,
	}.Build()
	File_featurelabd_proto = out.File
	file_featurelabd_proto_rawDesc = nil
	file_featurelabd_proto_goTypes = nil
	file_featurelabd_proto_depIdxs = nil
}