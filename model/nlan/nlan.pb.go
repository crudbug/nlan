// Code generated by protoc-gen-go.
// source: nlan.proto
// DO NOT EDIT!

/*
Package nlan is a generated protocol buffer package.

It is generated from these files:
	nlan.proto

It has these top-level messages:
	State
	Model
	Dvr
	Subnets
	IpDvr
	Vxlan
	Network
	L2Vpn
	Links
	Nodes
	Router
	Attrs
	Neighbor
	Ospf
	Vhosts
	VhostProps
	Interface
*/
package nlan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type State struct {
	Router map[string]*Model `protobuf:"bytes,1,rep,name=Router" json:"Router,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *State) GetRouter() map[string]*Model {
	if m != nil {
		return m.Router
	}
	return nil
}

type Model struct {
	Dvr        *Dvr                  `protobuf:"bytes,1,opt,name=Dvr" json:"Dvr,omitempty"`
	Ptn        map[string]*Network   `protobuf:"bytes,2,rep,name=Ptn" json:"Ptn,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Router     *Router               `protobuf:"bytes,3,opt,name=Router" json:"Router,omitempty"`
	Vhosts     *Vhosts               `protobuf:"bytes,4,opt,name=Vhosts" json:"Vhosts,omitempty"`
	Interfaces map[string]*Interface `protobuf:"bytes,5,rep,name=Interfaces" json:"Interfaces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Model) GetDvr() *Dvr {
	if m != nil {
		return m.Dvr
	}
	return nil
}

func (m *Model) GetPtn() map[string]*Network {
	if m != nil {
		return m.Ptn
	}
	return nil
}

func (m *Model) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

func (m *Model) GetVhosts() *Vhosts {
	if m != nil {
		return m.Vhosts
	}
	return nil
}

func (m *Model) GetInterfaces() map[string]*Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type Dvr struct {
	OvsBridges bool       `protobuf:"varint,1,opt,name=OvsBridges" json:"OvsBridges,omitempty"`
	Subnets    []*Subnets `protobuf:"bytes,2,rep,name=Subnets" json:"Subnets,omitempty"`
	Vxlan      []*Vxlan   `protobuf:"bytes,3,rep,name=Vxlan" json:"Vxlan,omitempty"`
}

func (m *Dvr) Reset()                    { *m = Dvr{} }
func (m *Dvr) String() string            { return proto.CompactTextString(m) }
func (*Dvr) ProtoMessage()               {}
func (*Dvr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Dvr) GetSubnets() []*Subnets {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *Dvr) GetVxlan() []*Vxlan {
	if m != nil {
		return m.Vxlan
	}
	return nil
}

type Subnets struct {
	IpDvr []*IpDvr `protobuf:"bytes,1,rep,name=IpDvr" json:"IpDvr,omitempty"`
	Peers []string `protobuf:"bytes,2,rep,name=Peers" json:"Peers,omitempty"`
	Ports []string `protobuf:"bytes,3,rep,name=Ports" json:"Ports,omitempty"`
	Vid   uint32   `protobuf:"varint,4,opt,name=Vid" json:"Vid,omitempty"`
	Vni   uint32   `protobuf:"varint,5,opt,name=Vni" json:"Vni,omitempty"`
}

func (m *Subnets) Reset()                    { *m = Subnets{} }
func (m *Subnets) String() string            { return proto.CompactTextString(m) }
func (*Subnets) ProtoMessage()               {}
func (*Subnets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Subnets) GetIpDvr() []*IpDvr {
	if m != nil {
		return m.IpDvr
	}
	return nil
}

type IpDvr struct {
	Addr string `protobuf:"bytes,1,opt,name=Addr" json:"Addr,omitempty"`
	Dhcp string `protobuf:"bytes,2,opt,name=Dhcp" json:"Dhcp,omitempty"`
	Mode string `protobuf:"bytes,3,opt,name=Mode" json:"Mode,omitempty"`
}

func (m *IpDvr) Reset()                    { *m = IpDvr{} }
func (m *IpDvr) String() string            { return proto.CompactTextString(m) }
func (*IpDvr) ProtoMessage()               {}
func (*IpDvr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Vxlan struct {
	LocalIp   string   `protobuf:"bytes,1,opt,name=LocalIp" json:"LocalIp,omitempty"`
	RemoteIps []string `protobuf:"bytes,2,rep,name=RemoteIps" json:"RemoteIps,omitempty"`
}

func (m *Vxlan) Reset()                    { *m = Vxlan{} }
func (m *Vxlan) String() string            { return proto.CompactTextString(m) }
func (*Vxlan) ProtoMessage()               {}
func (*Vxlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Network struct {
	L2Vpn []*L2Vpn `protobuf:"bytes,2,rep,name=L2Vpn" json:"L2Vpn,omitempty"`
	Links *Links   `protobuf:"bytes,3,opt,name=Links" json:"Links,omitempty"`
	Nodes *Nodes   `protobuf:"bytes,4,opt,name=Nodes" json:"Nodes,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Network) GetL2Vpn() []*L2Vpn {
	if m != nil {
		return m.L2Vpn
	}
	return nil
}

