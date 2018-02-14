package handler

import (
	"anla.io/hound/response"
	"github.com/houndgo/suuid"
	"github.com/kataras/iris"
)

type (
	// UUID is
	UUID struct{}
)

// Create is
func (ud UUID) Create(ctx iris.Context) {
	uid := suuid.New()
	response.JSON(ctx, uid.String())
}
