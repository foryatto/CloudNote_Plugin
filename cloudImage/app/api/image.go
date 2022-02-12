package api

import (
	"cloudImage/app/model"
	"cloudImage/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Image = imageApi{}

type imageApi struct{}

// Upload upload image files
func (*imageApi) Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("file")
	result, err := service.Image.Upload(files)
	if err != nil {
		r.Response.WriteJsonExit(model.StandResp{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	r.Response.WriteJsonExit(model.StandResp{
		Code: 0,
		Msg:  "success",
		Data: result,
	})
}
