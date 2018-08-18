package endpoint

import (
	"context"
)

// Endpoint is the fundamental building block of servers and clients.
// It represents a single RPC method.

// Endpoint是服务器和客户端的基础构件
// 它代表这一个单一的RPC方法

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Nop is an endpoint that does nothing and returns a nil error.
// Useful for tests.
// 这个方法就是用来测试的,Nop作为一个endpoint,没有任何业务逻辑就是返回一个nil
func Nop(context.Context, interface{}) (interface{}, error) { return struct{}{}, nil }

// Middleware是一个
// 中间件是端点的可链接行为修饰符
// Middleware is a chainable behavior modifier for endpoints.
type Middleware func(Endpoint) Endpoint

// Chain is a helper function for composing middlewares. Requests will
// traverse them in the order they're declared. That is, the first middleware
// is treated as the outermost middleware.

// 链是组成中间件的一个辅助函数。
// 请求会按照他们定义的顺序来遍历他们。也就是第一个中间件会被视为最外层的中间件
func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}
