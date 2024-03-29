package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/Apriil15/blog-server/global"
	"github.com/Apriil15/blog-server/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// Upload file
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSaveFile := upload.GetUploadSavePath()
	destination := uploadSaveFile + "/" + fileName

	if !upload.CheckContainExtension(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(destination) {
		err := upload.CreateSavePath(uploadSaveFile, os.ModePerm)
		if err != nil {
			return nil, errors.New("fail to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("over maximum file limit")
	}
	if upload.CheckPermission(uploadSaveFile) {
		return nil, errors.New("insufficient file permission")
	}

	if err := upload.SaveFile(fileHeader, destination); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
