// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/update/v1/message.proto

package update

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	v11 "go.temporal.io/api/common/v1"
	v1 "go.temporal.io/api/enums/v1"
	v12 "go.temporal.io/api/failure/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifies client's intent to wait for Update results.
type WaitPolicy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Indicates the Update lifecycle stage that the Update must reach before
	// API call is returned.
	// NOTE: This field works together with API call timeout which is limited by
	// server timeout (maximum wait time). If server timeout is expired before
	// user specified timeout, API call returns even if specified stage is not reached.
	LifecycleStage v1.UpdateWorkflowExecutionLifecycleStage `protobuf:"varint,1,opt,name=lifecycle_stage,json=lifecycleStage,proto3,enum=temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage" json:"lifecycle_stage,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WaitPolicy) Reset() {
	*x = WaitPolicy{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitPolicy) ProtoMessage() {}

func (x *WaitPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitPolicy.ProtoReflect.Descriptor instead.
func (*WaitPolicy) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WaitPolicy) GetLifecycleStage() v1.UpdateWorkflowExecutionLifecycleStage {
	if x != nil {
		return x.LifecycleStage
	}
	return v1.UpdateWorkflowExecutionLifecycleStage(0)
}

// The data needed by a client to refer to a previously invoked Workflow Update.
type UpdateRef struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	WorkflowExecution *v11.WorkflowExecution `protobuf:"bytes,1,opt,name=workflow_execution,json=workflowExecution,proto3" json:"workflow_execution,omitempty"`
	UpdateId          string                 `protobuf:"bytes,2,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UpdateRef) Reset() {
	*x = UpdateRef{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRef) ProtoMessage() {}

func (x *UpdateRef) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRef.ProtoReflect.Descriptor instead.
func (*UpdateRef) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRef) GetWorkflowExecution() *v11.WorkflowExecution {
	if x != nil {
		return x.WorkflowExecution
	}
	return nil
}

func (x *UpdateRef) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

// The outcome of a Workflow Update: success or failure.
type Outcome struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*Outcome_Success
	//	*Outcome_Failure
	Value         isOutcome_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetValue() isOutcome_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Outcome) GetSuccess() *v11.Payloads {
	if x != nil {
		if x, ok := x.Value.(*Outcome_Success); ok {
			return x.Success
		}
	}
	return nil
}

func (x *Outcome) GetFailure() *v12.Failure {
	if x != nil {
		if x, ok := x.Value.(*Outcome_Failure); ok {
			return x.Failure
		}
	}
	return nil
}

type isOutcome_Value interface {
	isOutcome_Value()
}

type Outcome_Success struct {
	Success *v11.Payloads `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type Outcome_Failure struct {
	Failure *v12.Failure `protobuf:"bytes,2,opt,name=failure,proto3,oneof"`
}

func (*Outcome_Success) isOutcome_Value() {}

func (*Outcome_Failure) isOutcome_Value() {}

// Metadata about a Workflow Update.
type Meta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An ID with workflow-scoped uniqueness for this Update.
	UpdateId string `protobuf:"bytes,1,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	// A string identifying the agent that requested this Update.
	Identity      string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Meta) Reset() {
	*x = Meta{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Meta) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

func (x *Meta) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type Input struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Headers that are passed with the Update from the requesting entity.
	// These can include things like auth or tracing tokens.
	Header *v11.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The name of the Update handler to invoke on the target Workflow.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The arguments to pass to the named Update handler.
	Args          *v11.Payloads `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Input) Reset() {
	*x = Input{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Input) GetHeader() *v11.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Input) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Input) GetArgs() *v11.Payloads {
	if x != nil {
		return x.Args
	}
	return nil
}

// The client request that triggers a Workflow Update.
type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Input         *Input                 `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *Request) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Request) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has been rejected.
type Rejection struct {
	state                            protoimpl.MessageState `protogen:"open.v1"`
	RejectedRequestMessageId         string                 `protobuf:"bytes,1,opt,name=rejected_request_message_id,json=rejectedRequestMessageId,proto3" json:"rejected_request_message_id,omitempty"`
	RejectedRequestSequencingEventId int64                  `protobuf:"varint,2,opt,name=rejected_request_sequencing_event_id,json=rejectedRequestSequencingEventId,proto3" json:"rejected_request_sequencing_event_id,omitempty"`
	RejectedRequest                  *Request               `protobuf:"bytes,3,opt,name=rejected_request,json=rejectedRequest,proto3" json:"rejected_request,omitempty"`
	Failure                          *v12.Failure           `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *Rejection) Reset() {
	*x = Rejection{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rejection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rejection) ProtoMessage() {}

func (x *Rejection) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rejection.ProtoReflect.Descriptor instead.
func (*Rejection) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *Rejection) GetRejectedRequestMessageId() string {
	if x != nil {
		return x.RejectedRequestMessageId
	}
	return ""
}

func (x *Rejection) GetRejectedRequestSequencingEventId() int64 {
	if x != nil {
		return x.RejectedRequestSequencingEventId
	}
	return 0
}

func (x *Rejection) GetRejectedRequest() *Request {
	if x != nil {
		return x.RejectedRequest
	}
	return nil
}

func (x *Rejection) GetFailure() *v12.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has
// been accepted (i.e. passed the worker-side validation phase).
type Acceptance struct {
	state                            protoimpl.MessageState `protogen:"open.v1"`
	AcceptedRequestMessageId         string                 `protobuf:"bytes,1,opt,name=accepted_request_message_id,json=acceptedRequestMessageId,proto3" json:"accepted_request_message_id,omitempty"`
	AcceptedRequestSequencingEventId int64                  `protobuf:"varint,2,opt,name=accepted_request_sequencing_event_id,json=acceptedRequestSequencingEventId,proto3" json:"accepted_request_sequencing_event_id,omitempty"`
	AcceptedRequest                  *Request               `protobuf:"bytes,3,opt,name=accepted_request,json=acceptedRequest,proto3" json:"accepted_request,omitempty"`
	unknownFields                    protoimpl.UnknownFields
	sizeCache                        protoimpl.SizeCache
}

func (x *Acceptance) Reset() {
	*x = Acceptance{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Acceptance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptance) ProtoMessage() {}

func (x *Acceptance) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptance.ProtoReflect.Descriptor instead.
func (*Acceptance) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{7}
}

func (x *Acceptance) GetAcceptedRequestMessageId() string {
	if x != nil {
		return x.AcceptedRequestMessageId
	}
	return ""
}

func (x *Acceptance) GetAcceptedRequestSequencingEventId() int64 {
	if x != nil {
		return x.AcceptedRequestSequencingEventId
	}
	return 0
}

func (x *Acceptance) GetAcceptedRequest() *Request {
	if x != nil {
		return x.AcceptedRequest
	}
	return nil
}

// An Update protocol message indicating that a Workflow Update has
// completed with the contained outcome.
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Outcome       *Outcome               `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_update_v1_message_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_temporal_api_update_v1_message_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Response) GetOutcome() *Outcome {
	if x != nil {
		return x.Outcome
	}
	return nil
}

