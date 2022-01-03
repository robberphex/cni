// Copyright 2019 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: cnigrpc.proto

//package cnigrpc;

package libcni

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

type ListNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNetworkRequest) Reset() {
	*x = ListNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkRequest) ProtoMessage() {}

func (x *ListNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkRequest.ProtoReflect.Descriptor instead.
func (*ListNetworkRequest) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{0}
}

type ListNetworkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Networks []*NetworkInfo `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
}

func (x *ListNetworkReply) Reset() {
	*x = ListNetworkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworkReply) ProtoMessage() {}

func (x *ListNetworkReply) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworkReply.ProtoReflect.Descriptor instead.
func (*ListNetworkReply) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ListNetworkReply) GetNetworks() []*NetworkInfo {
	if x != nil {
		return x.Networks
	}
	return nil
}

type NetworkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *NetworkInfo) Reset() {
	*x = NetworkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfo) ProtoMessage() {}

func (x *NetworkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfo.ProtoReflect.Descriptor instead.
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkInfo) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type CNIerror struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CNIerror) Reset() {
	*x = CNIerror{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIerror) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIerror) ProtoMessage() {}

func (x *CNIerror) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIerror.ProtoReflect.Descriptor instead.
func (*CNIerror) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{3}
}

func (x *CNIerror) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CNIcapArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortMappings []*CNIcapArgs_PORTMAPPINGS `protobuf:"bytes,1,rep,name=portMappings,proto3" json:"portMappings,omitempty"`
	FooMap       []*CNIcapArgs_FOOMAP       `protobuf:"bytes,2,rep,name=fooMap,proto3" json:"fooMap,omitempty"`
}

func (x *CNIcapArgs) Reset() {
	*x = CNIcapArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIcapArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIcapArgs) ProtoMessage() {}

func (x *CNIcapArgs) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIcapArgs.ProtoReflect.Descriptor instead.
func (*CNIcapArgs) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{4}
}

func (x *CNIcapArgs) GetPortMappings() []*CNIcapArgs_PORTMAPPINGS {
	if x != nil {
		return x.PortMappings
	}
	return nil
}

func (x *CNIcapArgs) GetFooMap() []*CNIcapArgs_FOOMAP {
	if x != nil {
		return x.FooMap
	}
	return nil
}

type ConfPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetDir  string `protobuf:"bytes,1,opt,name=netDir,proto3" json:"netDir,omitempty"`
	NetConf string `protobuf:"bytes,2,opt,name=netConf,proto3" json:"netConf,omitempty"`
}

func (x *ConfPath) Reset() {
	*x = ConfPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfPath) ProtoMessage() {}

func (x *ConfPath) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfPath.ProtoReflect.Descriptor instead.
func (*ConfPath) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{5}
}

func (x *ConfPath) GetNetDir() string {
	if x != nil {
		return x.NetDir
	}
	return ""
}

func (x *ConfPath) GetNetConf() string {
	if x != nil {
		return x.NetConf
	}
	return ""
}

type CNIaddMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf        string      `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	ContainerID string      `protobuf:"bytes,2,opt,name=containerID,proto3" json:"containerID,omitempty"`
	NetNS       string      `protobuf:"bytes,3,opt,name=netNS,proto3" json:"netNS,omitempty"`
	IfName      string      `protobuf:"bytes,4,opt,name=ifName,proto3" json:"ifName,omitempty"`
	CniArgs     string      `protobuf:"bytes,5,opt,name=cniArgs,proto3" json:"cniArgs,omitempty"`
	CapArgs     *CNIcapArgs `protobuf:"bytes,6,opt,name=capArgs,proto3" json:"capArgs,omitempty"` //string capArgs = 6;
}

func (x *CNIaddMsg) Reset() {
	*x = CNIaddMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIaddMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIaddMsg) ProtoMessage() {}

func (x *CNIaddMsg) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIaddMsg.ProtoReflect.Descriptor instead.
func (*CNIaddMsg) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{6}
}

