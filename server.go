package main

// 利用するパッケージを追加
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// main() は、実行されるプログラムが書かれた関数
func main() {
	// var r = gin.Default() とも書ける
	// := は省略形
	r := gin.Default()

	// ルーティング設定
	r.GET("/", handleRoot)

	// サーバを立ち上げる
	r.Run() // listen and serve on 0.0.0.0:8080
}

// GET "/" にアクセスが来たときに呼び出されている関数
func handleRoot(c *gin.Context) {
	// JSONを返す
	c.JSON(http.StatusOK, gin.H{
		"id":   0,
		"name": "asuka",
	})
}
