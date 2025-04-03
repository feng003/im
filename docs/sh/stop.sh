#!/bin/bash

JOB="job"
SCHEDULER="scheduler"

function stop(){

  # 检查是否存在正在运行的进程
  # shellcheck disable=SC2009
  if ps aux | grep -v grep | grep "$1" > /dev/null; then
      # 获取进程ID并杀死进程
      # shellcheck disable=SC2009
      PID=$(ps aux | grep -v grep | grep "$1" | awk '{print $2}')
      kill "$PID"
      echo "杀死进程成功"
  else
      echo "没有找到正在运行的进程"
  fi

  if lsof -i:"$3" | grep LISTEN > /dev/null; then
      PORTID=$(lsof -i:"$3" | grep LISTEN | awk '{print $2}')
      kill -9 "$PORTID"
      echo "杀死端口进程成功"
  fi
}

path=/home/zhang/go/src/yusheng/mqueue

stop ${JOB} ${path}/job 8672
stop ${SCHEDULER} ${path}/scheduler 8671