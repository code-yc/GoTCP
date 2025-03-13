package gnet

import "GoTCP/giface"

// Server 作为 IServer 接口实现，定义一个 Server 服务类
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// NewServer 创建一个服务器句柄
func NewServer(name string) giface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}