func (x *CNIaddMsg) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

func (x *CNIaddMsg) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *CNIaddMsg) GetNetNS() string {
	if x != nil {
		return x.NetNS
	}
	return ""
}

func (x *CNIaddMsg) GetIfName() string {
	if x != nil {
		return x.IfName
	}
	return ""
}

func (x *CNIaddMsg) GetCniArgs() string {
	if x != nil {
		return x.CniArgs
	}
	return ""
}

func (x *CNIaddMsg) GetCapArgs() *CNIcapArgs {
	if x != nil {
		return x.CapArgs
	}
	return nil
}

type ADDresult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	StdOut string `protobuf:"bytes,2,opt,name=stdOut,proto3" json:"stdOut,omitempty"`
}

func (x *ADDresult) Reset() {
	*x = ADDresult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADDresult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADDresult) ProtoMessage() {}

func (x *ADDresult) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADDresult.ProtoReflect.Descriptor instead.
func (*ADDresult) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{7}
}

func (x *ADDresult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ADDresult) GetStdOut() string {
	if x != nil {
		return x.StdOut
	}
	return ""
}

type CNIcheckMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf        string      `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	ContainerID string      `protobuf:"bytes,2,opt,name=containerID,proto3" json:"containerID,omitempty"`
	NetNS       string      `protobuf:"bytes,3,opt,name=netNS,proto3" json:"netNS,omitempty"`
	IfName      string      `protobuf:"bytes,4,opt,name=ifName,proto3" json:"ifName,omitempty"`
	CniArgs     string      `protobuf:"bytes,5,opt,name=cniArgs,proto3" json:"cniArgs,omitempty"`
	CapArgs     *CNIcapArgs `protobuf:"bytes,6,opt,name=capArgs,proto3" json:"capArgs,omitempty"` //string capArgs = 6;
}

func (x *CNIcheckMsg) Reset() {
	*x = CNIcheckMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIcheckMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIcheckMsg) ProtoMessage() {}

func (x *CNIcheckMsg) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIcheckMsg.ProtoReflect.Descriptor instead.
func (*CNIcheckMsg) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{8}
}

func (x *CNIcheckMsg) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

func (x *CNIcheckMsg) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *CNIcheckMsg) GetNetNS() string {
	if x != nil {
		return x.NetNS
	}
	return ""
}

func (x *CNIcheckMsg) GetIfName() string {
	if x != nil {
		return x.IfName
	}
	return ""
}

func (x *CNIcheckMsg) GetCniArgs() string {
	if x != nil {
		return x.CniArgs
	}
	return ""
}

func (x *CNIcheckMsg) GetCapArgs() *CNIcapArgs {
	if x != nil {
		return x.CapArgs
	}
	return nil
}

type CHECKresult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CHECKresult) Reset() {
	*x = CHECKresult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CHECKresult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CHECKresult) ProtoMessage() {}

func (x *CHECKresult) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CHECKresult.ProtoReflect.Descriptor instead.
func (*CHECKresult) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{9}
}

func (x *CHECKresult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CNIdelMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conf        string      `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	ContainerID string      `protobuf:"bytes,2,opt,name=containerID,proto3" json:"containerID,omitempty"`
	NetNS       string      `protobuf:"bytes,3,opt,name=netNS,proto3" json:"netNS,omitempty"`
	IfName      string      `protobuf:"bytes,4,opt,name=ifName,proto3" json:"ifName,omitempty"`
	CniArgs     string      `protobuf:"bytes,5,opt,name=cniArgs,proto3" json:"cniArgs,omitempty"`
	CapArgs     *CNIcapArgs `protobuf:"bytes,6,opt,name=capArgs,proto3" json:"capArgs,omitempty"` //string capArgs = 6;
}

func (x *CNIdelMsg) Reset() {
	*x = CNIdelMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIdelMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIdelMsg) ProtoMessage() {}

func (x *CNIdelMsg) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIdelMsg.ProtoReflect.Descriptor instead.
func (*CNIdelMsg) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{10}
}

