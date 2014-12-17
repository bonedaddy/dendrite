// Code generated by protoc-gen-go.
// source: pb_defs/chord.proto
// DO NOT EDIT!

/*
Package dendrite is a generated protocol buffer package.

It is generated from these files:
	pb_defs/chord.proto

It has these top-level messages:
	PBProtoVnode
	PBProtoPing
	PBProtoAck
	PBProtoErr
	PBProtoForward
	PBProtoJoin
	PBProtoLeave
	PBProtoListVnodes
	PBProtoListVnodesResp
*/
package dendrite

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// single vnode message
// usually included in other messages as a compact type
type PBProtoVnode struct {
	Id               []byte  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Host             *string `protobuf:"bytes,2,req,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PBProtoVnode) Reset()         { *m = PBProtoVnode{} }
func (m *PBProtoVnode) String() string { return proto.CompactTextString(m) }
func (*PBProtoVnode) ProtoMessage()    {}

func (m *PBProtoVnode) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PBProtoVnode) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

// ping message
type PBProtoPing struct {
	Version          *int64 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PBProtoPing) Reset()         { *m = PBProtoPing{} }
func (m *PBProtoPing) String() string { return proto.CompactTextString(m) }
func (*PBProtoPing) ProtoMessage()    {}

func (m *PBProtoPing) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

// dummy response, with additional ok state
type PBProtoAck struct {
	Version          *int64 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Ok               *bool  `protobuf:"varint,2,req,name=ok" json:"ok,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PBProtoAck) Reset()         { *m = PBProtoAck{} }
func (m *PBProtoAck) String() string { return proto.CompactTextString(m) }
func (*PBProtoAck) ProtoMessage()    {}

func (m *PBProtoAck) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *PBProtoAck) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

// error response
type PBProtoErr struct {
	Error            *string `protobuf:"bytes,2,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PBProtoErr) Reset()         { *m = PBProtoErr{} }
func (m *PBProtoErr) String() string { return proto.CompactTextString(m) }
func (*PBProtoErr) ProtoMessage()    {}

func (m *PBProtoErr) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// response which says request should be made to another host
type PBProtoForward struct {
	Version          *int64        `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Vnode            *PBProtoVnode `protobuf:"bytes,2,req,name=vnode" json:"vnode,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PBProtoForward) Reset()         { *m = PBProtoForward{} }
func (m *PBProtoForward) String() string { return proto.CompactTextString(m) }
func (*PBProtoForward) ProtoMessage()    {}

func (m *PBProtoForward) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *PBProtoForward) GetVnode() *PBProtoVnode {
	if m != nil {
		return m.Vnode
	}
	return nil
}

// request to join the cluster
type PBProtoJoin struct {
	ClusterName      *string       `protobuf:"bytes,1,req,name=clusterName" json:"clusterName,omitempty"`
	Source           *PBProtoVnode `protobuf:"bytes,2,req,name=source" json:"source,omitempty"`
	Target           *PBProtoVnode `protobuf:"bytes,3,req,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PBProtoJoin) Reset()         { *m = PBProtoJoin{} }
func (m *PBProtoJoin) String() string { return proto.CompactTextString(m) }
func (*PBProtoJoin) ProtoMessage()    {}

func (m *PBProtoJoin) GetClusterName() string {
	if m != nil && m.ClusterName != nil {
		return *m.ClusterName
	}
	return ""
}

func (m *PBProtoJoin) GetSource() *PBProtoVnode {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PBProtoJoin) GetTarget() *PBProtoVnode {
	if m != nil {
		return m.Target
	}
	return nil
}

// request to leave the cluster
type PBProtoLeave struct {
	Source           *PBProtoVnode `protobuf:"bytes,1,req,name=source" json:"source,omitempty"`
	Target           *PBProtoVnode `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PBProtoLeave) Reset()         { *m = PBProtoLeave{} }
func (m *PBProtoLeave) String() string { return proto.CompactTextString(m) }
func (*PBProtoLeave) ProtoMessage()    {}

func (m *PBProtoLeave) GetSource() *PBProtoVnode {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PBProtoLeave) GetTarget() *PBProtoVnode {
	if m != nil {
		return m.Target
	}
	return nil
}

// request the list of vnodes from a node
type PBProtoListVnodes struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PBProtoListVnodes) Reset()         { *m = PBProtoListVnodes{} }
func (m *PBProtoListVnodes) String() string { return proto.CompactTextString(m) }
func (*PBProtoListVnodes) ProtoMessage()    {}

// response with list of vnodes
type PBProtoListVnodesResp struct {
	Vnodes           []*PBProtoVnode `protobuf:"bytes,1,rep,name=vnodes" json:"vnodes,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PBProtoListVnodesResp) Reset()         { *m = PBProtoListVnodesResp{} }
func (m *PBProtoListVnodesResp) String() string { return proto.CompactTextString(m) }
func (*PBProtoListVnodesResp) ProtoMessage()    {}

func (m *PBProtoListVnodesResp) GetVnodes() []*PBProtoVnode {
	if m != nil {
		return m.Vnodes
	}
	return nil
}

func init() {
}