func (m *Network) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Network) GetNodes() *Nodes {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type L2Vpn struct {
	Ip    string   `protobuf:"bytes,1,opt,name=Ip" json:"Ip,omitempty"`
	Peers []string `protobuf:"bytes,2,rep,name=Peers" json:"Peers,omitempty"`
	Vid   uint32   `protobuf:"varint,3,opt,name=Vid" json:"Vid,omitempty"`
	Vni   uint32   `protobuf:"varint,4,opt,name=Vni" json:"Vni,omitempty"`
}

func (m *L2Vpn) Reset()                    { *m = L2Vpn{} }
func (m *L2Vpn) String() string            { return proto.CompactTextString(m) }
func (*L2Vpn) ProtoMessage()               {}
func (*L2Vpn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Links struct {
	LocalIp   string   `protobuf:"bytes,1,opt,name=LocalIp" json:"LocalIp,omitempty"`
	RemoteIps []string `protobuf:"bytes,2,rep,name=RemoteIps" json:"RemoteIps,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Nodes struct {
	L2Sw string `protobuf:"bytes,1,opt,name=L2Sw" json:"L2Sw,omitempty"`
	Ptn  string `protobuf:"bytes,2,opt,name=Ptn" json:"Ptn,omitempty"`
}

func (m *Nodes) Reset()                    { *m = Nodes{} }
func (m *Nodes) String() string            { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()               {}
func (*Nodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Router struct {
	Bgp         map[string]*Attrs `protobuf:"bytes,1,rep,name=Bgp" json:"Bgp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EmbeddedBgp bool              `protobuf:"varint,2,opt,name=EmbeddedBgp" json:"EmbeddedBgp,omitempty"`
	Loopback    string            `protobuf:"bytes,3,opt,name=Loopback" json:"Loopback,omitempty"`
	Ospf        []*Ospf           `protobuf:"bytes,4,rep,name=Ospf" json:"Ospf,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Router) GetBgp() map[string]*Attrs {
	if m != nil {
		return m.Bgp
	}
	return nil
}

func (m *Router) GetOspf() []*Ospf {
	if m != nil {
		return m.Ospf
	}
	return nil
}

type Attrs struct {
	Neighbors []*Neighbor `protobuf:"bytes,1,rep,name=Neighbors" json:"Neighbors,omitempty"`
}

func (m *Attrs) Reset()                    { *m = Attrs{} }
func (m *Attrs) String() string            { return proto.CompactTextString(m) }
func (*Attrs) ProtoMessage()               {}
func (*Attrs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Attrs) GetNeighbors() []*Neighbor {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

type Neighbor struct {
	NextHopSelf          bool   `protobuf:"varint,1,opt,name=NextHopSelf" json:"NextHopSelf,omitempty"`
	Peer                 string `protobuf:"bytes,2,opt,name=Peer" json:"Peer,omitempty"`
	RemoteAs             uint32 `protobuf:"varint,3,opt,name=RemoteAs" json:"RemoteAs,omitempty"`
	RouteReflectorClient bool   `protobuf:"varint,4,opt,name=RouteReflectorClient" json:"RouteReflectorClient,omitempty"`
}

func (m *Neighbor) Reset()                    { *m = Neighbor{} }
func (m *Neighbor) String() string            { return proto.CompactTextString(m) }
func (*Neighbor) ProtoMessage()               {}
func (*Neighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type Ospf struct {
	Area     string   `protobuf:"bytes,1,opt,name=Area" json:"Area,omitempty"`
	Networks []string `protobuf:"bytes,2,rep,name=Networks" json:"Networks,omitempty"`
}

func (m *Ospf) Reset()                    { *m = Ospf{} }
func (m *Ospf) String() string            { return proto.CompactTextString(m) }
func (*Ospf) ProtoMessage()               {}
func (*Ospf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type Vhosts struct {
	VhostProps []*VhostProps `protobuf:"bytes,1,rep,name=VhostProps" json:"VhostProps,omitempty"`
}

func (m *Vhosts) Reset()                    { *m = Vhosts{} }
func (m *Vhosts) String() string            { return proto.CompactTextString(m) }
func (*Vhosts) ProtoMessage()               {}
func (*Vhosts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Vhosts) GetVhostProps() []*VhostProps {
	if m != nil {
		return m.VhostProps
	}
	return nil
}

type VhostProps struct {
	Network string `protobuf:"bytes,1,opt,name=Network" json:"Network,omitempty"`
	Vhosts  uint32 `protobuf:"varint,2,opt,name=Vhosts" json:"Vhosts,omitempty"`
}

func (m *VhostProps) Reset()                    { *m = VhostProps{} }
func (m *VhostProps) String() string            { return proto.CompactTextString(m) }
func (*VhostProps) ProtoMessage()               {}
func (*VhostProps) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type Interface struct {
	Mode    string `protobuf:"bytes,1,opt,name=mode" json:"mode,omitempty"`
	Local   string `protobuf:"bytes,2,opt,name=local" json:"local,omitempty"`
	Remote  string `protobuf:"bytes,3,opt,name=remote" json:"remote,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
}

func (m *Interface) Reset()                    { *m = Interface{} }
func (m *Interface) String() string            { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()               {}
func (*Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func init() {
	proto.RegisterType((*State)(nil), "nlan.State")
	proto.RegisterType((*Model)(nil), "nlan.Model")
	proto.RegisterType((*Dvr)(nil), "nlan.Dvr")
	proto.RegisterType((*Subnets)(nil), "nlan.Subnets")
	proto.RegisterType((*IpDvr)(nil), "nlan.IpDvr")
	proto.RegisterType((*Vxlan)(nil), "nlan.Vxlan")
	proto.RegisterType((*Network)(nil), "nlan.Network")
	proto.RegisterType((*L2Vpn)(nil), "nlan.L2Vpn")
	proto.RegisterType((*Links)(nil), "nlan.Links")
	proto.RegisterType((*Nodes)(nil), "nlan.Nodes")
	proto.RegisterType((*Router)(nil), "nlan.Router")
	proto.RegisterType((*Attrs)(nil), "nlan.Attrs")
	proto.RegisterType((*Neighbor)(nil), "nlan.Neighbor")
	proto.RegisterType((*Ospf)(nil), "nlan.Ospf")
	proto.RegisterType((*Vhosts)(nil), "nlan.Vhosts")
	proto.RegisterType((*VhostProps)(nil), "nlan.VhostProps")
	proto.RegisterType((*Interface)(nil), "nlan.Interface")
}

var fileDescriptor0 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x56, 0xe2, 0xa4, 0x8d, 0xc7, 0x7d, 0xb1, 0x14, 0x88, 0x42, 0x85, 0xc0, 0x42, 0x50, 0x81,
	0x08, 0x22, 0x1c, 0x80, 0x8a, 0x4b, 0x4b, 0x2b, 0x11, 0x29, 0xa4, 0x55, 0x23, 0x72, 0xe3, 0xe0,
	0xc4, 0x9b, 0x36, 0x8a, 0xeb, 0xb5, 0xd6, 0xdb, 0xd7, 0x1f, 0xe2, 0xc8, 0x6f, 0x64, 0x76, 0x76,
	0xec, 0x3a, 0x3d, 0xf5, 0xe6, 0x79, 0xcf, 0xf7, 0xcd, 0xb7, 0x06, 0x48, 0x93, 0x28, 0xed, 0x66,
	0x5a, 0x19, 0x25, 0x1a, 0xf6, 0x3b, 0x4c, 0xa0, 0x39, 0x32, 0x91, 0x91, 0xe2, 0x2d, 0xac, 0x9c,
	0xaa, 0x4b, 0x23, 0x75, 0xbb, 0xf6, 0xd2, 0xdb, 0x0d, 0x7a, 0xcf, 0xba, 0x94, 0x4b, 0xc1, 0xae,
	0x8b, 0x1c, 0xa5, 0x46, 0xdf, 0x76, 0xf6, 0x20, 0xa8, 0x98, 0x22, 0x00, 0x6f, 0x21, 0x6f, 0xb1,
	0xa8, 0xb6, 0xeb, 0x8b, 0x0e, 0x34, 0xaf, 0xa2, 0xe4, 0x52, 0xb6, 0xeb, 0x68, 0x06, 0xbd, 0xc0,
	0xf5, 0xf8, 0xa5, 0x62, 0x99, 0xec, 0xd5, 0xbf, 0xd6, 0xc2, 0x7f, 0x75, 0x68, 0x92, 0x25, 0x9e,
	0x82, 0x77, 0x78, 0xa5, 0xa9, 0x2c, 0xe8, 0xf9, 0x2e, 0x0f, 0x1d, 0xe2, 0x15, 0x78, 0x27, 0x26,
	0xc5, 0x7a, 0xbb, 0xc3, 0x76, 0xa5, 0xbe, 0x8b, 0x6e, 0x37, 0x71, 0xa7, 0xdc, 0xd4, 0xa3, 0xea,
	0x35, 0x97, 0xe5, 0x7c, 0x36, 0x3a, 0x3e, 0x57, 0xb9, 0xc9, 0xdb, 0x8d, 0x6a, 0xd4, 0xf9, 0xc4,
	0x47, 0x80, 0x7e, 0x8a, 0x69, 0xb3, 0x68, 0x2a, 0xf3, 0x76, 0x93, 0xa6, 0x3c, 0xaf, 0x4e, 0xb9,
	0x8b, 0x3a, 0xb4, 0xdf, 0xa0, 0x55, 0x0e, 0x5e, 0x82, 0xba, 0xb3, 0x0c, 0x75, 0xdd, 0x35, 0x19,
	0x4a, 0x73, 0xad, 0xf4, 0xc2, 0x82, 0xed, 0x1c, 0xc0, 0xe6, 0xbd, 0x6e, 0xcb, 0x1d, 0x5e, 0x2c,
	0x77, 0xd8, 0x74, 0x1d, 0xca, 0x12, 0x22, 0xec, 0x37, 0xd1, 0x24, 0x04, 0xc0, 0xf1, 0x55, 0x7e,
	0xa0, 0xe7, 0xf1, 0x19, 0xae, 0x6d, 0xcb, 0x5b, 0x58, 0xbe, 0x3a, 0xba, 0x9c, 0xa4, 0x12, 0x91,
	0x3a, 0xb6, 0x78, 0x05, 0x76, 0xda, 0x5b, 0x8c, 0x6f, 0xd0, 0x81, 0x2c, 0x79, 0x77, 0xb7, 0x20,
	0x57, 0xf8, 0xa7, 0xac, 0xb5, 0x69, 0xfd, 0xcc, 0x9d, 0xa2, 0x92, 0x46, 0x2e, 0xb1, 0x0e, 0xcd,
	0x13, 0x29, 0xb5, 0x1b, 0xe0, 0x93, 0xa9, 0x34, 0xce, 0xf3, 0xc8, 0x44, 0x30, 0xe3, 0x79, 0x4c,
	0x34, 0xaf, 0x93, 0x91, 0xce, 0x91, 0x51, 0x34, 0xc2, 0x4f, 0xdc, 0x53, 0xac, 0x41, 0x63, 0x3f,
	0x8e, 0x35, 0x03, 0x46, 0xeb, 0xf0, 0x7c, 0x9a, 0x11, 0x5e, 0xb2, 0x2c, 0xe5, 0x74, 0x44, 0x3f,
	0x7c, 0xcf, 0xdb, 0x8a, 0x4d, 0x58, 0x1d, 0xa8, 0x69, 0x94, 0xf4, 0x33, 0xae, 0x7a, 0x04, 0xfe,
	0xa9, 0xbc, 0x50, 0x46, 0xf6, 0x33, 0x5e, 0xc4, 0xae, 0xcf, 0x44, 0xdb, 0xf5, 0x07, 0xbd, 0x71,
	0x56, 0x28, 0x86, 0xd7, 0x27, 0x17, 0xc5, 0xe6, 0xe9, 0x22, 0x67, 0x9d, 0x14, 0x31, 0xeb, 0xb2,
	0xb1, 0x21, 0x4e, 0x2f, 0x54, 0xc2, 0x31, 0x72, 0x85, 0xdf, 0xb9, 0x27, 0xb2, 0x5e, 0x2f, 0xd7,
	0xb8, 0xc7, 0x05, 0x83, 0xf7, 0xaa, 0xe0, 0x89, 0x09, 0x8b, 0xc4, 0x8d, 0x78, 0x08, 0x92, 0x90,
	0xd7, 0xb0, 0x6c, 0x0c, 0x7a, 0xa3, 0x6b, 0xce, 0x0c, 0x8a, 0x57, 0x60, 0xa9, 0xf9, 0x5b, 0x2b,
	0x04, 0x2f, 0x42, 0xf0, 0x0e, 0xce, 0x32, 0x3e, 0xd5, 0x93, 0xaa, 0xee, 0xbb, 0xe8, 0x77, 0x1a,
	0x7b, 0x0c, 0xc1, 0xd1, 0xc5, 0x44, 0xc6, 0xb1, 0x8c, 0x6d, 0x6e, 0x9d, 0xc4, 0xb2, 0x05, 0xad,
	0x81, 0x52, 0xd9, 0x24, 0x9a, 0x2e, 0x1c, 0xe1, 0xa2, 0x0d, 0x8d, 0xe3, 0x3c, 0x9b, 0xe1, 0xd2,
	0xb6, 0x17, 0xb8, 0x5e, 0xd6, 0xd3, 0xf9, 0x02, 0xad, 0xb2, 0xd9, 0x03, 0x5e, 0xf7, 0xbe, 0x31,
	0x3a, 0x27, 0xb1, 0xbe, 0x83, 0x26, 0x19, 0xf8, 0x88, 0xfd, 0xa1, 0x9c, 0x9f, 0x9d, 0x4f, 0x94,
	0xce, 0x79, 0xd9, 0x8d, 0xe2, 0x7d, 0x38, 0x37, 0x9e, 0xb0, 0x55, 0x7c, 0xdb, 0x8d, 0x87, 0xf2,
	0xc6, 0xfc, 0x54, 0xd9, 0x48, 0x26, 0x33, 0x96, 0x37, 0x12, 0x62, 0xf9, 0x66, 0xb1, 0xe0, 0xfe,
	0x8e, 0xba, 0xfd, 0x9c, 0x39, 0xdf, 0x81, 0x6d, 0x42, 0x7e, 0x2a, 0x67, 0x89, 0x9c, 0x1a, 0xa5,
	0x7f, 0x24, 0x73, 0x99, 0x1a, 0x3a, 0x42, 0x2b, 0x7c, 0xe3, 0xd0, 0x91, 0x00, 0xb5, 0x8c, 0x18,
	0xc0, 0x96, 0x1d, 0x4a, 0xba, 0x29, 0xf8, 0xef, 0x16, 0x7f, 0x0b, 0xf1, 0x1a, 0x80, 0xbe, 0x4e,
	0xb4, 0xca, 0x8a, 0xa5, 0xb7, 0x2a, 0xff, 0x0e, 0xf2, 0x87, 0x1f, 0xaa, 0x59, 0xf6, 0xc2, 0xdc,
	0x8f, 0x07, 0x6c, 0x94, 0x3f, 0x9f, 0x3a, 0x69, 0xa1, 0x0f, 0x7e, 0xf9, 0x9e, 0xed, 0x2e, 0x17,
	0x56, 0xf0, 0xa5, 0x9e, 0x12, 0xab, 0x0e, 0x06, 0x88, 0x95, 0x9a, 0x00, 0xf2, 0x79, 0xb0, 0x75,
	0x84, 0x2f, 0x47, 0xe6, 0x4e, 0xa1, 0xfe, 0x64, 0x85, 0xfe, 0xda, 0x9f, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0xb0, 0xa7, 0xa9, 0xc3, 0x05, 0x00, 0x00,
}
