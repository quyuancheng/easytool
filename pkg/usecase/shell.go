package usecase

import (
	"gorm.io/gorm"
	v1 "tool/pkg/controller/v1"
	"tool/pkg/models"
)

type ShellUsecase interface {
	GetShell() (*v1.Shell, error)
}

func NewShellUcase() ShellUsecase {
	return &ShellUcase{}
}

type ShellUcase struct {
	DB       *gorm.DB              `inject:""`
	ShellRepo models.ShellRepository `inject:""`
}

func (m *ShellUcase) GetShell() (*v1.Shell, error) {
	var shell *v1.Shell
	return shell, nil
}
