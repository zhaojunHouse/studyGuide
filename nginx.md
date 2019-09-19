# nginx配置
 * [nginx官网](http://nginx.org/en/docs/)
 * [nginx配置模块](https://blog.csdn.net/tan_zheng/article/details/83509071)
 
 
## main模块
	
	
	user www-data;
	
	worker_processes auto;   #auto默认为服务器的核心数
	
	pid /run/nginx.pid;
	
	#全局错误日志定义类型，[ debug | info | notice | warn | error | crit ]
	error_log /usr/local/nginx/logs/error.log info
	
	
### `user`
 * 主要是指定执行nginx的worker process的用户及组，linux里所有程序都是文件，都具有权限问题，这个指定的用户对特定的文件有没有权限访问或执行，就是这个用户的意义。
 * [用户及组/etc/group](https://www.cnblogs.com/peida/archive/2012/12/05/2802419.html)
 * [/etc/passwd](https://blog.csdn.net/a1154490629/article/details/52190801)
 * [用户/组/权限](https://blog.csdn.net/s740556472/article/details/78077453)
 * [Linux命令ll详解](https://www.cnblogs.com/kongzhongqijing/p/3488884.html)

### `worker_process`
 * 子进程数，建议和CPU核数一致
 * [`worker_proecss`](https://blog.csdn.net/fireroll/article/details/15756745)

### `worker_cpu_affinity`
 * 利用多核，多个进程对应多个CPU内核
 * [`worker_cpu_affinity`](https://blog.csdn.net/u011957758/article/details/50959823)

### `error_log`
 * 记录nginx的错误日志，可配置在http，server，main，location内
 * [`error_log`](https://blog.csdn.net/czlun/article/details/73251714)

### `pid`
 * 指定进程id的存储文件位置，nginx.pid存放的是nginx的master进程的进程号。可以使用kill -USR2重启
 * [`pid`](https://blog.51cto.com/meiling/2152547)

 
### `worker_rlimit_nofile`
 * 这个指令是指Nginx单个worker进程最大可用文件描述符数量，理论值应该是最多打开文件数（ulimit -n）与nginx进程数相除，但是nginx分配请求并不是那么均匀，所以最好与ulimit -n的值保持一致。
 * [`worker_rlimit_nofile `](https://www.jianshu.com/p/4fa08f2a04ed) 


## EVENTS模块
 * [`events`](https://blog.51cto.com/13229718/2080155)

		events {
	
		accept_mutex on; #设置网路连接序列化，防止惊群现象发生，默认为on
		
		multi_accept on; #设置一个进程是否同时接受多个网络连接，默认为off
		
		use epoll; #事件驱动模型，select|poll|kqueue|epoll|resig|/dev/poll|eventport
		
		worker_connections 1024; # 最大连接数
		
		client_header_buffer_size 4k;
		
		open_file_cache max=2000 inactive=60s;
		
		open_file_cache_valid 60s;
		
		open_file_cache_min_uses 1
		
		}
	
 
### `use `
 * 使用epoll的I/O 模型(值得注意的是如果你不知道Nginx该使用哪种轮询方法的话，它会选择一个最适合你操作系统的)
 
### `worker_connections` 
 * 工作进程的最大连接数量    理论上每台nginx服务器的最大连接数为`worker_processes`*`worker_connections`

### `accept_mutex`
 * 惊群现象：一个网路连接到来，多个睡眠的进程被同事叫醒，但只有一个进程能获得链接，这样会影响系统性能。设置网路连接序列化，防止惊群现象发生，默认为on

### `multi_accept`
 * 设置是否允许同时接受多个网络连接：

　　只能在events模块设置，Nginx服务器的每个工作进程可以同时接受多个新的网络连接，但是需要在配置文件中配置，此指令默认为关闭


### `keepalive_timeout`
 * 在处理完一个请求后保持这个 TCP 连接的打开状态
 * [`keepalive_timeout`](https://blog.csdn.net/weixin_42350212/article/details/81123932)


### `client_header_buffer_size`
 * 客户端请求头部的缓冲区大小，这个可以根据你的系统分页大小来设置，一般一个请求头的大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为系统分页大小。查看系统分页可以使用 getconf PAGESIZE命令

### `open_file_cache  max=2000 inactive=60s`
 * 为打开文件指定缓存，默认是没有启用的，max指定缓存最大数量，建议和打开文件数一致，inactive是指经过多长时间文件没被请求后删除缓存 打开文件最大数量为我们再main配置的worker_rlimit_nofile参数

### `open_file_cache_valid`
 * 这个是指多长时间检查一次缓存的有效信息。如果有一个文件在inactive时间内一次没被使用，它将被移除

### `open_file_cache_min_uses`
 *  open_file_cache指令中的inactive参数时间内文件的最少使用次数，如果超过这个数字，文件描述符一直是在缓存中打开的，如果有一个文件在inactive时间内一次没被使用，它将被移除。

 
## HTTP模块
	http{
       include       mime.types;
       default_type  application/octet-stream;
        
       server_tokens off;
        
       log_format main '$http_x_real_ip - $remote_addr - $remote_user [$time_local] "$request" '
                               '$status $body_bytes_sent "$http_referer" '
                               '"$http_user_agent" $http_x_forwarded_for'
                               ' "$request_time"';

        server_names_hash_bucket_size 128;
        client_header_buffer_size 128k;
        large_client_header_buffers 4 128k;
        client_max_body_size 100m;

        sendfile on;
        tcp_nopush on;
        tcp_nodelay on;

        keepalive_timeout 600s;

        fastcgi_connect_timeout 600s;
        fastcgi_send_timeout 600s;
        fastcgi_read_timeout 600s;
        fastcgi_buffer_size 128k;
        fastcgi_buffers 8 128k;
        fastcgi_busy_buffers_size 256k;
        fastcgi_temp_file_write_size 256k;

        gzip on;
        gzip_min_length  1k;
        gzip_buffers     4 16k;
        gzip_http_version 1.0;
        gzip_comp_level 2;
        gzip_types       text/plain application/x-javascript text/css application/xml application/json;
        gzip_vary on;

        # 加载sites-enabled下所有的站点配置
        include /usr/local/nginx/sites-enabled/*.conf;

	}

### `client_body_buffer_size`
 * 设置用于读取客户端请求体的缓冲区大小。在请求体大于缓冲区的情况下，将整个主体或仅其部分写入缓冲。默认情况下，缓冲区大小等于两个内存页。这是x86或其他的其他32位平台上的应为8K，在X86-64或其他64位平台上通常为16K。

### `client_body_temp_path`
 * 定义用于存储持有客户端请求主体的混存文件的目录。可在指定目录下使用三级子目录层次结构

### `client_body_timeout`
 * 定义用于读取客户端请求体的超时时间。超时仅设置在两个连续读取操作之间的时间段内，而不是用于整个请求体的传输。如果客户端在此时间内不发送任何内容，则请求终止（408）（请求超时）错误。

### `client_header_buffer_size`
 * 读取客户端请求头的缓冲区大小。

### `client_header_timeout`
 * 读取客户端请求头的超时

### `client_max_body_size`
 * 客户端请求主体的最大允许大小

### `connection_pool_size`
 * 允许精确调整每个连接内存分配。该指令对性能影响最小，一般不应使用。默认情况下，32位平台上的大小等于256字节，64位平台上的字节为512字节。

 
### `default_type`
 * 定义响应的默认MIME类型。可以使用类型指令来设置文件扩展名到MIME类型的映射。

### `large_client_header_buffers`
 * 设置用于读取大客户请求头的缓冲区的最大数量和大小。请求行不能超过一个缓冲区的大小，或者将414（请求URI过大）错误返回给客户端。请求头字段不能超过一个缓冲区的大小，或者400（坏请求）错误返回给客户端。仅根据需求分配缓冲器。默认情况下，缓冲区大小等于8K字节。如果在请求处理结束之后，连接被转换为保持活动状态，则释放这些缓冲器。

### `open_file_cache_errors`
 * 启用或禁用open_file_cache缓存文件查找错误。


### `send_timeout`
 * 这个超时时间是发送响应的超时时间，即Nginx服务器向客户端发送了数据包，但客户端一直没有去接收这个数据包。如果某个连接超过send_timeout定义的超时时间，那么Nginx将会关闭这个连接。

### `sendfile`
 * 设置为on可以启用Linux上的sendfile系统调用来发送文件，它减少了内核态与用户态之间的两次内存复制，这样就会从磁盘中读取文件后直接在内核态发送到网卡设备，提高了发送文件的效率。

### `tcp_nodelay`
 * 对keepalive连接是否使用TCP_NODELAY选项,此外，它在SSL连接上启用，用于无缓冲的代理，以及用于WebSocket代理。

 
### `tcp_nopush`
 * 在打开sendfile选项时，确定是否开启FreeBSD系统上的TCP_NOPUSH或Linux系统上的TCP_CORK功能。开启此选项允许在Linux和FreeBSD 4.x上将响应头和正文的开始部分一起发送；一次性发送整个文件。



### `gzip`
 * 禁启用gzip压缩响应。

### `gzip_buffers`
 * 设置用于压缩响应的缓冲区的数量和大小。默认情况下，缓冲区大小等于一个内存页。4K或8K，这取决于一个平台。

### `gzip_disable` `gzip_http_version` 

### `gzip_min_length`
 * 设置将被压缩的响应的最小长度。长度仅由响应头 “Content-Length”字段指定


### `fastcgi_connect_timeout`
 * 定义用于与FASCGI服务器建立连接的超时。应该注意的是，这个超时时间通常不能超过75秒。

### `fastcgi_send_timeout`
 * 设置用于向FASTCGI服务器发送请求的超时时间。超时仅设置在两个连续写入操作之间，而不是用于整个请求的传输。如果FASCGI服务器在此时间内没有接收到任何信息，则关闭连接。

### fastcgi_read_timeout
 * 定义用于从FASCGI服务器读取响应的超时。超时仅设置在两个连续读取操作之间，而不是用于整个响应的传输。如果FASCGI服务器在此时间内不发送任何内容，则连接被关闭。


### `fastcgi_buffer_size`
 * 设置用于读取从FASCGI服务器接收的响应的第一部分的缓冲区的大小。这部分通常包含一个小的响应头。默认情况下，缓冲区大小等于一个内存页。4K或8K，这取决于平台。然而，它也可以设置更小。


### `fastcgi_buffers`
 * 为单个连接设置用于从FASCGI服务器读取响应的缓冲区的数量和大小。默认情况下，缓冲区大小等于一个内存页。4K或8K，这取决于平台。

### `fastcgi_busy_buffers_size`
 * 当启用对来自FastCGI服务器的响应的缓冲时，限制可能忙于向客户端发送响应而响应尚未完全读取的缓冲区的总大小。同时，剩余的缓冲区可用于读取响应，如果需要，还可以将响应的一部分缓冲到临时文件。默认情况下，大小是由fastcgi_buffer_size和fastcgi_buffers指令所设置的两个缓冲区大小限制的。

### `fastcgi_temp_file_write_size`
 * 在启用从FastCGI服务器到临时文件的响应缓冲时，限制一次写入临时文件的数据的大小。默认情况下，大小是由fastcgi_buffer_size和fastcgi_buffers指令设置的两个缓冲区限制的。临时文件的最大大小由fastcgi_max_temp_file_size指令设置。



### upstream模块
 * [upstream](http://nginx.org/en/docs/http/ngx_http_upstream_module.html#example)

### server模块

### `location`
 
	
    等号类型（=）的优先级最高。一旦匹配成功，则不再查找其他匹配项。
    ^~类型表达式。一旦匹配成功，则不再查找其他匹配项。
    正则表达式类型（~ ~*）的优先级次之。如果有多个location的正则能匹配的话，则使用正则表达式最长的那个。
    常规字符串匹配类型。按前缀匹配。
    


 *  [location执行顺序](https://linux.cn/article-5690-rss.html)
 * [location](https://www.jianshu.com/p/5b4067f9fbcc)




	
### 配置文件详情nginx.conf

	user www www;
	worker_processes 8;
	
	//10000000表示启用第一个CPU内核，01000000表示启用第二个CPU内核
	worker_cpu_affinity 10000000 01000000 00100000 00010000 00001000 00000100 00000010 00000001;
	
	
	error_log /data/wwwlogs/error_nginx.log crit;
	pid /var/run/nginx.pid;
	worker_rlimit_nofile 51200;
	
	events {
	    use epoll;
	    worker_connections 51200;
	    }
	
	http {
	    include mime.types;
	    default_type application/octet-stream;
	    server_names_hash_bucket_size 128;
	    client_header_buffer_size 32k;
	    large_client_header_buffers 4 32k;
	    client_max_body_size 50m;
	    sendfile on;
	    tcp_nopush on;
	    keepalive_timeout 120;
	    server_tokens off;
	    tcp_nodelay on;
    
    fastcgi_connect_timeout 30;
    fastcgi_send_timeout 30;
    fastcgi_read_timeout 30;
    fastcgi_buffer_size 64k;
    fastcgi_buffers 4 64k;
    fastcgi_busy_buffers_size 128k;
    fastcgi_temp_file_write_size 128k;

    #Gzip Compression
    gzip on;
    gzip_buffers 16 8k;
    gzip_comp_level 6;
    gzip_http_version 1.1;
    gzip_min_length 256;
    gzip_proxied any;
    gzip_vary on;
    gzip_types
        text/xml application/xml application/atom+xml application/rss+xml application/xhtml+xml image/svg+xml
        text/javascript application/javascript application/x-javascript
        text/x-json application/json application/x-web-app-manifest+json
        text/css text/plain text/x-component
        font/opentype application/x-font-ttf application/vnd.ms-fontobject
        image/x-icon;
    gzip_disable  "msie6";

    #If you have a lot of static files to serve through Nginx then caching of the files' metadata (not the actual files' contents) can save some latency.
    open_file_cache max=1000 inactive=20s;
    open_file_cache_valid 30s;
    open_file_cache_min_uses 2;
    open_file_cache_errors on;


    server {
    listen 80;
    server_name _;
    access_log /data/wwwlogs/access_nginx.log combined;
    root /data/wwwroot/default;
    index index.html index.php;
    location ~ .*\.(php|php5)?$ {
        #fastcgi_pass remote_php_ip:9000;
        fastcgi_pass unix:/dev/shm/php-cgi.sock;
        fastcgi_index index.php;
        include fastcgi.conf;
        }
    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|flv|ico)$ {
        expires 30d;
        access_log off;
        }
    location ~ .*\.(js|css)?$ {
        expires 7d;
        access_log off;
        }
    }

	server {
	        listen      8080;
	        listen  443;
	        server_name cos.gameabc2.com;
	       # server_name 47.107.239.157;
	        ssl     on;
	        root    /data/release/mini-program/web;
	        index   index.php;
	        ssl_certificate /data/app/nginx/cert/gameabc2.com_combine.cer;
	        ssl_certificate_key     /data/app/nginx/cert/gameabc2.com.key;
	        ssl_session_timeout  5m;
	        ssl_session_cache shared:SSL:50m;
	        ssl_protocols TLSv1 TLSv1.1 TLSv1.2  SSLv2 SSLv3;
	        ssl_ciphers  ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
	        ssl_prefer_server_ciphers   on;

        location / {
            try_files $uri $uri/ /index.php$is_args$query_string;
        }
        location /nginx_status {
            default_type application/json;
            stub_status on;
            access_log off;
        }
        location = /status {
           # fastcgi_pass 127.0.0.1:9000;
            allow 127.0.0.1;
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
            include fastcgi_params;
           # fastcgi_pass 127.0.0.1:9000;
            fastcgi_pass unix:/dev/shm/php-cgi.sock;
        }

        location ~ \.php$ {
            try_files $uri =404;
           # fastcgi_pass 127.0.0.1:9000;
            fastcgi_pass unix:/dev/shm/php-cgi.sock;
            fastcgi_index index.php;
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
            fastcgi_param  YII_ENV 'prod';
            fastcgi_param  YII_DEBUG true;
            include fastcgi_params;
            include proxy.conf;
        }
    }


    include vhost/*.conf;
	}
	
	
	负载均衡配置
	upstream cos{
	      server 10.117.1.208:443 weight=100;
	      server 10.117.1.207:443 weight=100;
	      server 10.117.1.204:443 weight=1;
	}
	
	server {
	        listen      8080;
	        listen  443;
	        server_name cos.gameabc2.com;
	        ssl     on;
	        root    /data/release/mini-program/web;
	        index   index.php;
	        ssl_certificate /data/app/nginx/cert/gameabc2.com_combine.cer;
	        ssl_certificate_key     /data/app/nginx/cert/gameabc2.com.key;
	        ssl_session_timeout  5m;
	        ssl_session_cache shared:SSL:50m;
	        ssl_protocols TLSv1 TLSv1.1 TLSv1.2  SSLv2 SSLv3;
	        ssl_ciphers  ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
	        ssl_prefer_server_ciphers   on;

        location /cosplay/proof/test/ {
                proxy_pass https://cos;
        }

        location / {
            add_header Access-Control-Allow-Origin *;
            add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
            add_header Access-Control-Allow-Headers 'DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';
            proxy_request_buffering  off;
            proxy_pass https://cos/;

        }

        location /nginx_status {
            default_type application/json;
            stub_status on;
            access_log off;
        }

        location = /status {
           # fastcgi_pass 127.0.0.1:9000;
            allow 127.0.0.1;
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
            include fastcgi_params;
           # fastcgi_pass 127.0.0.1:9000;
            fastcgi_pass unix:/dev/shm/php-cgi.sock;
        }




    }
