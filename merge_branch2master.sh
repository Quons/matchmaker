#!/bin/bash
Branch=$1

echo $Branch

git checkout master

git merge $Branch
if  (($?==0)) then
 # 自动merge失败
 echo "you do not pass test"
 #reset
 git reset --hard
 #恢复到当前分支
 git checkout $Branch
 exit 1
fi
#merge成功，打tag,将代码推动到master
git tag
git push origin master