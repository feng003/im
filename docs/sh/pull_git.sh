#!/usr/bin/env bash

path=$(dirname $(dirname "$PWD"))

#ams=${path}/ams/rpc
#pms=${path}/pms/rpc
#ums=${path}/ums/rpc
#admin=${path}/admin
apiFront=${path}/api_front
apiAdmin=${path}/api_admin
rpcAdmin=${path}/rpc_admin
rpcFront=${path}/rpc_front
common=${path}/common
mqueue=${path}/mqueue
model=${path}/model
docs=${path}/docs
date=`date +%Y%m%d`

for dir in ${rpcFront} ${rpcAdmin}  ${apiFront} ${apiAdmin} ${common} ${model} ${docs} ${mqueue}
do
  cd ${dir}
#  git branch
  git pull
  echo "complete ${dir} - ${date}"
done


