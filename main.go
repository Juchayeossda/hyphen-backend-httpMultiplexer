package main

import (
	"hyphen-backend-httpMultiplexer/core"
	"net/http"
)

func main() {
	authorities := core.NewAuthority([]core.AuthoritySet{
		{Name: "auth", DestAuthority: "http://172.20.10.10:8080"},
	})

	r := core.NewRouter()

	r.CreateFlow(http.MethodPost, "/api/auth/signin", authorities.Get("auth"))
	r.CreateFlow(http.MethodPost, "/api/auth/signout", authorities.Get("auth"))
	r.CreateFlow(http.MethodPost, "/api/auth/signup", authorities.Get("auth"))

	r.Run(":9190")
}
