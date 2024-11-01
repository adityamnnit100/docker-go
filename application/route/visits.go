package route

import (
	"docker-go/application/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GetServerVisit(c *gin.Context) {
	totalVisits, err := db.RedisClient.Get(c, "myapp").Int()
	if err != nil {
		if err == redis.Nil {
			totalVisits = 0
		} else {
			fmt.Println("Failed to get total visits to server")
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

	}

	totalVisits++
	err = db.RedisClient.Set(c, "myapp", totalVisits, 0).Err()
	if err != nil {
		fmt.Println("Failed to update total visits: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Total Visits": totalVisits})

}
