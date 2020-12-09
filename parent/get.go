package parent

import (
	"fmt"
	"strings"

	zi "github.com/ashtyn3/zi/pkg"
	"github.com/goware/urlx"
)

// For making a parent node
func MakeParent(url string, pw string) Host {
	if strings.HasPrefix(url, "http") {
		u, _ := urlx.Parse(url)
		port := u.Port()
		host := u.Hostname()
		url := u.String()
		return Host{url: url, port: port, host: host}
	}
	return Host{}
}

func (h Host) Get() {
	zi.Zi(h.url, h.auth)
	fmt.Println(h)
}
