package model

import (
	"Gsky/utils"
	"Gsky/utils/errmsg"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func UploadFile(file multipart.File, fileName string) (string, int) {
	b, _ := ioutil.ReadAll(file)

	//把文件保存到指定位置
	folderPath, folderName := CreateDateDir(utils.SavePath)
	// 重新生成文件名
	newFileName := generateFileName()
	// 获取文件后缀名
	fileInfo := strings.Split(fileName, ".")
	fileExt := fileInfo[1]
	fullFileName := fmt.Sprintf("%s.%s", newFileName, fileExt)
	fullPath := filepath.Join(folderPath, fullFileName)
	err := ioutil.WriteFile(fullPath, b, 0777)
	if err != nil {
		return "",errmsg.ERROR_FILE_NOT_SAVE
	}
	//输出上传时文件名
	return filepath.Join(folderName, fullFileName), errmsg.SUCCESS
}

// 生成md5文件名
func generateFileName() string {
	md5Object := md5.New()
	md5Object.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hex.EncodeToString(md5Object.Sum(nil))[:16]
}

func CreateDateDir(Path string) (string, string) {
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath, folderName
}
