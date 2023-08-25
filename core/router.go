package core

import (
	"io"
	"log"
	"net/http"
)

type router struct {
	handlers map[string]map[string]http.HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]map[string]http.HandlerFunc)}
}

func (r *router) registerHandler(method, path string, handlerFunc http.HandlerFunc) {
	handler, isExist := r.handlers[method]

	if isExist == false {
		handler = make(map[string]http.HandlerFunc)
		r.handlers[method] = handler
	}

	handler[path] = handlerFunc
}

func (r *router) CreateFlow(method, path, destAuthority string, destPath ...string) {
	// 요청의 정보를 사용하여 흐름을 생성할 마이크로서비스의 URL을 생성합니다.
	var destURL string

	tempDestPath := ""
	for i, v := range destPath {
		if i >= 2 {
			log.Println("too much ... destPath")
			return
		}
		tempDestPath = v
	}

	if tempDestPath == "" {
		destURL = destAuthority + path
	} else {
		destURL = destAuthority + tempDestPath
	}

	r.registerHandler(method, path, func(w http.ResponseWriter, r *http.Request) {

		// 해당 URL로 요청합니다.
		req, err := http.NewRequest(method, destURL, r.Body)
		req.Header.Add("Content-type", "application/json")
		if err != nil {
			log.Println("http.NewRequest() error: ", err)
			return
		}

		// 응답을 수신합니다.
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("http.DefaultClient.Do() error: ", err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("io.ReadAll() error: ", err)
			return
		}

		// 마이크로서비스에서 응답을 클라이언트에게 전송합니다.(응답)
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, isHandlerExist := r.handlers[req.Method]; isHandlerExist {
		if handlerFunc, isHandlerFuncExist := handler[req.URL.Path]; isHandlerFuncExist {

			handlerFunc(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *router) Run(addr string) error {
	return http.ListenAndServe(addr, r)
}
