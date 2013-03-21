package main

import (
	"simpleweb"
	"testapp"
)

func main() {
	apps := []*simpleweb.AppUrls{
		testapp.MakeAppUrls(),
	}

	server := new(simpleweb.Simpleserver)
	server.Run(apps)
}
