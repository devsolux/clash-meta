package inbound

import (
	"net"

	C "github.com/devsolux/clash-meta/constant"
	"github.com/devsolux/clash-meta/transport/socks5"
)

// NewHTTP receive normal http request and return HTTPContext
func NewHTTP(target socks5.Addr, srcConn net.Conn, conn net.Conn, additions ...Addition) (net.Conn, *C.Metadata) {
	metadata := parseSocksAddr(target)
	metadata.NetWork = C.TCP
	metadata.Type = C.HTTP
	ApplyAdditions(metadata, WithSrcAddr(srcConn.RemoteAddr()), WithInAddr(conn.LocalAddr()))
	ApplyAdditions(metadata, additions...)
	return conn, metadata
}
