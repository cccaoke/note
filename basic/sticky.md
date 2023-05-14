# Sticky属性
Sticky属性只能应用在目录，当目录拥有Sticky属性所有在该目录中的文件或子目录无论是什么权限只有文件或子目录所有者和root用户能删除。

比如/tmp目录，因为机器上所有用户共享该目录，所以所有所有用户都有该目录的写权限，有了写权限就意味着可以删除别的用户的目录，这肯定是不允许的，因此
linux就引入了sticky属性，具体如下：
`
root@LAPTOP-N2QOFRKA:/# ls -al / | grep tmp
drwxrwxrwt   2 root root    4096 May 14 00:51 tmp
`