# 说明
>  ygser 接口文档: [待配置](https://github.com/yugef3h/blogserver)

>  如果对您对此项目有兴趣，可以点 "Star" 支持一下 谢谢！ ^_^

>  apidoc 生成 api 文档
```
1、npm install apidoc -g
2、在 package.json 中配置 apidoc
3、app.js 中设置生成文档的路径，若我放在 public/apidoc 中，则 app.use('/public',express.static('public'));
4、public 文件夹下 创建 apidoc
5、根目录下运行 apidoc -i routes -o public/apidoc，routes 为接口注释文件
```

>  crawler 思路v1
```
1、https://www.zwdu.com/search.php?keyword=择天记   
2、过滤文档，找到 https://www.zwdu.com/book/8634/
3、先查询数据库，有则读取后断开
4、若数据库无，则启动爬虫，这里会遇到一种数据丢失的情况，可能原因列举如下：
min1、cheerio 模块获取的元素而非文本导致无数据
min2、查询到一本数据库中有的小说，再查询一本数据库中无的，注意这之间需要关闭连接池，再启动爬虫。
5、是否配置swagger
6.后台数据库查询重构，例如查找article勿需查找全部，booktitles需要修改列的组成，类似article表
```

>  如有问题请直接在 Issues 中提，或者您发现问题并有非常好的解决方案，欢迎 PR 👍

>  相关项目地址：[前端项目地址](http://www.blackatall.cn)  
