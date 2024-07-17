package factory

/*

工厂模式分类：
1.简单工厂模式
2.工厂方法模式
3.抽象工厂模式
4.DI容器（依赖注入容器）

分类方法：
1.根据产品是具体产品还是具体工厂 分为简单工厂模式和工厂方法模式
2.根据工厂抽象程度分为工厂方法模式和抽象工厂模式

作用：
用于封装和管理对象创建，是一种创建模式


简单工厂模式
	定义：
		由于go本身没有构造函数，一般而言我们采用NewName的方式创建对象/接口， 当它返回的是接口的时候，其实就是简单工厂模式
	结构模型：
		工厂结构体
		产品接口
		产品结构体
		工厂方法模式
工厂方法模式
	定义：
		当对象的创建逻辑比较复杂，不只是简单的new一下就可以，而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式。将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类不至于过于发咋
	结构模型：
		工厂接口
		工厂结构体
		产品接口
		产品结构体 单产品

抽象工厂模式
	类比工厂方法模式，一个工厂创建不同类型的对象
	结构模型：
		工厂接口
		工厂结构体
		产品接口
		产品结构体 多产品


如何判断是否使用工厂模式？
	1.封装变化：创建逻辑有可能变化，封装工厂类之后，创建逻辑的变更对调用者更透明
	2.代码复用：创建代码抽离到独立 的工厂类之后可以复用
	3.隔离复杂性：封装复杂的创建逻辑，调用者无需了解如何创建对象
	4.控制复杂度： 将创建代码抽离出来，让原本的函数或类职责更单一，代码更简洁


DI容器
	golang 现有的依赖注入框架
	使用反射实现的: https://github.com/uber-go/dig
	使用 generate 实现的: https://github.com/google/wire

*/