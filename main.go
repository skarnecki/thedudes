package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skarnecki/thedudes/dude"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var Version = "1-UNSET"

func main() {

	engine, err := xorm.NewEngine("mysql", "root:pw@/thedudes?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = engine.Sync2(new(dude.Dude))
	if err != nil {
		panic(err)
	}

	services := dude.Routes{Engine: engine}

	r := gin.Default()
	r.Use(gin.Logger())
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"skarnecki":    "bar",
		"ben": "ten",
		"juan": "jr",
	}))

	v1 := authorized.Group("/v1")
	{
		v1.GET("/health", services.HealthEndpoint)
		v1.GET("/dudes", services.GetDudesEndpoint)
		v1.GET("/dude/:id", services.GetOneDudeEndpoint)
		v1.PUT("/dude", services.AddNewDudeEndpoint)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
