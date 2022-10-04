package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type movie struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Producer string `json:"producer"`
}

var movies = []movie{
	{Name: "WAR" + " " + "from invozone", Duration: 250, Producer: "William"},
	{Name: "Everest" + " " + "from invozone", Duration: 425, Producer: "Nikolas"},
	{Name: "Evil" + " " + "from invozone", Duration: 344, Producer: "Jhon"},
}

func getMovie(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func createMovie(c *gin.Context) {
	var newmovie movie
	if err := c.BindJSON(&newmovie); err != nil { //binding json element to new movie and we are checking about error
		return
	}
	newmovie.Name = newmovie.Name + " " + "From InvoZone" //cancatinating the string to variable Name
	movies = append(movies, newmovie)                     //append the new data from newmovie to movies slice
	c.IndentedJSON(http.StatusCreated, newmovie)          //indented JSON formated elemnts to newMovies and giving response 200

}
func main() {
	router := gin.Default()
	router.GET("/movies", getMovie)
	router.POST("/movies", createMovie)
	router.Run("localhost:8080")
}
