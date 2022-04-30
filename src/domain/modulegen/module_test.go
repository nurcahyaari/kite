package modulegen_test

import (
	"testing"

	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/stretchr/testify/assert"
)

func TestBuildModuleFile(t *testing.T) {

	testCases := []struct {
		name string
		exp  func() string
		act  func() (string, error)
	}{
		{
			name: "Test1",
			exp: func() string {
				return `package service

type Service interface {
}
type ServiceImpl struct {
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}
`
			},
			act: func() (string, error) {
				mg := modulegen.NewModuleGen()
				return mg.BuildModuleTemplate(modulegen.ModuleDto{
					ModuleName:  "Service",
					PackageName: "service",
					Path:        "",
					GomodName:   "test1",
				})

			},
		},
		{
			name: "Test2",
			exp: func() string {
				return `package repository

type Repository interface {
}
type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}
`
			},
			act: func() (string, error) {
				mg := modulegen.NewModuleGen()
				return mg.BuildModuleTemplate(modulegen.ModuleDto{
					ModuleName:  "Repository",
					PackageName: "repository",
					Path:        "",
					GomodName:   "test2",
				})

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			exp := tc.exp()
			act, err := tc.act()

			assert.NoError(t, err)
			assert.Equal(t, exp, act)
		})
	}
}
