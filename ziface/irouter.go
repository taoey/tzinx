package ziface

type IRouter interface {
	PreHandle(request IRequest)  // 处理业务之前的钩子方法hook
	Handle(request IRequest)     // 处理具体业务的方法hook
	PostHandle(request IRequest) // 处理业务之后的方法hook
}
