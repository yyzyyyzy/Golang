package sanbin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter //输出
	Req    *http.Request       //请求
	// request info
	Path   string            //请求路径
	Method string            //请求方法
	Params map[string]string //路由参数解析
	// response info
	StatusCode int //状态码
	// middleware
	handlers []HandlerFunc //中间件返回函数
	index    int           //索引
}

func newContext(w http.ResponseWriter, req *http.Request) *Context { //初始化Context
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) PostForm(key string) string { //请求方法
	return c.Req.FormValue(key) //原生go方法：得到Post方法请求的参数
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key) //原生go方法：得到get请求的参数
}

func (c *Context) Status(code int) { //设置状态码
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) { //设置header
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) { //设置输出字符串
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) { //设置输出json
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

// 在 HandlerFunc 中，希望能够访问到解析的参数，因此，需要对 Context 对象增加一个属性和方法，来提供对路由参数的访问
// 我们将解析后的参数存储到Params中，通过c.Param("lang")的方式获取到对应的值

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}
