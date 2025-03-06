package copy_helper

import (
	"log"

	"github.com/jinzhu/copier"
)

type (
	ModelConverter interface {
		Copy(to interface{}, from interface{})
	}
	modelConverter struct{}
)

func NewModelConverter() ModelConverter {
	return &modelConverter{}
}

func (m *modelConverter) Copy(to interface{}, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		log.Fatalln("failed to copy from model: ", err)
	}
}
