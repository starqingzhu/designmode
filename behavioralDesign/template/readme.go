package template

/*
定义：
	1.在一个方法中定义一个算法骨架，并将某些步骤推迟到子类中实现。
	2.可以让子类在不改变算法整体结构的情况下，重新定义算法的某些步骤。
应用场景：
	1.扩展 框架通过模版模式提供功能扩展点，让框架用户可以在不修改框架源码
	的情况下，基于扩展点定制化框架的功能
	2.复用 复用指的是，所有的子类可以服用父类中提供的模版方法的代码


与回调的区别：
	1.同步回调
		1）回调基于组合关系来实现，把一个对象传递给另一个对象，是一种
		对象之间的关系
		2）模版模式 基于继承关系来实现，子类重写父类的抽象方法，是一种
		类之间的关系
	2.异步回调
		1）更类似观察者模式

*/