#!/usr/bin/env bash
docker-compose -f mysql.yml up
docker-compose -f redis.yaml up
##如果需要用navicat连接mysql,可以授权改密码什么的