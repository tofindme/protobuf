// Code generated by protoc-gen-go.
// source: extend/extension.proto
// DO NOT EDIT!

package service

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// proto 文件信息
type FileSpec struct {
	FileName         *string                `protobuf:"bytes,1,req,name=file_name" json:"file_name,omitempty"`
	PackageName      *string                `protobuf:"bytes,2,req,name=package_name" json:"package_name,omitempty"`
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *FileSpec) Reset()         { *m = FileSpec{} }
func (m *FileSpec) String() string { return proto.CompactTextString(m) }
func (*FileSpec) ProtoMessage()    {}

func (m *FileSpec) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *FileSpec) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *FileSpec) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 服务信息
type ServiceSpec struct {
	ServiceName      *string                `protobuf:"bytes,1,req,name=service_name" json:"service_name,omitempty"`
	MethodList       []*ServiceMethodSpec   `protobuf:"bytes,2,rep,name=method_list" json:"method_list,omitempty"`
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}

func (m *ServiceSpec) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *ServiceSpec) GetMethodList() []*ServiceMethodSpec {
	if m != nil {
		return m.MethodList
	}
	return nil
}

func (m *ServiceSpec) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 服务方法列表
type ServiceMethodSpec struct {
	MethodName       *string                `protobuf:"bytes,1,req,name=method_name" json:"method_name,omitempty"`
	InputTypeName    *string                `protobuf:"bytes,2,req,name=input_type_name" json:"input_type_name,omitempty"`
	OutputTypeName   *string                `protobuf:"bytes,3,req,name=output_type_name" json:"output_type_name,omitempty"`
	HttpMethod       *string                `protobuf:"bytes,4,opt,name=http_method,def=GET" json:"http_method,omitempty"`
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,5,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ServiceMethodSpec) Reset()         { *m = ServiceMethodSpec{} }
func (m *ServiceMethodSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceMethodSpec) ProtoMessage()    {}

const Default_ServiceMethodSpec_HttpMethod string = "GET"

func (m *ServiceMethodSpec) GetMethodName() string {
	if m != nil && m.MethodName != nil {
		return *m.MethodName
	}
	return ""
}

func (m *ServiceMethodSpec) GetInputTypeName() string {
	if m != nil && m.InputTypeName != nil {
		return *m.InputTypeName
	}
	return ""
}

func (m *ServiceMethodSpec) GetOutputTypeName() string {
	if m != nil && m.OutputTypeName != nil {
		return *m.OutputTypeName
	}
	return ""
}

func (m *ServiceMethodSpec) GetHttpMethod() string {
	if m != nil && m.HttpMethod != nil {
		return *m.HttpMethod
	}
	return Default_ServiceMethodSpec_HttpMethod
}

func (m *ServiceMethodSpec) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 文件的扩展信息
type FileOption struct {
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *FileOption) Reset()         { *m = FileOption{} }
func (m *FileOption) String() string { return proto.CompactTextString(m) }
func (*FileOption) ProtoMessage()    {}

func (m *FileOption) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 服务的扩展信息
type ServiceOption struct {
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ServiceOption) Reset()         { *m = ServiceOption{} }
func (m *ServiceOption) String() string { return proto.CompactTextString(m) }
func (*ServiceOption) ProtoMessage()    {}

func (m *ServiceOption) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 服务方法的扩展信息
type MethodOption struct {
	Id               *uint32                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,2,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *MethodOption) Reset()         { *m = MethodOption{} }
func (m *MethodOption) String() string { return proto.CompactTextString(m) }
func (*MethodOption) ProtoMessage()    {}

func (m *MethodOption) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *MethodOption) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 消息的扩展信息
type MessageOption struct {
	ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs" json:"external_docs,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *MessageOption) Reset()         { *m = MessageOption{} }
func (m *MessageOption) String() string { return proto.CompactTextString(m) }
func (*MessageOption) ProtoMessage()    {}

func (m *MessageOption) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

