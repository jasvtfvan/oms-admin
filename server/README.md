
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
    username: mysql_admin
    password: Mysql123Admin456
    path: rm-2ze4vq08r30l8n032.mysql.rds.aliyuncs.com
redis:
    addr: 127.0.0.1:6379
    password: Redis123Admin456
    db: 0
system:
    username: oms_admin
    password: Oms123Admin456
    init-pwd: Oms123Admin456
    env: release
    addr: 8888
    use-tls: false
    tls-cert: ./resource/cert/server.pem
    tls-key: ./resource/cert/server.key
zap:
    level: info
version: "v0.0.1"
```
>cors: whitelist下保留添加真实域名，删掉多余域名，注意格式
>mysql: 数据库主要配置
>redis: 缓存主要配置
>system: env(debug/release)[debug允许所有跨域] tls[相关开启https中间件]
>zap: level(debug/info/warn/...)[请求信息输出到log/debug.log]


### 开发环境

#### 代码更新
* 特殊代码简介
1. initializer用来注册所有初始化程序，一个表（或复杂结构）对应一个go文件，
其中init_impl.go和register_tables.go不需要修改，属于封装内容
2. updater用来注册所有更新程序，一个表（或复杂结构）对应一个go文件，
其中register_tables.go不需要修改，属于封装内容
3. 在service/initialize中，实现了对initializer和updater的统一调度

* 数据库更新，需更新内容：
1. /model/模块文件夹/表结构.go 添加/修改要升级的内容
1. /global/order_vars.go 更新对应的初始化顺序
2. /initialize/updater/模块文件夹/表结构.go 添加/修改要升级的内容
3. /initialize/initializer/模块文件夹/表结构.go 更新的内容同步到初始化逻辑
4. /initialize/migrate/tables.go 没经过initializer/updater的，需初始化/更新的空表

#### RBAC模型说明
1. group为树结构设计，根group只有一个，其他公司/集团属于根的下级，所有业务数据根据group隔离
2. role设计比较特殊，依赖group存在，不能单独存在，一个group下可以创建多个role
3. user与group是多对多关系，user与role是多对多关系，user跟group关联决定当前组织，可以在多个group间来回切换，但不能同时操作多个group，而具体权限要根据role去匹配
```golang
[relationship]              [description]
group   1<-->n  group       tree结构
group   1<-->n  role        role属于group
user    n<-->n  group       user可以为多个group工作
user    n<-->n  role        user切换到指定group通过role确定权限
group   1<-->n  business    user切换到指定group，每条业务数据隶属于group
```

#### 权限策略
1. 通过JWT进行token验证，非cookie以便多客户端兼容
2. 权限模型使用RBAC with domains的casbin模型，将role、group、http_path、http_method放到policy中
3. 只将部分api（或资源）需要权限验证，针对操作基础数据api或特殊权限api，普通api不做严格限制
4. 请求时需要header加入group的code，便于组织隔离
5. 是否记录操作记录，根据用户的logOption标志

#### swagger
1. 安装全局swagger命令工具
```golang
go install github.com/swaggo/swag/cmd/swag@latest
```
2. 查看swagger命令工具版本
```golang
swag -v
```
3. 修改main.go的注释信息，修改各个接口的注释信息
4. 在/server目录下，执行一键生成docs命令
```golang
swag init
``` 
5. 根据控制台输出的文档地址，打开浏览器调试，例如：
http://127.0.0.1:8888/swagger/index.html

#### docker-redis启动
```sh
docker run --name oms-redis \
-p 127.0.0.1:6379:6379 \
-v ~/Documents/data/redis/data:/data \
-d --restart=always redis:6.2.14
```

#### docker-mysql启动
```sh
docker run --name oms-mysql \
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
* 15、查看连接情况 ```show processlist;```
* 16、根据用户查看连接情况 ```SELECT * FROM information_schema.processlist WHERE USER='mysql_admin';```
* 17、查看连接超时时间 ```show global variables like 'wait_timeout'; -- 单位秒```
* 18、设置连接超时是时间 ```set global wait_timeout = 7200; -- 设置为7200秒，即2小时```

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


### 常见解决方案

#### 忘记密码
登录管理员，重置用户密码
默认密码 username123456 ，登录后发现密码小于6位、没有大小写字母、没有数字，强制修改

#### 超级管理员忘记密码
1. 本地环境运行 generator_test.go TestBcryptHash方法
2. 将所得密文手动更新到mysql
3. 手动连接redis，找到jwt_管理员username，清除
4. 告诉管理员新密码，尝试登录


### 附录

#### 参考文献
[gin](https://learnku.com/docs/gin-gonic/1.7)
[gorm](https://learnku.com/docs/gorm/v2)
[casbin](https://casbin.org/zh/)

#### 特别鸣谢
[piexlmax](https://github.com/piexlmax)  https://www.gin-vue-admin.com/
[ZHOUYI](https://gitee.com/Z568_568)  https://gitee.com/Z568_568/ZY-Admin-template
[https://www.vue3js.cn/](https://www.vue3js.cn/)
[https://buqiyuan.github.io/vue3-antdv-admin-docs/](https://buqiyuan.github.io/vue3-antdv-admin-docs/)
