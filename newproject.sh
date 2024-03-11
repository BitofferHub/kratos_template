#!/bin/bash

# 获取参数
shortUrlX=$1

# 创建目标目录
mkdir "$shortUrlX"

# 拷贝 kratos_template 内容到 shortUrlX 目录
cp -R kratos_template/* "$shortUrlX"

# 修改目录名或文件名前缀为 shortUrlX
find "$shortUrlX" -depth -name "*userX*" -execdir bash -c 'mv "$1" "${1//userX/$2}"' bash {} "$shortUrlX" \;

# 修改文件内容中的 userX 字符串为 shortUrlX
find "$shortUrlX" -type f -exec sed -i 's/userX/'"$shortUrlX"'/g' {} +


echo "任务完成！"