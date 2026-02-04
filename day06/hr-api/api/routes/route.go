package routes

import (
	"net/http"

	handlers "github.com/HendraaaIrwn/hr-restapi/hr-api/internal/handler"
	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/repositories"
	"github.com/HendraaaIrwn/hr-restapi/hr-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	//5. Initialize repositories
	regionRepo := repositories.NewRegionRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)

	//6. Initialize services
	regionService := services.NewRegionService(regionRepo)
	departmentService := services.NewDepartmentService(departmentRepo)

	//7. Initialize handlers
	regionHandler := handlers.NewRegionHandler(regionService)
	departmentHandler := handlers.NewDepartmentHandler(departmentService)
	employeeHandler := handlers.NewEmployeeHandler(services.NewEmployeeService(repositories.NewEmployeeRepository(db)))

	//3.1 call handler
	router.GET("/", welcomeHandler)

	//9.call basepath
	basePath := viper.GetString("SERVER.BASE_PATH")

	//8. grouping subroute with prefix /api
	api := router.Group(basePath)
	{

		// region routes endpoints
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.GET("/countries", regionHandler.GetRegionsWithCountry)
			regions.GET("/:id/countries", regionHandler.GetRegionByIdWithCountry)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}

		department := api.Group("/departments")
		{
			department.POST("", departmentHandler.CreateDepartment)
			department.GET("/:id", departmentHandler.GetDepartmentByID)
			department.GET("", departmentHandler.GetAllDepartments)
			department.PUT("/:id", departmentHandler.UpdateDepartment)
			department.DELETE("/:id", departmentHandler.DeleteDepartment)
			department.GET("/search", departmentHandler.SearchDepartments) // filter search
		}

		employee := api.Group("/employees")
		{
			employee.POST("", employeeHandler.CreateEmployee)
			employee.GET("/:id", employeeHandler.GetEmployeeByID)
			employee.GET("", employeeHandler.GetAllEmployees)
			employee.PUT("/:id", employeeHandler.UpdateEmployee)
			employee.DELETE("/:id", employeeHandler.DeleteEmployee)
			employee.GET("/search", employeeHandler.SearchEmployees) // filter search
		}
	}

}

// 2. create first handler
func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to gin framework",
		"status":  "running",
	})
}