func (x *CNIdelMsg) GetConf() string {
	if x != nil {
		return x.Conf
	}
	return ""
}

func (x *CNIdelMsg) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *CNIdelMsg) GetNetNS() string {
	if x != nil {
		return x.NetNS
	}
	return ""
}

func (x *CNIdelMsg) GetIfName() string {
	if x != nil {
		return x.IfName
	}
	return ""
}

func (x *CNIdelMsg) GetCniArgs() string {
	if x != nil {
		return x.CniArgs
	}
	return ""
}

func (x *CNIdelMsg) GetCapArgs() *CNIcapArgs {
	if x != nil {
		return x.CapArgs
	}
	return nil
}

type DELresult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	StdOut string `protobuf:"bytes,2,opt,name=stdOut,proto3" json:"stdOut,omitempty"`
}

func (x *DELresult) Reset() {
	*x = DELresult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DELresult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DELresult) ProtoMessage() {}

func (x *DELresult) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DELresult.ProtoReflect.Descriptor instead.
func (*DELresult) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{11}
}

func (x *DELresult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DELresult) GetStdOut() string {
	if x != nil {
		return x.StdOut
	}
	return ""
}

type CNIcapArgs_PORTMAPPINGS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostPort      float64 `protobuf:"fixed64,1,opt,name=hostPort,proto3" json:"hostPort,omitempty"`
	ContainerPort float64 `protobuf:"fixed64,2,opt,name=containerPort,proto3" json:"containerPort,omitempty"`
	Protocol      string  `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *CNIcapArgs_PORTMAPPINGS) Reset() {
	*x = CNIcapArgs_PORTMAPPINGS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIcapArgs_PORTMAPPINGS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIcapArgs_PORTMAPPINGS) ProtoMessage() {}

func (x *CNIcapArgs_PORTMAPPINGS) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIcapArgs_PORTMAPPINGS.ProtoReflect.Descriptor instead.
func (*CNIcapArgs_PORTMAPPINGS) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CNIcapArgs_PORTMAPPINGS) GetHostPort() float64 {
	if x != nil {
		return x.HostPort
	}
	return 0
}

func (x *CNIcapArgs_PORTMAPPINGS) GetContainerPort() float64 {
	if x != nil {
		return x.ContainerPort
	}
	return 0
}

func (x *CNIcapArgs_PORTMAPPINGS) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type CNIcapArgs_FOOMAP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thing1 string `protobuf:"bytes,1,opt,name=thing1,proto3" json:"thing1,omitempty"`
	Thing2 string `protobuf:"bytes,2,opt,name=thing2,proto3" json:"thing2,omitempty"`
}

func (x *CNIcapArgs_FOOMAP) Reset() {
	*x = CNIcapArgs_FOOMAP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cnigrpc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CNIcapArgs_FOOMAP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CNIcapArgs_FOOMAP) ProtoMessage() {}

func (x *CNIcapArgs_FOOMAP) ProtoReflect() protoreflect.Message {
	mi := &file_cnigrpc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CNIcapArgs_FOOMAP.ProtoReflect.Descriptor instead.
func (*CNIcapArgs_FOOMAP) Descriptor() ([]byte, []int) {
	return file_cnigrpc_proto_rawDescGZIP(), []int{4, 1}
}

func (x *CNIcapArgs_FOOMAP) GetThing1() string {
	if x != nil {
		return x.Thing1
	}
	return ""
}

func (x *CNIcapArgs_FOOMAP) GetThing2() string {
	if x != nil {
		return x.Thing2
	}
	return ""
}

var File_cnigrpc_proto protoreflect.FileDescriptor

var file_cnigrpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6e, 0x69, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2f, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x4e, 0x49,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xac, 0x02, 0x0a, 0x0a,
	0x43, 0x4e, 0x49, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x70, 0x6f,
	0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x63, 0x61, 0x70,
	0x41, 0x72, 0x67, 0x73, 0x2e, 0x50, 0x4f, 0x52, 0x54, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47,
	0x53, 0x52, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x31, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x4d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x63, 0x61, 0x70, 0x41,
	0x72, 0x67, 0x73, 0x2e, 0x46, 0x4f, 0x4f, 0x4d, 0x41, 0x50, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x4d,
	0x61, 0x70, 0x1a, 0x6c, 0x0a, 0x0c, 0x50, 0x4f, 0x52, 0x54, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x1a, 0x38, 0x0a, 0x06, 0x46, 0x4f, 0x4f, 0x4d, 0x41, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x32, 0x22, 0x3c, 0x0a, 0x08, 0x43, 0x6f,
	0x6e, 0x66, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x74, 0x44, 0x69, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x74, 0x44, 0x69, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x43, 0x4e, 0x49,
	0x61, 0x64, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x65, 0x74, 0x4e, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74,
	0x4e, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6e,
	0x69, 0x41, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6e, 0x69,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43,
	0x4e, 0x49, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x52, 0x07, 0x63, 0x61, 0x70, 0x41, 0x72,
	0x67, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x41, 0x44, 0x44, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x4f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x4f, 0x75, 0x74, 0x22, 0xb9, 0x01,
	0x0a, 0x0b, 0x43, 0x4e, 0x49, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6e,
	0x66, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x4e, 0x53, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x4e, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x66, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x66, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6e, 0x69, 0x41, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6e, 0x69, 0x41, 0x72, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x63,
	0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c,
	0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73,
	0x52, 0x07, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x22, 0x23, 0x0a, 0x0b, 0x43, 0x48, 0x45,
	0x43, 0x4b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xb7,
	0x01, 0x0a, 0x09, 0x43, 0x4e, 0x49, 0x64, 0x65, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x4e, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x4e, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x66, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x66, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6e, 0x69, 0x41, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6e, 0x69, 0x41, 0x72, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x61,
	0x70, 0x41, 0x72, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x07, 0x63, 0x61, 0x70, 0x41, 0x72, 0x67, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x64, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64,
	0x4f, 0x75, 0x74, 0x32, 0x51, 0x0a, 0x0a, 0x43, 0x4e, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c,
	0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xda, 0x01, 0x0a, 0x09, 0x43, 0x4e, 0x49, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x09, 0x43, 0x4e, 0x49, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x50,
	0x61, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x4e, 0x49, 0x61, 0x64,
	0x64, 0x12, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x61, 0x64,
	0x64, 0x4d, 0x73, 0x67, 0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x41, 0x44,
	0x44, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x4e, 0x49,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43,
	0x4e, 0x49, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x62,
	0x63, 0x6e, 0x69, 0x2e, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x4e, 0x49, 0x64, 0x65, 0x6c, 0x12, 0x11, 0x2e, 0x6c, 0x69,
	0x62, 0x63, 0x6e, 0x69, 0x2e, 0x43, 0x4e, 0x49, 0x64, 0x65, 0x6c, 0x4d, 0x73, 0x67, 0x1a, 0x11,
	0x2e, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69, 0x2e, 0x44, 0x45, 0x4c, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6e, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x63, 0x6e, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cnigrpc_proto_rawDescOnce sync.Once
	file_cnigrpc_proto_rawDescData = file_cnigrpc_proto_rawDesc
)

func file_cnigrpc_proto_rawDescGZIP() []byte {
	file_cnigrpc_proto_rawDescOnce.Do(func() {
		file_cnigrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cnigrpc_proto_rawDescData)
	})
	return file_cnigrpc_proto_rawDescData
}

var file_cnigrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_cnigrpc_proto_goTypes = []interface{}{
	(*ListNetworkRequest)(nil),      // 0: libcni.ListNetworkRequest
	(*ListNetworkReply)(nil),        // 1: libcni.ListNetworkReply
	(*NetworkInfo)(nil),             // 2: libcni.NetworkInfo
	(*CNIerror)(nil),                // 3: libcni.CNIerror
	(*CNIcapArgs)(nil),              // 4: libcni.CNIcapArgs
	(*ConfPath)(nil),                // 5: libcni.ConfPath
	(*CNIaddMsg)(nil),               // 6: libcni.CNIaddMsg
	(*ADDresult)(nil),               // 7: libcni.ADDresult
	(*CNIcheckMsg)(nil),             // 8: libcni.CNIcheckMsg
	(*CHECKresult)(nil),             // 9: libcni.CHECKresult
	(*CNIdelMsg)(nil),               // 10: libcni.CNIdelMsg
	(*DELresult)(nil),               // 11: libcni.DELresult
	(*CNIcapArgs_PORTMAPPINGS)(nil), // 12: libcni.CNIcapArgs.PORTMAPPINGS
	(*CNIcapArgs_FOOMAP)(nil),       // 13: libcni.CNIcapArgs.FOOMAP
}
var file_cnigrpc_proto_depIdxs = []int32{
	2,  // 0: libcni.ListNetworkReply.networks:type_name -> libcni.NetworkInfo
	12, // 1: libcni.CNIcapArgs.portMappings:type_name -> libcni.CNIcapArgs.PORTMAPPINGS
	13, // 2: libcni.CNIcapArgs.fooMap:type_name -> libcni.CNIcapArgs.FOOMAP
	4,  // 3: libcni.CNIaddMsg.capArgs:type_name -> libcni.CNIcapArgs
	4,  // 4: libcni.CNIcheckMsg.capArgs:type_name -> libcni.CNIcapArgs
	4,  // 5: libcni.CNIdelMsg.capArgs:type_name -> libcni.CNIcapArgs
	0,  // 6: libcni.CNIService.ListNetwork:input_type -> libcni.ListNetworkRequest
	5,  // 7: libcni.CNIserver.CNIconfig:input_type -> libcni.ConfPath
	6,  // 8: libcni.CNIserver.CNIadd:input_type -> libcni.CNIaddMsg
	8,  // 9: libcni.CNIserver.CNIcheck:input_type -> libcni.CNIcheckMsg
	10, // 10: libcni.CNIserver.CNIdel:input_type -> libcni.CNIdelMsg
	1,  // 11: libcni.CNIService.ListNetwork:output_type -> libcni.ListNetworkReply
	3,  // 12: libcni.CNIserver.CNIconfig:output_type -> libcni.CNIerror
	7,  // 13: libcni.CNIserver.CNIadd:output_type -> libcni.ADDresult
	9,  // 14: libcni.CNIserver.CNIcheck:output_type -> libcni.CHECKresult
	11, // 15: libcni.CNIserver.CNIdel:output_type -> libcni.DELresult
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cnigrpc_proto_init() }
func file_cnigrpc_proto_init() {
	if File_cnigrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cnigrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworkRequest); i {
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
		file_cnigrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworkReply); i {
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
		file_cnigrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfo); i {
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
		file_cnigrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIerror); i {
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
		file_cnigrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIcapArgs); i {
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
		file_cnigrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfPath); i {
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
		file_cnigrpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIaddMsg); i {
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
		file_cnigrpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADDresult); i {
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
		file_cnigrpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIcheckMsg); i {
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
		file_cnigrpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CHECKresult); i {
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
		file_cnigrpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIdelMsg); i {
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
		file_cnigrpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DELresult); i {
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
		file_cnigrpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIcapArgs_PORTMAPPINGS); i {
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
		file_cnigrpc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CNIcapArgs_FOOMAP); i {
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
			RawDescriptor: file_cnigrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cnigrpc_proto_goTypes,
		DependencyIndexes: file_cnigrpc_proto_depIdxs,
		MessageInfos:      file_cnigrpc_proto_msgTypes,
	}.Build()
	File_cnigrpc_proto = out.File
	file_cnigrpc_proto_rawDesc = nil
	file_cnigrpc_proto_goTypes = nil
	file_cnigrpc_proto_depIdxs = nil
}
