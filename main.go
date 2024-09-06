package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/find-pairs", func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"err": "invalid request"})
			return
		}

		if errs := req.validate(); len(errs) > 0 {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"err": errs})
			return
		}

		ans := findPairs(req.Numbers, req.Target)
		c.JSON(200, map[string]interface{}{
			"solutins": ans,
		})
	})

	if err := engine.Run(":8081"); err != nil {
		log.Fatalf("failed to start server err: %s", err.Error())
	}
}
