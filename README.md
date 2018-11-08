<!-- TOC -->

- [网络](#网络)
    - [1 网络协议](#1-网络协议)  
            &nbsp;&nbsp;- [1.1 TCP/IP](#11-tcpip)  
            &nbsp;&nbsp;- [1.2 HTTP](#12-http)  
            &nbsp;&nbsp;- [1.3 LOT七大通信协议](#13-lot七大通信协议)
    - [2 网络模型](#2-网络模型)  
            &nbsp;&nbsp;- [2.1 I/O模型](#21-io模型)  
            &nbsp;&nbsp;- [2.2 select，poll，epoll](#22-selectpollepoll)  
            &nbsp;&nbsp;- [2.3 BIO NIO AIO](#23-bio-nio-aio)
            &nbsp;&nbsp;- [2.4 kqueue](#24-kqueue)
- [中间件](#中间件)  
    - [1 Web Server](#1-web-server)  
            &nbsp;&nbsp;- [1.1 nginx](#11-nginx)  
            &nbsp;&nbsp;- [1.2 apache](#12-apache)  
            &nbsp;&nbsp;- [1.3 OpenResty](#13-openresty)
    - [2 缓存](#2-缓存)
    - [3 消息队列](#3-消息队列)
    - [4 定时调度](#4-定时调度)
    - [5 API网关](#5-api网关)
- [设计模式](#设计模式)
- [操作系统](#操作系统)
    - [1 计算机原理](#1-计算机原理)
    - [2 CPU](#2-cpu)
    - [3 进程](#3-进程)
    - [4 线程](#4-线程)
    - [5 协程](#5-协程)
    - [6 Linux](#6-linux)
- [架构&运维&统计&技术支持](#架构运维统计技术支持)
    - [1 日志服务](#1-日志服务)
    - [2 微服务架构](#2-微服务架构)
    - [3 RPC](#3-rpc)
    - [4 docker](#4-docker)
- [安全](#安全)
    - [1 堡垒机（跳板机）](#1-堡垒机跳板机)
- [并发](#并发)
    - [1 并发概念](#1-并发概念)
    - [2 锁](#2-锁)
- [数据库](#数据库)
- [数据结构&算法](#数据结构算法)
    - [1 数据结构](#1-数据结构)
    - [2 算法](#2-算法)
- [测试](#测试)
    - [1 压力测试](#1-压力测试)

<!-- /TOC -->
# 网络

## 1 网络协议

#### 1.1 TCP/IP
- 《[OSI七层网络协议模型](https://www.cnblogs.com/Robin-YB/p/6668762.html)》
   * OSI七层：物理层（IEEE 802.1A, IEEE 802.2到IEEE 802.11）、数据链路层（FDDI, Ethernet, Arpanet, PDN, SLIP, PPP）、网络层（IP, ICMP, ARP, RARP, AKP, UUCP）、传输层（TCP, UDP）、会话层（SMTP, DNS）、表示层（Telnet, Rlogin, SNMP, Gopher）、应用层（HTTP、TFTP, FTP, NFS, WAIS、SMTP）
   * TCP/IP四层：数据链路层、网络层、传输层、应用层

- 《[点对点 ，端到端](https://blog.csdn.net/qq_34940959/article/details/78583993)》
   * 端到端：针对传输层，发送端与接收端之间建立连接再发数据。
   * 点到点：针对数据链路层或网络层，数据发给直接相连的设备。

- 《[TCP/IP四层协议理解](https://www.cnblogs.com/onepixel/p/7092302.html)》
   * 链路层：对电信号进行分组并形成具有特定意义的数据帧，然后以广播的形式通过物理介质发送给接收方
   * 网络层：（IP协议，ARP协议，路由协议）定义网络地址，区分网段，子网内MAC寻址，对于不同子网的数据包进行路由
   * 传输层：定义端口，标识应用程序身份，实现端口到端口的通信，TCP协议可以保证数据传输的可靠性。
   * 应用层：定义数据格式并按照对应的格式解读数据

- 《[TCP三次握手，四次挥手](https://www.cnblogs.com/huhuuu/p/3572485.html)》《[TCP三次握手，四次挥手详解](https://blog.csdn.net/qzcsu/article/details/72861891)》
   * URG	紧急指针是否有效。为1，表示某一位需要被优先处理。
   * ACK	确认号是否有效，一般置为1。
   * PSH	提示接收端应用程序立即从TCP缓冲区把数据读走。
   * RST	对方要求重新建立连接，复位。
   * SYN	请求建立连接，并在其序列号的字段进行序列号的初始值设定。建立连接，设置为1.
   * FIN	希望断开连接。
   * 三次握手： 1. SYN = 1，seq = j  2. SYN=1，ACK=1,ack= j+1,seq= k    3.ACK=1,ack=k+1
   * 四次挥手： 1. FIN=m 2.ACK=m+1 3.FIN=n   4 ACK=1,ack=n+1

- 《[TCP客户端服务端，连接断开，示例](https://github.com/zhaojunHouse/demo)》

#### 1.2 HTTP
- 《[HTTP协议详解，抓包分析](https://www.cnblogs.com/wangning528/p/6388464.html)》《[HTTP2.0原理详解](https://blog.csdn.net/zhuyiquan/article/details/69257126)》《[HTTP2.0二进制帧](https://blog.csdn.net/u012657197/article/details/77877840)》
   * HTTP2.0二进制桢
   * HTTP2.0:多路复用、压缩头信息、请求划分优先级、支持服务器端主动推送   
   * HTTP1.0:每个请求都需单独建立连接（keep-alive能解决部分问题单不能交叉推送）、每个请求和响应都需要完整的头信息、数据未加密


- 《[HTTPS原理](https://www.cnblogs.com/zhangshitong/p/6478721.html)》《[一个故事讲清HTTPS](https://mp.weixin.qq.com/s/StqqafHePlBkWAPQZg3NrA)》
   * 1. 浏览器发送HTTPS请求   2 服务器发送数字证书给客户端    3 浏览器CA列表验证证书    4 浏览器产生对称密钥，用公钥加密   5 用私钥揭秘获得对称密钥
   6 用对称密钥加密通信。

#### 1.3 LOT七大通信协议
- 《[LOT七大通信协议](https://www.jianshu.com/p/f5a6977e9fef)》
   

## 2 网络模型
#### 2.1 I/O模型
- 《[web优化必须了解的原理之I/o的五种模型和web的三种工作模式](http://blog.51cto.com/litaotao/1289790)》
   * I/O步骤：进程向操作系统请求数据 ;操作系统把外部数据加载到内核的缓冲区中;操作系统把内核的缓冲区拷贝到进程的缓冲区 ;进程获得数据完成自己的功能 ;
   * 五种I/O模型：阻塞I/O，非阻塞I/O，I/O复用、事件(信号)驱动I/O、异步I/O，前四种I/O属于同步操作，I/O的第一阶段不同、第二阶段相同，最后的一种则属于异步操作。
   * 三种 Web Server 工作方式：Prefork(多进程)、Worker方式(线程方式)、Event方式。
- 《[一文读懂I/O复用](https://blog.csdn.net/wangxindong11/article/details/78591308)》
   * 第一阶段select阻塞，第二阶段recvfrom阻塞
#### 2.2 select，poll，epoll
- 《[select、poll、epoll之间的区别总结[整理]](http://www.cnblogs.com/Anker/p/3265058.html)》《[select、poll、epoll区别](http://xingyunbaijunwei.blog.163.com/blog/static/76538067201241685556302/)》《[epoll使用详解](https://www.cnblogs.com/fnlingnzb-learner/p/5835573.html)》
   * 每次调用select，都需要把fd集合从用户态拷贝到内核态，这个开销在fd很多时会很大
   * 同时每次调用select都需要在内核遍历传递进来的所有fd，这个开销在fd很多时也很大
   * select支持的文件描述符数量太小了，默认是1024

#### 2.3 BIO NIO AIO
- 《[BIO与NIO、AIO的区别](https://blog.csdn.net/skiof007/article/details/52873421)》   
- 《[两种高效的服务器设计模型：Reactor和Proactor模型](https://blog.csdn.net/u013074465/article/details/46276967)》  

# 中间件 
## 1 Web Server
#### 1.1 nginx
- 《[nginx入门教程](http://tengine.taobao.org/book/chapter_02.html)》《[nginx中文文档](http://www.nginx.cn/doc/)》
   * 进程模型：一个master进程，管理多个worker进程。一个请求只能一个worker进程处理。worker进程抢acceptMutex锁,然后接受请求，处理请求，返回请求，断开连接
   * 事件模型：高并发，异步非阻塞事件机制（例如epoll，循环处理准备好的事件请求，IO事件没有准备好就放回epoll)。
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


- 《[一个HTTP请求过程](https://blog.csdn.net/zhouziyu2011/article/details/60468703)》《[HTTP请求过程](https://blog.csdn.net/qq_37187976/article/details/79167168)》
   * DNS域名解析；TCP连接；接受请求；处理请求；返回响应；关闭TCP；记录日志；解析HTML展示页面；


#### 1.2 apache
- 《[apache与nginx原理详解及对比](https://blog.csdn.net/yf3585595511/article/details/54933646)》《[apache原理](https://www.cnblogs.com/ghosth/p/7502197.html)》《[apache与nginx对比及优缺点](https://www.cnblogs.com/cunkouzh/p/5410154.html)》
   * apache一个子进程处理一个请求。nginx一个子进程通过异步非阻塞事件模型可以处理多个请求，所以并发能力强。
   * nginx轻量级，占用资源少。反向代理负载均衡。
   * apache模块丰富
  
- 《[apache两种工作模式](https://blog.csdn.net/wangyunzhong123/article/details/77469643)》
   * perfork MPM多个进程，无线程设计，每个进程处理一个请求。worker MPM一个进程下有多个线程。每个线程处理一个请求。

- 《[Nginx+Apache动静分离](https://www.cnblogs.com/mangood/p/6048720.html)》
   * nginx处理静态请求，apache处理动态请求

#### 1.3 OpenResty
- 《[浅谈OpenResty](http://www.linkedkeeper.com/detail/blog.action?bid=1034)》 《[官方网站](http://openresty.org/cn/)》
   *  通过lua模块在nginx上开发。黑名单，防刷，限流等
- 《[openResty作者agentzh的Nginx教程](https://openresty.org/download/agentzh-nginx-tutorials-zhcn.html)》
- 《[跟着开涛学openResty](http://jinnianshilongnian.iteye.com/blog/2190344)》
   * TODO 《亿级流量网站》

## 2 缓存
- 《[HTTP缓存机制](https://www.cnblogs.com/chenqf/p/6386163.html)》
   * TODO
  
## 3 消息队列
- 《[消息队列简介及好处](https://yq.aliyun.com/articles/606240)》
   * 提高系统响应速度、提高系统稳定性、异步化、解耦、消除峰值
- 《[Kafka、RabbitMQ、RocketMQ等消息中间件的对比 —— 消息发送性能和区别](https://blog.csdn.net/yunfeng482/article/details/72856762)》
   * Push方式：优点是可以尽可能快地将消息发送给消费者，缺点是如果消费者处理能力跟不上，消费者的缓冲区可能会溢出。 RabbitMQ 消费者默认是推模式（也支持拉模式）。
   * Pull方式：优点是消费端可以按处理能力进行拉去，缺点是会增加消息延迟。 Kafka 默认是拉模式。
- 《[消息顺序和重复问题](https://www.jianshu.com/p/453c6e7ff81c)》

- 《[RabbitMQ应用场景以及基本原理介绍](https://blog.csdn.net/whoamiyang/article/details/54954780)》《[消息队列之 RabbitMQ](https://www.jianshu.com/p/79ca08116d57)》- 《[RabbitMQ之消息确认机制（事务+Confirm）](https://blog.csdn.net/u013256816/article/details/55515234)》
   * 支持事务，推拉模式都是支持、适合需要可靠性消息传输的场景。

- 《[RocketMQ 实战之快速入门](https://www.jianshu.com/p/824066d70da8)》《[RocketMQ源码分析](http://www.iocoder.cn/categories/RocketMQ/?vip&architect-awesome)》
   * Java实现，推拉模式都是支持，吞吐量逊于Kafka。可以保证消息顺序。
   
- 《[ActiveMQ消息队列介绍](https://www.cnblogs.com/wintersun/p/3962302.html)》
   * 纯Java实现，兼容JMS，可以内嵌于Java应用中。
   
- 《[Kafka深度解析，众人推荐，精彩好文](https://www.cnblogs.com/wintersun/p/3962302.html)》《[Kafka分区机制介绍与示例](http://lxw1234.com/archives/2015/10/538.htm)》
   * 高吞吐量、采用拉模式。适合高IO场景，比如日志同步。

- 《[Redis用作消息队列](https://blog.csdn.net/qq_34212276/article/details/78455004)》
   * 生产者、消费者模式完全是客户端行为，list 和 拉模式实现，阻塞等待采用 blpop 指令。

## 4 定时调度
   * TODO
  
## 5 API网关
   * TODO


# 设计模式
TODO
# 操作系统
## 1 计算机原理
## 2 CPU
## 3 进程
## 4 线程
## 5 协程
## 6 Linux
TODO


   

# 架构&运维&统计&技术支持
## 1 日志服务
- 《[ELKB日志收集系统搭建](http://cjting.me/misc/build-log-system-with-elkb/)》《[ELK日志收集系统搭建](https://blog.csdn.net/lzw_2006/article/details/51280058)》《[日志收集系统](https://www.cnblogs.com/beginmind/p/6058194.html)》
   * Beats日志搬运工。安装在每台需要收集日志的服务器上，将日志发送给Logstash进行处理。[Filebeat配置安装](https://www.cnblogs.com/jingmoxukong/p/8185321.html)
   * LogStash把没条日志解析成每个字段[LogStash配置安装](https://www.cnblogs.com/moonlightL/p/7760512.html)
   * ElasticSearch全文搜索日志 [ElasticSearch全文搜索引擎](http://www.ruanyifeng.com/blog/2017/08/elasticsearch.html)《[索引搜索查询及增删改查](https://www.cnblogs.com/pilihaotian/p/5830754.html)》《[ElasticSearch权威指南](https://es.xiaoleilu.com/010_Intro/00_README.html)》
   * Kibana是ElasticSearch全文搜索图形化界面 [Kibana安装](https://www.cnblogs.com/moonlightL/p/7764919.html)
-  《[FileBeat+kafka进行日志实时传输](https://blog.csdn.net/paicMis/article/details/78888750)》
   * kafka[教程](https://www.w3cschool.cn/apache_kafka/apache_kafka_introduction.html)
- logrotate日志切割
   * [linux日志切割](https://blog.csdn.net/junli_chen/article/details/77193438)
   * [日志切割配置]( http://blog.51cto.com/wn2100/2074048)



## 2 微服务架构
- 《[微服务架构设计](https://www.cnblogs.com/wintersun/p/6219259.html)》《[微服务架构选型](http://www.infoq.com/cn/articles/micro-service-technology-stack)》
   * API网关，RESTFUL，微服务
- 《[RESTFUL API架构](https://www.runoob.com/w3cnote/restful-architecture.html)》
- 《[API网关](https://www.cnblogs.com/savorboard/p/api-gateway.html)》
   * [kong](git地址https://github.com/Kong/kong) 
-  项目实例
   * TODO
   
## 3 RPC
- 《[RPC概述](https://blog.csdn.net/top_code/article/details/54615853)》
   * 远程调用程序并返回
- 《[分布式RPC框架性能大比拼dubbo、motan、rpcx、gRPC、thrift的性能比较](https://blog.csdn.net/testcs_dn/article/details/78050590)》
- 《[Dubbo官网](http://dubbo.apache.org/en-us/)》《[Dubbo实现原理简单介绍](https://www.cnblogs.com/steven520213/p/7606598.html)》
- 《[Thrift官网](http://thrift.apache.org/)》《[Thrift详解](https://blog.csdn.net/kesonyk/article/details/50924489)》
- 《[Grpc官网](https://grpc.io/)》《[RPC原理介绍](https://www.cnblogs.com/LBSer/p/4853234.html)》

   


## 4 docker
- 《[几张图帮你理解 docker 基本原理及快速入门](https://www.cnblogs.com/SzeCheng/p/6822905.html)》
   * 镜像，容器。一个镜像可以建立多个容器，一个容器最多127层。
- 《[docker核心技术与实现](https://draveness.me/docker)》
- 《[docker教程](http://www.runoob.com/docker/docker-tutorial.html)》
-  项目实例
   * TODO





# 安全

## 1 堡垒机（跳板机）
- 《[搭建简易堡垒机](http://blog.51cto.com/zero01/2062618)》



 
# 并发
## 1 并发概念
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

   
## 2 锁


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




# 数据库

- 《[三大范式五大约束](https://www.cnblogs.com/waj6511988/p/7027127.html)》
- 《[数据库事务特性](https://blog.csdn.net/u012440687/article/details/52116108)》
- 《[数据库事务隔离级别](https://blog.csdn.net/qq_33290787/article/details/51924963)》
- 《[InnoDB幻读问题](http://blog.sina.com.cn/s/blog_499740cb0100ugs7.html)》

# 数据结构&算法
## 1 数据结构
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
## 2 算法
TODO


#  测试
## 1 压力测试
- 《[apache ab测试使用指南](https://blog.csdn.net/blueheart20/article/details/52170790)》
   * too many open files (解决方案：ulimit -n; nginx配置events同级 worker_rlimit_nofile 15360;)
   * apr_socket_recv:Operation timed out  (解决方法：ab加上-k 开启keepAlive)
   * apr_socket_connect(): Operation already in progress (解决方法：apache/conf/extra/httpd-mpm.conf  修改 ThreadsPerChild)
   * apr_socket_recv: Connection reset by peer (54)
   * setsockopt(TCP_NODELAY) failed (22: Invalid argument) while keepalive
 
 - 《[大型网站压力测试及优化方案](https://www.cnblogs.com/binyue/p/6141088.html)》
    *  jmeter
    * 《[全链路压测](https://www.jianshu.com/p/27060fd61f72)》
    *  资源服务监控《[20个命令行工具监控](http://blog.jobbole.com/96846/)》

