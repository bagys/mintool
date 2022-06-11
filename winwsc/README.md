# wsc
window service 练习小工具

> E:\wsc
```sh
Usage:
  wsc [command] [service]

Commands:
  ls
  restart
  start
  startall
  stop
  stopall

Examples:
  wsc ls
  wsc startall nginx
  wsc start nginx
```

> E:\wsc start
```sh
Usage:
  wsc start
  wsc start <service>

Service:
  mysql
  nginx
  phpfpm
  redis
```

> E:\wsc start nginx
```sh
Nginx 服务正在启动 .
Nginx 服务已经启动成功。
```

> E:\wsc ls
```sh
SERVICE          DESCRIBE             STATUS
nginx            nginx serer          running
redis            redis serer          running
phpfpm           nginx serer          running
mysql            mysql serer          running
```