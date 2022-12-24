package repository

import (
	"WeddingUtilities/model"
)

type AddressRepository interface {
	All() ([]model.Address, error)

	FindByID(Id string) (*model.Address, error)
	FindByCodeName(CodeName string) (*model.Address, error)
	FindByName(name string) (*model.Address, error)
	FindByDivisionType(Divisiontype string) (*model.Address, error)
	FindByPhoneCode(phoneCode string) (*model.Address, error)
	FindByLevel(level int) (*model.Address, error)
	FindByParentId(ParentId string) (*model.Address, error)

	Save(address model.Address) error
}
