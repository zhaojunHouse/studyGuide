#nginx配置

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