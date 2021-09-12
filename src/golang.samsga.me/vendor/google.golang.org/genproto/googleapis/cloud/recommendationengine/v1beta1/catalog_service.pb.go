// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/catalog_service.proto

package recommendationengine

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for CreateCatalogItem method.
type CreateCatalogItemRequest struct {
	// Required. The parent catalog resource name, such as
	// "projects/*/locations/global/catalogs/default_catalog".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The catalog item to create.
	CatalogItem          *CatalogItem `protobuf:"bytes,2,opt,name=catalog_item,json=catalogItem,proto3" json:"catalog_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateCatalogItemRequest) Reset()         { *m = CreateCatalogItemRequest{} }
func (m *CreateCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCatalogItemRequest) ProtoMessage()    {}
func (*CreateCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{0}
}

func (m *CreateCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCatalogItemRequest.Unmarshal(m, b)
}
func (m *CreateCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCatalogItemRequest.Merge(m, src)
}
func (m *CreateCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCatalogItemRequest.Size(m)
}
func (m *CreateCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCatalogItemRequest proto.InternalMessageInfo

func (m *CreateCatalogItemRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCatalogItemRequest) GetCatalogItem() *CatalogItem {
	if m != nil {
		return m.CatalogItem
	}
	return nil
}

// Request message for GetCatalogItem method.
type GetCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogitems/some_catalog_item_id".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCatalogItemRequest) Reset()         { *m = GetCatalogItemRequest{} }
func (m *GetCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetCatalogItemRequest) ProtoMessage()    {}
func (*GetCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{1}
}

func (m *GetCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCatalogItemRequest.Unmarshal(m, b)
}
func (m *GetCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *GetCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCatalogItemRequest.Merge(m, src)
}
func (m *GetCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetCatalogItemRequest.Size(m)
}
func (m *GetCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCatalogItemRequest proto.InternalMessageInfo

func (m *GetCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for ListCatalogItems method.
type ListCatalogItemsRequest struct {
	// Required. The parent catalog resource name, such as
	// "projects/*/locations/global/catalogs/default_catalog".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Maximum number of results to return per page. If zero, the
	// service will choose a reasonable default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The previous ListCatalogItemsResponse.next_page_token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. A filter to apply on the list results.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCatalogItemsRequest) Reset()         { *m = ListCatalogItemsRequest{} }
func (m *ListCatalogItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCatalogItemsRequest) ProtoMessage()    {}
func (*ListCatalogItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{2}
}

func (m *ListCatalogItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCatalogItemsRequest.Unmarshal(m, b)
}
func (m *ListCatalogItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCatalogItemsRequest.Marshal(b, m, deterministic)
}
func (m *ListCatalogItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCatalogItemsRequest.Merge(m, src)
}
func (m *ListCatalogItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCatalogItemsRequest.Size(m)
}
func (m *ListCatalogItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCatalogItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCatalogItemsRequest proto.InternalMessageInfo

func (m *ListCatalogItemsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCatalogItemsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCatalogItemsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCatalogItemsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Response message for ListCatalogItems method.
type ListCatalogItemsResponse struct {
	// The catalog items.
	CatalogItems []*CatalogItem `protobuf:"bytes,1,rep,name=catalog_items,json=catalogItems,proto3" json:"catalog_items,omitempty"`
	// If empty, the list is complete. If nonempty, the token to pass to the next
	// request's ListCatalogItemRequest.page_token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCatalogItemsResponse) Reset()         { *m = ListCatalogItemsResponse{} }
func (m *ListCatalogItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCatalogItemsResponse) ProtoMessage()    {}
func (*ListCatalogItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{3}
}

func (m *ListCatalogItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCatalogItemsResponse.Unmarshal(m, b)
}
func (m *ListCatalogItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCatalogItemsResponse.Marshal(b, m, deterministic)
}
func (m *ListCatalogItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCatalogItemsResponse.Merge(m, src)
}
func (m *ListCatalogItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCatalogItemsResponse.Size(m)
}
func (m *ListCatalogItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCatalogItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCatalogItemsResponse proto.InternalMessageInfo

func (m *ListCatalogItemsResponse) GetCatalogItems() []*CatalogItem {
	if m != nil {
		return m.CatalogItems
	}
	return nil
}

func (m *ListCatalogItemsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for UpdateCatalogItem method.
type UpdateCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogItems/some_catalog_item_id".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The catalog item to update/create. The 'catalog_item_id' field
	// has to match that in the 'name'.
	CatalogItem *CatalogItem `protobuf:"bytes,2,opt,name=catalog_item,json=catalogItem,proto3" json:"catalog_item,omitempty"`
	// Optional. Indicates which fields in the provided 'item' to update. If not
	// set, will by default update all fields.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCatalogItemRequest) Reset()         { *m = UpdateCatalogItemRequest{} }
func (m *UpdateCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCatalogItemRequest) ProtoMessage()    {}
func (*UpdateCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{4}
}

func (m *UpdateCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCatalogItemRequest.Unmarshal(m, b)
}
func (m *UpdateCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCatalogItemRequest.Merge(m, src)
}
func (m *UpdateCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCatalogItemRequest.Size(m)
}
func (m *UpdateCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCatalogItemRequest proto.InternalMessageInfo

func (m *UpdateCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCatalogItemRequest) GetCatalogItem() *CatalogItem {
	if m != nil {
		return m.CatalogItem
	}
	return nil
}

func (m *UpdateCatalogItemRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request message for DeleteCatalogItem method.
type DeleteCatalogItemRequest struct {
	// Required. Full resource name of catalog item, such as
	// "projects/*/locations/global/catalogs/default_catalog/catalogItems/some_catalog_item_id".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCatalogItemRequest) Reset()         { *m = DeleteCatalogItemRequest{} }
func (m *DeleteCatalogItemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCatalogItemRequest) ProtoMessage()    {}
func (*DeleteCatalogItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d69b5a822db872f, []int{5}
}

func (m *DeleteCatalogItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCatalogItemRequest.Unmarshal(m, b)
}
func (m *DeleteCatalogItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCatalogItemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCatalogItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCatalogItemRequest.Merge(m, src)
}
func (m *DeleteCatalogItemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCatalogItemRequest.Size(m)
}
func (m *DeleteCatalogItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCatalogItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCatalogItemRequest proto.InternalMessageInfo

func (m *DeleteCatalogItemRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.CreateCatalogItemRequest")
	proto.RegisterType((*GetCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.GetCatalogItemRequest")
	proto.RegisterType((*ListCatalogItemsRequest)(nil), "google.cloud.recommendationengine.v1beta1.ListCatalogItemsRequest")
	proto.RegisterType((*ListCatalogItemsResponse)(nil), "google.cloud.recommendationengine.v1beta1.ListCatalogItemsResponse")
	proto.RegisterType((*UpdateCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.UpdateCatalogItemRequest")
	proto.RegisterType((*DeleteCatalogItemRequest)(nil), "google.cloud.recommendationengine.v1beta1.DeleteCatalogItemRequest")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/catalog_service.proto", fileDescriptor_5d69b5a822db872f)
}

var fileDescriptor_5d69b5a822db872f = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xd7, 0x38, 0x6d, 0x45, 0x26, 0x69, 0x21, 0x23, 0x41, 0xac, 0x2d, 0x88, 0x95, 0x2b, 0xa1,
	0x76, 0xd5, 0x7a, 0x48, 0x2a, 0x15, 0x64, 0x84, 0xc8, 0xee, 0x26, 0x84, 0x48, 0xad, 0x08, 0x5b,
	0x08, 0x12, 0x04, 0xa2, 0x89, 0xf7, 0xc5, 0x31, 0xb1, 0x67, 0x8c, 0x67, 0x36, 0x81, 0xa2, 0x4a,
	0x08, 0xf1, 0x0d, 0x38, 0x20, 0x4e, 0x95, 0x38, 0xf2, 0x19, 0x38, 0x71, 0xcc, 0xb1, 0x08, 0x0e,
	0x11, 0x07, 0x0e, 0x20, 0x21, 0x3e, 0x05, 0xf2, 0x8c, 0x4d, 0xbc, 0xff, 0xc4, 0x7a, 0x97, 0x1e,
	0xfd, 0xde, 0x6f, 0xde, 0xbc, 0xdf, 0x7b, 0xbf, 0xf7, 0x3c, 0xf8, 0x8d, 0x40, 0x88, 0x20, 0x02,
	0xea, 0x47, 0xa2, 0xd7, 0xa5, 0x29, 0xf8, 0x22, 0x8e, 0x81, 0x77, 0x99, 0x0a, 0x05, 0x07, 0x1e,
	0x84, 0x1c, 0xe8, 0xf1, 0xca, 0x3e, 0x28, 0xb6, 0x42, 0x7d, 0xa6, 0x58, 0x24, 0x82, 0x3d, 0x09,
	0xe9, 0x71, 0xe8, 0x83, 0x9b, 0xa4, 0x42, 0x09, 0x72, 0xc3, 0x04, 0x70, 0x75, 0x00, 0x77, 0x54,
	0x00, 0x37, 0x0f, 0x50, 0x7b, 0x3e, 0xbf, 0x8b, 0x25, 0x21, 0x65, 0x9c, 0x0b, 0xa5, 0x41, 0xd2,
	0x04, 0xaa, 0x2d, 0x97, 0xbc, 0x7e, 0x14, 0x02, 0x57, 0xb9, 0xe3, 0xc5, 0x92, 0xe3, 0x20, 0x84,
	0xa8, 0xbb, 0xb7, 0x0f, 0x87, 0xec, 0x38, 0x14, 0x69, 0x0e, 0x78, 0xa5, 0x32, 0x87, 0xfc, 0xe0,
	0x9d, 0xc9, 0x0f, 0x86, 0x71, 0x22, 0xd2, 0x22, 0xa3, 0x6b, 0xf9, 0xb9, 0x48, 0xf0, 0x20, 0xed,
	0x71, 0x1e, 0xf2, 0x80, 0x8a, 0x04, 0xd2, 0x3e, 0x3e, 0x57, 0x73, 0x90, 0xfe, 0xda, 0xef, 0x1d,
	0x50, 0x88, 0x13, 0xf5, 0x79, 0xee, 0xac, 0x0f, 0x3a, 0x0d, 0xb1, 0x98, 0xc9, 0x23, 0x83, 0x70,
	0xbe, 0x45, 0xd8, 0x6e, 0xa7, 0xc0, 0x14, 0xb4, 0x4d, 0xce, 0x5b, 0x0a, 0xe2, 0x0e, 0x7c, 0xda,
	0x03, 0xa9, 0xc8, 0x55, 0x7c, 0x29, 0x61, 0x29, 0x70, 0x65, 0xa3, 0x3a, 0xba, 0x3e, 0xdf, 0x9a,
	0xfb, 0xbd, 0x69, 0x75, 0x72, 0x13, 0xf9, 0x18, 0x2f, 0x16, 0xad, 0x0a, 0x15, 0xc4, 0xb6, 0x55,
	0x47, 0xd7, 0x17, 0x56, 0xef, 0xb8, 0x13, 0x37, 0xca, 0x2d, 0xdd, 0x68, 0x42, 0x2f, 0xf8, 0xe7,
	0x16, 0xe7, 0x65, 0xfc, 0xec, 0x26, 0xa8, 0x11, 0x59, 0x2d, 0xe3, 0x0b, 0x9c, 0xc5, 0x50, 0xce,
	0x49, 0x1b, 0x9c, 0xef, 0x10, 0x5e, 0xbe, 0x1b, 0xca, 0xf2, 0x19, 0x39, 0x11, 0x95, 0x3a, 0x9e,
	0x4f, 0x58, 0x00, 0x7b, 0x32, 0x7c, 0x00, 0x9a, 0xc7, 0xc5, 0xcc, 0x8f, 0x3a, 0x4f, 0x65, 0xd6,
	0xfb, 0xe1, 0x03, 0x20, 0x0e, 0xc6, 0x1a, 0xa1, 0xc4, 0x11, 0x70, 0x7b, 0xae, 0x08, 0x81, 0x3a,
	0xfa, 0xe0, 0xbb, 0x99, 0x35, 0xbb, 0xe2, 0x20, 0x8c, 0x14, 0xa4, 0xf6, 0x85, 0x73, 0x7f, 0x6e,
	0x72, 0x1e, 0x21, 0x6c, 0x0f, 0xe7, 0x26, 0x13, 0xc1, 0x25, 0x90, 0x0f, 0xf1, 0xe5, 0x72, 0x29,
	0xa5, 0x8d, 0xea, 0x73, 0xd3, 0xd7, 0xb2, 0xb3, 0x58, 0x2a, 0xa3, 0x24, 0x2f, 0xe1, 0xa7, 0x39,
	0x7c, 0xa6, 0xf6, 0x4a, 0xf9, 0x67, 0x14, 0xe7, 0x3b, 0x97, 0x33, 0xf3, 0x76, 0x91, 0xbe, 0xf3,
	0x0b, 0xc2, 0xf6, 0x7b, 0x49, 0x77, 0xb4, 0x12, 0xc6, 0xd5, 0xfc, 0x49, 0xab, 0x80, 0xac, 0xe1,
	0x85, 0x9e, 0x4e, 0x4a, 0x8b, 0x56, 0x57, 0x7e, 0x61, 0xb5, 0x56, 0x84, 0x2f, 0x74, 0xed, 0xbe,
	0x99, 0xe9, 0xfa, 0x1e, 0x93, 0x47, 0xa6, 0xea, 0xd8, 0x9c, 0xc9, 0x0c, 0xce, 0x6d, 0x6c, 0xaf,
	0x43, 0x04, 0x95, 0x68, 0xad, 0xfe, 0xb8, 0x88, 0xaf, 0xe4, 0xf8, 0xfb, 0x66, 0x0f, 0x91, 0xbf,
	0x10, 0x5e, 0x1a, 0x9a, 0x14, 0xd2, 0xae, 0xc2, 0x74, 0xcc, 0x9c, 0xd5, 0xa6, 0x2c, 0x97, 0xb3,
	0xf3, 0xd5, 0xcf, 0x7f, 0x7c, 0x63, 0x6d, 0x3b, 0x6b, 0xff, 0xee, 0x8f, 0x2f, 0x8c, 0xa2, 0x5f,
	0x4f, 0x52, 0xf1, 0x09, 0xf8, 0x4a, 0xd2, 0x06, 0x8d, 0x84, 0x6f, 0x16, 0x06, 0x6d, 0x14, 0x7b,
	0x49, 0xd2, 0xc6, 0x43, 0x5a, 0x16, 0x89, 0xd7, 0xd7, 0x44, 0xf2, 0x1b, 0xc2, 0x57, 0xfa, 0x67,
	0x8f, 0xac, 0x55, 0x48, 0x71, 0xe4, 0xd8, 0x4e, 0x4d, 0x72, 0xfb, 0xac, 0xa9, 0x7b, 0xa2, 0xb9,
	0xb6, 0x49, 0xf3, 0x9c, 0x6b, 0x66, 0xfd, 0x6f, 0xa6, 0x7d, 0x44, 0x69, 0xa3, 0xf1, 0x90, 0xfc,
	0x89, 0xf0, 0x33, 0x83, 0xb3, 0x48, 0x5a, 0x15, 0xd2, 0x1b, 0xb3, 0x64, 0x6a, 0xed, 0x99, 0x62,
	0x98, 0x65, 0xe0, 0xbc, 0xa5, 0x89, 0xb6, 0xc8, 0xcc, 0x4d, 0x25, 0x5f, 0x5b, 0x78, 0x69, 0x68,
	0xa2, 0x2b, 0x29, 0x76, 0xdc, 0x3e, 0x98, 0xba, 0x99, 0xe9, 0x59, 0xd3, 0x2e, 0x4b, 0xed, 0x66,
	0x69, 0xb8, 0x35, 0xef, 0x77, 0x56, 0x67, 0x6f, 0xf0, 0x80, 0x9a, 0x7f, 0x42, 0x78, 0x69, 0x68,
	0x03, 0x54, 0x2a, 0xc3, 0xb8, 0xfd, 0x51, 0x7b, 0x6e, 0x68, 0x11, 0x6d, 0x64, 0x7f, 0xdf, 0x01,
	0xcd, 0x36, 0xfe, 0x07, 0xcd, 0x3e, 0xb6, 0x30, 0xd9, 0xd2, 0x8f, 0x83, 0x3e, 0xd5, 0xae, 0x57,
	0x60, 0x31, 0x7c, 0xbc, 0xa0, 0xf1, 0x42, 0x11, 0xa5, 0xf4, 0xd2, 0x70, 0xdf, 0x2e, 0x5e, 0x1a,
	0xce, 0xaf, 0xe8, 0xb4, 0xf9, 0x25, 0xc2, 0xeb, 0xb3, 0x5d, 0x95, 0xff, 0xeb, 0x5e, 0xad, 0x1a,
	0xe5, 0x1e, 0x28, 0xd6, 0x65, 0x8a, 0xe9, 0x5a, 0xde, 0x75, 0x36, 0x67, 0xde, 0x75, 0xe6, 0x8d,
	0xe5, 0xa1, 0x46, 0xed, 0xfd, 0xd3, 0xe6, 0xb5, 0x91, 0x19, 0x98, 0x1c, 0x59, 0x12, 0x4a, 0xd7,
	0x17, 0xf1, 0xe3, 0xa6, 0x7b, 0xa8, 0x54, 0x22, 0x3d, 0x4a, 0x4f, 0x4e, 0x4e, 0x06, 0x9c, 0x94,
	0xf5, 0xd4, 0xa1, 0x79, 0xda, 0xdd, 0x4a, 0x22, 0xa6, 0x0e, 0x44, 0x1a, 0xb7, 0x1e, 0x59, 0xf8,
	0x96, 0x2f, 0xe2, 0xc9, 0x9b, 0xb3, 0x8d, 0x3e, 0xf8, 0x28, 0x07, 0x07, 0x22, 0x62, 0x3c, 0x70,
	0x45, 0x1a, 0xd0, 0x00, 0xb8, 0x16, 0x16, 0x3d, 0xbf, 0x72, 0x82, 0x47, 0xe4, 0x6b, 0xa3, 0x9c,
	0xdf, 0x5b, 0x17, 0x3b, 0x1b, 0xed, 0xe6, 0xd6, 0x0f, 0xd6, 0x8d, 0x4d, 0x73, 0x4f, 0x5b, 0x27,
	0xd5, 0xe9, 0xc3, 0x6e, 0x98, 0xa4, 0x76, 0x56, 0x5a, 0x59, 0xa0, 0xd3, 0x02, 0xbb, 0xab, 0xb1,
	0xbb, 0xa3, 0xb0, 0xbb, 0x3b, 0xe6, 0xd2, 0xbf, 0xad, 0x9b, 0x06, 0xeb, 0x79, 0x1a, 0xec, 0x79,
	0xa3, 0xd0, 0x9e, 0x97, 0xc3, 0xf7, 0x2f, 0x69, 0x62, 0xb7, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0x6c, 0x95, 0x59, 0x19, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	// Creates a catalog item.
	CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Gets a specific catalog item.
	GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Gets a list of catalog items.
	ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error)
	// Updates a catalog item. Partial updating is supported. Non-existing
	// items will be created.
	UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error)
	// Deletes a catalog item.
	DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Bulk import of multiple catalog items. Request processing may be
	// synchronous. No partial updating supported. Non-existing items will be
	// created.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully updated.
	ImportCatalogItems(ctx context.Context, in *ImportCatalogItemsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateCatalogItem(ctx context.Context, in *CreateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/CreateCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCatalogItem(ctx context.Context, in *GetCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/GetCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListCatalogItems(ctx context.Context, in *ListCatalogItemsRequest, opts ...grpc.CallOption) (*ListCatalogItemsResponse, error) {
	out := new(ListCatalogItemsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/ListCatalogItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCatalogItem(ctx context.Context, in *UpdateCatalogItemRequest, opts ...grpc.CallOption) (*CatalogItem, error) {
	out := new(CatalogItem)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/UpdateCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCatalogItem(ctx context.Context, in *DeleteCatalogItemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/DeleteCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ImportCatalogItems(ctx context.Context, in *ImportCatalogItemsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.recommendationengine.v1beta1.CatalogService/ImportCatalogItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	// Creates a catalog item.
	CreateCatalogItem(context.Context, *CreateCatalogItemRequest) (*CatalogItem, error)
	// Gets a specific catalog item.
	GetCatalogItem(context.Context, *GetCatalogItemRequest) (*CatalogItem, error)
	// Gets a list of catalog items.
	ListCatalogItems(context.Context, *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error)
	// Updates a catalog item. Partial updating is supported. Non-existing
	// items will be created.
	UpdateCatalogItem(context.Context, *UpdateCatalogItemRequest) (*CatalogItem, error)
	// Deletes a catalog item.
	DeleteCatalogItem(context.Context, *DeleteCatalogItemRequest) (*empty.Empty, error)
	// Bulk import of multiple catalog items. Request processing may be
	// synchronous. No partial updating supported. Non-existing items will be
	// created.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully updated.
	ImportCatalogItems(context.Context, *ImportCatalogItemsRequest) (*longrunning.Operation, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) CreateCatalogItem(ctx context.Context, req *CreateCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCatalogItem(ctx context.Context, req *GetCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) ListCatalogItems(ctx context.Context, req *ListCatalogItemsRequest) (*ListCatalogItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogItems not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateCatalogItem(ctx context.Context, req *UpdateCatalogItemRequest) (*CatalogItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteCatalogItem(ctx context.Context, req *DeleteCatalogItemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCatalogItem not implemented")
}
func (*UnimplementedCatalogServiceServer) ImportCatalogItems(ctx context.Context, req *ImportCatalogItemsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCatalogItems not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_CreateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/CreateCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCatalogItem(ctx, req.(*CreateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/GetCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCatalogItem(ctx, req.(*GetCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/ListCatalogItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListCatalogItems(ctx, req.(*ListCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/UpdateCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCatalogItem(ctx, req.(*UpdateCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/DeleteCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCatalogItem(ctx, req.(*DeleteCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ImportCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ImportCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.recommendationengine.v1beta1.CatalogService/ImportCatalogItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ImportCatalogItems(ctx, req.(*ImportCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.recommendationengine.v1beta1.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCatalogItem",
			Handler:    _CatalogService_CreateCatalogItem_Handler,
		},
		{
			MethodName: "GetCatalogItem",
			Handler:    _CatalogService_GetCatalogItem_Handler,
		},
		{
			MethodName: "ListCatalogItems",
			Handler:    _CatalogService_ListCatalogItems_Handler,
		},
		{
			MethodName: "UpdateCatalogItem",
			Handler:    _CatalogService_UpdateCatalogItem_Handler,
		},
		{
			MethodName: "DeleteCatalogItem",
			Handler:    _CatalogService_DeleteCatalogItem_Handler,
		},
		{
			MethodName: "ImportCatalogItems",
			Handler:    _CatalogService_ImportCatalogItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/recommendationengine/v1beta1/catalog_service.proto",
}
