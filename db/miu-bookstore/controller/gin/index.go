package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"head": "index_head",
	})
}