// 消息成员的扩展信息
// 这是重要信息, 在运行时可动态获取改信息对 message 做合法性验证
type FieldOption struct {
	MinValue          *float64               `protobuf:"fixed64,1,opt,name=min_value" json:"min_value,omitempty"`
	MaxValue          *float64               `protobuf:"fixed64,2,opt,name=max_value" json:"max_value,omitempty"`
	MultipleOfValue   *float64               `protobuf:"fixed64,3,opt,name=multiple_of_value" json:"multiple_of_value,omitempty"`
	ExclusiveMinValue *float64               `protobuf:"fixed64,4,opt,name=exclusive_min_value" json:"exclusive_min_value,omitempty"`
	ExclusiveMaxValue *float64               `protobuf:"fixed64,5,opt,name=exclusive_max_value" json:"exclusive_max_value,omitempty"`
	MinLength         *int32                 `protobuf:"varint,6,opt,name=min_length" json:"min_length,omitempty"`
	MaxLength         *int32                 `protobuf:"varint,7,opt,name=max_length" json:"max_length,omitempty"`
	RegexpValue       *string                `protobuf:"bytes,8,opt,name=regexp_value" json:"regexp_value,omitempty"`
	EnumValue         []string               `protobuf:"bytes,9,rep,name=enum_value" json:"enum_value,omitempty"`
	ExternalDocs      *ExternalDocumentation `protobuf:"bytes,10,opt,name=external_docs" json:"external_docs,omitempty"`
	StructTag         *string                `protobuf:"bytes,11,opt,name=struct_tag" json:"struct_tag,omitempty"`
	XXX_unrecognized  []byte                 `json:"-"`
}

func (m *FieldOption) Reset()         { *m = FieldOption{} }
func (m *FieldOption) String() string { return proto.CompactTextString(m) }
func (*FieldOption) ProtoMessage()    {}

func (m *FieldOption) GetMinValue() float64 {
	if m != nil && m.MinValue != nil {
		return *m.MinValue
	}
	return 0
}

func (m *FieldOption) GetMaxValue() float64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *FieldOption) GetMultipleOfValue() float64 {
	if m != nil && m.MultipleOfValue != nil {
		return *m.MultipleOfValue
	}
	return 0
}

func (m *FieldOption) GetExclusiveMinValue() float64 {
	if m != nil && m.ExclusiveMinValue != nil {
		return *m.ExclusiveMinValue
	}
	return 0
}

func (m *FieldOption) GetExclusiveMaxValue() float64 {
	if m != nil && m.ExclusiveMaxValue != nil {
		return *m.ExclusiveMaxValue
	}
	return 0
}

func (m *FieldOption) GetMinLength() int32 {
	if m != nil && m.MinLength != nil {
		return *m.MinLength
	}
	return 0
}

func (m *FieldOption) GetMaxLength() int32 {
	if m != nil && m.MaxLength != nil {
		return *m.MaxLength
	}
	return 0
}

func (m *FieldOption) GetRegexpValue() string {
	if m != nil && m.RegexpValue != nil {
		return *m.RegexpValue
	}
	return ""
}

func (m *FieldOption) GetEnumValue() []string {
	if m != nil {
		return m.EnumValue
	}
	return nil
}

func (m *FieldOption) GetExternalDocs() *ExternalDocumentation {
	if m != nil {
		return m.ExternalDocs
	}
	return nil
}

func (m *FieldOption) GetStructTag() string {
	if m != nil && m.StructTag != nil {
		return *m.StructTag
	}
	return ""
}

// 扩展文档
type ExternalDocumentation struct {
	Title            *string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description      *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Url              *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalDocumentation) Reset()         { *m = ExternalDocumentation{} }
func (m *ExternalDocumentation) String() string { return proto.CompactTextString(m) }
func (*ExternalDocumentation) ProtoMessage()    {}

func (m *ExternalDocumentation) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *ExternalDocumentation) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ExternalDocumentation) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

var E_FileOption = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*FileOption)(nil),
	Field:         50001,
	Name:          "file_option",
	Tag:           "bytes,50001,opt,name=file_option",
}

var E_ServiceOption = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ServiceOption)(nil),
	Field:         50002,
	Name:          "service_option",
	Tag:           "bytes,50002,opt,name=service_option",
}

var E_MethodOption = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodOption)(nil),
	Field:         50003,
	Name:          "method_option",
	Tag:           "bytes,50003,opt,name=method_option",
}

var E_MessageOption = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MessageOption)(nil),
	Field:         50004,
	Name:          "message_option",
	Tag:           "bytes,50004,opt,name=message_option",
}

var E_FieldOption = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldOption)(nil),
	Field:         50005,
	Name:          "field_option",
	Tag:           "bytes,50005,opt,name=field_option",
}

func init() {
	proto.RegisterExtension(E_FileOption)
	proto.RegisterExtension(E_ServiceOption)
	proto.RegisterExtension(E_MethodOption)
	proto.RegisterExtension(E_MessageOption)
	proto.RegisterExtension(E_FieldOption)
}
