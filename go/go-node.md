### 数据结构

[slice](slice.md)

### 优雅的代码实现

给定两个协程，一个打印奇数，一个打印偶数，但是交替输出奇数和偶数  
只用一个chan就可以实现: [multi_thread.go](elegant-code/print_nums.go)