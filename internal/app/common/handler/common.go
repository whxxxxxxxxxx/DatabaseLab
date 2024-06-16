package handler

import (
	"DatabaseLab/middleware/response"
	"DatabaseLab/pkg/errorx"
	"DatabaseLab/pkg/imagex"
	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, _, _ := c.Request.FormFile("file")
	name, err := imagex.UploadImg(file)
	var res response.JsonResponse
	if err != nil {
		code := errorx.ErrorImgUpload
		res = response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	res = response.ReturnResponse(errorx.Success, errorx.GetMsg(errorx.Success), name, nil)
	response.HttpResponse(c, res)
}
