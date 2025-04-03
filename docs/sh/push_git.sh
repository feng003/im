#!/usr/bin/env bash


path=$(dirname $(dirname "$PWD"))

ams=${path}/ams/rpc
adminAms=${path}/ams/admin/
pms=${path}/pms/rpc
adminPms=${path}/pms/admin/
ums=${path}/ums/rpc
apiFront=${path}/api_front
apiAdmin=${path}/api_admin
common=${path}/common
model=${path}/model
date=`date +%Y%m%d`

for dir in ${ams} ${adminAms} ${pms} ${adminPms} ${ums} ${apiFront} ${common} ${model} ${apiAdmin}
do
  cd ${dir}
  git add .
  git commit -m "add ${date}"
  git push
  echo "complete ${dir} - ${date}"
done


