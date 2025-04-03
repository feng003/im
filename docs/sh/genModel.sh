#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
path=$(dirname "$PWD")
echo path: ${path}

modeldir=${path}/model/$1/

# 数据库配置
host=106.14.15.23
port=3306
dbname=main_$1
username=yusheng
passwd=Shys@1130

if test $1 ;then
  echo "开始创建库：$dbname 的表：$1"
else
  echo "loss proto params $1 for example: ams pms"
  exit
fi

goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true -i '' --home=${path}/1.6.0/  --style=gozero