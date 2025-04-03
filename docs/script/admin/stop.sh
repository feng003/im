#!/bin/bash

# 设置Go程序的名称
#AMS_NAME="admin_ams"
#PMS_NAME="admin_pms"
#UMS_NAME="admin_ums"
RPC_CMS="rpc_cms"
RPC_NAME="rpc_admin"
API_NAME="api_admin"


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

#stop ${AMS_NAME} ${path}/ams/admin 8611
#stop ${PMS_NAME} ${path}/pms/admin 8612
#stop ${UMS_NAME} ${path}/ums/admin 8613
stop ${RPC_NAME} ${path}/rpc_admin 8651
stop ${RPC_CMS} ${path}/rpc_cms 8653
stop ${API_NAME} ${path}/api_admin 8688

