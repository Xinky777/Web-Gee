package gee

import (
	"fmt"
	"net/http"
)

//HandlerFunc defines the request handler used by gee_web
//定义gee使用的请求处理程序
type HandlerFunc func(http.ResponseWriter, *http.Request)

//Engine implement the interface of ServeHTTP
//实现ServeHTTP的接口
type Engine struct {
	router map[string]HandlerFunc
}

//New the constructor of gee_web
//gee的构造函数
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

//addRoute 增添 [方法]-模式 处理函数
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//GET define the method to add GET request
//定义增加 GET 请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST define the method to add POST request
//定义增加 POST 请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//ServeHTTP engine方法实现类型handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}
