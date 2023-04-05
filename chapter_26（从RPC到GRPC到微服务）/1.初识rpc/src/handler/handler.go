package handler

// 可以解决微服务名称冲突问题
const HelloServiceName = "handler/HelloService"

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
