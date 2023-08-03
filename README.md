#- [Basic Blockchain](#basic-blockchain)
- [Ascii Art](#ascii-art)
- [Docker](#docker)
  - [Build](#build)

# Ascii Art

https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=SteamX

```
███████╗████████╗███████╗ █████╗ ███╗   ███╗██╗  ██╗
██╔════╝╚══██╔══╝██╔════╝██╔══██╗████╗ ████║╚██╗██╔╝
███████╗   ██║   █████╗  ███████║██╔████╔██║ ╚███╔╝ 
╚════██║   ██║   ██╔══╝  ██╔══██║██║╚██╔╝██║ ██╔██╗ 
███████║   ██║   ███████╗██║  ██║██║ ╚═╝ ██║██╔╝ ██╗
╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝
```

# Docker

Download and install the latest version of [Docker Desktop](https://www.docker.com/products/docker-desktop)

## Build

First build the client image

```bash
$ docker image ls
REPOSITORY           TAG       IMAGE ID       CREATED       SIZE

$ docker build -t go-rails/client:v1.10.21 .
...

$ docker image ls
REPOSITORY         TAG        IMAGE ID       CREATED          SIZE
go-rails/client    v1.10.21   4e8873698c07   2 minutes ago    49.3MB
```