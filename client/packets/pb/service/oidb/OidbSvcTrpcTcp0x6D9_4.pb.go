// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0x6D9_4.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

// Group Send File
type OidbSvcTrpcTcp0X6D9_4 struct {
	Body *OidbSvcTrpcTcp0X6D9_4Body `protobuf:"bytes,5,opt"`
	_    [0]func()
}

type OidbSvcTrpcTcp0X6D9_4Body struct {
	GroupUin uint32                     `protobuf:"varint,1,opt"`
	Type     uint32                     `protobuf:"varint,2,opt"` // 2
	Info     *OidbSvcTrpcTcp0X6D9_4Info `protobuf:"bytes,3,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0X6D9_4Info struct {
	BusiType uint32               `protobuf:"varint,1,opt"` // 102
	FileId   string               `protobuf:"bytes,2,opt"`
	Field3   uint32               `protobuf:"varint,3,opt"` // random
	Field4   proto.Option[string] `protobuf:"bytes,4,opt"`
	Field5   bool                 `protobuf:"varint,5,opt"`
	_        [0]func()
}
