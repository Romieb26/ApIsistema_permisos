package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/application"
)

type DeleteRolController struct {
	deleteUseCase *application.DeleteRolUseCase
}

func NewDeleteRolController(deleteUseCase *application.DeleteRolUseCase) *DeleteRolController {
	return &DeleteRolController{deleteUseCase: deleteUseCase}
}

func (ctrl *DeleteRolController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	if err := ctrl.deleteUseCase.Run(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo eliminar el rol", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Rol eliminado exitosamente"})
}
