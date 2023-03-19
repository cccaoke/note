**不是所有的特征都拥有特征对象，只有对象安全的特征才有，需要满足两个限制**：
- 方法不能返回Self类型（特征对象是动态分发，无法获知最初的类型了）
- 方法不能有任何泛型参数（猜测：rust会根据泛型函数的调用情况把所有的泛型函数具形化，但是如果是特征对象，那么特征对象去哪里找该具形的函数呢？）

**&[str]类型中底层数据是如何被回收的？**  
比如let a="hello world"，那么hello world的生命周期结束之后就会被drop掉，a只是对hello world字符串的借用，如果在函数中传递，那么就需要生命周期规则，
编译器检查的是生命周期，为了不出现悬垂引用，要保证引用的生命周期比引用的数据要短。