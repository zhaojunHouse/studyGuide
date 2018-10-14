#  安全
#### 堡垒机（跳板机）
- 《[搭建简易堡垒机](http://blog.51cto.com/zero01/2062618)》


#  Web Server
#### nginx
- 《[nginx入门教程](http://tengine.taobao.org/book/chapter_02.html#id13)》
   * 进程模型：一个master进程，管理多个worker进程。一个请求只能一个worker进程处理。worker进程抢acceptMutex锁,然后接受请求，处理请求，返回请求，断开连接
   * 事件模型：高并发，异步非阻塞事件机制。epoll，循环处理准备好的事件请求。没有准备好就放回对列。
   * connection: TCP三次握手
   * request:  请求行、请求头、请求体、响应行、响应头、响应体
   
- 《[nginx配置](https://www.cnblogs.com/knowledgesea/p/5175711.html)》
   * 负载均衡 [配置文档](https://www.cnblogs.com/andyfengzp/p/6434125.html)
   * 反向代理和负载均衡 [配置文档](https://www.cnblogs.com/Miss-mickey/p/6734831.html)


- 《[nginx正向代理和反向代理](https://www.cnblogs.com/Anker/p/6056540.html)》
   * 正向代理：访问无法访问的服务器；缓存，加速访问；登陆授权；访问日志记录；
   * 反向代理：负载均衡；保证原始服务器安全；
   * 正向代理和反向代理 [配置文档](https://blog.csdn.net/baidu_19473529/article/details/79435069)
   * 正向代理和反向代理 [配置示例](https://blog.csdn.net/hiyun9/article/details/51602428)
   
- 《[nginx中文文档](http://www.nginx.cn/doc/)》



#  容器技术
#### docker
- 《[几张图帮你理解 docker 基本原理及快速入门](https://www.cnblogs.com/SzeCheng/p/6822905.html)》
   * 镜像，容器。一个镜像可以建立多个容器，一个容器最多127层。
- 《[docker核心技术与实现](https://draveness.me/docker)》
- 《[docker教程](http://www.runoob.com/docker/docker-tutorial.html)》
   
   
   

#  并发
##### 概念
- 《[为什么需要并发？](https://juejin.im/post/5ae6c3ef6fb9a07ab508ac85)》
  * 充分利用多核CPU的计算能力
  * 业务拆分，提升性能
   
- 《[并发与并行，并发与异步区别](https://www.cnblogs.com/xqn2017/p/8029472.html)》
  * 并发：通过切换时间片的方式多个任务交替进行；并行：真正意义上的“同时进行”
   
- 《[高并发实现需要考虑什么？](https://www.cnblogs.com/PerkinsZhu/p/7242247.html)》 
  * 硬件，网络，架构设计（各种中间件，缓存），开发语言，数据结构，算法，数据库优化，代码层（多线程编程，同步异步）
  
- 《[高并发实现中一种方式：多线程开发及优缺点](https://www.cnblogs.com/xiejw/p/5259437.html)》
  *  优点：耗时任务放入子线程防止阻塞，提升响应速度；充分利用CPU，提高性能；改善程序结构；
  *  缺点：频繁切换上下文；线程安全问题；
   
- 《[线程安全问题及解决方案](https://www.cnblogs.com/zhanht/p/5450325.html)》
  * 产生原因：CPU切换时间片执行多线程产生竞态条件
  * 原子性，可见性，时序性；原子性：要么都执行，要么都不执行。a++是先查，后加，不是原子操作。
  * 加锁

   
#### 锁


- 《[锁的分类](https://www.cnblogs.com/qifengshi/p/6831055.html)》
   * 乐观锁&悲观锁，互斥锁/写锁，共享锁/读锁，公平锁/非公平锁，重入锁/非重入锁，分段锁等。
   
- 《[乐观锁](http://www.importnew.com/20472.html)》
   * 是一种思想，提交更新时对冲突检测，让用户决定重试或是其他。
   * ABA问题，解决方法对每次更新加上版本号
  
- 《[悲观锁与乐观锁](http://www.importnew.com/20472.html)》
   * select for update
   
- 《[对mysql乐观锁、悲观锁、共享锁、排它锁、行锁、表锁概念的理解](https://blog.csdn.net/puhaiyang/article/details/72284702)》
   * 对于update,insert,delete语句会自动加排它锁,只共享读。
   
- 《[select.. for update导致的数据库死锁分析](https://www.cnblogs.com/Lawson/p/5008741.html)》
   * mysql的innodb存储引擎实务锁虽然是锁行，但它内部是锁索引的。
   * 锁相同数据的不同索引条件可能会引起死锁。

- 《[常见的死锁原因与解决方法](https://www.cnblogs.com/zejin2008/p/5262751.html)》《[delete造成的死锁](http://hedengcheng.com/?p=844)》
   
- 《[mysql加锁分析](http://hedengcheng.com/?p=771)》  《[聚簇索引与非聚簇索引](http://www.cnblogs.com/tuyile006/archive/2009/08/28/1555615.html)》《[MySQL索引原理及慢查询优化](https://tech.meituan.com/mysql_index.html)》



# 数据结构
- 队列 
  * [实现](https://github.com/zhaojunHouse/studyGuide/blob/master/queue.go)
- 链表 
  * [实现](https://github.com/zhaojunHouse/studyGuide/blob/master/node.go)
- 集合 
  * [实现](https://github.com/zhaojunHouse/studyGuide/blob/master/set.go)
- 数组
- 关联数组
- 字典
- 哈希表
   * [定义](https://blog.csdn.net/duan19920101/article/details/51579136/)
   * [实现](https://github.com/zhaojunHouse/studyGuide/blob/master/hash.go)
- 栈
   * [实现](https://github.com/zhaojunHouse/studyGuide/blob/master/stack.go)
- 树
- 图
- 堆


# TCP/HTTP

- 《[OSI七层网络协议模型](https://www.cnblogs.com/Robin-YB/p/6668762.html)》

- 《[点对点 ，端到端](https://blog.csdn.net/qq_34940959/article/details/78583993)》

- 《[TCP/IP四层协议理解](https://www.cnblogs.com/onepixel/p/7092302.html)》

- 《[TCP三次握手，四次挥手](https://www.cnblogs.com/huhuuu/p/3572485.html)》

- 《[TCP客户端服务端，连接断开，示例](https://www.cnblogs.com/huhuuu/p/3572485.html)》

- 《[HTTP协议详解，抓包分析](https://www.cnblogs.com/wangning528/p/6388464.html)》

- 《[HTTP2.0原理详解](https://blog.csdn.net/zhuyiquan/article/details/69257126)》

- 《[HTTP2.0二进制帧](https://blog.csdn.net/u012657197/article/details/77877840)》

- 《[HTTPS原理](https://www.cnblogs.com/zhangshitong/p/6478721.html)》《[一个故事讲清HTTPS](https://mp.weixin.qq.com/s/StqqafHePlBkWAPQZg3NrA)》


# 数据库

- 《[三大范式五大约束](https://www.cnblogs.com/waj6511988/p/7027127.html)》
- 《[数据库事务特性](https://blog.csdn.net/u012440687/article/details/52116108)》
- 《[数据库事务隔离级别](https://blog.csdn.net/qq_33290787/article/details/51924963)》
- 《[InnoDB幻读问题](http://blog.sina.com.cn/s/blog_499740cb0100ugs7.html)》



# 参考文档

- logrotate日志切割
   * [linux日志切割](https://blog.csdn.net/junli_chen/article/details/77193438)
   * [日志切割配置]( http://blog.51cto.com/wn2100/2074048)
