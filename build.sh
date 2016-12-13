#!/bin/bash -x

# 部署系统会将项目在部署系统中的名字作为第一个参数传递给 build.sh
# 因为 httpprocess 所在的 git 源中包含多个模块,
# 所以我们需要通过这个参数区分不同模块的编译动作
name=$1
cluster=$(echo $name |sed -n 's/.*_cluster\.\([^._]*\)_.*/\1/p')
job=$(echo $name |sed -n 's/^job\.\([^._]*\)_.*/\1/p')
BUILDDIR=$( cd $(dirname $0) && pwd)

# make release and src dir
rm -rf release src ; mkdir -p release src/${job}
cp -r * src/${job}/ && rm -rf src/${job}/{release,src,build.sh}
export GOPATH=/home/work/tmp/go_${name}:${BUILDDIR}
rm -rf /home/work/tmp/go_${name}
export GOBIN=${BUILDDIR}/bin
export GOROOT=/usr/local/go1.5.3
export GOTOOLDIR=${GOROOT}/pkg/tool/linux_amd64

cd ${BUILDDIR}/src/${job}  # 进入子工程目录
gpm-git # 读取当前目录下Godeps-Git下载相关依赖
${GOROOT}/bin/go version
${GOROOT}/bin/go env
${GOROOT}/bin/go install # 生成可执行文件(会生成在$GOBIN中)
cp $GOBIN/${job} ${BUILDDIR}/release/  || exit 1 # 从$GOBIN拷贝可执行文件到release目录下
cd conf && for f in $(ls *_${cluster});do mv $f ${f%%_$cluster};done  && cd ..
cp -r conf/ ${BUILDDIR}/release/ # 拷贝配置文件到release目录下
cp -r deploy/ ${BUILDDIR}/release/
