package pexels

import (
	"github.com/gin-gonic/gin"
	. "gwt/pkg"
)

type SearchPhotosRequest struct {
	Query string `json:"query" form:"query" binding:"required"`
	PerPage int `json:"per_page" form:"per_page" binding:"gte=1"`
	Page int `json:"page" form:"page" binding:"gte=1"`
}

func SearchPhotos(c *gin.Context) {
	g := Gin{C:c}
	var req SearchPhotosRequest
	err := c.ShouldBind(&req)
	if err != nil {
		g.InvalidParam(err)
		return
	}

	r, err := pc.SearchPhotos(req.Query, req.PerPage, req.Page)
	if err != nil {
		g.Response(200, Error, nil)
		return
	}

	g.Response(200, Success, r)
	return
}