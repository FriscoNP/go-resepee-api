package controller

import (
	"go-resepee-api/app/controller/response"
	"go-resepee-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FileController struct {
	DB *gorm.DB
}

func NewFileController(db *gorm.DB) *FileController {
	return &FileController{
		DB: db,
	}
}

func (controller *FileController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	fileType := c.FormValue("file_type")
	file, err := c.FormFile("file")
	if err != nil {
		log.Warn(err.Error())
		return SendError(c, http.StatusBadRequest, err)
	}

	fileUC := usecase.NewFileUC(ctx, controller.DB)
	res, err := fileUC.Store(fileType, file)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	return SendSuccess(c, response.CreateFileResponse(&res), "file_uploaded")
}
