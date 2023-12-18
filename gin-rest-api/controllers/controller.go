package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"
	_ "net/http/httputil"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// GetAllStudents godoc
// @Sumary List all students
// @Description Access database to retrieve all students
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} []models.Student
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students [get]
func GetAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// GetStudentById godoc
// @Summary Find a student by ID and return the information information
// @Description Receive the student ID and return his information
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Used to find the student"
// @Success 200 {object} models.Student
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students/{id} [get]
func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Find(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// FindStudentByCPF godoc
// @Summary Find a student by CPF and return the information information
// @Description Receive the student CPF and return his information
// @Tags students
// @Accept json
// @Produce json
// @Param cpf path string true "Used to find the student"
// @Success 200 {object} models.Student
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students/cpf/{cpf} [get]
func FindStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student CPF not Found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

// CreateNewStudent godoc
// @Summary Create a new Student
// @Description Create and add a new student to the database
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student Information"
// @Success 200 {object} models.Student
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students [post]
func CreateNewStudent(c *gin.Context) {
	var newStudent models.Student

	err := c.ShouldBindJSON(&newStudent)

	if err := models.ValidateStudentData(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&newStudent)
	c.JSON(http.StatusOK, newStudent)
}

// UpdateStudent godoc
// @Summary Update student information
// @Description Receive and update the information of given student
// @Tags students
// @Accept json
// @Produce json
// @Param id query int true "Used to find the student"
// @Param student body models.Student true "Student Information"
// @Success 200 {object} models.Student
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students [patch]
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	err := c.ShouldBindJSON(&student)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

// DeleteStudent godoc
// @Summary Delete student information
// @Description Receive and ID and remove the student information from the database
// @Tags students
// @Accept json
// @Produce json
// @Param id query int true "Used to find the student"
// @Success 200 {object} string
// @Failure 400 {object} httputil.HTTPError
// @Router /api/students [delete]
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student removed successfully",
	})
}

func ShowIndex(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func NotFoundRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
