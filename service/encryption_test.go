package service

import (
	"github.com/DalvinCodes/digital-commerce/users/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEncryptionService(t *testing.T) {
	assertions := assert.New(t)

	// Given
	want := &EncryptionService{&Encryption{}}

	// When
	got := NewEncryptionService(&Encryption{})

	// Then
	assertions.Equal(want, got)
}

func TestEncryption_Encrypt(t *testing.T) {
	assertions := assert.New(t)

	// Given
	es := NewEncryptionService(&Encryption{
		GeneratePwd:       func(bytes []byte, i int) ([]byte, error) { return []byte("test"), nil },
		CompareHashAndPwd: mockCompareHashAndPassword,
	})

	pwString := "test"
	pwSlice := []byte(pwString)

	// When
	hashedPW, err := es.EncryptionI.Encrypt(pwString)

	// Then
	assertions.NoError(err)
	assertions.Equal(pwSlice, hashedPW)
}

func TestEncryption_Encrypt_PasswordIsEmpty(t *testing.T) {
	assertions := assert.New(t)

	// Given
	es := NewEncryptionService(&Encryption{
		GeneratePwd:       mockGenerateFromPassword,
		CompareHashAndPwd: mockCompareHashAndPassword,
	})

	pwString := ""

	// When
	hashedPW, err := es.EncryptionI.Encrypt(pwString)

	// Then
	assertions.ErrorIs(err, pkg.ErrInvalidInput)
	assertions.Nil(hashedPW)
}

func TestEncryption_Encrypt_InvalidCostIndex(t *testing.T) {
	assertions := assert.New(t)

	// Given
	es := NewEncryptionService(&Encryption{
		GeneratePwd: func(bytes []byte, i int) ([]byte, error) {
			return nil, pkg.ErrInvalidInput
		},
		CompareHashAndPwd: mockCompareHashAndPassword,
	})

	pwString := "test"

	// When
	hashedPW, err := es.EncryptionI.Encrypt(pwString)

	// Then
	assertions.Nil(hashedPW)
	assertions.ErrorIs(err, pkg.ErrInvalidInput)
}

func TestEncryption_IsPassword(t *testing.T) {
	assertions := assert.New(t)

	// Given
	es := NewEncryptionService(&Encryption{
		GeneratePwd:       mockGenerateFromPassword,
		CompareHashAndPwd: mockCompareHashAndPassword,
	})

	pwHash, mockHash := []byte("test"), []byte("test")

	// When
	isPassword, err := es.EncryptionI.IsPassword(pwHash, mockHash)

	// Then
	assertions.Nil(err)
	assertions.True(isPassword)
}

func TestEncryption_IsPassword_ReturnFalse(t *testing.T) {
	assertions := assert.New(t)

	// Given
	es := NewEncryptionService(&Encryption{
		GeneratePwd: mockGenerateFromPassword,
		CompareHashAndPwd: func(pwHash, password []byte) error {
			return pkg.ErrInvalidInput
		}})

	pwHash, mockHash := []byte("te__"), []byte("test")

	// When
	isPassword, err := es.EncryptionI.IsPassword(pwHash, mockHash)

	// Then
	assertions.Error(err)
	assertions.False(isPassword)
}
func mockGenerateFromPassword(pHash []byte, cost int) ([]byte, error) {
	return pHash, nil
}
func mockCompareHashAndPassword(pwHash, password []byte) error {
	return nil
}
