### docker
+ yum install -y yum-utils
+ yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
+ sudo yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin
+ sudo systemctl start docker
#### mysql
+ docker pull mysql:5.7
+ mkdir -p  /data/mysql/data  /data/mysql/conf /data/mysql/logs
+ ` [client]
  port=3306
  default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
character_set_server=utf8
secure_file_priv=/var/lib/mysql
log_bin_trust_function_creators=1
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
log-bin=/var/lib/mysql/mysql-bin # 开启log-bin
server-id=123654
expire_logs_days=30
lower_case_table_names=1 # 配置大小写不敏感
innodb_buffer_pool_size=2G # 性能优化：具数值根据自己内存而定`
+ cd /data/mysql/conf
+ chmod 644 my.cnf
+ docker run -it --name mysql5.7 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root --privileged=true -v /data/mysql/conf/my.cnf:/etc/mysql/my.cnf -v /data/mysql/data:/var/lib/mysql -v /data/mysql/logs:/var/log/mysql -v /etc/localtime:/etc/localtime:ro -d --restart=always mysql:5.7
+ docker exec -it mysql /bin/bash
+ mysql -uroot -p
#### redis
+ docker pull redis
+ docker run -d -p 6379:6379 --restart always --name some-redis \
  -v $PWD/conf/redis.conf:/etc/redis/redis.conf \
  -v $PWD/data:/data \
  redis redis-server /etc/redis/redis.conf \
  --requirepass "123456" --appendonly yes
+ `requirepass 123456  #默认空, 连接时需要输入的密码
  appendonly yes  # redis持久化（可选）
  databases 16    # 数据库个数（可选），可以改改看，看看能不能生效
  port 6379   # redis监听的端口号
  daemonize no    # 默认no，改为yes意为以守护进程方式启动，可后台运行，除非kill进程，改为yes会使配置文件方式启动redis失败
  protected-mode no   # 默认yes，开启保护模式会限制为本地访问
  bind 127.0.0.1        # 注释掉这部分，这是限制redis只能本地访问`
+ docker restart some-redis
+ docker exec -it redis redis-cli
#### mongo
+ docker pull mongo:4.2.2
+ docker run -d --name mongo -v mongo_data_configdb:/data/configdb -v mongo_data_db:/data/db -p 27017:27017 mongo:4.2.2 --auth
+ docker exec -it mymongo mongo admin 进入后创建用户 
+ db.createUser({ user: 'admin', pwd: 'qwe258ZXC', roles: [ { role: "root", db: "admin" } ] });
