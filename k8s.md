# k8s

	K8S，就是基于容器的集群管理平台，它的全称，是kubernetes。
	
	
# docker LNMP

	docker inspect php --format='{{.NetworkSettings.IPAddress}}'

	docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v /docker/mysql:/var/lib/mysql --name mysql5.7 mysql:5.7

	docker run -p 8081:80 --name nginx  -v /data/app/entry_api:/usr/share/nginx/html   -v /docker/nginx/conf.d:/etc/nginx/conf.d   --privileged=true  -d nginx

	docker run -p 9000:9000 --name php -v /data/app/entry_api:/var/www/html -d php:7.1-fpm

	docker run -p 6379:6379 —name redis -v /docker/data:/data  -d redis  redis-server --appendonly yes
	
## docker nginx配置
	server {
	
	  listen  80  default_server;
	
	  server_name _;
	
	
	  location / {
	      root   /usr/share/nginx/html/app/public;
	      index index.html index.htm index.php;
	      try_files $uri $uri/ /index.php?$query_string;
	  }
	
	
	
	  location ~ \.php(.*)$ {
	
	   root   /var/www/html/app/public;
	
	   fastcgi_pass 172.17.0.4:9000;   #php容器的IP地址
	
	   fastcgi_index index.php;
	
	   fastcgi_split_path_info ^((?U).+\.php)(/?.+)$;
	
	   fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
	
	   fastcgi_param PATH_INFO $fastcgi_path_info;
	
	   fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
	
	   include  fastcgi_params;
	
	  }
	
	}