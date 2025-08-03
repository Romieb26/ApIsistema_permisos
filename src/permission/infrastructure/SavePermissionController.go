package infrastructure

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type SavePermissionController struct {
	useCase *application.SavePermissionUseCase
}

func NewSavePermissionController(useCase *application.SavePermissionUseCase) *SavePermissionController {
	return &SavePermissionController{useCase: useCase}
}

func (ctrl *SavePermissionController) Run(c *gin.Context) {
	tutoradoID, _ := strconv.Atoi(c.PostForm("id_tutorado_fk"))
	docenteID, _ := strconv.Atoi(c.PostForm("id_docente_fk"))
	date := c.PostForm("date")
	motivo := c.PostForm("motivo")
	estatus := c.PostForm("estatus")

	// Manejo del archivo
	file, err := c.FormFile("evidencia")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo de evidencia requerido"})
		return
	}

	// Crear nombre único para la imagen
	timestamp := time.Now().Unix()
	filename := filepath.Base(file.Filename)
	imagePath := filepath.Join("uploads", strconv.FormatInt(timestamp, 10)+"_"+filename)

	// Guardar archivo
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
		return
	}

	// Crear entidad
	permission := entities.NewPermission(
		int32(tutoradoID),
		int32(docenteID),
		imagePath, // solo ruta
		date,
		motivo,
		estatus,
	)

	result, err := ctrl.useCase.Run(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
