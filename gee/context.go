package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//origin objects 原始对象
	Writer http.ResponseWriter //响应
	Req    *http.Request       //请求
	//request info 请求接口
	Path   string //路径
	Method string //方法
	//response info 响应请求
	statusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//PostFrom 返回查询的命名组件的第一个值
func (c *Context) PostFrom(key string) string {
	return c.Req.FormValue(key)
}

//Query 查询网络
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//Status 设置响应
func (c *Context) Status(code int) {
	c.statusCode = code
	c.Writer.WriteHeader(code)
}

//SetHeader 设置头文件
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//String 构造String
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//JSON 构造JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "text/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//Data 构造Data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//HTML 构造HTML响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
