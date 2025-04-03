#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package



#表生成的genmodel目录
path=$(dirname "$PWD")
protodir=${path}/proto

# 数据库配置
host=106.14.15.23
port=3306
dbname=main_$1
username=yusheng
passwd=Shys@1130

if test $1 ;then
  echo "开始创建proto : $1"
else
  echo "loss proto params $1 for example: ams pms"
  exit
fi

sql2pb -go_package ./pb -host ${host} -package pb -user ${username} -password ${passwd} -port ${port} -schema ${dbname} -service_name $1  > ${protodir}/$1.proto
