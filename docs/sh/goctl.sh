#!/usr/bin/env bash

generate_code() {
    local mode=$1
    local api_param=$2
    local rpc_param=$3

    # 获取路径并设置变量
    # shellcheck disable=SC2155
    local path=$(dirname $(dirname "$PWD"))
    local dir
    local homeDir=${path}/docs/1.6.0/
    local apiDir=${path}/${api_param}
    local rpcDir=${path}/${rpc_param}
    local rpcName=$(echo ${rpc_param} | sed 's/\//_/g')

    if [[ "$mode" == "front" ]]; then
        dir=${path}/docs/script/front
    elif [[ "$mode" == "admin" ]]; then
        dir=${path}/docs/script/admin
    else
        echo "Invalid mode: $mode. Use 'front' or 'admin'."
        exit 1
    fi

    echo "mode is ${mode} and apiDir: ${apiDir} and rpcDir: ${rpcDir} and dir is ${dir}"

    # 检查参数 $2 和 $3 是否存在
    if [[ -z "$api_param" || -z "$rpc_param" ]]; then
        echo "loss api params $api_param or rpc params $rpc_param; for example api_${mode} rpc_${mode}"
        exit 1
    else
        echo "generate api $api_param and rpc $rpc_param"
    fi

    # 确保 proto 文件存在
    local proto_file=${dir}/proto/${rpcName}.proto
    if [[ ! -f "$proto_file" ]]; then
        echo "Proto file not found: $proto_file"
        exit 1
    fi

    # 生成 API 和 RPC 代码
    goctl api go -api "${dir}"/${mode}.api -dir ${apiDir} -style gozero --home ${homeDir}
    goctl rpc protoc --proto_path="${dir}"/proto --go_out=${rpcDir} --go-grpc_out=${rpcDir} --zrpc_out=${rpcDir} --style gozero ${proto_file} -m
    goctl api plugin -plugin goctl-swagger="swagger -filename api_${mode}.json" -api ${dir}/${mode}.api -dir ./${mode}/
}

generate_code "admin" "api_admin" "rpc_admin"
generate_code "front" "api_front" "rpc_front"