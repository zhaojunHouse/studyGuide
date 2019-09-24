#php&fpm

[`php-fpm`详细配置](https://www.jb51.net/article/148550.htm)

	pid = /opt/remi/php56/root/var/run/php-fpm/php-fpm.pid 
	#pid进程文件，默认为none。
	
	error_log = /opt/remi/php56/root/var/log/php-fpm/error.log 
	#错误日志位置，默认：安装路径 #INSTALL_PREFIX#/log/php-fpm.log。如果设置为syslog，log就会发送给syslogd服务而不会写进文件里。
	
	syslog.facility = daemon 
	#把日志写进系统log，linux还不够熟悉，暂时不用理会。
	
	syslog.ident = php-fpm 
	#系统日志标示，如果跑了多个fpm进程，需要用这个来区分日志是谁的。
	
	log_level = notice 
	#日志等级，默认notice，可选：alert, error, warning, notice, debug
	
	emergency_restart_threshold = 60 
	#配合下面emergency_restart_interval参数
	
	emergency_restart_interval = 60s 
	#如果在此参数设置的时间内，出现SIGSEGV或SIGBUS的子进程数超过emergency_restart_threshold参数设置的值，那么fpm就会优雅的重启，值是0表示off这个功能，可用的单位有：s秒,m分,h时,d天。
	
	process_control_timeout = 0 
	#设置子进程接受主进程复用信号的超时时间。这个每天明白，是过了这个时间就不能复用了？
	
	process.max = 128 
	#当动态管理子进程时，fpm最多能fork多少个进程，0表示无限制，这是所有进程池能启动子进程的总和，谨慎使用。
	
	process.priority = -19 
	#设置子进程的优先级，在master进程以root用户启动时有效；如果没有设置，子进程会继承master进程的优先级，值范围-19(最高)到20(最低)，默认不设置。
	
	daemonize = yes 
	#设置成no用于调试bug，默认为yes。
	
	rlimit_files = 1024 
	#设置master进程最多能打开的文件，默认为系统的值。
	
	rlimit_core = 0 
	#master进程核心rlimit限制值；可选unlimited或>=0的整数，默认为系统的值。
	
	events.mechanism = epoll 
	#事件处理机制，默认自动检测，可选值：select，poll，epoll(linux>=2.5.44)，kqueue，/dev/poll，port
	
	systemd_interval = 10s 
	#当fpm被设置为系统服务时，多久向服务器报告一次状态，单位有s,m,h。
	
	
#


	

    user = apache

    group = apache

    #以什么用户什么组的权限来运行池fpm。

    用apache可以像httpd服务一样去访问某些目录

    listen = 127.0.0.1:9000 
    #监听的ip和端口，可以 /path/to/unix/socket 来监听unix socket，性能更好。

    listen.backlog = 65535 
    #未accept处理的socket队列大小，-1 on FreeBSD and OpenBSD，其他平台默认65535，高并发时重要，合理设置会及时处理排队的请求；太大会积压太多，处理完后nginx在前面都等超时断开这个和fpm的socket连接了，就杯具了。不要用-1，建议1024以上，最好是2的幂值。

    #一个池共用一个backlog队列，所有的池进程都去这个队列里accept连接。

    #最大数量受限于系统配置 cat /proc/sys/net/core/somaxconn，系统配置修改：vim /etc/sysctl.conf，增加 net.core.somaxconn = 2000 则最大为2000，然后php最大的backlog可以到2000。

    listen.owner = apache

    listen.group = apache

    listen.mode = 0660

    #用socket连接方式时，指定拥有unix socket权限的用户，默认和运行的用户一样；用tcp连接可以注释掉

    listen.allowed_clients = 127.0.0.1 
    #设置允许连接fpm的地址，比如nginx就要来连，多个地址用逗号隔开，如果不配置，则默认任意地址都能来连。

    process.priority = -19 
    #池进程的权限，同样要master进程是root用户才有效，和全局那个一样，不设置的话会继承master进程的优先级。

    pm = dynamic 
    #启动时子进程管理方式，可选值：static(启动时创建指定个数)， dynamic(启动时根据情况创建，至少有一个)， ondemand(启动时不创建子进程，有需求才创建)

    pm.max_children = 5 
    #该池同时最多存在5个进程， 三种管理方式都要配置

    pm.start_servers = 2 
    #fpm启动时创建2个子进程，只适用动态dynamic管理方式

    pm.min_spare_servers = 2 
    #服务器闲置时最少保持2个子进程，不够这个数就会创建，只适用动态dynamic管理方式

    pm.max_spare_servers = 3 
    #服务器闲置时最多要有几个，多了会kill，只适用动态dynamic管理方式

    pm.process_idle_timeout = 10s
    #子进程闲置10s后就会被杀掉。

    pm.max_requests = 500 
    #每个子进程最大处理500请求就被回收，可防止内存泄露。

    pm.status_path string

    #FPM 状态页面的网址。如果没有设置，则无法访问状态页面，默认值：无。

    ping.path string

    #FPM 监控页面的 ping 网址。如果没有设置，则无法访问 ping 页面。该页面用于外部检测 FPM 是否存活并且可以响应请求。请注意必须以斜线开头（/）。

    ping.response string
    #用于定义 ping 请求的返回响应。返回为 HTTP 200 的 text/plain 格式文本。默认值：pong。

    process.priority int
    #设置 worker 的 nice(2)优先级（如果设置了的话）。 该值从 -19（最高优先级） 到 20（更低优先级）。 默认值：不设置

    prefix string
    #检测路径时使用的前缀

    access.log = var/log/$pool.access.log 
    #访问文件日志，没啥用处，比如yii2每次都记录访问index.php，只是记录真实的PHP文件。

    slowlog = var/log/$pool.log.slow 
    #PHP文件执行过慢的日志，会准确的记录具体哪一行代码太慢，这个非常有用，在设置了时间时生效。

    request_slowlog_timeout = 2s 
    #超过这个运行时间就会写慢日志

    request_terminate_timeout = 3s 
    #单个请求的超时时间，有时候php.ini设置的最大执行时间未生效，这个就会来干掉那个执行太久的请求。

    rlimit_files = 1024 
    #最大打开句柄数，默认为系统值。

    rlimit_core = 0 
    #最多的核心使用数，默认为系统分配。

    chroot = /path 
    #路径必须是绝对路径，改变子进程的跟目录，可以把进程对文件系统的读写与实际的操作系统文件系统隔离，对安全有好处。

    chdir = /var/www 

    #改变当前工作目录，可以用相对路径，默认是当前目录或者chroot。

    catch_workers_output = yes 
    #重定向标准输出stdout和标准错误stderr到主错误日志，如果不设置，这两个日志就会定向到/dev/null，在高负载情况下，这个配置会引起页面延迟几毫秒，默认不开启。

    clear_env = no 
    #创建work进程时是否清除环境变量，如果是yes，那么该子进程 getenv() 就访问不到 $_ENV 和$_SERVER 了。

    security.limit_extensions = .php .php3 .php4 .php5 
    #为了安全，限制能执行的脚本后缀

    #为当前池指定另外的 php.ini 配置，比如指定当前池的错误日志写在哪个地方

    php_value/php_flag 
    #可以设置php.ini的内容，可以被ini_set覆盖

    php_admin_value/php_admin_flag 
    #这个同上，但是不会被ini_set覆盖。

    #其中flag设置的，值只能是on, off, 1, 0, true, false, yes or no，其他类型的值需要用value。

    php_flag[display_errors] = off

    php_admin_value[error_log] = /var/log/fpm-php.www.log

    php_admin_flag[log_errors] = on

    php_admin_value[memory_limit] = 32M

    #这种方法设置 `disable_functions` 和 `disable_classes` 时，不会覆盖 php.ini 的设置，只会追加。
