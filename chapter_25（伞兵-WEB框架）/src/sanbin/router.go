package sanbin

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       //roots 来存储每种请求方式的Trie 树根节点
	handlers map[string]HandlerFunc //使用 handlers 存储每种请求方式的 HandlerFunc
}

func newRouter() *router { //新建路由
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

func parsePattern(pattern string) []string { //解析:和*两种匹配符，返回一个map
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

//解析了:和*两种匹配符的参数，返回一个 map
//例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

//在调用匹配到的handler前，将解析出来的路由参数赋值给了c.Params， 这样就能够在handler中，通过Context对象访问到具体的值了。
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
