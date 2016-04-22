package main

import (
	"fmt"
	"github.com/mholt/caddy/middleware/fastcgi"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

type Proxy struct {
	server fastcgi.Handler
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := p.server.ServeHTTP(w, r)
	fmt.Println(status, err)
}

func main() {
	root := os.Args[1]
	source := os.Args[2]
	dest := os.Args[3]

	_ = os.Remove(source)

	listener, err := net.Listen("unix", source)
	if err != nil {
		panic(err)
	}
	err = os.Chmod(source, 0770)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	handler := fastcgi.Handler{
		Next:    nil,
		Root:    root,
		AbsRoot: root,
		Rules: []fastcgi.Rule{{
			Path:    "/",
			Address: dest,
			Ext:     "php",
		},
		},
	}

	p := Proxy{handler}

	err = fcgi.Serve(listener, p)
	if err != nil {
		panic(err)
	}

}
