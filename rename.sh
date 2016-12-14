#!/bin/bash
if [ $# != 1 ]; then
echo usage: $0 NewProjectName
exit 1;
fi

sed -i "" "s/sample/$1/g" `grep sample -rlI .`