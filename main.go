package main

import (
	"flag"
	"fmt"
	"github.com/mholt/caddy/middleware"
	"github.com/mholt/caddy/middleware/fastcgi"
	logg "github.com/mholt/caddy/middleware/log"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

var listen string
var target string
var log_path string
var root string

func init() {
	flag.StringVar(&listen, "listen", "/var/run/fcgiproxy.sock", "What Apache or Nginx see")
	flag.StringVar(&target, "target", "/var/run/php5-fpm.sock", "Where is php-fpm")
	flag.StringVar(&log_path, "log", "/var/log/fcgiproxy.log", "Log path")
	flag.StringVar(&root, "root", "/var/www", "Where the PHP files live")
}

type Proxy struct {
	server middleware.Handler
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := p.server.ServeHTTP(w, r)
	fmt.Println(status, err)
}

func main() {
	flag.Parse()

	_ = os.Remove(listen)

	listener, err := net.Listen("unix", listen)
	if err != nil {
		panic(err)
	}
	err = os.Chmod(listen, 0770)
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
			Address: target,
			Ext:     "php",
		},
		},
	}

	var file *os.File
	file, err = os.OpenFile(log_path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := logg.Logger{
		Next: handler,
		Rules: []logg.Rule{{
			PathScope:  "/",
			OutputFile: log_path,
			Format:     "{when} {host} {method} {path} {latency} {size} {status}",
			Log:        log.New(file, "", 0),
		},
		},
	}

	p := Proxy{logger}

	err = fcgi.Serve(listener, p)
	if err != nil {
		panic(err)
	}

}
