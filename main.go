package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// レコードのアルバムに関するデータ
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// レコードのあるベムのデータの素となるスライス
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// JSON形式のすべてのアルバムリストを返す
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//受け取ったJSONをnewAlbumにバインドするためにBindJSONを呼び出す
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//スライスへ新しいアルバムを追加する
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
