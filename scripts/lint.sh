#!/bin/bash

# vendor および テストファイルを除く全ての Go ファイルを検索
FILES=$(find . -name "*.go" ! -path "./vendor/*" ! -name "*_test.go")

error_found=false

for FILE in $FILES
do
  echo "===================="
  echo "$FILE をチェック中"
  echo "===================="

  # コメントとそれに続く関数を抽出
  awk '/\/\/.*$/{comment=$0; next} /^func/{print comment "\n" $0}' $FILE | while read -r line
  do
    # 関数名を取得
    func_name=$(echo "$line" | awk '/^func/{print $2}' | awk -F '(' '{print $1}')
    # コメントが存在し、それが関数名と一致しているかチェック
    if echo "$line" | awk -v fn="$func_name" '/\/\// && $0 ~ fn {print}' &> /dev/null; then
      echo "関数 $func_name には適切な関数説明文が定義されています"
    else
      echo "エラー: 関数 $func_name には適切な関数説明文が定義されていません"
      error_found=true
    fi
  done
  echo ""
done

if [ "$error_found" = false ] ; then
    echo "=================================="
    echo "すべての関数に適切な紹介文が定義されています！！"
else
    echo "=================================="
    echo "一部の関数に適切な紹介文が定義されていません。修正してください。"
fi