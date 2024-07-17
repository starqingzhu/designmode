package chain

/*
职责链模式
定义：
	将请求的发送和接收解耦，让多个接收对象都有机会处理这个请求。将这些接收
对象串成一条链，并沿着这条链传递这个请求，直到链上的某个接收对象能够处理它为止。
	有一种变体是每个对象都能处理。

场景作用：
	复用和扩展，在实际的项目开发中比较常用，特别是框架开发中，我们可以利用他们
来提供框架的扩展点，能够让框架的使用者在不修改框架源码的情况下，基于扩展点定制化
框架的功能。
	eg:
		过滤器
		中间件

实现方式：
	链式
	数组
*/
