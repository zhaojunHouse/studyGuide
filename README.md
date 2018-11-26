<!-- TOC -->
- [数据结构&算法](#数据结构算法)
    - [1 数据结构](#1-数据结构)
    - [2 算法](#2-算法)
- [操作系统](#操作系统)
    - [1 计算机原理](#1-计算机原理)
    - [2 CPU](#2-cpu)
    - [3 进程](#3-进程)
    - [4 线程](#4-线程)
    - [5 协程](#5-协程)
    - [6 Linux](#6-linux)
- [网络](#网络)
    - [1 网络协议](#1-网络协议)  
            &nbsp;&nbsp; [1.1  TCP/IP](#11-tcpip)  
            &nbsp;&nbsp; [1.2  HTTP](#12-http)  
            &nbsp;&nbsp; [1.3  LOT七大通信协议](#13-lot七大通信协议)
    - [2 网络模型](#2-网络模型)  
            &nbsp;&nbsp; [2.1  I/O模型](#21-io模型)  
            &nbsp;&nbsp; [2.2  select，poll，epoll](#22-selectpollepoll)  
            &nbsp;&nbsp; [2.3  BIO NIO AIO](#23-bio-nio-aio)  
            &nbsp;&nbsp; [2.4  kqueue](#24-kqueue)  
            &nbsp;&nbsp; [2.5  长连接短链接](#25-长连接短链接)
- [设计模式](#设计模式)
   - [1 设计模式六大原则](#1-设计模式六大原则)    
   - [2 23种设计模式](#2-23种设计模式)    
  
- [设计思想&开发模式](#设计思想&开发模式)
- [中间件](#中间件)  
    - [1 Web Server](#1-web-server)  
            &nbsp;&nbsp; [1.1  nginx](#11-nginx)  
            &nbsp;&nbsp; [1.2  apache](#12-apache)  
            &nbsp;&nbsp; [1.3  OpenResty](#13-openresty)
    - [2 缓存](#2-缓存)
    - [3 消息队列](#3-消息队列)
    - [4 定时调度](#4-定时调度)
    - [5 API网关](#5-api网关)
- [数据库](#数据库)
- [测试](#测试)
    - [1 压力测试](#1-压力测试)    
- [并发](#并发)
    - [1 并发概念](#1-并发概念)    
            &nbsp;&nbsp; [1.1  并发](#11-并发)     
            &nbsp;&nbsp; [1.2  多线程](#12-多线程)  
            &nbsp;&nbsp; [1.3  线程安全](#13-线程安全)     
    - [2 锁](#2-锁)    
            &nbsp;&nbsp; [2.1  悲观锁](#21-悲观锁)     
            &nbsp;&nbsp; [2.2  乐观锁ABA问题](#22-乐观锁)     
            &nbsp;&nbsp; [2.3  死锁](#23-死锁)     
- [性能](#性能)
    - [1 性能优化方法](#1-性能优化方法) 
    - [2 容量评估](#2-容量评估) 
    - [3 CDN加速](#3-cdn加速) 
    - [4 连接池](#4-连接池) 
    - [5 优化工具](#5-优化工具) 
- [安全](#安全)
    - [1 堡垒机（跳板机）](#1-堡垒机跳板机) 
- [运维&统计&技术支持](#运维统计技术支持)  
    - [1 常规监控](#1-常规监控)    
             &nbsp;&nbsp; [1.1 命令行监控工具](#11-命令行监控工具)         
    - [2 日志服务](#2-日志服务)
    - [3 RPC](#3-rpc)
    - [4 docker](#4-docker) 
    - [5 APM](#5-APM) 
    - [6 统计分析](#6-统计分析) 
- [架构&框架](#架构框架)    
    - [1 架构](#1-架构)         
             &nbsp;&nbsp; [1.1  架构师修炼图](#11-架构师修炼图)        
             &nbsp;&nbsp; [1.2  微服务架构](#12-微服务架构)      
- [分布式设计](#分布式设计)
    - [1 扩展性设计](#1-扩展性设计)       
            &nbsp;&nbsp; [1.1 分布式&异步](#11-分布式异步)        
            &nbsp;&nbsp; [1.2 分布式之数据切分](#12-分布式之数据切分)         
            &nbsp;&nbsp; [1.3 分布式事务一致性](#13-分布式事务一致性)      
    - [2 稳定性&高可用](#2-稳定性高可用)  
           &nbsp;&nbsp; [2.1 负载均衡](#21-负载均衡)      
           &nbsp;&nbsp; [2.2 限流](#22-限流)  
           &nbsp;&nbsp; [2.3 容灾 ](#23-容灾)  
           &nbsp;&nbsp; [2.4 平滑重启 ](#24-平滑重启)  
    - [3 数据库扩展](#3-数据库扩展)   
            &nbsp;&nbsp; [3.1 读写分离](#31-读写分离)      
            &nbsp;&nbsp; [3.2 分片模式](#32-分片模式)      
    - [4 服务治理](#4-服务治理)   
            &nbsp;&nbsp; [4.1 服务发现](#41-服务发现)  
            &nbsp;&nbsp; [4.2 服务路由](#42-服务路由)  
    - [5 分布式一致性](#5-分布式一致性)   
            &nbsp;&nbsp; [5.1 CAP与BASE理论](#51-CAP与BASE理论)  
            &nbsp;&nbsp; [5.2 分布式锁](#52-分布式锁)  
            &nbsp;&nbsp; [5.3 分布式一致性算法](#53-分布式一致性算法)  
            &nbsp;&nbsp; [5.4 幂等](#54-幂等)  
    - [6 分布式文件系统](#6-分布式文件系统)   
    - [7 唯一ID生成](#7-唯一ID生成)  
    - [8 一致性HASH算法](#8-一致性HASH算法)  
- [大数据](#大数据)
- [搜索引擎](#搜索引擎)
- [项目管理](#项目管理)
- [资讯&技术资源](#资讯技术资源)  
    - [1 行业资讯](#1-行业资讯)
    - [2 博客](#2-博客)

<!-- /TOC -->
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

# 操作系统
## 1 计算机原理
- 《[操作系统基础知识——操作系统的原理，类型和结构](https://segmentfault.com/a/1190000003692840)》

## 2 CPU
- 《[操作系统学习笔记-CPU](http://blog.51cto.com/10827670/2170890)》
- 《[操作系统之CPU调度](https://blog.csdn.net/qq_36838191/article/details/83383226)》

## 3 进程
- 《[多线程与多进程的概念和实例讲解](https://www.2cto.com/kf/201802/719224.html)》

## 4 线程
- 《[线程的生命周期及状态转换详解](https://blog.csdn.net/asdf_1024/article/details/78978437)》

## 5 协程
- 《[Golang 之协程详解](https://www.cnblogs.com/liang1101/p/7285955.html)》

## 6 Linux
- 《[Linux 命令大全](http://www.runoob.com/linux/linux-command-manual.html)》

   
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

#### 2.4 kqueue
- 《[kqueue用法简介](http://www.cnblogs.com/luminocean/p/5631336.html)》

#### 2.5 长连接短链接
- 《[长连接短链接](https://www.cnblogs.com/pangguoping/p/5571422.html)》



# 设计模式
## 1 设计模式六大原则
- 《[设计模式的六大原则](https://blog.csdn.net/q291611265/article/details/48465113)》
   * 开闭原则：对扩展开放,对修改关闭，多使用抽象类和接口。
   * 里氏替换原则：基类可以被子类替换，使用抽象类继承,不使用具体类继承。
   * 依赖倒转原则：要依赖于抽象,不要依赖于具体，针对接口编程,不针对实现编程。
   * 接口隔离原则：使用多个隔离的接口,比使用单个接口好，建立最小的接口。
   * 迪米特法则：一个软件实体应当尽可能少地与其他实体发生相互作用，通过中间类建立联系。
   * 合成复用原则：尽量使用合成/聚合,而不是使用继承。
## 2 23种设计模式
- 《[23种设计模式](http://www.runoob.com/design-pattern/design-pattern-tutorial.html)》
- 《[23种设计模式全解析](https://www.cnblogs.com/susanws/p/5510229.html)》
- 《[23种设计模式类图和示例](https://github.com/ToryZhou/design-pattern)》





# 设计思想&开发模式
TODO

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


# 数据库

- 《[三大范式五大约束](https://www.cnblogs.com/waj6511988/p/7027127.html)》
- 《[数据库事务特性](https://blog.csdn.net/u012440687/article/details/52116108)》
- 《[数据库事务隔离级别](https://blog.csdn.net/qq_33290787/article/details/51924963)》
- 《[InnoDB幻读问题](http://blog.sina.com.cn/s/blog_499740cb0100ugs7.html)》



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

 
# 并发
## 1 并发概念
#### 1.1 并发
- 《[为什么需要并发？](https://juejin.im/post/5ae6c3ef6fb9a07ab508ac85)》
  * 充分利用多核CPU的计算能力
  * 业务拆分，提升性能
   
- 《[并发与并行，并发与异步区别](https://www.cnblogs.com/xqn2017/p/8029472.html)》
  * 并发：通过切换时间片的方式多个任务交替进行；并行：真正意义上的“同时进行”
   
- 《[高并发实现需要考虑什么？](https://www.cnblogs.com/PerkinsZhu/p/7242247.html)》 
  * 硬件，网络，架构设计（各种分布式，缓存，中间件，异步），开发语言，数据结构，算法，数据库优化，代码层（多线程编程，同步异步）
#### 1.2 多线程
- 《[高并发实现中一种方式：多线程开发及优缺点](https://www.cnblogs.com/xiejw/p/5259437.html)》
  *  优点：耗时任务放入子线程防止阻塞，提升响应速度；充分利用CPU，提高性能；改善程序结构；
  *  缺点：频繁切换上下文；线程安全问题；

#### 1.3 线程安全
- 《[线程安全问题及解决方案](https://www.cnblogs.com/zhanht/p/5450325.html)》
  * 产生原因：CPU切换时间片执行多线程产生竞态条件
  * 原子性，可见性，时序性；原子性：要么都执行，要么都不执行。a++是先查，后加，不是原子操作。
  * 加锁

   
## 2 锁


- 《[锁的分类](https://www.cnblogs.com/qifengshi/p/6831055.html)》
   * 乐观锁&悲观锁，互斥锁/写锁，共享锁/读锁，公平锁/非公平锁，重入锁/非重入锁，分段锁等。
#### 2.1 悲观锁   
- 《[悲观锁与乐观锁](http://www.importnew.com/20472.html)》
   * select for update     
#### 2.2 乐观锁
- 《[乐观锁](http://www.importnew.com/20472.html)》
   * 是一种思想，提交更新时对冲突检测，让用户决定重试或是其他。
   * ABA问题，解决方法对每次更新加上版本号
  
 
- 《[对mysql乐观锁、悲观锁、共享锁、排它锁、行锁、表锁概念的理解](https://blog.csdn.net/puhaiyang/article/details/72284702)》
   * 对于update,insert,delete语句会自动加排它锁,只共享读。
#### 2.3 死锁   
- 《[select.. for update导致的数据库死锁分析](https://www.cnblogs.com/Lawson/p/5008741.html)》
   * mysql的innodb存储引擎实务锁虽然是锁行，但它内部是锁索引的。
   * 锁相同数据的不同索引条件可能会引起死锁。

- 《[常见的死锁原因与解决方法](https://www.cnblogs.com/zejin2008/p/5262751.html)》《[delete造成的死锁](http://hedengcheng.com/?p=844)》
   
- 《[mysql加锁分析](http://hedengcheng.com/?p=771)》  《[聚簇索引与非聚簇索引](http://www.cnblogs.com/tuyile006/archive/2009/08/28/1555615.html)》《[MySQL索引原理及慢查询优化](https://tech.meituan.com/mysql_index.html)》



# 性能
## 1 性能优化方法
- 《[15天的性能优化工作，5方面的调优经验](https://blog.csdn.net/huangwenyi1010/article/details/72673447?ref=myread)》
    * 代码层面、业务层面、数据库层面、服务器层面、前端优化。  
- 《[系统优化的几个方面](https://blog.csdn.net/tenglizhe/article/details/44563135)》

## 2 容量评估 
- 《[互联网架构，如何进行容量设计？](https://mp.weixin.qq.com/s?__biz=MjM5ODYxMDA5OQ==&mid=2651959542&idx=1&sn=2494bbea9a855e0e1c3ccd6d2562a600&scene=21#wechat_redirect)》
   * 总访问量，平均访问量，高峰QPS，压测获取单机极限QPS，机器配置。
- 《[互联网性能与容量评估的方法论和典型案例](https://blog.csdn.net/u012528360/article/details/70054156)》

## 3 CDN加速
- 《[CDN加速原理](https://www.cnblogs.com/wxiaona/p/5867685.html)》
- 《[国内有哪些比较好的 CDN？](https://www.zhihu.com/question/20536932)》

## 4 连接池
- 《[golang通用连接池实现](https://blog.csdn.net/mengxinghuiku/article/details/79730871)》

## 5 优化工具
- 《[php性能优化工具xhprof安装](https://www.cnblogs.com/etata/p/5177844.html)》
- 《[golang性能优化工具](https://blog.csdn.net/WaltonWang/article/details/54019891)》



# 安全

## 1 堡垒机（跳板机）
- 《[搭建简易堡垒机](http://blog.51cto.com/zero01/2062618)》



   

# 运维&统计&技术支持
## 1 常规监控
- 《[【分享】腾讯业务系统监控的修炼之路](https://blog.csdn.net/enweitech/article/details/77849205)》
   * 监控的方式：主动、被动、旁路(比如舆情监控)
   * 监控类型： 基础监控、服务端监控、客户端监控、 监控、用户端监控
   * 监控的目标：全、快、准
   * 核心指标：请求量、成功率、耗时
- 《[开源还是商用？十大云运维监控工具横评](https://www.oschina.net/news/67525/monitoring-tools)》
   * Zabbix、Nagios、Ganglia、Zenoss、Open-falcon、监控宝、 360网站服务监控、阿里云监控、百度云观测、小蜜蜂网站监测等。
- 《[监控报警系统搭建及二次开发经验](http://developer.51cto.com/art/201612/525373.htm)》
#### 1.1 命令行监控工具
- 《[20个命令行工具监控 Linux 系统性能](http://blog.jobbole.com/96846/)》
- 《[命令行工具监控](https://coderxing.gitbooks.io/architecture-evolution/di-er-pian-ff1a-feng-kuang-yuan-shi-ren/44-an-quan-yu-yun-wei/445-fu-wu-qi-zhuang-tai-jian-ce/4451-ming-ling-xing-gong-ju.html)》

## 2 日志服务
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
   
## 5 APM 
#### application performance management
- 《[Dapper，大规模分布式系统的跟踪系统](http://bigbully.github.io/Dapper-translation/)》
- 《[OpenTracing官方标准-中文版](https://github.com/opentracing-contrib/opentracing-specification-zh)》
- 《[开源 APM 技术选型与实战](https://www.infoq.cn/article/apm-Pinpoint-practice)》


## 6 统计分析
- 《[流量统计的基础：埋点](https://zhuanlan.zhihu.com/p/25195217)》
- 《[APP埋点常用的统计工具、埋点目标和埋点内容](https://www.25xt.com/company/17066.html)》
- 《[美团点评前端无痕埋点实践](https://tech.meituan.com/mt_mobile_analytics_practice.html)》




# 架构&框架
## 1 架构
#### 1.1 架构师修炼图
- 《[架构师技能修炼图](https://blog.csdn.net/hemin1003/article/details/53633926)》
#### 1.2 微服务架构
- 《[微服务架构设计](https://www.cnblogs.com/wintersun/p/6219259.html)》《[微服务架构选型](http://www.infoq.com/cn/articles/micro-service-technology-stack)》
   * API网关，RESTFUL，微服务
- 《[RESTFUL API架构](https://www.runoob.com/w3cnote/restful-architecture.html)》
- 《[API网关](https://www.cnblogs.com/savorboard/p/api-gateway.html)》
   * [kong](git地址https://github.com/Kong/kong) 
-  项目实例
   * TODO
 


# 分布式设计
## 1 扩展性设计
#### 1.1 分布式&异步
- 《[说说如何实现可扩展性的大型网站架构](https://blog.csdn.net/deniro_li/article/details/78458306)》
   * 分布式+异步任务
   * 分布式要求：负载均衡，失效转移，高效远程调用，整合异构系统，版本控制，实时监控
#### 1.2 分布式之数据切分 
- 《[可扩展性设计之数据切分](https://yq.aliyun.com/articles/38119)》
   * 垂直切分，水平切分，联合切分
   * 垂直切分缺点： join表分页搜索需要程序完成；事务处理复杂；
   * 水平切分缺点： 维护数据变复杂；切分无具体标准；
   * 切分与整合带来的问题：分布式事务；跨节点join（全局表，冗余，数据组装）；跨节点合并分页排序（先排序分页，再数据组装）；
#### 1.3 分布式事务一致性  
- 《[分布式事务一致性问题](https://blog.csdn.net/dinglang_2009/article/details/53195835)》《[分库分表事务一致性](https://www.infoq.cn/article/solution-of-distributed-system-transaction-consistency)》《[保证分布式系统数据一致性的6种方案](https://weibo.com/ttarticle/p/show?id=2309403965965003062676)》
   * 提供回滚接口
   * 《[本地消息表](https://segmentfault.com/a/1190000012415698)》
   * 非事务MQ
   * 事务MQ
   

## 2 稳定性&高可用
- 《[系统设计：关于高可用系统的一些技术方案](https://blog.csdn.net/hustspy1990/article/details/78008324)》
   * 可扩展：水平扩展、垂直扩展。 通过冗余部署，避免单点故障。一主多从。
   * 隔离：避免单一业务占用全部资源。避免业务之间的相互影响 2. 机房隔离避免单点故障。
   * 解耦：降低维护成本，降低耦合风险。减少依赖，减少相互间的影响。
   * 限流：遇到突发流量时，保证系统稳定。
   * 降级：紧急情况下释放非核心功能的资源。牺牲非核心业务，保证核心业务的高可用。
   * 熔断：异常情况超出阈值进入熔断状态，快速失败。减少不稳定的外部依赖对核心服务的影响。
   * 自动化测试：通过完善的测试，减少发布引起的故障。
   * 灰度发布：灰度发布是速度与安全性作为妥协，能够有效减少发布故障。
   
- 《[关于高可用的系统](https://coolshell.cn/articles/17459.html)》
   * 数据不丢，就必需要持久化；服务高可用，就必需要有备用；复制，就会有数据一致性
#### 2.1 负载均衡
- 《[软/硬件负载均衡产品 你知多少？](https://www.cnblogs.com/lcword/p/5773296.html)》《[转！！负载均衡器技术Nginx和F5的优缺点对比](https://www.cnblogs.com/wuyun-blog/p/6186198.html)》
   * 硬件，软件负载均衡
- 《[负载均衡的几种算法](https://www.cnblogs.com/tianzhiliang/articles/2317808.html)》   
   * 轮流；比例；优先权；   最小连接；响应最快；服务器性能模式；观察模式 请求数响应时间；
- 《[DNS做负载均衡](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/611-dns-fang-shi.html)》
   * 配置简单；生效慢。
- 《[NGINX做负载均衡](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/613-nginx-fu-zai-jun-heng.html)》
   * 成本低。主要适用web应用。
- 《[借助LVS+Keepalived实现负载均衡](https://www.cnblogs.com/edisonchou/p/4281978.html)》
   * 支持到OSI四层，性能较高
- 《[HAProxy用法详解 全网最详细中文文档](http://www.ttlsa.com/linux/haproxy-study-tutorial/)》《[Haproxy+Keepalived+MySQL实现读均衡负载](http://blog.itpub.net/25704976/viewspace-1319781/)》《[rabbitmq+haproxy+keepalived实现高可用集群搭建](https://www.cnblogs.com/lylife/p/5584019.html)》
   * HAProxy性能和LVS差不多。

#### 2.2  限流
- 《[谈谈高并发系统的限流](https://www.cnblogs.com/haoxinyue/p/6792309.html)》
   * 计数器：通过滑动窗口计数器，控制单位时间内的请求次数，简单粗暴。
   * 漏桶算法：固定容量的漏桶，漏桶满了就丢弃请求，比较常用。
   * 令牌桶算法：固定容量的令牌桶，按照一定速率添加令牌，处理请求前需要拿到令牌，拿不到令牌则丢弃请求，或进入丢队列，可以通过控制添加令牌的速率，来控制整体速度。Guava 中的 RateLimiter 是令牌桶的实现。
   * Nginx 限流：通过 limit_req 等模块限制并发连接数。
   
####  2.3 容灾
- 《[防雪崩利器：熔断器 Hystrix 的原理与使用](https://segmentfault.com/a/1190000005988895)》
   * 雪崩效应原因：硬件故障、程序Bug、缓存击穿、用户大量请求。
   * 雪崩的对策：限流、改进缓存模式(缓存预加载、同步调用改异步)、自动扩容、降级。
   * Hystrix设计原则：
   * 资源隔离：Hystrix通过将每个依赖服务分配独立的线程池进行资源隔离, 从而避免服务雪崩。
   * 熔断开关：服务的健康状况 = 请求失败数 / 请求总数，通过阈值设定和滑动窗口控制开关。
   * 命令模式：通过继承 HystrixCommand 来包装服务调用逻辑。
   
- 《[缓存穿透，缓存击穿，缓存雪崩解决方案分析](https://blog.csdn.net/zeb_perfect/article/details/54135506)》
   * 主要策略：失效瞬间：单机使用锁；使用分布式锁；不过期；
   * 热点数据：热点数据单独存储；使用本地缓存；分成多个子key；

- 《[“异地多活”多机房部署经验谈](http://dc.idcquan.com/ywgl/71559.shtml)》《[异地多活（异地双活）实践经验](https://blog.csdn.net/jeffreynicole/article/details/48135093)》《[依赖治理、灰度发布、故障演练，阿里电商故障演练系统的设计与实战经验](https://mp.weixin.qq.com/s?__biz=MjM5MDE0Mjc4MA==&mid=2650996320&idx=1&sn=0ed3be190bbee4a9277886ef88cbb2e5)》

####  2.4 平滑重启
- 《[Golang学习--平滑重启](https://www.cnblogs.com/CraryPrimitiveMan/p/8560839.html)》


## 3 数据库扩展
#### 3.1 读写分离
- 《[【大型网站技术实践】初级篇：搭建MySQL主从复制经典架构](https://www.cnblogs.com/edisonchou/p/4133148.html)》《[MySQL主从复制(Master-Slave)实践](https://www.cnblogs.com/gl-developer/p/6170423.html)》
   * 一主n从；自动同步数据。
   
- 《[Haproxy+多台MySQL从服务器(Slave) 实现负载均衡](https://blog.csdn.net/nimasike/article/details/48048341)》

- 《[DRBD+Heartbeat+Mysql高可用读写分离架构](https://www.cnblogs.com/zhangsubai/p/6801764.html)》
  
- 《[MySQL Cluster 集群解决方案](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/62-ke-kuo-zhan-de-shu-ju-ku-jia-gou/621-gao-ke-yong-mysql-de-ji-zhong-fang-an/6214-mysql-cluster-fang-an.html)》  

#### 3.2 分片模式
- 《[分库分表需要考虑的问题及方案](https://www.jianshu.com/p/32b3e91aa22c)》
- 《[mysql分表和表分区详解](https://www.2cto.com/database/201503/380348.html)》

## 4 服务治理
#### 4.1 服务发现
- 《[永不失联！如何实现微服务架构中的服务发现](https://blog.csdn.net/jiaolongdy/article/details/51188798)》
   * 客户端服务发现模式：客户端直接查询注册表，同时自己负责负载均衡。Eureka 采用这种方式。
   * 服务器端服务发现模式：客户端通过负载均衡查询服务实例。
- 《[SpringCloud服务注册中心比较:Consul vs Zookeeper vs Etcd vs Eureka](https://blog.csdn.net/u010963948/article/details/71730165)》
- 《[基于Zookeeper的服务注册与发现](http://mobile.51cto.com/news-502394.htm)》

#### 4.2 服务路由
- 《[分布式服务框架学习笔记4 服务路由](https://blog.csdn.net/xundh/article/details/59492750)》


## 5 分布式一致性
#### 5.1 CAP与BASE理论
- 《[从分布式一致性谈到CAP理论、BASE理论](http://www.cnblogs.com/szlbm/p/5588543.html)》
   * 一致性分类：强一致(立即一致)；弱一致(可在单位时间内实现一致，比如秒级)；最终一致(弱一致的一种，一定时间内最终一致)
   * CAP：一致性、可用性、分区容错性(网络故障引起)
   * BASE：Basically Available（基本可用）、Soft state（软状态）和Eventually consistent（最终一致性）
   * BASE理论的核心思想是：即使无法做到强一致性，但每个应用都可以根据自身业务特点，采用适当的方式来使系统达到最终一致性。

#### 5.2 分布式锁
- 《[分布式锁的几种实现方式~](http://www.hollischuang.com/archives/1716)》
   * 数据库；缓存；zookeeper；
- 《[基于Zookeeper的分布式锁](https://www.tuicool.com/articles/VZJr6fY)》
- 《[jedisLock—redis分布式锁实现](https://www.cnblogs.com/0201zcr/p/5942748.html)》
   * 利用setnx
- 《[Memcached 和 Redis 分布式锁方案](https://blog.csdn.net/albertfly/article/details/77412333)》
   * 利用 memcached 的 add（有别于set）操作，当key存在时，返回false。
#### 5.3 分布式一致性
#### 5.4 幂等
- 《[分布式系统---幂等性设计](https://www.cnblogs.com/wxgblogs/p/6639272.html)》
- 《[传统事务与柔性事务](https://www.jianshu.com/p/ab1a1c6b08a1)》
   
## 6 分布式文件系统
- 《[分布式文件系统：原理、问题与方法(转)](https://www.cnblogs.com/rainy-shurun/p/5335747.html)》

## 7 唯一ID生成 
- 《[高并发分布式系统中生成全局唯一Id汇总](https://www.cnblogs.com/baiwa/p/5318432.html)》
- 《[TDDL 在分布式下的SEQUENCE原理](https://blog.csdn.net/hdu09075340/article/details/79103851)》

## 8 一致性HASH算法
- 《[什么是一致性Hash算法？](https://blog.csdn.net/bntX2jSQfEHy7/article/details/79549368)》


# 大数据
TODO


# 搜索引擎
TODO

# 项目管理
TODO

# 资讯&技术资源
## 1 行业资讯
- 《[36Kr](https://36kr.com/)》
- 《[TeahWeb](http://www.techweb.com.cn/)》
## 2 博客
- 《[阿里中间件团队](http://jm.taobao.org/)》
- 《[美团技术团队](https://tech.meituan.com/)》
