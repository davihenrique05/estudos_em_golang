package routes

import (
	"gin-rest-api/controllers"

	docs "gin-rest-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Version = "1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
