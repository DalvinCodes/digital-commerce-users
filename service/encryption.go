package service

import (
	"github.com/DalvinCodes/digital-commerce/users/pkg"
	"go.uber.org/zap"
)

type EncryptionI interface {
	Encrypt(password string) ([]byte, error)
	IsPassword(password, pwHash []byte) (bool, error)
}

type EncryptionService struct {
	EncryptionI EncryptionI
}

var encryptionCost = 14

var logger = zap.L()

type Encryption struct {
	GeneratePwd       func([]byte, int) ([]byte, error)
	CompareHashAndPwd func([]byte, []byte) error
}

func NewEncryptionService(encryptSrv EncryptionI) *EncryptionService {
	es := &EncryptionService{
		EncryptionI: encryptSrv,
	}
	return es
}

func (e *Encryption) Encrypt(password string) ([]byte, error) {
	if password == "" {
		logger.Error("error: ", zap.Error(pkg.ErrInternalEmptyPassword))
		return nil, pkg.ErrInvalidInput
	}

	pSlice := []byte(password)

	pBytes, err := e.GeneratePwd(pSlice, encryptionCost)
	if err != nil {
		return nil, err
	}

	return pBytes, nil
}

func (e *Encryption) IsPassword(pwHash, password []byte) (bool, error) {
	if err := e.CompareHashAndPwd(pwHash, password); err != nil {
		logger.Error("error", zap.Error(err))
		return false, pkg.ErrInvalidInput
	}
	return true, nil
}
