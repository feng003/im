#!/bin/bash

JOB="job"
SCHEDULER="scheduler"

#go mod tidy

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

  if lsof -i:$3 | grep LISTEN > /dev/null; then
      PORTID=$(lsof -i:$3 | grep LISTEN | awk '{print $2}')
      kill -9 $PORTID
  fi

  # shellcheck disable=SC2034
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64
  go build -o $1 -buildvcs=false

  echo "编译成功 $1"
  sleep 1
  ./$1 &
  echo "重新启动程序 $1"
}

# shellcheck disable=SC2046
path=$(dirname $(dirname "$PWD"))

#echo ${path}

restart ${SCHEDULER} "${path}"/mqueue/scheduler 8671
restart ${JOB}  "${path}"/mqueue/job 8672
