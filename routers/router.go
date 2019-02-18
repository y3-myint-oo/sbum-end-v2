package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitConnection - Create Routers
func InitConnection() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode) //
	//router.Use(gin.Logger(//))
	//router.Use(gin.Recovery())
	router.Use(cors.Default())
	v1 := router.Group("/api/v1/")
	{
		//Network Check
		v1.GET("/network", networkStatus)

		//Serial Number Generator
		v1.GET("/serial", serialNumber)

		//Data CURD
		v1.POST("/data", createData)
		v1.GET("/data", fetchAllData)

		//Stock CURD
		v1.POST("/stock", createStock)
		v1.DELETE("/delete/stock", deleteUnit)
		v1.GET("/stock", fetchAllStock)
		v1.POST("/stock/totalprce", feacthTotalPrice)
		v1.POST("/stock/historyprice", featchHistoryPrice)

		//Unit CURD
		v1.POST("/unit", createUnit)
		v1.POST("/delete/unit", deleteUnit)
		v1.GET("/unit", fetchAllUnit)
		v1.POST("/unit/add/converter", addUnitConverter)
		v1.POST("/unit/delete/converter", deleteUnitConverter)

	}
	router.Run(":8081")

	//http.HandleFunc("/delete/unit", handler)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}
