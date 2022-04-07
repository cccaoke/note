### DeletedFinalStateUnknown这个状态是怎么来的

#### 用法

```
podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			nc.podUpdated(nil, pod)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(nil, pod)
			}
		},
		UpdateFunc: func(prev, obj interface{}) {
			prevPod := prev.(*v1.Pod)
			newPod := obj.(*v1.Pod)
			nc.podUpdated(prevPod, newPod)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(prevPod, newPod)
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod, isPod := obj.(*v1.Pod)
			// We can get DeletedFinalStateUnknown instead of *v1.Pod here and we need to handle that correctly.
			if !isPod {
				deletedState, ok := obj.(cache.DeletedFinalStateUnknown)
				if !ok {
					klog.Errorf("Received unexpected object: %v", obj)
					return
				}
				pod, ok = deletedState.Obj.(*v1.Pod)
				if !ok {
					klog.Errorf("DeletedFinalStateUnknown contained non-Pod object: %v", deletedState.Obj)
					return
				}
			}
			nc.podUpdated(pod, nil)
			if nc.taintManager != nil {
				nc.taintManager.PodUpdated(pod, nil)
			}
		},
	})
```

上面即我们常见的给pod添加监听，这里面有一个问题，在我们的DeleteFunc回调函数中我们首先拿着v1.Pod做类型断言，如果不满足这时候我们就会使用DeletedFinalStateUnknown做类型断言，也就是说，DeleteFunc回调函数中有可能会产生该类型对象，那么类型对象是怎么来的呢？

#### 原因

在整个Informer中，我们认为DeltaFifo是我们的消息队列，那么这个消息队列的生产者就是Reflector，同样ListAndWatch的操作也是在该对象中完成的，在List-Watch中，我们说先会执行一把全量list，结束后我们会进行watch操作，所有的这些数据都会被放到DeltaFifo中，
在我们list完成后，我们就需要把我们list的全量数据都放到DeltaFifo中，这里具体实现就是Reflector的syncWith方法，参考如下：

```
// syncWith replaces the store's items with the given list.
func (r *Reflector) syncWith(items []runtime.Object, resourceVersion string) error {
	found := make([]interface{}, 0, len(items))
	for _, item := range items {
		found = append(found, item)
	}
	return r.store.Replace(found, resourceVersion)
}
```

从注释中我们可以看到该方法替换store中所有的items，事实上如果items中的对象不存在，这里也会执行插入动作，这里我们摘出Repleace方法中一个小片段说明问题

```
...
for k, oldItem := range f.items {
    if keys.Has(k) {
        continue
    }
    var deletedObj interface{}
    if n := oldItem.Newest(); n != nil {
        deletedObj = n.Object
    }
    queuedDeletions++
    if err := f.queueActionLocked(Deleted, DeletedFinalStateUnknown{k, deletedObj}); err != nil {
        return err
    }
}
...
```

上述代码片段我们可以看到首先我们便利当前内存中的store中的items，然后从会从待替换（插入）的items中寻找是否存在，如果不存在，这个时候说明该对象已经被删除，
但是我们不知道为什么被删除了，因为如果是正常的删除动作会被我们watch到，那就说明发生了断连，在断连期间该对象被删除了，这是我们并不直到该对象在被删除之前发生过什么，
但是按照k8s面向终态的思想，这是应该给客户端一个机会去处理该对象的删除事件，但是我们又不知道删除时的前一个对象是什么样（由于断连我们不知道版本），
所以k8s就构造了一个DeletedFinalStateUnknown类型的对象来通知所有的监听者，该对象被删除了，我只有我记录的上一个版本的对象，请你们自行选择处理。