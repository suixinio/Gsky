package v1

import (
	"Gsky/model"
	"Gsky/utils"
	"Gsky/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// todo 错误代码统一
// todo 先用起来再优化
func Upload(c *gin.Context) {

	file, fileHander, err := c.Request.FormFile("file")
	//
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_FILE_EMPTY,
			"message": errmsg.GetErrMsg(errmsg.ERROR_FILE_EMPTY),
		})
		return
	}
	// 文件大小
	if fileHander.Size > utils.FileSize || fileHander.Size < 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_FILE_SIZE,
			"message": errmsg.GetErrMsg(errmsg.ERROR_FILE_SIZE),
		})
		return
	}
	// 校验图片类型
	if judgeFileMimeTypes(fileHander.Header.Get("Content-Type")) == false {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_FILE_TYPE,
			"message": errmsg.GetErrMsg(errmsg.ERROR_FILE_TYPE),
		})
		return
	}
	fileName := fileHander.Filename

	filePath, code := model.UploadFile(file, fileName)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     strings.Replace(filePath, "\\", "/", -1),
	})
}

// 判断文件类型
func judgeFileMimeTypes(mimeTypes string) bool {
	var isAllowMimeTypes bool
	switch mimeTypes {
	case "image/jpeg":
		isAllowMimeTypes = true
	case "image/jpg":
		isAllowMimeTypes = true
	case "image/png":
		isAllowMimeTypes = true
	default:
		isAllowMimeTypes = false
	}
	return isAllowMimeTypes
}
