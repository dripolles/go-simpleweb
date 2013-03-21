package simpleweb

import (
	"net/http"
)

type AppUrls struct {
	Prefix string
	Urls   []*Url
}

type Url struct {
	Pattern string
	Handler (func(w http.ResponseWriter, r *http.Request))
}
