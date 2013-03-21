package testapp

import (
	"net/http"
	"simpleweb"
)

func MakeAppUrls() *simpleweb.AppUrls {
	tv := &TestView{Value: 0}

	urls := []*simpleweb.Url{
		&simpleweb.Url{
			Pattern: "/{filename}",
			Handler: func(w http.ResponseWriter, r *http.Request) { tv.TestHandler(w, r) },
		},
	}

	appUrls := &simpleweb.AppUrls{Prefix: "/test", simpleweb.Urls: urls}
	return appUrls
}
