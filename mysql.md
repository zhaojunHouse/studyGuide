# mysql

### 读写分离
   * (why)降低读写加锁冲突。提高读性能。
   * (what)数据库分为主库，从库。主库只写，从库只读。
   * (how）代码区分读还是写；代理服务器；
   * (when）适合读压力大，少量写的情况
   
# redis

###  一台机器开启多个服务配置
   ```python
    port 6381
    pidfile /var/run/redis_6381.pid
    logfile /var/log/redis/redis.6381.log
    dbfilename dump.6381.rdb
    appendfilename "appendonly.6381.aof"
   ```
   
   
### 主从复制
    ```python
      slaveof host port
      slaveof no one
      master_repl_offset // 相同说明同步完成
   ```
   
   
   
