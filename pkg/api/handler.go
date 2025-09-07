package api

import (
	"github.com/Archii0/lumo/pkg/store"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", pong)

	r.GET("/key/:key", getValue)
	r.PUT("/key/:key", putValue)
	r.DELETE("/key/:key", deleteValue)

	return r

}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getValue(c *gin.Context) {
	key := c.Param("key")
	value, ok := store.Get(key)

	if !ok {
		c.JSON(404, gin.H{"error": "key not found"})
		return
	}
	c.JSON(200, gin.H{"key": key, "value": value})
}

func putValue(c *gin.Context) {
	key := c.Param("key")

	var json struct {
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
	}

	store.Set(key, json.Value)

	c.JSON(200, gin.H{"message": "ok"})
}

func deleteValue(c *gin.Context) {
	key := c.Param("key")
	store.Delete(key)
	c.JSON(200, gin.H{"message": "deleted"})
}
