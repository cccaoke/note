### 让人耳目一新的算法题

+ [字典序的第K小数字](https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/)
+ [区域和检索 - 数组可修改](https://leetcode-cn.com/problems/range-sum-query-mutable/)
  使用到了多种优化手段，比如二级索引，这个在k8s中也有相关题先，比如EndpointSlice，线段树等
+ [随机数索引](https://leetcode-cn.com/problems/random-pick-index/)
  该题目可以使用蓄水池抽样算法，更一般的，我们从n个数字中选出m个数字，使得每个数字被选中的概率都是m/n，
我们逐个读取数字，对于第i个元素，我们以m/i的概率选择该元素，当i<m时，此概率为1，当i>m时，我们再以1/m的概率选择一个前边的元素做替换，
这样每个元素被选中的概率就是m/n，证明参考：https://blog.csdn.net/anshuai_aw1/article/details/88750673
问题：蓄水池算法有什么实际的应用场景？ 
+ 