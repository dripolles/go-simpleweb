package simpleweb

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http/fcgi"
)

type Simpleserver struct{}

func (s *Simpleserver) Run(apps []*AppUrls) {
	fmt.Println("OK")
	l, err := net.Listen("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Println("Noooo")
	}

	router := mux.NewRouter()
	RegisterAll(router, apps)
	fcgi.Serve(l, router)

}

func RegisterAll(router *mux.Router, apps []*AppUrls) {
	for i := 0; i < len(apps); i++ {
		Register(router, apps[i])
	}
}

func Register(router *mux.Router, app *AppUrls) {
	sr := router.PathPrefix(app.Prefix).Subrouter()
	urls := app.Urls
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		sr.HandleFunc(url.Pattern, url.Handler)
	}
}
