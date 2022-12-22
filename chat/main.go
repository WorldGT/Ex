package main

import (
	"chat/internal/handler"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	// router := mux.NewRouter()
	// go h.run()
	// router.HandleFunc("/ws", myws)
	// if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
	// 	fmt.Println("err:", err)
	// }

	RegisterHandlers()
}

var (
	port    = flag.Int("port", 8080, "the port to listen")
	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
	cpu     = flag.Int64("cpu", 500, "cpu threshold")
)

func RegisterHandlers() {

	flag.Parse()
	fmt.Println("port=", *port)
	logx.Disable()

	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host:         "127.0.0.1",
		Port:         *port,
		Timeout:      *timeout,
		CpuThreshold: *cpu,
	})
	defer engine.Stop()
	go h.run()
	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			if r.Method != "GET" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			http.ServeFile(w, r, "local.html")
		},
	})

	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			myws(w, r)
		},
	})

	engine.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/user/login",
		Handler: handler.UserLoginHandler(),
	},
	)

	engine.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/user/register",
		Handler: handler.UserRegisterHandler(),
	},
	)

	engine.Start()
}
