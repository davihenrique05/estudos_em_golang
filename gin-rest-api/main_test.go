package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"gin-rest-api/controllers"
	"gin-rest-api/database"
	"gin-rest-api/models"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var CPF string

func SetupRoutesConfig() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func SetupDatabase() *sql.DB {
	database.ConnectWithDatabase()
	db, _ := database.DB.DB()
	return db
}

func CreateMockStudent() {
	ranNumber := fmt.Sprint(10000000000 + rand.Intn(90000000))

	student := models.Student{
		Name: "Student",
		CPF:  ranNumber,
		RG:   "123456789"}

	database.DB.Create(&student)
	ID = int(student.ID)
	CPF = student.CPF
}

func RemoveMockStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetStudentsById(t *testing.T) {
	db := SetupDatabase()
	defer db.Close()

	CreateMockStudent()
	defer RemoveMockStudent()

	var responseStudent models.Student

	r := SetupRoutesConfig()
	r.GET("/api/students/:id", controllers.GetStudentById)
	reqUrl := fmt.Sprintf("/api/students/%d", ID)
	req, _ := http.NewRequest("GET", reqUrl, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	json.Unmarshal(response.Body.Bytes(), &responseStudent)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, ID, int(responseStudent.ID))
}

func TestGetStudentByCPF(t *testing.T) {
	db := SetupDatabase()
	defer db.Close()

	CreateMockStudent()
	defer RemoveMockStudent()

	var responseStudent models.Student

	r := SetupRoutesConfig()
	r.GET("/api/students/cpf/:cpf", controllers.FindStudentByCPF)
	reqUrl := fmt.Sprintf("/api/students/cpf/%s", CPF)
	req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	json.Unmarshal(response.Body.Bytes(), &responseStudent)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, ID, int(responseStudent.ID))
	assert.Equal(t, CPF, responseStudent.CPF)
}

func TestUpdateStudent(t *testing.T) {
	db := SetupDatabase()
	defer db.Close()

	CreateMockStudent()
	defer RemoveMockStudent()

	r := SetupRoutesConfig()
	r.PATCH("/api/students/:id", controllers.UpdateStudent)
	reqUrl := fmt.Sprintf("/api/students/%d", ID)

	toUpdateStudent := models.Student{Name: "Student Name", CPF: "12345678901", RG: "123456700"}
	updatedStudentJson, _ := json.Marshal(toUpdateStudent)

	req, _ := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(updatedStudentJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var responseStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &responseStudent)

	assert.Equal(t, toUpdateStudent.CPF, responseStudent.CPF)
	assert.Equal(t, toUpdateStudent.RG, responseStudent.RG)
}

func TestDeleteStudent(t *testing.T) {
	db := SetupDatabase()
	defer db.Close()

	CreateMockStudent()
	r := SetupRoutesConfig()
	r.DELETE("/api/students/:id", controllers.DeleteStudent)

	reqUrl := fmt.Sprintf("/api/students/%d", ID)
	req, _ := http.NewRequest("DELETE", reqUrl, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
