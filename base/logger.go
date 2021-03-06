/*------------------------
name        logger
describe    logger middle ware for gin
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package base

import (
	//golang official package
	"fmt"
	"net/http/httputil"
	"time"

	//third party package
	"github.com/gin-gonic/gin"

	//project package
	. "github.com/ailncode/gorgw/config"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(time.Now().String() + " " + c.Request.Method + " " + c.Request.URL.String() + " " + c.Request.RemoteAddr)
		if Conf["debug"] == "true" {
			dump, err := httputil.DumpRequest(c.Request, true)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(dump))
			}
		}
	}
}
