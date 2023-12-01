package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// データベース接続の設定
	config := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "myapp",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("データベースに接続しました。")

	// ルーターの設定
	router := gin.Default()
	setupRouter(router)

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func setupRouter(router *gin.Engine) {
	// CORS設定
	router.Use(corsMiddleware())

	// 静的ファイルのルートを設定
	router.Static("/static", "./view")

	// ルートの設定
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PATCH("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 環境変数からCORS許可オリジンを取得
		corsAllowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
		if corsAllowedOrigin == "" {
			corsAllowedOrigin = "*" // デフォルトはすべてのオリジンを許可
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", corsAllowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		// 他のミドルウェアやルーターハンドラの呼び出しを続行
		c.Next()
	}
}
