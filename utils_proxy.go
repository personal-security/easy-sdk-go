package easysdk

import (
	"fmt"
	"os"

	"golang.org/x/net/proxy"
)

type ProxyModel struct {
	Protocol string
	Addr     string
	Port     int
	User     string
	Pass     string
}

func GetProxyDial(proxyModel *ProxyModel) proxy.Dialer {
	dialer, err := proxy.SOCKS5("tcp", proxyModel.Addr, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		return nil
	}
	return dialer
}
