#!/bin/bash
echo 开始执行时间 `date +"%Y-%m-%d %H:%M:%S"`
echo "现在拉取分支更新"
git fetch -p > run.log
echo `date +"%Y-%m-%d %H:%M:%S"`
cat run.log
echo "拉取代码"
git pull > run.log
echo `date +"%Y-%m-%d %H:%M:%S"`
cat run.log

local_dir=`pwd`
module_path=./../go-share-library
cd ${module_path}
path_dir=`pwd`
chmod 777 "${local_dir}/go.mod"
echo replace" "github.com/jackSpanrrows/go-share-library "=>" "${path_dir}" >> "${local_dir}/go.mod"

cd ${local_dir}
echo "更新模块版本"
go mod tidy
go mod download
echo All update complete `date +"%Y-%m-%d %H:%M:%S"`
