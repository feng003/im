#!/usr/bin/env bash

generate_code() {
    local mode=$1
    local api_param=$2
#    local rpc_param=$3

    # 获取路径并设置变量
    local path=$(dirname $(dirname $(dirname "$PWD")))
    local dir
    local homeDir=${path}/docs/1.6.0/
    local apiDir=${path}/${api_param}
    local rpcDir=${path}/${MAIN_RPC_PROTO}
#    local rpcName=$(echo ${MAIN_RPC_PROTO} | sed 's/\//_/g')

    if [[ "$mode" == "front" ]]; then
        dir=${path}/docs/script/front/
    elif [[ "$mode" == "admin" ]]; then
        dir=${path}/docs/script/admin/
    else
        echo "Invalid mode: $mode. Use 'front' or 'admin'."
        exit 1
    fi

    echo "mode is ${mode} and apiDir: ${apiDir} and rpcDir: ${rpcDir}"

    # 检查参数 $2 和 $3 是否存在
    if [[ -z "$api_param" || -z "$MAIN_RPC_PROTO" ]]; then
        echo "loss api params $api_param or rpc params $MAIN_RPC_PROTO; for example api_${mode} rpc_${mode}"
        exit 1
    else
        echo "generate api $api_param and rpc $MAIN_RPC_PROTO"
    fi

    # 生成 API 和 RPC 代码
    goctl api go -api ${dir}${mode}.api -dir ${apiDir} -style gozero --home ${homeDir}
    # 生成主proto
    cd ./proto
    goctl rpc protoc ./$MAIN_RPC_PROTO.proto --go_out=${rpcDir} --go-grpc_out=${rpcDir} --zrpc_out=${rpcDir} --style gozero -m
    # 生成引入的其他proto
    for OTHER_PROTO in "${OTHER_PROTO_LIST[@]}"; do
      # 生成二级proto
      goctl rpc protoc $OTHER_PROTO.proto --go_out=${rpcDir} --go-grpc_out=${rpcDir} --zrpc_out=${rpcDir} --style gozero -m
      # 删除生成的二级入口文件
      MAIN_FIAL=${OTHER_PROTO//_/}
      MAIN_GO_FILE=${rpcDir}/$MAIN_FIAL.go
      echo "删除生成的二级入口文件: ${MAIN_GO_FILE}"
      rm $MAIN_GO_FILE
    done

    # 生成api
    cd ../
    goctl api plugin -plugin goctl-swagger="swagger -filename api_${mode}.json" -api ${mode}.api -dir .
}
# 主proto和引入的其他proto用于支持，rpc下多个proto文件
# 主proto
MAIN_RPC_PROTO="rpc_admin"
# 引入的其他proto
OTHER_PROTO_LIST=("rpc_admin_heepay" "rpc_admin_antchain")
generate_code "admin" "api_admin"
