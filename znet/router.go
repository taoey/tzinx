package znet

import "github.com/taoey/tzinx/ziface"

// 实现router，先嵌入baseRouter基类，用户根据这个基类进行重写
// 用户只需要重新需要重新的方法，没必要重写全部方法
type BaseRouter struct{}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
