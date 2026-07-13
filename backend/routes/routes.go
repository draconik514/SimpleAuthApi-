package routes

import(
	"backend/controllers"
	"github.com/gin-gonic/gin"

)

func SetupRoutes(r *gin.Engine){
	api := r.Group("/api")

	{
		api.POST("/register", controllers.Register)
	}
}