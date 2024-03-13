#### docker-mysql启动
```sh
docker run --name gva-mysql \
-e MYSQL_DATABASE=qmPlus \
-e MYSQL_USER=gva \
-e MYSQL_ROOT_PASSWORD=Aa@6447985 \
-p 127.0.0.1:3306:3306 \
--restart always \
-v ~/Documents/data/mysql:/var/lib/mysql \
-d mysql:8.0.36 \
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
