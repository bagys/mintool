[TOC]

# winhosts

> E:\> hosts

```sh
Usage:
  hosts [command]

command:
  ls
  add

Examples:
  hosts ls
```

> E:\> hosts ls

```sh
 - IP                      HOST
 * 192.168.31.165          host.docker.internal
 * 192.168.31.165          gateway.docker.internal
 * 127.0.0.1               kubernetes.docker.internal
 * 127.0.0.1               activate.navicat.com
 * 127.0.0.1               ban.com
 * 127.0.0.1               localhost
```

> E:\> hosts add

```sh
Error: required flag(s) "ip", "host" not set
Usage:
  hosts add

Examples:
  hosts add --ip=127.0.0.1 --host=localhost
  hosts add -i 127.0.0.1 -h localhost

Flags:
  -i, --ip string     register ip
  -h, --host string   register host
```



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
