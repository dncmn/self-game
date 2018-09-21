#!/usr/bin/env bash

##关闭数据库
docker-compose -f mysql.yml down
docker-compose -f redis.yml down
docker-compose -f pg.yml down

##打开数据库
docker-compose -f mysql.yml up
docker-compose -f redis.yml up
docker-compose -f pg.yml up
##如果需要用navicat连接mysql,可以授权改密码什么的