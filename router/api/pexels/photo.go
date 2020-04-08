package pexels

import (
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gwt/pkg"
)

type searchPhotosRequest struct {
	Query string `json:"query" form:"query" binding:"required"`
	PerPage int `json:"per_page" form:"per_page" binding:"gte=1"`
	Page int `json:"page" form:"page" binding:"gte=1"`
}

// SearchPhotos 搜索照片
func SearchPhotos(c *gin.Context) {
	g := pkg.RC{C:c}
	var req searchPhotosRequest
	err := c.ShouldBind(&req)
	if err != nil {
		g.InvalidParam(err)
		return
	}

	r, err := pc.SearchPhotos(req.Query, req.PerPage, req.Page)
	if err != nil {
		g.Response(200, pkg.Error, nil)
		return
	}

	g.Response(200, pkg.Success, r)
	return
}