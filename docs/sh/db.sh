#!/bin/bash

# 配置数据库连接信息
DB_HOST="106.14.15.23"
DB_PORT=3306
DB_USER="yusheng"
DB_PASS="Shys@1130"

path=$(dirname "$PWD")
echo path: ${path}

db_update() {
    # SQL 文件路径
    local DB_NAME=$1_$2
    SQL_FILE=${path}/sql/$1/$2.sql
    echo "SQL FILE: $SQL_FILE"
    # 检查 SQL 文件是否存在
    if [ ! -f "$SQL_FILE" ]; then
      echo "SQL NOT EXIST: $SQL_FILE"
      exit 1
    fi

    # 执行 SQL 文件
    mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$SQL_FILE"

    # 检查执行结果
    if [ $? -eq 0 ]; then
      echo "db update success"
    else
      echo "db update fail"
    fi
}

#db_update ys admin
#db_update ys front
#db_update main admin
#db_update main front

ysFront=ys_front
ysAdmin=ys_admin
mainFront=main_front
mainAdmin=main_admin

goCtl(){
  local DB_NAME=$1
  local modelDir=${path}/model/$1/

  # 获取所有表名
  TABLES=$(mysql -h "$DB_HOST" -u "$DB_USER" -p"$DB_PASS" -D "$DB_NAME" -e "SHOW TABLES;" | awk '{if(NR>1) print $1}')

  if test $1 ;then
    echo "开始创建库：$dbname 的表：$1"
  else
    echo "loss proto params $1 for example: ams pms"
    exit
  fi

#  goctl model mysql datasource -url="${$DB_USER}:${$DB_PASS}@tcp(${$DB_HOST}:${DB_PORT})/$1" -table="${tables}"  -dir="${modeldir}" -cache=true -i '' --home=${path}/1.6.0/  --style=gozero
  # 输出表名
  echo $1:"数据库中的表名有："
  for table in $TABLES; do
    echo "$table"
    goctl model mysql datasource -url="${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/$1" -table="${table}"  -dir="${modelDir}" -cache=true -i '' --home=${path}/1.6.0/  --style=gozero
  done
}

#goCtl ys_admin

# List of directories
#directories=("${ysFront}" "${ysAdmin}" "${mainFront}" "${mainAdmin}")
directories=("${mainFront}" "${mainAdmin}")
for dir in "${directories[@]}"; do
  goCtl "$dir"
done


