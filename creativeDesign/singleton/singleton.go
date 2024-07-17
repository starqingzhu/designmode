package singleton

/*
定义：
	一个类只允许创建一个对象（或者实例），那么这个类叫单例类。这种设计模式叫单例模式。
用处：
	从业务概念上，有些数据在系统中只应该保存一份，脚比较适合设计为单例类。

唯一性：
	1)进程间唯一
	2）线程间唯一（go 协程间唯一）
	3）集群唯一

存在问题：
	1）对oop的特性支持不友好
	2）隐藏类之间的关系
	3）对代码扩展性不好 （创建多个实例）
	4）对代码的可测性不友好
	5）不支持有参数构造函数

如何实现：
	1）构造函数private访问权限
	2）对象创建时线程安全问题
	3）是否支持延迟加载
	4）getInstance的性能问题

实现方式：
	1)懒汉
	2)饿汉
	3)double check
	4)sync.Once
*/

// type Singleton interface {
// }
