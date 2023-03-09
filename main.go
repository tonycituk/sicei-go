package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Student struct {
    Nombre string  `json:"nombre"`
    Matricula  string  `json:"matricula"`
}

// students slice
var students = []Student{
    {Nombre: "Manuel Antonio Cituk Martínez", Matricula: "20216392"},
    {Nombre: "Julian Chan Palomo", Matricula: "20216390"},
    {Nombre: "Jorge Teodoro Dawn Rodríguez", Matricula: "17000972"},

}
func getStudents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, students)
}

func main() {
  gin.SetMode(gin.ReleaseMode)
  r := gin.Default()

  
  r.GET("/students", getStudents)
  http.ListenAndServe(":80", r)
}
