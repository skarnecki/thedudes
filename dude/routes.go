package dude

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-xorm/xorm"
	"strconv"
	"net/http"
)

type Routes struct {
	Engine *xorm.Engine
}

func (r * Routes) HealthEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"healthy": "true",
	})
}

func (r * Routes) GetDudesEndpoint(c *gin.Context)  {
	var everyone []Dude
	err := r.Engine.Find(&everyone)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, &everyone)
	return
}


func (r * Routes) GetOneDudeEndpoint(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("name"))
	if err != nil {
		panic(err)
	}

	var dude = Dude{Id: int64(id)}
	has, err := r.Engine.Get(&dude)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	if has {
		c.JSON(200, &dude)
		return
	}
	c.Status(404)
	return
}

func (r * Routes) AddNewDudeEndpoint(c *gin.Context)  {
	var json Dude
	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {
		has, err := r.Engine.Get(&json)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		if has {
			c.Status(409)
			return
		}
		_, err = r.Engine.Insert(&json)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(201, json)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
