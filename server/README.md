
### 生产环境

#### config.yaml部分修改
**不要全局直接复制，要每一行进行配置修改**
```yaml
cors:
    mode: strict-whitelist
    whitelist:
        - allow-origin: example1.com
        - allow-origin: example2.com
mysql:
    port: "3306"
    db-name: oms
    username: root
    password: Mysql123Admin456
    path: 127.0.0.1
redis:
    addr: 127.0.0.1:6379
    password: ""
    db: 0
system:
    env: release
    addr: 8888
    use-tls: false
    tls-cert: ./resource/cert/server.pem
    tls-key: ./resource/cert/server.key
zap:
    level: info
version: "v0.0.3"
```
>cors: whitelist下保留添加真实域名，删掉多余域名，注意格式
>mysql: 数据库主要配置
>redis: 缓存主要配置
>system: env(debug/release)[debug允许所有跨域] tls[相关开启https中间件]
>zap: level(debug/info/warn/...)[请求信息输出到log/debug.log]


### 开发环境

#### docker-redis启动
```sh
docker run --name redis \
-p 127.0.0.1:6379:6379 \
-v ~/Documents/data/redis/data:/data \
-d --restart=always redis:6.2.14
```

#### docker-mysql启动
```sh
docker run --name mysql \
-e MYSQL_ROOT_PASSWORD=Mysql123Admin456 \
-p 127.0.0.1:3306:3306 \
-v ~/Documents/data/mysql/data:/var/lib/mysql \
-v ~/Documents/data/mysql/log:/var/log/mysql \
-v ~/Documents/data/mysql/init/init.sql:/docker-entrypoint-initdb.d/init.sql \
-d --restart always mysql:8.0.36 \
mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

#### mysql最佳实践
* 1、sql一律小写
* 2、建表时先判断是否存在 if not exists
* 3、所有的列和表都加comment
* 4、字符串长度比较短时尽量使用char，定长有利于内存对齐，读写性能更好，<br>
varchar频繁修改容易产生内存碎片
* 5、满足需求前提下尽量使用短数据类型：<br>
tinyint vs int，float vs double，date vs datetime
* 6、null有别于''和0
* 7、尽量设为not null，比如default ''或者default 0<br>
有些DB索引列不允许包含null，对null统计时结果不符合预期，null值有时候拖慢系统性能
* 8、explain + select语句，查看是否走索引
* 9、mysql如何没有选择最优索引方案，在where前加force index ([index_name])
* 10、大部分的慢查询是因为没有正确使用索引（一般控制在50ms之内，有些公司要求100ms内）
* 11、一次select不要超过1000行
* 12、分页查询limit m,n会检索前m+n行，只返回n行，通常用id>x来代替这种分页方式
* 13、批量操作时最好是一条sql语句搞定；其次是打包成一个事务，一次性提交<br>
（高并发情况下减少对共享资源的争用）
* 14、不要使用连表操作，join逻辑在业务代码里完成
* 15、查看连接数 ```show processlist;```
* 16、查看连接超时时间 ```show global variables like 'wait_timeout'; -- 单位秒```
* 17、设置连接超时是时间 ```set global wait_timeout = 28800; -- 设置为28800秒，即8小时```

#### mysql防止sql注入
* 1、检查select username from user where username='"+username+"'
>lilei' or '1'='1注入后 select username from user where username='lilei' or '1'='1'
>查询到了所有用户
* 2、检查insert into student(name)values('"+username+"')
>lilei');drop table student;--注入后 insert into student(name)values('lilei');drop table student;--')
>--后边变成了注释，删掉整个表
* 3、入参加正则校验、长度限制
* 4、对特殊符号（<>&*;'"等）进行转义或编码转换，go的text/template包里边的HTMLEscapeString函数可以对字符串进行转义
* 5、不要将用户输入直接嵌入到sql语句中，不要使用字符串拼接，应该使用参数化查询接口
* 6、使用专业SQL注入检测工具检测，如sqlmap、SQLninja
* 7、避免将SQL错误信息返回到前端，以防止攻击者利用这些错误信息进行SQL注入

#### 特别鸣谢
[piexlmax](https://github.com/piexlmax)  https://www.gin-vue-admin.com/
[ZHOUYI](https://gitee.com/Z568_568)  https://gitee.com/Z568_568/ZY-Admin-template
