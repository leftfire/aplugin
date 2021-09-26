package arpc

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

//对外的接口   返回到resp
type Arpc interface {
	UseFunc(req Req, resp *Resp) error
}

//请求对象
type Req struct {
	//方法名
	FuncName string
	//插件配置
	Config map[string]string
	//请求参数
	Para map[string]string
}

//返回对象
type Resp struct {
	Code string
	Data string
}

//client
type ArpcClient struct {
	client *rpc.Client
}

func (a *ArpcClient) UseFunc(req Req, resp *Resp) error {
	err := a.client.Call("Plugin.UseFunc", req, resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return err
	}
	return nil
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type ArpcRPCServer struct {
	// This is the real implementation
	Impl Arpc
}

func (s *ArpcRPCServer) UseFunc(req Req, resp *Resp) error {
	err := s.Impl.UseFunc(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
type ArpcPlugin struct {
	// Impl Injection
	Impl Arpc
}

func (p *ArpcPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ArpcRPCServer{Impl: p.Impl}, nil
}

func (p *ArpcPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ArpcClient{client: c}, nil
}
