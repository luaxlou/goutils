package gindefault

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Run(addr string, initial func(engine *gin.Engine)) {

	time.LoadLocation("Asia/Shanghai")

	engine := gin.New()

	engine.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "X-Auth-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	initial(engine)

	log.Println("Starting http server at address:", addr)

	engine.Run(addr)
}
