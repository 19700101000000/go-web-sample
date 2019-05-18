# go-web-sample
`gin-gonic`
を利用したgolangのapiサーバのサンプルソースです。  
こんな感じで書くよ〜くらいの気持ちで見てね。  
導入の流れをやってみようｂ

## 前提条件
1. [golang](https://golang.org/)
が入っていないと動きません
1. webサーバーライブラリとして、
[gin-gonic](https://github.com/gin-gonic/gin)
を利用します

## 導入
1. [golang](https://golang.org/)
を入れていない人は入れてください
※最新バージョンを入れてください
1. `export GO111MODULE=on`
と環境パスを追加するコマンドを実行します
1. `go build` とコマンドを実行すると、
`./sample` という実行ファイルが出来上がります
1. `./sample` を実行すると、webサーバーが建ちます
※ `localhost:8080` で立ち上がったことを確認できます
