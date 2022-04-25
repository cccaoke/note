### 数据结构

[slice](slice.md)

### 优雅的代码实现

给定两个协程，一个打印奇数，一个打印偶数，但是交替输出奇数和偶数  
只用一个chan就可以实现: [multi_thread.go](elegant-code/print_nums.go)

辗转相除求两数的最大公约数: [gcd.go](elegant-code/gcd.go)

### 待补充  
chan，select原理  
golang调度系统  
gc基本原理  
golang内存管理，包括逃逸分析  
如何优化golang程序，pprof等，编译参数，gc调优  
cgo以及plan9汇编 
