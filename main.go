package main

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/db"
	"my-postcrossing/m"
	"my-postcrossing/router"
)

func main() {
	db.InitDatabase()
	m.Migrate(db.DBConn)

	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8899")
}
