package main

import (
	"gin-demo/initRouter"
)

func main() {
	// router := gin.Default()
	// router.Run()

	router := initRouter.SetupRouter()
	_ = router.Run(":9000")
}
