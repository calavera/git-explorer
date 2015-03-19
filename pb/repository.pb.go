// Code generated by protoc-gen-go.
// source: repository.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	repository.proto

It has these top-level messages:
	Repository
	Commit
	TreeEntry
	Blob
	Empty
*/
package pb

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type ObjectType int32

const (
	ObjectType_AnyType    ObjectType = 0
	ObjectType_BadType    ObjectType = 1
	ObjectType_CommitType ObjectType = 3
	ObjectType_TreeType   ObjectType = 4
	ObjectType_BlobType   ObjectType = 5
	ObjectType_TagType    ObjectType = 6
)

var ObjectType_name = map[int32]string{
	0: "AnyType",
	1: "BadType",
	3: "CommitType",
	4: "TreeType",
	5: "BlobType",
	6: "TagType",
}
var ObjectType_value = map[string]int32{
	"AnyType":    0,
	"BadType":    1,
	"CommitType": 3,
	"TreeType":   4,
	"BlobType":   5,
	"TagType":    6,
}

func (x ObjectType) String() string {
	return proto.EnumName(ObjectType_name, int32(x))
}

type Repository struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ref  string `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	Base string `protobuf:"bytes,3,opt,name=base" json:"base,omitempty"`
	Url  string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}

type Commit struct {
	Oid     string `protobuf:"bytes,1,opt,name=oid" json:"oid,omitempty"`
	Summary string `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}

type TreeEntry struct {
	Oid  string     `protobuf:"bytes,1,opt,name=oid" json:"oid,omitempty"`
	Name string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Path string     `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Type ObjectType `protobuf:"varint,4,opt,name=type,enum=pb.ObjectType" json:"type,omitempty"`
}

func (m *TreeEntry) Reset()         { *m = TreeEntry{} }
func (m *TreeEntry) String() string { return proto.CompactTextString(m) }
func (*TreeEntry) ProtoMessage()    {}

type Blob struct {
	Oid  string `protobuf:"bytes,1,opt,name=oid" json:"oid,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("pb.ObjectType", ObjectType_name, ObjectType_value)
}

// Client API for RepositoryExplorer service

type RepositoryExplorerClient interface {
	CreateRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Repository, error)
	FetchRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Commit, error)
	ListTreeEntries(ctx context.Context, in *Repository, opts ...grpc.CallOption) (RepositoryExplorer_ListTreeEntriesClient, error)
	GetBlobData(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Blob, error)
	DeleteRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Empty, error)
}

type repositoryExplorerClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryExplorerClient(cc *grpc.ClientConn) RepositoryExplorerClient {
	return &repositoryExplorerClient{cc}
}

func (c *repositoryExplorerClient) CreateRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := grpc.Invoke(ctx, "/pb.RepositoryExplorer/CreateRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryExplorerClient) FetchRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Commit, error) {
	out := new(Commit)
	err := grpc.Invoke(ctx, "/pb.RepositoryExplorer/FetchRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryExplorerClient) ListTreeEntries(ctx context.Context, in *Repository, opts ...grpc.CallOption) (RepositoryExplorer_ListTreeEntriesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RepositoryExplorer_serviceDesc.Streams[0], c.cc, "/pb.RepositoryExplorer/ListTreeEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryExplorerListTreeEntriesClient{stream}
	if err := x.ClientStream.SendProto(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RepositoryExplorer_ListTreeEntriesClient interface {
	Recv() (*TreeEntry, error)
	grpc.ClientStream
}

type repositoryExplorerListTreeEntriesClient struct {
	grpc.ClientStream
}

func (x *repositoryExplorerListTreeEntriesClient) Recv() (*TreeEntry, error) {
	m := new(TreeEntry)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryExplorerClient) GetBlobData(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Blob, error) {
	out := new(Blob)
	err := grpc.Invoke(ctx, "/pb.RepositoryExplorer/GetBlobData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryExplorerClient) DeleteRepo(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.RepositoryExplorer/DeleteRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepositoryExplorer service

type RepositoryExplorerServer interface {
	CreateRepo(context.Context, *Repository) (*Repository, error)
	FetchRepo(context.Context, *Repository) (*Commit, error)
	ListTreeEntries(*Repository, RepositoryExplorer_ListTreeEntriesServer) error
	GetBlobData(context.Context, *Repository) (*Blob, error)
	DeleteRepo(context.Context, *Repository) (*Empty, error)
}

func RegisterRepositoryExplorerServer(s *grpc.Server, srv RepositoryExplorerServer) {
	s.RegisterService(&_RepositoryExplorer_serviceDesc, srv)
}

func _RepositoryExplorer_CreateRepo_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Repository)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RepositoryExplorerServer).CreateRepo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RepositoryExplorer_FetchRepo_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Repository)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RepositoryExplorerServer).FetchRepo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RepositoryExplorer_ListTreeEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Repository)
	if err := stream.RecvProto(m); err != nil {
		return err
	}
	return srv.(RepositoryExplorerServer).ListTreeEntries(m, &repositoryExplorerListTreeEntriesServer{stream})
}

type RepositoryExplorer_ListTreeEntriesServer interface {
	Send(*TreeEntry) error
	grpc.ServerStream
}

type repositoryExplorerListTreeEntriesServer struct {
	grpc.ServerStream
}

func (x *repositoryExplorerListTreeEntriesServer) Send(m *TreeEntry) error {
	return x.ServerStream.SendProto(m)
}

func _RepositoryExplorer_GetBlobData_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Repository)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RepositoryExplorerServer).GetBlobData(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RepositoryExplorer_DeleteRepo_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Repository)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RepositoryExplorerServer).DeleteRepo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RepositoryExplorer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RepositoryExplorer",
	HandlerType: (*RepositoryExplorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _RepositoryExplorer_CreateRepo_Handler,
		},
		{
			MethodName: "FetchRepo",
			Handler:    _RepositoryExplorer_FetchRepo_Handler,
		},
		{
			MethodName: "GetBlobData",
			Handler:    _RepositoryExplorer_GetBlobData_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _RepositoryExplorer_DeleteRepo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTreeEntries",
			Handler:       _RepositoryExplorer_ListTreeEntries_Handler,
			ServerStreams: true,
		},
	},
}