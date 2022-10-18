## list-watch中bookmark事件？

在reflector，首先会执行全量list动作，然后执行watch动作，假设一个对象它长时间没有更新，恰巧此时又发生了断链，
在我们现有的watch中，我们就会重新建立从上一次拿到的resourceVersion开始watch，那么因为长时间没有更新动作，所以我们手里的resourceVersion太老了，
以至于etcd中已经没有从我们指定的resourceVersion后的数据了（etcd的event一般保留五分钟，以免数据膨胀），这时kube-api就会告诉我们too old resourceVersion，
那么我们就需要从0开始重新list，这样就会增加kube-api的负载。因此k8s引入了bookmark，翻译成中文就是书签，经常记录一下我们watch到哪里了，也即是我们在watch的过程中，
kube-api可能不定期的给我们发送bookmark事件，然后里面加入了最新的resourceVersion，当我们观察到bookmark事件以后，我们就更新我们内存中记录的resourceVersion，
那么当我们发生了断链以后，不至于我们的resourceVersion太老而不被kube-api接受。
