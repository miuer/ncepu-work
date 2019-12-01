package gin

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	initFrontEnd(router)

	initBackEnd(router)

	router.Run()

}

func initFrontEnd(engine *gin.Engine) {
	engine.LoadHTMLGlob("./miu-bookstore/view/web/**/*")
	engine.Static("assets", "./miu-bookstore/static")
}

func initBackEnd(engine *gin.Engine) {
	engine.GET("/home", home)
	engine.GET("/search", novelSearch)
}
