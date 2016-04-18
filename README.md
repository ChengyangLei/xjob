## xjob
### 代码地址：https://github.com/vwufeng/xjob

### 代码目录
crawler 拉勾网Go语言职位爬取  
web 数据展示  
util 公有函数  
assets 图片、css、js文件  
templates html模板、高德地图地铁路线api js请求    

### 运行代码
1、配置config.ini文件，格式如下：  
[mysql]  
user = xx  
pwd = xx  
host = xx  
port = xx  
db = xx  

[gaode]  
key = xxxxx  

[mysql]是数据库配置参数，[gaode]是高德地图开发者账户的key。  

2、执行爬虫程序：  
```bash
go build -o spider github.com/vwufeng/xjob/crawler
crawler
```

3、执行web程序：  
```bash
go build -o site github.com/vwufeng/xjob/web
web
```

4、查看效果：  
访问http://localhost:8080/job