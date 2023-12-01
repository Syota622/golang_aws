package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 以下はhandler.goに含まれる関数の例です。

// getAlbums はアルバムの一覧を取得します。
func getAlbums(c *gin.Context) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums は新しいアルバムを追加します。
func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newAlbum.ID = id
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID は特定のIDを持つアルバムを取得します。
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album Album

	if err := db.QueryRow("SELECT * FROM album WHERE id = ?", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// updateAlbumByID は特定のIDを持つアルバムの情報を更新します。
func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updateAlbum Album

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bind error"})
		return
	}

	if _, err := db.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", updateAlbum.Title, updateAlbum.Artist, updateAlbum.Price, id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "update error"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

// deleteAlbumByID は特定のIDを持つアルバムを削除します。
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	if _, err := db.Exec("DELETE FROM album WHERE id = ?", id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
