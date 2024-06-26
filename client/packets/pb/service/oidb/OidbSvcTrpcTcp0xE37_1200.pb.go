// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/service/oidb/OidbSvcTrpcTcp0xE37_1200.proto

package oidb

// FileDownload
type OidbSvcTrpcTcp0XE37_1200 struct {
	SubCommand uint32                        `protobuf:"varint,1,opt"` // 1200
	Field2     int32                         `protobuf:"varint,2,opt"` // 1
	Body       *OidbSvcTrpcTcp0XE37_1200Body `protobuf:"bytes,14,opt"`
	Field101   int32                         `protobuf:"varint,101,opt"`  // 3
	Field102   int32                         `protobuf:"varint,102,opt"`  // 103
	Field200   int32                         `protobuf:"varint,200,opt"`  // 1
	Field99999 []byte                        `protobuf:"bytes,99999,opt"` // 0xc0, 0x85, 0x2c, 0x01
}

type OidbSvcTrpcTcp0XE37_1200Body struct {
	ReceiverUid string `protobuf:"bytes,10,opt"`
	FileUuid    string `protobuf:"bytes,20,opt"`
	Type        int32  `protobuf:"varint,30,opt"` // 2
	FileHash    string `protobuf:"bytes,60,opt"`
	T2          int32  `protobuf:"varint,601,opt"` // 0
	_           [0]func()
}

type OidbSvcTrpcTcp0XE37_1200Response struct {
	Command    uint32                                `protobuf:"varint,1,opt"`
	SubCommand uint32                                `protobuf:"varint,2,opt"`
	Body       *OidbSvcTrpcTcp0XE37_1200ResponseBody `protobuf:"bytes,14,opt"`
	Field50    uint32                                `protobuf:"varint,50,opt"`
	_          [0]func()
}

type OidbSvcTrpcTcp0XE37_1200ResponseBody struct {
	Field10  uint32                            `protobuf:"varint,10,opt"`
	State    string                            `protobuf:"bytes,20,opt"`
	Result   *OidbSvcTrpcTcp0XE37_1200Result   `protobuf:"bytes,30,opt"`
	Metadata *OidbSvcTrpcTcp0XE37_1200Metadata `protobuf:"bytes,40,opt"`
	_        [0]func()
}

type OidbSvcTrpcTcp0XE37_1200Result struct {
	Server           string   `protobuf:"bytes,20,opt"`
	Port             uint32   `protobuf:"varint,40,opt"`
	Url              string   `protobuf:"bytes,50,opt"`
	AdditionalServer []string `protobuf:"bytes,60,rep"`
	SsoPort          uint32   `protobuf:"varint,80,opt"`
	SsoUrl           string   `protobuf:"bytes,90,opt"`
	Extra            []byte   `protobuf:"bytes,120,opt"`
}

type OidbSvcTrpcTcp0XE37_1200Metadata struct {
	Uin        uint32 `protobuf:"varint,1,opt"`
	Field2     uint32 `protobuf:"varint,2,opt"`
	Field3     uint32 `protobuf:"varint,3,opt"`
	Size       uint32 `protobuf:"varint,4,opt"`
	Timestamp  uint32 `protobuf:"varint,5,opt"`
	FileUuid   string `protobuf:"bytes,6,opt"`
	FileName   string `protobuf:"bytes,7,opt"`
	Field100   []byte `protobuf:"bytes,100,opt"`
	Field101   []byte `protobuf:"bytes,101,opt"`
	Field110   uint32 `protobuf:"varint,110,opt"`
	Timestamp1 uint32 `protobuf:"varint,130,opt"`
	FileHash   string `protobuf:"bytes,140,opt"`
	Field141   []byte `protobuf:"bytes,141,opt"`
	Field142   []byte `protobuf:"bytes,142,opt"`
}
