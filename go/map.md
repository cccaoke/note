#### map

- map中通过mapaccess函数来找到给定key的value，其函数定义如下所示：func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer，
  在hash了给定key以后，取低B位来索引桶，高8位索引HOB
- 为了加速索引过程，golang在源码中实现了很多的mapaccess系列函数
- 上面通过hash值来计算HOB时，有一个判断，如果<1或者>5的话，就会给HOB加上一个偏移minTopHash，因为前边几个是保留字段，标识该值是否发生了迁移。
- bucket数目的扩容应该分为两种，同等大小扩容，这样就会扩容出来一批新的bucket，其大小和老的这一批一模一样，这时是没有必要执行迁移的，
  因为迁移了还是这么多，这时候hmap的oldBuckets和buckets的数目就是一样的，还有一种可能是扩容以后buckets是oldBuckets的两倍。
- 当遍历新的bucket时，需要考虑该bucket是否是由老的oldBucket迁移而来的，迁移分为同等大小迁移或者扩大两倍迁移，都需要找到老的bucket，从而进行check，
  判断的算法逻辑为，首先选定该checkBucket，该checkBucket是搬迁以后的bucket num，如果其最高位和tophash中的最低位一致，说明应该在本次iternext流程中check，
  否则就会被判定为应该在另一新的bucket中check
- 当给某个key分配value时，会先找到对应的bucket，如果此bucket关联的bucket还没有完全迁移到新的bucket，那么需要先执行迁移操作，而后才会进行赋值，
  