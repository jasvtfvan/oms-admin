# oms-admin

#### 介绍
全域运行管理系统

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)


## 配置2套密钥对，用于个人和企业分开
1. 生成新的密钥对
```ssh-keygen -t rsa -C fan.z@snapinspect.com```
2. 修改密钥对文件名，不要覆盖默认的
```sh
Enter file in which to save the key (/Users/jasvtfvan/.ssh/id_rsa): /Users/jasvtfvan/.ssh/id_rsa_snapinspect
```
3. 查看`ssh agent`
```ssh-add -l```
4. 添加新密钥到`agent`，其中`-K`放到keychain中，再次查看`agent`
```ssh-add -K ~/.ssh/id_rsa_snapinspect```
5. 复制对应的公钥到`github`，推拉代码
