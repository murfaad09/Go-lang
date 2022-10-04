package main

import (
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "52", Title: "Frozen", Artist: "Maclern", Price: 52.22},
	{ID: "55", Title: "Dozen", Artist: "MaclernF1", Price: 58.22},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
