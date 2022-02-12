package service

import (
	"errors"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guid"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var Image = imageService{}

type imageService struct {}

var ValidFileType = []string{".jpg", ".png"}

// Upload upload image files
func (*imageService) Upload(files ghttp.UploadFiles) ([]string, error) {

	fileCount := len(files)

	if fileCount == 0 {
		return nil, errors.New("no files")
	}

	newFileList := make([]string, 0, fileCount)

	for i:=0; i<fileCount; i++ {

		fileType := path.Ext(files[i].Filename)

		valid := false

		// 校验文件类型的合法性
		for i:=0; i<len(ValidFileType); i++ {
			if strings.ToLower(ValidFileType[i]) == strings.ToLower(fileType) {
				valid = true
				break
			}
		}
		if !valid {
			return nil, errors.New("invalid file")
		}

		// 文件重命名, 存储
		rootPath, _ := os.Getwd()
		newFileName := guid.S() + fileType
		files[i].Filename = newFileName
		saveName, err := files[i].Save(filepath.Join(rootPath, "public", "static"))
		if err != nil {
			return nil, err
		}
		newFileList = append(newFileList, saveName)
	}

	return newFileList, nil
}
