#!/bin/bash

# 设置Go程序的名称
#AMS_NAME="ys_ams"
#PMS_NAME="ys_pms"
#UMS_NAME="ys_ums"
#OMS_NAME="ys_oms"
RPC_NAME="rpc_im"
API_NAME="api_front"


function stop(){

  # 检查是否存在正在运行的进程
  if ps aux | grep -v grep | grep "$1" > /dev/null; then
      # 获取进程ID并杀死进程
      PID=$(ps aux | grep -v grep | grep "$1" | awk '{print $2}')
      kill "$PID"
      echo "杀死进程成功"
  else
      echo "没有找到正在运行的进程"
  fi


  if lsof -i:$3 | grep LISTEN > /dev/null; then
      PORTID=$(lsof -i:$3 | grep LISTEN | awk '{print $2}')
      kill -9 $PORTID
      echo "杀死端口进程成功"
  fi
}

path=$(dirname $(dirname $(dirname "$PWD")))

stop ${RPC_NAME} ${path}/rpc_im 8653
stop ${API_NAME} ${path}/api_front 8686

