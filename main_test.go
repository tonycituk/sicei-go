package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
	gin.SetMode(gin.TestMode)
    router := gin.Default()
    return router
}

func TestGetStudentsRoute(t *testing.T) {
	r := SetUpRouter()
	r.GET("/students", getStudents)
    req, _ := http.NewRequest("GET", "/students", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var students_response []Student
    json.Unmarshal(w.Body.Bytes(), &students_response)
	//fmt.Println(students_response)
    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, students_response)
}