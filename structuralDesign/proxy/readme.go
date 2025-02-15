package proxy

/*
作用：
	在不改变原始类（或叫被代理类）代码的情况下，通过引入代理类来给原始类附加功能（但是附加的功能和原有业务功能没有强关联）

应用场景：
	业务系统的非功能性需求开发（监控、统计、鉴权、限流、事务等）
	RPC
	缓存

分类：
	动态代理
		1）一般采用反射实现
		2）不事先为每个原始类编写代理类，而是在运行的时候，动态的创建原始类对应的代理类
	静态代理
		1）代理类实现和源类相同的接口，每个类都单独编辑一个代理类
		2）实现源类中所有方法，并为每个方法都附加相似的代码逻辑
		3）如果添加的功能类有不止一个，我们需要针对每个类创建一个代理类

*/
