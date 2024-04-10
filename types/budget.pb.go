// PATH="${PATH}:${HOME}/go/bin"

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: types/budget.proto

package types

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Amount   string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	ExpireAt string `protobuf:"bytes,4,opt,name=expireAt,proto3" json:"expireAt,omitempty"`
	SetAt    string `protobuf:"bytes,5,opt,name=setAt,proto3" json:"setAt,omitempty"`
	UserId   string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_budget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_types_budget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_types_budget_proto_rawDescGZIP(), []int{0}
}

func (x *Budget) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Budget) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Budget) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Budget) GetExpireAt() string {
	if x != nil {
		return x.ExpireAt
	}
	return ""
}

func (x *Budget) GetSetAt() string {
	if x != nil {
		return x.SetAt
	}
	return ""
}

func (x *Budget) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Income struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Amount     string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	ReceivedAt string `protobuf:"bytes,4,opt,name=receivedAt,proto3" json:"receivedAt,omitempty"`
	UserId     string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Income) Reset() {
	*x = Income{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_budget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Income) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Income) ProtoMessage() {}

func (x *Income) ProtoReflect() protoreflect.Message {
	mi := &file_types_budget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Income.ProtoReflect.Descriptor instead.
func (*Income) Descriptor() ([]byte, []int) {
	return file_types_budget_proto_rawDescGZIP(), []int{1}
}

func (x *Income) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Income) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Income) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Income) GetReceivedAt() string {
	if x != nil {
		return x.ReceivedAt
	}
	return ""
}

func (x *Income) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
}

func (x *BudgetRequest) Reset() {
	*x = BudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_budget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetRequest) ProtoMessage() {}

func (x *BudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_budget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetRequest.ProtoReflect.Descriptor instead.
func (*BudgetRequest) Descriptor() ([]byte, []int) {
	return file_types_budget_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type IncomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=iD,proto3" json:"iD,omitempty"`
}

func (x *IncomeRequest) Reset() {
	*x = IncomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_budget_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeRequest) ProtoMessage() {}

func (x *IncomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_budget_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeRequest.ProtoReflect.Descriptor instead.
func (*IncomeRequest) Descriptor() ([]byte, []int) {
	return file_types_budget_proto_rawDescGZIP(), []int{3}
}

func (x *IncomeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_types_budget_proto protoreflect.FileDescriptor

var file_types_budget_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x65, 0x74, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x65,
	0x74, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x06,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a,
	0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x22, 0x1f,
	0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x32,
	0xd2, 0x03, 0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x40, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x40, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x69, 0x2d, 0x41, 0x73, 0x73, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x73,
	0x68, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_budget_proto_rawDescOnce sync.Once
	file_types_budget_proto_rawDescData = file_types_budget_proto_rawDesc
)

func file_types_budget_proto_rawDescGZIP() []byte {
	file_types_budget_proto_rawDescOnce.Do(func() {
		file_types_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_budget_proto_rawDescData)
	})
	return file_types_budget_proto_rawDescData
}

var file_types_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_budget_proto_goTypes = []interface{}{
	(*Budget)(nil),        // 0: users.Budget
	(*Income)(nil),        // 1: users.Income
	(*BudgetRequest)(nil), // 2: users.BudgetRequest
	(*IncomeRequest)(nil), // 3: users.IncomeRequest
	(*empty.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_types_budget_proto_depIdxs = []int32{
	0, // 0: users.BudgetService.InsertBudget:input_type -> users.Budget
	2, // 1: users.BudgetService.GetBudgetByID:input_type -> users.BudgetRequest
	0, // 2: users.BudgetService.UpdateBudgetrByID:input_type -> users.Budget
	2, // 3: users.BudgetService.DeleteBudgetByID:input_type -> users.BudgetRequest
	1, // 4: users.BudgetService.InsertIncome:input_type -> users.Income
	3, // 5: users.BudgetService.GetIncomeByID:input_type -> users.IncomeRequest
	1, // 6: users.BudgetService.UpdateIncomeByID:input_type -> users.Income
	3, // 7: users.BudgetService.DeleteIncomeByID:input_type -> users.IncomeRequest
	0, // 8: users.BudgetService.InsertBudget:output_type -> users.Budget
	0, // 9: users.BudgetService.GetBudgetByID:output_type -> users.Budget
	4, // 10: users.BudgetService.UpdateBudgetrByID:output_type -> google.protobuf.Empty
	4, // 11: users.BudgetService.DeleteBudgetByID:output_type -> google.protobuf.Empty
	1, // 12: users.BudgetService.InsertIncome:output_type -> users.Income
	1, // 13: users.BudgetService.GetIncomeByID:output_type -> users.Income
	4, // 14: users.BudgetService.UpdateIncomeByID:output_type -> google.protobuf.Empty
	4, // 15: users.BudgetService.DeleteIncomeByID:output_type -> google.protobuf.Empty
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_budget_proto_init() }
func file_types_budget_proto_init() {
	if File_types_budget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_budget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
		file_types_budget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Income); i {
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
		file_types_budget_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetRequest); i {
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
		file_types_budget_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomeRequest); i {
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
			RawDescriptor: file_types_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_budget_proto_goTypes,
		DependencyIndexes: file_types_budget_proto_depIdxs,
		MessageInfos:      file_types_budget_proto_msgTypes,
	}.Build()
	File_types_budget_proto = out.File
	file_types_budget_proto_rawDesc = nil
	file_types_budget_proto_goTypes = nil
	file_types_budget_proto_depIdxs = nil
}