var File_temporal_api_update_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_update_v1_message_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x65, 0x0a,
	0x0f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x66, 0x12, 0x58, 0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x05,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x70, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xa2, 0x02, 0x0a, 0x09, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x24, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x20, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0xe7,
	0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x1b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x18, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x24,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x10,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x42, 0x89, 0x01, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x23, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0xaa, 0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69,
	0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xea,
	0x02, 0x1b, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_temporal_api_update_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_update_v1_message_proto_rawDescData []byte
)

func file_temporal_api_update_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_update_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_update_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_update_v1_message_proto_rawDesc), len(file_temporal_api_update_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_update_v1_message_proto_rawDescData
}

var file_temporal_api_update_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_temporal_api_update_v1_message_proto_goTypes = []any{
	(*WaitPolicy)(nil), // 0: temporal.api.update.v1.WaitPolicy
	(*UpdateRef)(nil),  // 1: temporal.api.update.v1.UpdateRef
	(*Outcome)(nil),    // 2: temporal.api.update.v1.Outcome
	(*Meta)(nil),       // 3: temporal.api.update.v1.Meta
	(*Input)(nil),      // 4: temporal.api.update.v1.Input
	(*Request)(nil),    // 5: temporal.api.update.v1.Request
	(*Rejection)(nil),  // 6: temporal.api.update.v1.Rejection
	(*Acceptance)(nil), // 7: temporal.api.update.v1.Acceptance
	(*Response)(nil),   // 8: temporal.api.update.v1.Response
	(v1.UpdateWorkflowExecutionLifecycleStage)(0), // 9: temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage
	(*v11.WorkflowExecution)(nil),                 // 10: temporal.api.common.v1.WorkflowExecution
	(*v11.Payloads)(nil),                          // 11: temporal.api.common.v1.Payloads
	(*v12.Failure)(nil),                           // 12: temporal.api.failure.v1.Failure
	(*v11.Header)(nil),                            // 13: temporal.api.common.v1.Header
}
var file_temporal_api_update_v1_message_proto_depIdxs = []int32{
	9,  // 0: temporal.api.update.v1.WaitPolicy.lifecycle_stage:type_name -> temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage
	10, // 1: temporal.api.update.v1.UpdateRef.workflow_execution:type_name -> temporal.api.common.v1.WorkflowExecution
	11, // 2: temporal.api.update.v1.Outcome.success:type_name -> temporal.api.common.v1.Payloads
	12, // 3: temporal.api.update.v1.Outcome.failure:type_name -> temporal.api.failure.v1.Failure
	13, // 4: temporal.api.update.v1.Input.header:type_name -> temporal.api.common.v1.Header
	11, // 5: temporal.api.update.v1.Input.args:type_name -> temporal.api.common.v1.Payloads
	3,  // 6: temporal.api.update.v1.Request.meta:type_name -> temporal.api.update.v1.Meta
	4,  // 7: temporal.api.update.v1.Request.input:type_name -> temporal.api.update.v1.Input
	5,  // 8: temporal.api.update.v1.Rejection.rejected_request:type_name -> temporal.api.update.v1.Request
	12, // 9: temporal.api.update.v1.Rejection.failure:type_name -> temporal.api.failure.v1.Failure
	5,  // 10: temporal.api.update.v1.Acceptance.accepted_request:type_name -> temporal.api.update.v1.Request
	3,  // 11: temporal.api.update.v1.Response.meta:type_name -> temporal.api.update.v1.Meta
	2,  // 12: temporal.api.update.v1.Response.outcome:type_name -> temporal.api.update.v1.Outcome
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_temporal_api_update_v1_message_proto_init() }
func file_temporal_api_update_v1_message_proto_init() {
	if File_temporal_api_update_v1_message_proto != nil {
		return
	}
	file_temporal_api_update_v1_message_proto_msgTypes[2].OneofWrappers = []any{
		(*Outcome_Success)(nil),
		(*Outcome_Failure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_update_v1_message_proto_rawDesc), len(file_temporal_api_update_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_update_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_update_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_update_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_update_v1_message_proto = out.File
	file_temporal_api_update_v1_message_proto_goTypes = nil
	file_temporal_api_update_v1_message_proto_depIdxs = nil
}
