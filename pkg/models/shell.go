package models

import (
	"gorm.io/gorm"
)

type Shell struct {
	Model
}

type ShellRepository interface {
	List() ([]*Shell, int64, error)
}

type ShellRepo struct {
	DB *gorm.DB `inject:""`
}

func NewShellRepo() ShellRepository {
	return &ShellRepo{}
}

func (m *ShellRepo) WithTransaction(db *gorm.DB) ShellRepository {
	return &ShellRepo{DB: db}
}

func (m *ShellRepo) List() ([]*Shell, int64, error) {
	return nil, 0, nil
}
