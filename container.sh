#!/bin/bash

args=`getopt pgx $*`
if [ $? != 0 ]
then
  echo "Failed to getopt."
  exit 2
fi
set -- $args
for i
do
  case "$i"
  in
    -p)
      echo build postgres image.; build_postgres=1
      shift;;
    -g)
      echo build golang image.; build_golang=1
      shift;;
    -x)
      echo build nginx image.; build_nginx=1
      shift;;
    --)
      shift; break;;
  esac
done

if [ ${build_postgres:=0} -gt 0 ]; then
  docker build -t takasing/postgres:0.1 /home/core/share/vagrant/postgres
fi

if [ ${build_golang:=0} -gt 0 ]; then
  docker build -t takasing/golang:0.1.1 /home/core/share
fi

if [ ${build_nginx:=0} -gt 0 ]; then
  docker build -t takasing/nginx:0.1 /home/core/share/vagrant/nginx
fi

psgr=`docker ps -a | grep postgres | wc -l`
if [ $psgr -lt 1 ]; then
  docker run \
    --name postgres \
    -itd \
    takasing/postgres:0.1
fi

glg=`docker ps -a | grep golang | wc -l`
if [ $glg -lt 1 ]; then
  docker run \
    --name golang \
    --link postgres:db \
    --privileged \
    -v /home/core/share:/go/src/golang-gin-api \
    -itd \
    takasing/golang:0.1.1
  docker exec golang sysctl -w fs.inotify.max_user_watches=524288
fi
 
ngx=`docker ps -a | grep web | wc -l`
if [ $ngx -lt 1 ]; then
docker run \
  --name web \
  --link golang:api \
  -p 80:80 \
  -d \
  takasing/nginx:0.1
fi

