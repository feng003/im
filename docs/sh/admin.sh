#!/bin/bash

path=$(dirname $(dirname "$PWD"))
echo "path" ${path}
exit

GO_PATH=(
    ${path}"/rpc_admin/"
    ${path}"/rpc_front/"
    ${path}"/api_admin/"
    ${path}"/api_front/"
)

GO_NAME=(
  "rpcadmin.go"
  "rpcfront.go"
  "admin.go"
  "front.go"
)

# 定义端口号数组
PORTS=(
    "8651"
    "8652"
    "8688"
    "8686"
)

# 启动Go进程函数
start_processes() {
    for ((i=0; i<${#GO_PATH[@]}; i++)); do
        cd ${GO_PATH[$i]} && pwd
        file=${GO_NAME[$i]}
        port="${PORTS[$i]}"
        if pgrep -f "$file" > /dev/null; then
            echo "Process for $file is already running."
        else
            echo "Starting Go process for $file on port $port..."

            go build -o $1 -buildvcs=false
            echo "编译成功"
#            nohup go run "$file" > /dev/null 2>&1 &
#            # 等待进程执行完成
#            pid=$(ps -ef | grep -v grep | grep "$file" | awk '{print $2}')
#            while ps -p $pid > /dev/null; do
#                sleep 1
#            done
            echo "Go process for $file started."
        fi
    done
}

# 停止Go进程函数
stop_processes() {
    for port in "${PORTS[@]}"; do
        process_id=$(lsof -t -i:$port)
        if [ -n "$process_id" ]; then
            echo "Stopping process listening on port $port..."
            if kill "$process_id"; then
                echo "Process listening on port $port stopped successfully."
            else
                echo "Failed to stop process listening on port $port."
            fi
        else
            echo "No process listening on port $port."
        fi
    done
}



# 根据命令参数执行相应操作
case "$1" in
    start)
        start_processes
        ;;
    stop)
        stop_processes
        ;;
    *)
        echo "Usage: $0 {start|stop}"
        exit 1
        ;;
esac

exit 0