#!/usr/bin/env bash

generate_code() {
    # 获取路径并设置变量
    # shellcheck disable=SC2155
    # shellcheck disable=SC2046
    local path=$(dirname $(dirname $(dirname "$PWD")))
    local dir=${path}/docs/script/front/
    local apiDir=${path}/$2
    local homeDir=${path}/docs/1.6.0/
    local rpcDir=${path}/$MAIN_RPC_PROTO
#    local rpcName=$(echo $MAIN_RPC_PROTO | cut -d'/' -f1)

    echo "path: ${path}"

    # 检查参数 $2 和 $MAIN_RPC_PROTO 是否存在
    if [[ -z "$2" || -z "$MAIN_RPC_PROTO" ]]; then
        echo "loss api params $2 or rpc params $MAIN_RPC_PROTO ;for example api_front oms/rpc"
        exit 1
    else
        echo "generate api $2 and rpc $MAIN_RPC_PROTO"
    fi

    # 生成 API 和 RPC 代码
    goctl api go -api ${dir}/front.api -dir ${apiDir} -style gozero --home ${homeDir}
    # 生成主proto
    cd ./proto
    goctl rpc protoc $MAIN_RPC_PROTO.proto --go_out=${rpcDir} --go-grpc_out=${rpcDir} --zrpc_out=${rpcDir} --style gozero -m
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
    goctl api plugin -plugin goctl-swagger="swagger -filename api_front.json" -api front.api -dir .
}
# 主proto和引入的其他proto用于支持，rpc下多个proto文件
# 主proto
MAIN_RPC_PROTO="rpc_front"
# 引入的其他proto
OTHER_PROTO_LIST=("rpc_front_heepay" "rpc_front_oms" "rpc_front_pms")
generate_code "front" "api_front"
