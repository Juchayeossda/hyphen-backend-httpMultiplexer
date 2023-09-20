package main

import "github.com/JunBeomHan/hmux"

const User = "http://localhost:8080"
const Mail = "http://localhost:8081"
const Inquery = "http://localhost:8082"
const SISS = "http://localhost:8083"
const Hellog = "http://localhost:8084"

func main() {

	s := hmux.NewServer()

	s.CreateFlow("GET", "/api/inquiry/myInquirys", Inquery)
	s.CreateFlow("GET", "/api/inquiry/inquirys/:id", Inquery)
	s.CreateFlow("POST", "/api/inquiry", Inquery)

	// Auth controller
	s.CreateFlow("POST", "/api/auth/signin", User)
	s.CreateFlow("POST", "/api/auth/signout", User)
	s.CreateFlow("POST", "/api/auth/signup", User)

	// UserController
	s.CreateFlow("DELETE", "/api/user/drop", User)
	s.CreateFlow("GET", "/api/user/image", User)
	s.CreateFlow("FATCH", "/api/user/image", User)
	s.CreateFlow("DELETE", "/api/user/image", User)
	s.CreateFlow("GET", "/api/user/info", User)
	s.CreateFlow("PATCH", "/api/user/name", User)
	s.CreateFlow("PATCH", "/api/user/password", User)
	s.CreateFlow("DELETE", "/api/user/drop", User)

	// Validate Controller
	s.CreateFlow("POST", "/api/auth/jwt/refresh", User)
	s.CreateFlow("POST", "/api/auth/jwt/validate", User)

	// Mail Controller
	s.CreateFlow("POST", "/api/auth/email", Mail)

	// SISS Controller
	s.CreateFlow("POST", "/api/siss/upload/image", SISS)
	s.CreateFlow("GET", "/api/siss/extract/image/:imageName", SISS)

	s.Run(":7777")
}
