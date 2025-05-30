// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: service.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Username          string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName          string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName      *string                `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Email         *string                `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Password      *string                `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type LoginUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	mi := &file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	User                  *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	SessionId             string                 `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	AccessToken           string                 `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AccessTokenExpiresAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=access_token_expires_at,json=accessTokenExpiresAt,proto3" json:"access_token_expires_at,omitempty"`
	RefreshTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=refresh_token_expires_at,json=refreshTokenExpiresAt,proto3" json:"refresh_token_expires_at,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	mi := &file_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *LoginUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *LoginUserResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LoginUserResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginUserResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginUserResponse) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

func (x *LoginUserResponse) GetRefreshTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiresAt
	}
	return nil
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EmailId       int64                  `protobuf:"varint,1,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	SecretCode    string                 `protobuf:"bytes,2,opt,name=secret_code,json=secretCode,proto3" json:"secret_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	mi := &file_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyEmailRequest) GetEmailId() int64 {
	if x != nil {
		return x.EmailId
	}
	return 0
}

func (x *VerifyEmailRequest) GetSecretCode() string {
	if x != nil {
		return x.SecretCode
	}
	return ""
}

type VerifyEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsVerified    bool                   `protobuf:"varint,1,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailResponse) Reset() {
	*x = VerifyEmailResponse{}
	mi := &file_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResponse) ProtoMessage() {}

func (x *VerifyEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResponse.ProtoReflect.Descriptor instead.
func (*VerifyEmailResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *VerifyEmailResponse) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

var File_service_proto protoreflect.FileDescriptor

const file_service_proto_rawDesc = "" +
	"\n" +
	"\rservice.proto\x12\x03gen\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/rpc/error_details.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xdc\x01\n" +
	"\x04User\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12J\n" +
	"\x13password_changed_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x11passwordChangedAt\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"~\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\"3\n" +
	"\x12CreateUserResponse\x12\x1d\n" +
	"\x04user\x18\x01 \x01(\v2\t.gen.UserR\x04user\"\xb2\x01\n" +
	"\x11UpdateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12 \n" +
	"\tfull_name\x18\x02 \x01(\tH\x00R\bfullName\x88\x01\x01\x12\x19\n" +
	"\x05email\x18\x03 \x01(\tH\x01R\x05email\x88\x01\x01\x12\x1f\n" +
	"\bpassword\x18\x04 \x01(\tH\x02R\bpassword\x88\x01\x01B\f\n" +
	"\n" +
	"_full_nameB\b\n" +
	"\x06_emailB\v\n" +
	"\t_password\"3\n" +
	"\x12UpdateUserResponse\x12\x1d\n" +
	"\x04user\x18\x01 \x01(\v2\t.gen.UserR\x04user\"J\n" +
	"\x10LoginUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\xc1\x02\n" +
	"\x11LoginUserResponse\x12\x1d\n" +
	"\x04user\x18\x01 \x01(\v2\t.gen.UserR\x04user\x12\x1d\n" +
	"\n" +
	"session_id\x18\x02 \x01(\tR\tsessionId\x12!\n" +
	"\faccess_token\x18\x03 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x04 \x01(\tR\frefreshToken\x12Q\n" +
	"\x17access_token_expires_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x14accessTokenExpiresAt\x12S\n" +
	"\x18refresh_token_expires_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\x15refreshTokenExpiresAt\"P\n" +
	"\x12VerifyEmailRequest\x12\x19\n" +
	"\bemail_id\x18\x01 \x01(\x03R\aemailId\x12\x1f\n" +
	"\vsecret_code\x18\x02 \x01(\tR\n" +
	"secretCode\"6\n" +
	"\x13VerifyEmailResponse\x12\x1f\n" +
	"\vis_verified\x18\x01 \x01(\bR\n" +
	"isVerified2\xe2\x04\n" +
	"\vAuthService\x12\x89\x01\n" +
	"\n" +
	"CreateUser\x12\x16.gen.CreateUserRequest\x1a\x17.gen.CreateUserResponse\"J\x92A4\x12\x0fCreate new user\x1a!Use this API to create a new user\x82\xd3\xe4\x93\x02\r:\x01*\"\b/v1/user\x12\x81\x01\n" +
	"\n" +
	"UpdateUser\x12\x16.gen.UpdateUserRequest\x1a\x17.gen.UpdateUserResponse\"B\x92A,\x12\vUpdate user\x1a\x1dUse this API to update a user\x82\xd3\xe4\x93\x02\r:\x01*2\b/v1/user\x12\xa2\x01\n" +
	"\tLoginUser\x12\x15.gen.LoginUserRequest\x1a\x16.gen.LoginUserResponse\"f\x92AJ\x12\n" +
	"Login user\x1a<Use this API to login user and get access and refresh tokens\x82\xd3\xe4\x93\x02\x13:\x01*\"\x0e/v1/user/login\x12\x9d\x01\n" +
	"\vVerifyEmail\x12\x17.gen.VerifyEmailRequest\x1a\x18.gen.VerifyEmailResponse\"[\x92A;\x12\fVerify Email\x1a+Use this API to verify user's email address\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/user/verify-emailB|\x92AM\x12K\n" +
	"\fAuth Service\"6\n" +
	"\x11Lucas H. Santiago\x12!https://github.com/lucashsantiago2\x031.0Z*github.com/lucasHSantiago/gobank/proto/genb\x06proto3"

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData []byte
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)))
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_proto_goTypes = []any{
	(*User)(nil),                  // 0: gen.User
	(*CreateUserRequest)(nil),     // 1: gen.CreateUserRequest
	(*CreateUserResponse)(nil),    // 2: gen.CreateUserResponse
	(*UpdateUserRequest)(nil),     // 3: gen.UpdateUserRequest
	(*UpdateUserResponse)(nil),    // 4: gen.UpdateUserResponse
	(*LoginUserRequest)(nil),      // 5: gen.LoginUserRequest
	(*LoginUserResponse)(nil),     // 6: gen.LoginUserResponse
	(*VerifyEmailRequest)(nil),    // 7: gen.VerifyEmailRequest
	(*VerifyEmailResponse)(nil),   // 8: gen.VerifyEmailResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_service_proto_depIdxs = []int32{
	9,  // 0: gen.User.password_changed_at:type_name -> google.protobuf.Timestamp
	9,  // 1: gen.User.created_at:type_name -> google.protobuf.Timestamp
	0,  // 2: gen.CreateUserResponse.user:type_name -> gen.User
	0,  // 3: gen.UpdateUserResponse.user:type_name -> gen.User
	0,  // 4: gen.LoginUserResponse.user:type_name -> gen.User
	9,  // 5: gen.LoginUserResponse.access_token_expires_at:type_name -> google.protobuf.Timestamp
	9,  // 6: gen.LoginUserResponse.refresh_token_expires_at:type_name -> google.protobuf.Timestamp
	1,  // 7: gen.AuthService.CreateUser:input_type -> gen.CreateUserRequest
	3,  // 8: gen.AuthService.UpdateUser:input_type -> gen.UpdateUserRequest
	5,  // 9: gen.AuthService.LoginUser:input_type -> gen.LoginUserRequest
	7,  // 10: gen.AuthService.VerifyEmail:input_type -> gen.VerifyEmailRequest
	2,  // 11: gen.AuthService.CreateUser:output_type -> gen.CreateUserResponse
	4,  // 12: gen.AuthService.UpdateUser:output_type -> gen.UpdateUserResponse
	6,  // 13: gen.AuthService.LoginUser:output_type -> gen.LoginUserResponse
	8,  // 14: gen.AuthService.VerifyEmail:output_type -> gen.VerifyEmailResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
