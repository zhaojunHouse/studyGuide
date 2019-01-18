<!-- TOC -->
- [数据结构&算法](#数据结构算法)
    - [1 复杂度分析](#1-复杂度分析)   
       &nbsp;&nbsp; [1.1  时间复杂度](#11-时间复杂度)  
       &nbsp;&nbsp; [1.2  时间复杂度量级](#12-时间复杂度量级)   
       &nbsp;&nbsp; [1.3  空间复杂度](#13-空间复杂度)   
       &nbsp;&nbsp; [1.4  最好最坏平均均摊时间复杂度](#14-最好最坏平均均摊时间复杂度) 
    - [2 数组](#2-数组)  
      &nbsp;&nbsp; [2.1  定义](#21-定义)   
      &nbsp;&nbsp; [2.2  插入删除](#22-插入删除)   
      &nbsp;&nbsp; [2.3  查找](#23-查找)   
    - [3 链表](#3-链表)  
      &nbsp;&nbsp; [3.1  插入删除](#31-插入删除)  
      &nbsp;&nbsp; [3.2  查找](#32-查找)   
      &nbsp;&nbsp; [3.3  单链&双向链表&循环链表](#33-单链双向链表循环链表)   
      &nbsp;&nbsp; [3.4  链表操作](#34-链表操作)      
    - [4 栈](#4-栈)  
      &nbsp;&nbsp; [4.1  入栈出栈](#41-时间复杂度)   
      &nbsp;&nbsp; [4.2  顺序栈&链式栈](#42-顺序栈链式栈)   
      &nbsp;&nbsp; [4.3  栈应用](#43-栈应用)     
    - [5 队列](#5-队列)   
      &nbsp;&nbsp; [5.1  队头队尾](#51-队头队尾)   
      &nbsp;&nbsp; [5.2  顺序队列&链式队列](#52-顺序队列链式队列)  
      &nbsp;&nbsp; [5.3  循环队列](#53-循环队列)   
      &nbsp;&nbsp; [5.4  阻塞队列&并发队列](#54-阻塞队列并发队列)    
    - [6 递归](#6-递归)   
      &nbsp;&nbsp; [6.1  堆栈溢出](#61-堆栈溢出)     
    - [7 排序](#7-排序)   
      &nbsp;&nbsp; [7.1  冒泡排序](#71-冒泡排序)     
      &nbsp;&nbsp; [7.2  插入排序](#72-插入排序)    
      &nbsp;&nbsp; [7.3  选择排序](#73-选择排序)    
      &nbsp;&nbsp; [7.4  归并排序](#74-归并排序)    
      &nbsp;&nbsp; [7.5  快排](#75-快排)   
      &nbsp;&nbsp; [7.6  桶排序](#76-桶排序)   
      &nbsp;&nbsp; [7.7  计数排序](#77-计数排序)   
      &nbsp;&nbsp; [7.8  基数排序](#78-基数排序)   
    - [8 二分查找](#8-二分查找)   
    - [9 跳表](#9-跳表)  
    - [10 散列表](#10-散列表)   
      &nbsp;&nbsp; [10.1  散列冲突](#101-散列冲突)  
      &nbsp;&nbsp; [10.2  设计工业级散列表](#102-设计工业级散列表)      
      &nbsp;&nbsp; [10.3  LRU缓存淘汰算法](#103-LRU缓存淘汰算法)     
      &nbsp;&nbsp; [10.4  Redis有序集合](#104-Redis有序集合)     
      &nbsp;&nbsp; [10.5  哈希算法](#105-哈希算法)     
      &nbsp;&nbsp; [10.6  一致性哈希算法](#106-一致性哈希算法)    
    - [11 树](#11-树)    
      &nbsp;&nbsp; [11.1  高度深度层](#111-高度深度层)  
      &nbsp;&nbsp; [11.2  二叉树&满二叉树&完全二叉树](#112-二叉树满二叉树完全二叉树)     
      &nbsp;&nbsp; [11.3  二叉查找树](#113-二叉查找树)     
      &nbsp;&nbsp; [11.4  红黑树](#114-红黑树)     
      &nbsp;&nbsp; [11.5  递归树](#115-递归树)    
      &nbsp;&nbsp; [11.6  B树&B+树](#116-B树B+树)    
    - [12 堆](#12-堆)     
    - [13 图](#13-图)    
    - [14 算法](#14-算法)    
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
            &nbsp;&nbsp; [2.1  本地缓存](#21-本地缓存)  
            &nbsp;&nbsp; [2.2  客户端缓存](#22-客户端缓存)  
            &nbsp;&nbsp; [2.3  服务端缓存](#23-服务端缓存)  
            &nbsp;&nbsp; [2.3.1  Web缓存](#231-Web缓存)  
            &nbsp;&nbsp; [2.3.2  Memcache](#232-Memcache)  
            &nbsp;&nbsp; [2.3.3  Redis](#233-Redis)  
            &nbsp;&nbsp; [2.3.4  Tair](#234-Tair)  
    - [3 消息队列](#3-消息队列)    
            &nbsp;&nbsp; [3.1 消息队列常见问题  ](#31-消息队列常见问题)    
            &nbsp;&nbsp; [3.2 RabbitMQ](#32-RabbitMQ)    
            &nbsp;&nbsp; [3.3 RocketMQ](#33-RocketMQ)    
            &nbsp;&nbsp; [3.4 ActiveMQ](#34-ActiveMQ)    
            &nbsp;&nbsp; [3.5 Kafka](#35-Kafka)    
            &nbsp;&nbsp; [3.6 Redis](#36-Redis)           
    - [4 定时调度](#4-定时调度)
    - [5 API网关](#5-api网关)  
    - [6 配置中心](#6-配置中心)  
- [数据库](#数据库)
    - [1 基本理论](1-基本理论)      
        &nbsp;&nbsp; [1.1 三大范式](#11-三大范式)    
        &nbsp;&nbsp; [1.2 事物隔离级别](#12-事物隔离级别)    
    - [2 原理](2-原理)      
        &nbsp;&nbsp; [2.1 架构设计](#21-架构设计)    
        &nbsp;&nbsp; [2.2 数据存储](#22-数据存储)    
    - [3 索引](3-索引)      
        &nbsp;&nbsp; [3.0 索引原理](#30-索引原理)    
        &nbsp;&nbsp; [3.1 索引类型](#31-索引类型)   
        &nbsp;&nbsp; [3.2 索引原则](#32-索引原则)    
        &nbsp;&nbsp; [3.3 索引失效](#33-索引失效)    
        &nbsp;&nbsp; [3.4 explain](#34-explain)            
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
	- [2 web安全](#2-web安全)   
          &nbsp;&nbsp; [2.1  XSS](#21-XSS)         
          &nbsp;&nbsp; [2.2  CSRF](#22-CSRF)         
          &nbsp;&nbsp; [2.3  SQL注入](#23-SQL注入)         
          &nbsp;&nbsp; [2.4  HashDos](#24-HashDos)         
          &nbsp;&nbsp; [2.5  脚本注入](#25-脚本注入)         
          &nbsp;&nbsp; [2.6  漏洞扫描工具](#26-漏洞扫描工具)    
          &nbsp;&nbsp; [2.7  验证码](#27-验证码)         
	- [3 DDoS 防范](#3-ddos防范)
	- [4 用户隐私信息保护](#4-用户隐私信息保护)
	- [5 序列化漏洞](#5-序列化漏洞)
	- [6 加密解密](#6-加密解密)   
        &nbsp;&nbsp; [6.1  对称加密](#61-对称加密)        
        &nbsp;&nbsp; [6.2  哈希算法](#62-哈希算法)        
        &nbsp;&nbsp; [6.3  非对称加密](#63-非对称加密)        
	- [7 服务器安全](#7-服务器安全)
	- [8 数据安全](#8-数据安全)   
        &nbsp;&nbsp; [8.1  数据备份](#81-数据备份)        
	- [9 网络隔离](#9-网络隔离)    
        &nbsp;&nbsp; [9.1  内外网分离](#91-内外网分离)        
        &nbsp;&nbsp; [9.2  登录跳板机](#92-登录跳板机)        
	- [10 授权、认证](#10-授权认证)   
        &nbsp;&nbsp; [10.1  RBAC](#101-RBAC)        
        &nbsp;&nbsp; [10.2  OAuth2.0](#102-OAuth2.0)        
        &nbsp;&nbsp; [10.3  双因素认证](#103-双因素认证)        
        &nbsp;&nbsp; [10.4  单点登录](#104-单点登录)        
- [运维&统计&技术支持](#运维统计技术支持)  
    - [1 常规监控](#1-常规监控)    
             &nbsp;&nbsp; [1.1 命令行监控工具](#11-命令行监控工具)         
    - [2 日志服务](#2-日志服务)
    - [3 RPC](#3-rpc)
    - [4 docker](#4-docker) 
    - [5 APM](#5-APM) 
    - [6 统计分析](#6-统计分析) 
    - [7 持续集成](#7-持续集成) 
    - [8 自动化运维](#8-自动化运维) 
    - [9 灰度&蓝绿&AB](#9-灰度蓝绿AB) 
    - [10 虚拟化](#10-虚拟化) 
    - [11 devops&openstack](#11-devopsopenstack) 
    - [12 测试](#12-测试)   
          &nbsp;&nbsp; [12.1 压力测试](#121-压力测试)    
          &nbsp;&nbsp; [12.2 测试驱动开发](#122-测试驱动开发)    
          &nbsp;&nbsp; [12.3 单元测试](#123-单元测试)   
    - [13 CGI](#13-CGI)    
 
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
    - [1 流式计算](#1-流式计算)   
    - [2 Hadoop](#2-Hadoop)  
    - [3 Spark](#3-Spark)  
- [搜索引擎](#搜索引擎)
    - [1 搜索引擎原理](#1-搜索引擎原理)
    - [2 Lucene](#2-Lucene)   
    - [3 Elasticsearch](#3-Elasticsearch)   
    - [4 Solr](#4-Solr)   
    - [5 sphinx](#5-sphinx)   
- [golang](#golang) 
- [NSQ](#NSQ) 
- [YII](#YII)
- [项目管理](#项目管理)
- [资讯&技术资源](#资讯技术资源)  
    - [1 行业资讯](#1-行业资讯)
    - [2 博客](#2-博客)    
- [TODO](#TODO) 
<!-- /TOC -->
# 数据结构&算法
* 建立时间复杂度、空间复杂度意识，写出高质量的代码
* 性能更优，核心竞争力
* [数据结构算法导图](https://static001.geekbang.org/resource/image/91/a7/913e0ababe43a2d57267df5c5f0832a7.jpg)

## 1 复杂度分析
#### 1.1 时间复杂度

 * 所有代码执行时间与数据规模增长的趋势，叫时间复杂度。
 * 可以理解每行代码执行时间都是unit_time,执行次数和数据量的关系。知道执行次数，就能知道复杂度。

#### 1.2 时间复杂度量级
 * 常数阶、对数阶、线性阶、线性对数阶、指数阶，平方阶，立方阶，k次方阶，阶乘阶。


#### 1.3 空间复杂度
* 算法存储空间与数据量的关系。

#### 1.4 最好最坏平均均摊时间复杂度
* 平均：执行次数*概率


## 2 数组
#### 2.1 定义
* 内存连续（随机访问），数据类型相同，线性表。

#### 2.2 插入删除

* 插入删除    平均复杂度O（n）

#### 2.3 查找

* 查找 todo
* 随机访问    下标访问复杂度O（1）


## 3 链表
#### 3.1 插入删除
* 插入删除  复杂度为O（1）

#### 3.2 查找
* 查找     复杂度O（n）

#### 3.3 单链&双向链表&循环链表
* 单链表，双向链表，循环链表. 带头链表，不带头链表

#### 3.4 链表操作
* 反转，合并有序链表，是否有环，删除倒数第n个节点，查找中间节点。


## 4 栈
#### 4.1 入栈出栈
* 入栈，出栈，栈顶元素。

#### 4.2 顺序栈&链式栈
* 数组实现顺序栈，链表实现链式栈。

#### 4.3 栈应用
* 用栈实现：表达式求值


## 5 队列
#### 5.1 队头队尾
* 队头出队，队尾入队

#### 5.2 顺序队列&链式队列
* 顺序对列，链式队列

#### 5.3 循环队列
* 队满：（tail+1）% n = head；
* 队空： tail== head；

#### 5.4 阻塞队列&并发队列

## 6 递归
#### 6.1 堆栈溢出
* 递归用循环代替
* 堆栈溢出，重复计算，空间度高，函数调用耗时多。


## 7 排序
#### 7.1 冒泡排序
* 原地排序，稳定性
* 有序度，逆序度，满有序度 n*（n-1）/2
* 时间复杂度。  O（n*n）

#### 7.2 插入排序
* 优于冒泡


#### 7.3 选择排序
* 低于冒泡和插入

#### 7.4 归并排序
* 时间复杂度：都为 nlogn
* 空间负责度： n


#### 7.5 快排


#### 7.6 桶排序


#### 7.7 计数排序


#### 7.8 基数排序

## 8 二分查找

* 只适合数组
* 只适合有序
* 数据太小或太大都不适合
* 类似二分查找示例（如果有重复的数据），第一个小于等于某个值。。。

## 9 跳表
* 空间复杂度 O（n）
* 时间复杂度 O（logn）
* 跳表索引动态更新


## 10 散列表
#### 10.1 散列冲突
* 哈希表，键值，哈希函数，散列值
* 哈希冲突： 链表法；开放寻址法；重复哈希；

####  10.2 设计工业级散列表
* 设计工业级散列表：
* hash函数不能太复杂；
* 动态扩容；
* 高效扩容先扩容，然后等插入数据时再拷贝数据到新的散列表里；
* 散列冲突链表法可以用红黑树替代；

#### 10.3 LRU缓存淘汰算法
* 散列表+链表；复杂度都是O（1）

#### 10.4 Redis有序集合
* 散列表+跳表
* 可以按照添加顺序 顺序访问


#### 10.5 哈希算法

* 应用： 安全加密、唯一标识、数据校验、散列函数、负载均衡、数据分片、分布式存储


#### 10.6 一致性哈希算法


## 11 树
#### 11.1 高度深度层


#### 11.2 二叉树&满二叉树&完全二叉树
* 存储方法：数组，链表；
* 前序遍历，中序遍历，后续遍历；

#### 11.3 二叉查找树
* 数字，有序
* 插入，查找，删除（如果有两个子节点，找到右节点的最小值，替换。）
* 重复： 链表法；放在右子树；
* 查找时间复杂度：logn
* 查找树和散列表对比

#### 11.4 红黑树
#### 11.5 递归树
#### 11.6 B树&B+树
* [B树&B+树示例](https://www.cnblogs.com/nullzx/p/8729425.html)


## 12 堆
* 堆的实现，插入，查找
* 堆排序

## 13 图
* 有向图，无向图
* 临接表（可以把链表换成散列表，跳表，红黑树），临接矩阵存储
* 深度，广度优先搜索


## 14 算法

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
- 《[缓存失效策略](https://blog.csdn.net/clementad/article/details/48229243)》

#### 2.1 本地缓存  
- 《[Nginx本地缓存](https://coderxing.gitbooks.io/architecture-evolution/di-er-pian-ff1a-feng-kuang-yuan-shi-ren/42-xing-neng-zhi-ben-di-huan-cun/422-fu-wu-duan-ben-di-huan-cun/nginx-ben-di-huan-cun.html)》
- 《[PageSpeed 自动优化网站的神器](https://coderxing.gitbooks.io/architecture-evolution/di-er-pian-ff1a-feng-kuang-yuan-shi-ren/42-xing-neng-zhi-ben-di-huan-cun/422-fu-wu-duan-ben-di-huan-cun/4222-pagespeed.html)》

#### 2.2 客户端缓存 
- 《[H5 和移动端 WebView 缓存机制解析与实战](https://mp.weixin.qq.com/s/qHm_dJBhVbv0pJs8Crp77w)》

#### 2.3 服务端缓存 
##### 2.3.1 web缓存
- 《[基于HAProxy的高性能HTTP缓存服务器和RESTful NoSQL缓存服务器](https://github.com/jiangwenyuan/nuster/blob/master/README-CN.md)》
- 《[Varnish Cache](https://github.com/varnishcache/varnish-cache)》
- 《[squid-cache](https://github.com/squid-cache/squid)》  

##### 2.3.2 Memcache
- 《[Memcached 教程](http://www.runoob.com/memcached/memcached-tutorial.html)》 
- 《[深入理解Memcached原理](https://www.jianshu.com/p/36e5cd400580)》 
- 《[Memcached工作原理](https://www.jianshu.com/p/36e5cd400580)》 
- 《[Memcache技术分享：介绍、使用、存储、算法、优化、命中率](http://zhihuzeye.com/archives/2361)》 
- 《[memcache 中 add 、 set 、replace 的区别](https://blog.csdn.net/liu251890347/article/details/37690045)》 

##### 2.3.3 Redis
- 《[Redis 教程](http://www.runoob.com/redis/redis-tutorial.html)》
- 《[redis底层原理](https://blog.csdn.net/wcf373722432/article/details/78678504)》
- 《[Redis持久化](http://doc.redisfans.com/topic/persistence.html)》
   * RDB方式：定期备份快照，常用于灾难恢复。优点：通过fork出的进程进行备份，不影响主进程、RDB 在恢复大数据集时的速度比 AOF 的恢复速度要快。缺点：会丢数据。
   * AOF方式：保存操作日志方式。优点：恢复时数据丢失少，缺点：文件大，回复慢。
   * 也可以两者结合使用。
- 《[分布式缓存--序列3--原子操作与CAS乐观锁](https://blog.csdn.net/chunlongyu/article/details/53346436)》
- 《[Redis单线程架构](https://blog.csdn.net/sunhuiliang85/article/details/73656830)》
- 《[redis的回收策略](https://blog.csdn.net/qq_29108585/article/details/63251491)》

##### 2.3.4 Tair
- 《[tair源码](https://github.com/alibaba/tair)》
- 《[tair与redis比较总结](https://blog.csdn.net/farphone/article/details/53522383)》



  
## 3 消息队列
#### 3.1 消息队列常见问题
- 《[消息队列简介及好处](https://yq.aliyun.com/articles/606240)》
   * 提高系统响应速度、提高系统稳定性、异步化、解耦、消除峰值
- 《[Kafka、RabbitMQ、RocketMQ等消息中间件的对比 —— 消息发送性能和区别](https://blog.csdn.net/yunfeng482/article/details/72856762)》
   * Push方式：优点是可以尽可能快地将消息发送给消费者，缺点是如果消费者处理能力跟不上，消费者的缓冲区可能会溢出。 RabbitMQ 消费者默认是推模式（也支持拉模式）。
   * Pull方式：优点是消费端可以按处理能力进行拉去，缺点是会增加消息延迟。 Kafka 默认是拉模式。
- 《[消息顺序和重复问题](https://www.jianshu.com/p/453c6e7ff81c)》

#### 3.2 RabbitMQ
- 《[RabbitMQ应用场景以及基本原理介绍](https://blog.csdn.net/whoamiyang/article/details/54954780)》《[消息队列之 RabbitMQ](https://www.jianshu.com/p/79ca08116d57)》- 《[RabbitMQ之消息确认机制（事务+Confirm）](https://blog.csdn.net/u013256816/article/details/55515234)》
   * 支持事务，推拉模式都是支持、适合需要可靠性消息传输的场景。

#### 3.3 RocketMQ
- 《[RocketMQ 实战之快速入门](https://www.jianshu.com/p/824066d70da8)》《[RocketMQ源码分析](http://www.iocoder.cn/categories/RocketMQ/?vip&architect-awesome)》
   * Java实现，推拉模式都是支持，吞吐量逊于Kafka。可以保证消息顺序。

#### 3.4 ActiveMQ
- 《[ActiveMQ消息队列介绍](https://www.cnblogs.com/wintersun/p/3962302.html)》
   * 纯Java实现，兼容JMS，可以内嵌于Java应用中。
#### 3.5 Kafka
- 《[Kafka深度解析，众人推荐，精彩好文](https://www.cnblogs.com/wintersun/p/3962302.html)》《[Kafka分区机制介绍与示例](http://lxw1234.com/archives/2015/10/538.htm)》
   * 高吞吐量、采用拉模式。适合高IO场景，比如日志同步。

#### 3.6 Redis
- 《[Redis用作消息队列](https://blog.csdn.net/qq_34212276/article/details/78455004)》
   * 生产者、消费者模式完全是客户端行为，list 和 拉模式实现，阻塞等待采用 blpop 指令。

## 4 定时调度
- 《[linux定时任务cron配置](https://www.cnblogs.com/shuaiqing/p/7742382.html)》
- 《[Linux cron运行原理](https://my.oschina.net/daquan/blog/483305)》
- 《[这些优秀的国产分布式任务调度系统，你用过几个？](https://blog.csdn.net/qq_16216221/article/details/70314337)》

  
## 5 API网关
请求转发、安全认证、协议转换、容灾
- 《[API网关那些儿](http://ylzheng.com/2017/03/14/the-things-about-api-gateway/)》
- 《[谈 API 网关的背景、架构以及落地方案](https://www.infoq.cn/article/2016%2F07%2FAPI-background-architecture-floo)》
- 《[使用Zuul构建API Gateway](https://blog.csdn.net/zhanglh046/article/details/78651993)》
- 《[HTTP API网关选择之一Kong介绍](https://mp.weixin.qq.com/s/LIq2CiXJQmmjBC0yvYLY5A)》


## 6 配置中心
- 《[Apollo（阿波罗）是携程框架部门研发的分布式配置中心](https://github.com/ctripcorp/apollo)》
- 《[基于zookeeper实现统一配置管理理](https://blog.csdn.net/u011320740/article/details/78742625)》

# 数据库
## 1 基本理论
#### 1.1 三大范式
- 《[三大范式五大约束](https://www.cnblogs.com/waj6511988/p/7027127.html)》

#### 1.2 事物隔离级别
- 《[事务隔离级别](https://www.cnblogs.com/huanongying/p/7021555.html)》
- 《[数据库事务特性](https://blog.csdn.net/u012440687/article/details/52116108)》
- 《[数据库事务隔离级别](https://blog.csdn.net/qq_33290787/article/details/51924963)》
- 《[InnoDB幻读问题](http://blog.sina.com.cn/s/blog_499740cb0100ugs7.html)》

 * 原子性，一致性，隔离性，持久性
 * 读不提交，出现脏读
 * 读提交，出现不可重复读
 * 重复读，出现幻读
 * 序列化，--


## 2 原理
#### 2.1 架构设计
 * 连接层，优化/缓存层，存储引擎

#### 2.2 数据存储
 * frm文件表空间描述
 * ibd文件索引和数据
 * page页：头双向链表，body链表存储记录
 * 索引：B+树



## 3 索引
#### 3.0 索引原理
- 《[索引原理](https://www.cnblogs.com/bypp/p/7755307.html)》
 * 聚族索引与非聚簇索引
 * B+树，非叶子结点不存数据，只存索引键值，叶子结点存数据。


#### 3.1 索引类型
 * 聚族索引与非聚簇索引
 * 单列索引
 * 索引覆盖
 * 组合索引，索引合并
 * 全文索引

#### 3.2 索引原则
- 《[mysql索引](https://www.cnblogs.com/chenshishuo/p/5030029.html)》《[mysql索引](https://www.cnblogs.com/bypp/p/7755307.html)》《[mysql索引](https://www.jb51.net/article/133626.htm)》
 * 1.在经常需要搜索的列上,可以加快索引的速度
 * 2.主键列上可以确保列的唯一性
 * 3.在表与表的而连接条件上加上索引,可以加快连接查询的速度
 * 4.在经常需要排序(order by),分组(group by)和的distinct 列上加索引 可以加快排序查询的时间,  (单独order by 用不了索引，索引考虑加where 或加limit)
 * 5.在一些where 之后的 < <= > >= BETWEEN IN 以及某个情况下的like 建立字段的索引(B-TREE)

 * 6.like语句的 如果你对nickname字段建立了一个索引.当查询的时候的语句是 nickname lick '%ABC%' 那么这个索引讲不会起到作用.而nickname lick 'ABC%' 那么将可以用到索引

 * 7.索引不会包含NULL列,如果列中包含NULL值都将不会被包含在索引中,复合索引中如果有一列含有NULL值那么这个组合索引都将失效,一般需要给默认值0或者 ' '字符串

 * 8.使用短索引,如果你的一个字段是Char(32)或者int(32),在创建索引的时候指定前缀长度 比如前10个字符 (前提是多数值是唯一的..)那么短索引可以提高查询速度,并且可以减少磁盘的空间,也可以减少I/0操作.

 * 9.不要在列上进行运算,这样会使得mysql索引失效,也会进行全表扫描

 * 10.选择越小的数据类型越好,因为通常越小的数据类型通常在磁盘,内存,cpu,缓存中 占用的空间很少,处理起来更快
 * 尽量选择区分度高的列作索引


#### 3.3 索引失效
* 索引列 like "%xx"
* 使用函数计算
* or 有没有索引的列
* 类型不一致
* order by 查询的列不是主键排序就不走索引


#### 3.4 explain
- 《[explain分析SQL](https://www.cnblogs.com/butterfly100/archive/2018/01/15/8287569.html)》
  * select_type：simple，primary，subquery，derived，union，union result
  * type：system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL





 
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

## 2 web 安全

#### 2.1 XSS
* [《xss攻击原理与解决方法》](https://blog.csdn.net/qq_21956483/article/details/54377947)
#### 2.2 CSRF
* [《CSRF原理及防范》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/641-web-an-quan-fang-fan/6412-csrf.html)

#### 2.3 SQL 注入

* [《SQL注入》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/641-web-an-quan-fang-fan/6413-sql-zhu-ru.html)

#### 2.4 HashDos


* [《邪恶的JAVA HASH DOS攻击》](http://www.freebuf.com/articles/web/14199.html)
	* 利用JsonObject 上传大Json，JsonObject 底层使用HashMap；不同的数据产生相同的hash值，使得构建Hash速度变慢，耗尽CPU。
* [《一种高级的DoS攻击-Hash碰撞攻击》](http://blog.it2048.cn/article_hash-collision.html )
* [《关于Hash Collision DoS漏洞：解析与解决方案》](http://www.iteye.com/news/23939/)

#### 2.5 脚本注入

* [《上传文件漏洞原理及防范》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/641-web-an-quan-fang-fan/6414-shang-chuan-wen-jian-guo-lv.html)

#### 2.6 漏洞扫描工具
* [《DVWA》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/6421-dvwa.html)
* [W3af](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/w3af.html)
* [OpenVAS详解](https://blog.csdn.net/xygg0801/article/details/53610640)

#### 2.7 验证码

* [《验证码原理分析及实现》](https://blog.csdn.net/niaonao/article/details/51112686)

* [《详解滑动验证码的实现原理》](https://my.oschina.net/jiangbianwanghai/blog/1031031)
	* 滑动验证码是根据人在滑动滑块的响应时间，拖拽速度，时间，位置，轨迹，重试次数等来评估风险。

* [《淘宝滑动验证码研究》](https://www.cnblogs.com/xcj26/p/5242758.html)

## 3 DDoS防范
* [《学习手册：DDoS的攻击方式及防御手段》](http://netsecurity.51cto.com/art/201601/503799.htm)
* [《免费DDoS攻击测试工具大合集》](http://netsecurity.51cto.com/art/201406/442756.htm)

## 4 用户隐私信息保护

1. 用户密码非明文保存，加动态salt。
2. 身份证号，手机号如果要显示，用 “\*” 替代部分字符。
3. 联系方式在的显示与否由用户自己控制。
4. TODO

* [《个人隐私包括哪些》](https://zhidao.baidu.com/question/1988017976673661587.html)
* [《在互联网上，隐私的范围包括哪些？》](https://www.zhihu.com/question/20137108)

* [《用户密码保存》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/642-shu-ju-jia-mi/6425-jia-mi-chang-jing-ff1a-yong-hu-mi-ma-bao-cun.html)

## 5 序列化漏洞
* [《Lib之过？Java反序列化漏洞通用利用分析》](https://blog.chaitin.cn/2015-11-11_java_unserialize_rce/)

## 6 加密解密

#### 6.1 对称加密

* [《常见对称加密算法》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/642-shu-ju-jia-mi/6421-chang-jian-dui-cheng-jia-mi-suan-fa.html)
	* DES、3DES、Blowfish、AES
	* DES 采用 56位秘钥，Blowfish 采用1到448位变长秘钥，AES 128，192和256位长度的秘钥。
	* DES 秘钥太短（只有56位）算法目前已经被 AES 取代，并且 AES 有硬件加速，性能很好。
	
#### 6.2 哈希算法
* [《常用的哈希算法》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/642-shu-ju-jia-mi/6422-chang-jian-ha-xi-suan-fa-and-hmac.html)
	* MD5 和 SHA-1 已经不再安全，已被弃用。
	* 目前 SHA-256 是比较安全的。
	
* [《基于Hash摘要签名的公网URL签名验证设计方案》](https://blog.csdn.net/zhangruhong168/article/details/78033202)

#### 6.3 非对称加密
* [《常见非对称加密算法》](https://coderxing.gitbooks.io/architecture-evolution/di-san-pian-ff1a-bu-luo/642-shu-ju-jia-mi/6424-chang-yong-fei-dui-cheng-jia-mi-suan-fa.html)
	* RSA、DSA、ECDSA(螺旋曲线加密算法)
	* 和 RSA 不同的是 DSA 仅能用于数字签名，不能进行数据加密解密，其安全性和RSA相当，但其性能要比RSA快。
	* 256位的ECC秘钥的安全性等同于3072位的RSA秘钥。

		[《区块链的加密技术》](http://baijiahao.baidu.com/s?id=1578348858092033763&wfr=spider&for=pc)	


## 7 服务器安全
* [《Linux强化论：15步打造一个安全的Linux服务器》](http://www.freebuf.com/articles/system/121540.html)

## 8 数据安全

#### 8.1 数据备份

TODO

## 9 网络隔离

#### 9.1 内外网分离

TODO

#### 9.2 登录跳板机
在内外环境中通过跳板机登录到线上主机。
* [《搭建简易堡垒机》](http://blog.51cto.com/zero01/2062618)

## 10  授权、认证
#### 10.1 RBAC 
* [《基于组织角色的权限设计》](https://www.cnblogs.com/zq8024/p/5003050.html)
* [《权限系统与RBAC模型概述》](https://www.cnblogs.com/shijiaqi1066/p/3793894.html)
* [《Spring整合Shiro做权限控制模块详细案例分析》](https://blog.csdn.net/he90227/article/details/38663553)

#### 10.2 OAuth2.0
* [《理解OAuth 2.0》](http://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html)
* [《一张图搞定OAuth2.0》](https://www.cnblogs.com/flashsun/p/7424071.html)

#### 10.3 双因素认证（2FA）

2FA - Two-factor authentication，用于加强登录验证

常用做法是 登录密码 + 手机验证码（或者令牌Key，类似于与网银的 USB key）

* 【《双因素认证（2FA）教程》】(http://www.ruanyifeng.com/blog/2017/11/2fa-tutorial.html)

#### 10.4 单点登录(SSO)

* [《单点登录原理与简单实现》](https://www.cnblogs.com/ywlaker/p/6113927.html)

* [CAS单点登录框架](https://github.com/apereo/cas)





   

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


## 7 持续集成
- 《[持续集成是什么？](http://www.ruanyifeng.com/blog/2015/09/continuous-integration.html)》
   * 合并；交付；测试；部署；
- 《[8个流行的持续集成工具](https://www.testwo.com/article/1170)》
- 《[使用Jenkins进行持续集成](https://www.liaoxuefeng.com/article/001463233913442cdb2d1bd1b1b42e3b0b29eb1ba736c5e000)》

## 8 自动化运维
- 《[Ansible基础配置和企业级项目实用案例](https://www.cnblogs.com/heiye123/articles/7855890.html)》
   * 批量系统配置、批量程序部署、批量运行命令等功能
- 《[Ansible中文权威指南](http://www.ansible.com.cn/index.html)》   
- 《[自动化运维工具——puppet详解（一）](https://www.cnblogs.com/keerya/p/8040071.html)》 
- 《[Chef 的安装与使用](https://www.ibm.com/developerworks/cn/cloud/library/1407_caomd_chef/)》 

## 9 灰度蓝绿AB
- 《[蓝绿部署](http://www.cnblogs.com/renzaijianghu/p/9197075.html)》 
- 《[蓝绿部署、A/B 测试以及灰度发布](https://www.v2ex.com/t/344341)》
- 《[技术干货 | AB 测试和灰度发布探索及实践](https://testerhome.com/topics/11165)》

## 10 虚拟化
- 《[VPS的三种虚拟技术OpenVZ、Xen、KVM优缺点比较](https://blog.csdn.net/enweitech/article/details/52910082)》
- 《[KVM详解，太详细太深入了，经典](http://blog.chinaunix.net/uid-20201831-id-5775661.html)》
- 《[KVM 虚拟机安装详解](https://www.coderxing.com/kvm-install.html)》
- 《[Xen虚拟化基本原理详解](https://www.cnblogs.com/sddai/p/5931201.html)》
- 《[开源Linux容器 OpenVZ 快速上手指南](https://blog.csdn.net/longerzone/article/details/44829255)》

## 11 devops&openstack
- 《[DevOps 详解](https://www.infoq.cn/article/detail-analysis-of-devops)》
- 《[OpenStack构架知识梳理](https://www.cnblogs.com/klb561/p/8660264.html)》



##  12 测试
#### 12.1 压力测试
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

#### 12.2 测试驱动开发
- 《[深度解读 - TDD（测试驱动开发）](https://www.jianshu.com/p/62f16cd4fef3)》


####  12.3 单元测试
- 《[单元测试](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/00140137128705556022982cfd844b38d050add8565dcb9000)》
   * 单元测试可以有效地测试某个程序模块的行为，是未来重构代码的信心保证。
   * 单元测试的测试用例要覆盖常用的输入组合、边界条件和异常。
   * 单元测试代码要非常简单，如果测试代码太复杂，那么测试代码本身就可能有bug。
   * 单元测试通过了并不意味着程序就没有bug了，但是不通过程序肯定有bug。

## 13 CGI
- 《[简单说明CGI和动态请求是什么](https://www.cnblogs.com/f-ck-need-u/p/7627035.html)》
   * CGI,fastCGI
   * php-cgi,php-fpm
   * web-server,php-fpm,zend引擎，CGI脚本



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

## 1 流式计算

#### 1.1 Storm
* [官方网站](http://storm.apache.org/)
* [《最详细的Storm入门教程》](https://blog.csdn.net/uisoul/article/details/77989927)

#### 1.2 Flink
* [《Flink之一 Flink基本原理介绍》](https://blog.csdn.net/lisi1129/article/details/54844919)

#### 1.3 Kafka Stream
* [《Kafka Stream调研：一种轻量级流计算模式》](https://yq.aliyun.com/articles/58382)

#### 1.4 应用场景

例如：

* 广告相关实时统计；
* 推荐系统用户画像标签实时更新；
* 线上服务健康状况实时监测；
* 实时榜单；
* 实时数据统计。

## 2 Hadoop

* [《用通俗易懂的话说下hadoop是什么,能做什么》](https://blog.csdn.net/houbin0912/article/details/72967178)
* [《史上最详细的Hadoop环境搭建》](http://gitbook.cn/books/5954c9600326c7705af8a92a/index.html)

#### 2.1 HDFS
* [《【Hadoop学习】HDFS基本原理》](https://segmentfault.com/a/1190000011575458)

#### 2.2 MapReduce
* [《用通俗易懂的大白话讲解Map/Reduce原理》](https://blog.csdn.net/oppo62258801/article/details/72884633)
* [《 简单的map-reduce的java例子》](https://blog.csdn.net/foye12/article/details/78358292)

#### 2.3 Yarn
* [《初步掌握Yarn的架构及原理》](http://www.cnblogs.com/codeOfLife/p/5492740.html)

## 3 Spark
* [《Spark(一): 基本架构及原理》](http://www.cnblogs.com/tgzhu/p/5818374.html)



# 搜索引擎
## 1 搜索引擎原理

* [《倒排索引--搜索引擎入门》](https://www.jianshu.com/p/0193dc44135b)

## 2 Lucene
* [《Lucene入门简介》](https://www.cnblogs.com/rodge-run/p/6551152.html)

## 3 Elasticsearch

* [《Elasticsearch学习，请先看这一篇！》](https://blog.csdn.net/laoyang360/article/details/52244917)
* [《Elasticsearch索引原理》](https://blog.csdn.net/cyony/article/details/65437708)

## 4 Solr
* [《 Apache Solr入门教程》](https://blog.csdn.net/u011936655/article/details/51960005)
* [《elasticsearch与solr比较》](https://blog.csdn.net/convict_eva/article/details/53537837)

## 5 sphinx 
* [《Sphinx 的介绍和原理探索》](http://blog.jobbole.com/101672/)

# golang
## 1 基本概念
#### 1.1 GOPATH和工作区
- 《[GO环境安装及目录](https://www.jianshu.com/p/8a87eeec15f2)》《[Go语言之讲解GOROOT、GOPATH、GOBIN](https://www.cnblogs.com/pyyu/p/8032257.html)》
 * GOROOT,GOPATH,GOBIN
 * src,bin,pkg
 * go build; go install ; go get;

#### 1.2 命令源码文件
 * 如果一个源码文件声明属于main包，并且包含一个无参数声明且无结果声明的main函数，那么它就是命令源码文件

#### 1.3 库源码文件
 * 库源码文件是不能被直接运行的源码文件，它仅用于存放程序实体，这些程序实体可以被其他代码使用
 * internal
 * 函数导出，访问权限

#### 1.4 变量声明
 * var s string
 * s := "name"
 * 短变量声明
 
#### 1.5 作用域
 * 包级私有的、模块级私有的和公开的

## 2 数组切片
 * 在无需扩容时，append函数返回的是指向原底层数组的新切片,而在需要扩容时，append函数返回的是指向新底层数组的新切片。
 * slice扩容 2倍，1.25倍。
 * 底层数组不会被替换，直接生产新的数组

 
## 3 链表
 * container/list

## 4 字典哈希

## 5 channel

#### 5.0 注意点
 * 通道不能一直阻塞
 * goroutine执行完再退出
 * goroutine泄漏
 * channel关闭
 * goroutine不能打开太多，应该有限制

#### 5.1 通道规则
 * 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
 * 发送操作和接收操作中对元素值的处理都是不可分割的
 * 发送操作在完全完成之前会被阻塞。接收操作也是如此


#### 5.2 无缓存通道
 * Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作
 * 接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作

 
#### 5.3 close
 * 当一个channel被关闭后，再向该channel发送数据将导致panic异常
 * 当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值
 * 只需要关闭接受者goroutine
 * range是安全接收的，它会判断是否关闭且没有收据接收。

 
#### 5.4 单方向channel
 * 只有在发送者所在的goroutine才会调用close函数
 * 任何双向channel向单向channel变量的赋值操作都将导致该隐式转换，没有反向转换的语法

#### 5.5 缓存通道
 * 内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作
 * channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。
 * 可以保证go程执行完再退出（？？）



#### 5.6 并发循环
 * 循环时显式的将参数传给go程。
 * wg.wait为什么要在go func(){}中
 * for range channel当channel关闭且流干的时候才退出循环

#### 5.7 select
 * 同时都满足条件，会随机选择一个执行。是不是有bug
 * Loop goto  break

#### 5.8 并发的退出
 * 通过close广播，select去接收。 
 
 
## 6 基于共享变量的并发

#### 6.1 数据竞争
 * 数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生。
 * 三种方法避免数据竞争：不要去写变量；避免从多个goroutine访问变量；加锁；

#### 6.2 互斥锁
 * sync.Mutex互斥锁

#### 6.3 读写锁 
 * sync.RMutex读写锁

#### 6.4 内存同步
 * 内存同步
 * 不使用channel且不使用mutex这样的显式同步操作时，我们就没法保证事件在不同的goroutine中看到的执行顺序是一致的

#### 6.5 初始化
 * sync.once
 * 保证loadIcons的变量能够对所有goroutine可见

#### 6.6 并发的非阻塞缓存
 * close(ch)会广播到每个goroutine
 * 两种方式：加锁；channel ；
 
# NSQ
TODO

# YII


# 《计算机组成原理》
## 1 计算机概述
#### 1.1 软硬件
 * 软件，硬件
 * 系统软件；应用软件

#### 1.2 层次结构
 * 汇编--机器语言-解释操作系统-微程序解释机器语言-执行机器语言
 * 翻译程序两种：编译程序；解释程序；

#### 1.3 计算机组成
 * 运算器，存储器，控制器，输入输出
 * 指令和数据相同地位存放在存储器，可按地址寻访。
 * 指令和数据由二进制码表示
 * 指令由操作码和地址码组成
 * 指令是顺序存放的
 * 以运算器为中心
 * 现代计算机：MM+CPU+IO

#### 1.4 计算机工作过程
 * 数学建模；确定计算方法；编程序


# 项目管理
TODO

# 资讯&技术资源
## 1 行业资讯
- 《[36Kr](https://36kr.com/)》
- 《[TeahWeb](http://www.techweb.com.cn/)》
## 2 博客
- 《[阿里中间件团队](http://jm.taobao.org/)》
- 《[美团技术团队](https://tech.meituan.com/)》

## TODO

* 计算机组成原理
* 编译原理和技术
* 操作系统原理与设计
* 数据结构算法基础
* 计算机组成原理实验
* 数据库系统及应用
* 并行计算
* 软件工程
* 计算机导论
* 计算机网络