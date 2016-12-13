#!/bin/bash
sed -i "" "s/sample/$1/g" `grep sample -rlI .`