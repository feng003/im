#!/bin/bash

# 设置Go程序的名称
RPC_ADMIN="rpc_admin"
RPC_CMS="rpc_cms"
API_ADMIN="api_admin"

function restart(){

  # 检查是否存在正在运行的进程
  if ps aux | grep -v grep | grep "$1" > /dev/null; then
      # 获取进程ID并杀死进程
      PID=$(ps aux | grep -v grep | grep "$1" | awk '{print $2}')
      kill "$PID"
      echo "杀死进程成功"
  else
      echo "没有找到正在运行的进程"
  fi

  cd $2 && pwd
  rm -rf $2/logs/*

#  CGO_ENABLED=0 GOOS=linux GOARCH=amd64

  if lsof -i:$3 | grep LISTEN > /dev/null; then
      PORTID=$(lsof -i:$3 | grep LISTEN | awk '{print $2}')
      kill -9 $PORTID
  fi

  go build -o $1 -buildvcs=false -ldflags="-s -w"
  echo "编译成功"

  ./$1 &
  echo "重新启动程序 $1"
}

path=$(dirname $(dirname $(dirname "$PWD")))


restart ${RPC_ADMIN} ${path}/rpc_admin 8651
#restart ${RPC_CMS} ${path}/rpc_cms 8653
restart ${API_ADMIN} ${path}/api_admin 8688



