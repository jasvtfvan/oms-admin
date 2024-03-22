#!/bin/sh

# 柔性停止debug
echo "#!/bin/sh --------:柔性停止debug"

CURRENT_DIR=$(dirname "$0")

echo "#!/bin/sh --------:当前目录: $CURRENT_DIR"

# 从pid.txt文件中读取PID
pid=$(cat $CURRENT_DIR/pid.txt)

# 检查PID是否为空
if [ -z "$pid" ]; then
  echo "#!/bin/sh --------:Error: PID is empty or pid.txt file is not found."
  exit 1
fi

# 使用kill命令结束进程
kill "$pid"

# 检查kill命令是否成功
if [ $? -eq 0 ]; then
  echo "#!/bin/sh --------:Process with PID $pid has been killed."
else
  echo "#!/bin/sh --------:Failed to kill process with PID $pid."
fi
