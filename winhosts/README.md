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
