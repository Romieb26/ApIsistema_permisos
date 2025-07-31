package application

import repositories "github.com/Romieb26/ApIsistema_permisos/src/rol/domain"

type DeleteRolUseCase struct {
	repo repositories.IRol
}

func NewDeleteRolUseCase(repo repositories.IRol) *DeleteRolUseCase {
	return &DeleteRolUseCase{repo: repo}
}

func (uc *DeleteRolUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
