#!/bin/bash

# commit 前更新一遍 vendor
go mod tidy
go mod vendor
git add vendor/

PASS=true

# 静态错误检查
go vet ./...
if [[ $? != 0 ]]; then
  PASS=false
fi

STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep ".go$")

for FILE in $STAGED_GO_FILES
do
  # 跳过vendor文件
  if [[ $FILE == "vendor"* ]];then
    continue
  fi

  # goimport 检查/修正包导入
  goimports -w $FILE
  if [[ $? != 0 ]]; then
    PASS=false
  fi

  # go vet 静态错误检查
  # todo 上面做了全局 go vet 这里是否还需要对文件vet检查？
  go vet $FILE
  if [[ $? != 0 ]]; then
    PASS=false
  fi

  # gofmt 调整代码格式
  UNFORMATED=$(gofmt -l $FILE)
  if [[ $UNFORMATED != "" ]]; then
    gofmt -w $PWD/$UNFORMATED
    if [[ $? != 0 ]]; then
      PASS=false
    fi
  fi

done

if ! $PASS; then
    printf "\033[31m COMMIT FAILED \033[0m\n"
    exit 1
else
    printf "\033[32m COMMIT SUCCEEDED \033[0m\n"
fi

exit 0
