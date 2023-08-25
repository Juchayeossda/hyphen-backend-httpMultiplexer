package main

import (
	"hyphen-backend-httpMultiplexer/core"
	"net/http"
)

func main() {
	authorities := core.NewAuthority([]core.AuthoritySet{
		{Name: "auth", DestAuthority: "http://localhost:8000"},
	})

	r := core.NewRouter()

	r.CreateFlow(http.MethodGet, "/login", authorities.Get("auth"))
	r.CreateFlow(http.MethodGet, "/signup", authorities.Get("auth"), "/login")

	r.Run(":9190")
}
