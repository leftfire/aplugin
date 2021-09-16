package proto

import (
	context "context"
	"log"

	"github.com/hashicorp/go-plugin"
	grpc "google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"itsm": &MyPlugin{},
}

//对外接口 两个方法
type ExposeFunc interface {
	UseFunc(name string, config map[string]string) []byte
	UseStremFunc(req *FuncReq, pb ITS_UseStremFuncServer) ITS_UseStremFuncClient
}

//定义自己的plugin实现grpc的结构体
type MyPlugin struct {
	plugin.Plugin
	Impl ExposeFunc
}

//实现 Grpc的服务端方法
type GRPCServer struct {
	// This is the real implementation
	Broker *plugin.GRPCBroker
	Impl   ExposeFunc
}

//服务端逻辑
func (m *GRPCServer) UseStremFunc(req *FuncReq, pb ITS_UseStremFuncServer) error {
	m.Impl.UseStremFunc(req, pb)
	return nil
}
func (m *GRPCServer) UseFunc(c context.Context, req *FuncReq) (*FuncResp, error) {
	x := m.Impl.UseFunc(req.FuncName, req.Config)
	return &FuncResp{Data: x}, nil
}
func (m *GRPCServer) mustEmbedUnimplementedITSServer() {}

//实现Grpc proto的客户端的方法
type GRPCClient struct {
	Broker *plugin.GRPCBroker
	client ITSClient
}

// func (m *GRPCClient) UseFunc(ctx context.Context, in *FuncReq) (*FuncResp, error) {
// 	return &FuncResp{}, nil
// }

func (m *GRPCClient) UseFunc(name string, config map[string]string) []byte {
	res, err := m.client.UseFunc(context.Background(), &FuncReq{FuncName: name, Config: config})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res.Data
}

//处理客户端逻辑
func (m *GRPCClient) UseStremFunc(fr *FuncReq, pb ITS_UseStremFuncServer) ITS_UseStremFuncClient {
	cli, _ := m.client.UseStremFunc(context.Background(), fr)
	return cli
}

func (p *MyPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterITSServer(s, &GRPCServer{Impl: p.Impl, Broker: broker})
	return nil
}

func (p *MyPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewITSClient(c), Broker: broker}, nil
}
