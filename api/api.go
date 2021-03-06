/*------------------------
name        api
describe    api library
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package api

import (
	//golang official package
	"fmt"
	"time"

	//third party package
	"github.com/gin-gonic/gin"

	//project package
	"github.com/ailncode/gorgw/api/bucket"
	"github.com/ailncode/gorgw/api/object"
	"github.com/ailncode/gorgw/base"
	. "github.com/ailncode/gorgw/config"
)

//type of Action
type Action func(c *gin.Context)

//api struct
type Api struct {
	Listen string
}

//method of Api to start up api
func (a *Api) Run() {
	fmt.Println(time.Now().String(), "Run Api with Config:")
	fmt.Println("------------------config------------------")
	for k, v := range Conf {
		fmt.Println(k, "\t\t\t\t\t:\t\t\t", v)
	}
	fmt.Println("------------------------------------------")
	if Conf["debug"] != "true" {
		//switch gin mode to release
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	//use logger middle ware
	router.Use(base.Logger())
	authorized := router.Group("")
	authorized.Use(base.Authorizer())
	{
		//bucket
		authorized.POST("/", bucket.Post)
		bucketGroup := router.Group("")
		bucketGroup.Use(base.Authorizer())
		{
			bucketGroup.PUT("/:bucketname", bucket.Put)
			bucketGroup.GET("/", bucket.Get)
			bucketGroup.GET("/:bucketname", bucket.List)
		}
		objectGroup := router.Group("")
		objectGroup.Use(base.Authorizer(), base.CheckBucket())
		{
			objectGroup.POST("/:bucketname", object.Post)
			objectGroup.PUT("/:bucketname/:objectkey", object.Put)
			objectGroup.GET("/:bucketname/:objectkey", object.Get)
		}
	}
	router.Run(a.Listen)
}
