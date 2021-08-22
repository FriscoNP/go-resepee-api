package usecase

import (
	"context"
	"go-resepee-api/db/repository"
	"go-resepee-api/entity"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FileUC struct {
	Context context.Context
	DB      *gorm.DB
}

type FileUCInterface interface {
	Store(fileType string, file *multipart.FileHeader) (res entity.File, err error)
}

func NewFileUC(ctx context.Context, db *gorm.DB) FileUCInterface {
	return &FileUC{
		Context: ctx,
		DB:      db,
	}
}

func (uc *FileUC) Store(fileType string, file *multipart.FileHeader) (res entity.File, err error) {
	fileRepo := repository.NewFileRepository(uc.DB)

	// source
	src, err := file.Open()
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}
	defer src.Close()

	folderPath := "public/" + fileType
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	// destination
	dest := folderPath + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
	dst, err := os.Create(dest)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}
	defer dst.Close()

	// copy file to folder
	if _, err = io.Copy(dst, src); err != nil {
		log.Warn(err.Error())
		return res, err
	}

	fileEntity := entity.File{
		Type: fileType,
		Path: dest,
	}
	res, err = fileRepo.Store(&fileEntity)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return res, err
}
