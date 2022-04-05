#### slice结构体定义

```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

![slice示例](./images/img.png)

#### slice的初始化

```
func makeslice(et *_type, len, cap int) unsafe.Pointer {
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}

	return mallocgc(mem, et, true)
}
```

上面判断了len和cap够不够用，有没有越界，然后直接调用mallocgc分配内存。 问题：从make是怎么到makeslice函数的

#### 扩容

扩容应该会调用growslice函数，该函数讲old slice扩容为cap大小，代码在: runtime/slice.go

```
func growslice(et *_type, old slice, cap int) slice
```

前置校验

```
// slice只能扩容，不能缩容
if cap < old.cap {
	panic(errorString("growslice: cap out of range"))
}

// 如果type的size为0，那么只用改变其cap大小即可
if et.size == 0 {
	// append should not create a slice with nil pointer but non-zero len.
	// We assume that append doesn't need to preserve old.array in this case.
	return slice{unsafe.Pointer(&zerobase), old.len, cap}
}
```
扩容策略：如果给定的cap大于old的两倍，那么扩容后newcap就是给定的cap，否则按照如下策略扩容

- 如果old.cap<1024，那么newcap=old.cap * 2
- 否则按照每次增长1/4 * old.cap的策略，一直累加，直到newcap大于给定的cap

在选定cap以后就会根据不同的类型计算len，cap，以及是否溢出，溢出的标准就是待分配的内存是否大于
机器可以可被分配的最大内存

小知识：在判断一个数是否是2的倍数时，slice中计算方式为，也即是x中只能有一个1
```
func isPowerOfTwo(x uintptr) bool {
	return x&(x-1) == 0
}
```



