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

type UpdatePermissionController struct {
	useCase *application.UpdatePermissionUseCase
}

func NewUpdatePermissionController(useCase *application.UpdatePermissionUseCase) *UpdatePermissionController {
	return &UpdatePermissionController{useCase: useCase}
}

func (ctrl *UpdatePermissionController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	tutoradoID, _ := strconv.Atoi(c.PostForm("id_tutorado_fk"))
	docenteID, _ := strconv.Atoi(c.PostForm("id_docente_fk"))
	date := c.PostForm("date")
	motivo := c.PostForm("motivo")
	estatus := c.PostForm("estatus")

	// Manejo opcional de imagen
	var imagePath string
	file, err := c.FormFile("evidencia")
	if err == nil {
		timestamp := time.Now().Unix()
		filename := filepath.Base(file.Filename)
		imagePath = filepath.Join("uploads", strconv.FormatInt(timestamp, 10)+"_"+filename)

		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la nueva imagen"})
			return
		}
	}

	// Si no hay nueva imagen, se puede dejar el campo vacío o actualizar con la anterior desde DB si se desea (requiere modificación del repositorio).

	permission := entities.NewPermission(
		int32(tutoradoID),
		int32(docenteID),
		imagePath,
		date,
		motivo,
		estatus,
	)
	permission.ID = int32(id)

	result, err := ctrl.useCase.Run(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
