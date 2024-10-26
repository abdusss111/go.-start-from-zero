package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type user struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Age int  `json:"age"`
    Course  int `json:"course"`
}

type course struct {
    Title string `json:"title"`
    Grade float64 `json:"grade"`
    Description string `json:"description"`
    Prerequisites string `json: "prerequisites"`
}

// users 
var users = []user{
    {ID: "1", Name: "Abdussalam", Age: 19, Course: 3},
    {ID: "2", Name: "Abdussalyam", Age: 19, Course: 2},
    {ID: "3", Name: "Abdusalam", Age: 19, Course: 1},
}

var courses = []course{
    {Title: "OOP", Grade: 4.0, Description: "Object Oriented Programming using Java", Prerequisites: "PP1, PP2" },
    {Title: "Angular", Grade: 4.0, Description: "Frontend Framework Angular TS", Prerequisites: "PP1, PP2, WD" },
    {Title: "DSA", Grade: 4.0, Description: "Data Storage and Analysis", Prerequisites: "no prerequisites" },

}

// getusers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

func getCourses(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, courses)
}

func main() {
    router := gin.Default()
    router.GET("courses", getCourses)

    router.GET("users", getUsers)

    router.Run("localhost:8080")
}
