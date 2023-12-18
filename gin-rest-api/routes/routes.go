package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.ShowIndex)

	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.GET("/api/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.PATCH("/api/students/:id", controllers.UpdateStudent)
	r.POST("/api/students", controllers.CreateNewStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudent)

	r.NoRoute(controllers.NotFoundRoute)

	r.Run()
}
