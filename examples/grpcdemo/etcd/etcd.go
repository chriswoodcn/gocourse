package etcd

import "strconv"

type ServerEndPoint struct {
	addr string
	port int
}

const DefaultServerEndPointAddr = "127.0.0.1"
const DefaultServerEndPointPort = 2379

func (p *ServerEndPoint) Url() string {
	if p.addr == "" {
		p.addr = DefaultServerEndPointAddr
	}
	if p.port == 0 {
		p.port = DefaultServerEndPointPort
	}
	return p.addr + ":" + strconv.Itoa(p.port)
}
