# dockerfile は dockerimage を作成するための設定ファイル
# Goの公式イメージからダウンロードするgoのバージョン指定
# 指定したバージョンでコンテナ上でGoが実行される
FROM golang:1.20

# 下記の命令で、Dockerfile の後続の命令(COPY, RUN..)
# がその指定したディレクトリを対象とする
# つまりgo-todo-backendを対象に COPYやRUNが行われる
WORKDIR /go-todo-backend

# 前提として、dockerfileはdockerimageを作成するための指示書やルール
# 実際にコンテナを作成するのはdockeriamgeであり、dockeriamgeをビルドすることでコンテナが作成される
# 下記でCOPYやRUNを記載しているのは、dockerfileをもとに作成した
# dockerimageをビルド(コンテナ作成)するタイミングでCOPY,ダウンロードしてねという
# 指示をしているだけ
# 実際にCOPYやRUNを行うのはdockerimage
# dockerfileはあくまでdockerimageの元となるもの
COPY go.mod .
COPY go.sum .
RUN go mod download

# docker は独立した環境で動作するため、goをdocker上で起動させたい場合、
# 実行させるためのソースコードも必要
# 下記はワークングディレクトリ配下全てをコンテナにコピーするコマンド
# コンテナ上にソースコードをコピーすることでプログラムを
# コンテナ上で動作させられる
COPY . .

# RUN go build -o main .このコマンドはdockerimageがビルド(コンテナ作成)
# される際に、コンテナ上でソースコードをコンパイルし、バイナリファイル(main)として
# 作成される命令
# つまりコンテナ上にはソースコードがそのまま配置されるのではなく、コンパイルした
# バイナリファイルが配置される
WORKDIR /go-todo-backend/cmd
RUN go build -o main

# 下記はコンテナが実行される際にデフォルトで実行される
# コマンドどを指定している
# つまり、コンパイルして作成したバイナリファイル名がmain
# の場合、下記で./mainを指定しているため、コンテナ起動時に
# プログラムが自動で実行されることを意味する
CMD ["./main"]