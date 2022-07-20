package mysql

import (
	"macaiki/pkg/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSuccessfullDeleteComment(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	threadRepo := CreateNewThreadRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(utils.AnyTime{}, uint(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = threadRepo.DeleteComment(uint(1))
	assert.NoError(t, err)
}

func TestNoRowsAffectedDeleteComment(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	threadRepo := CreateNewThreadRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(utils.AnyTime{}, uint(1)).WillReturnResult(sqlmock.NewResult(0, 0))
	mockObj.ExpectCommit()

	err = threadRepo.DeleteComment(uint(1))
	assert.Error(t, err)
}

func TestSuccessfullDeleteThread(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	threadRepo := CreateNewThreadRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(utils.AnyTime{}, uint(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	mockObj.ExpectCommit()

	err = threadRepo.DeleteThread(uint(1))
	assert.NoError(t, err)
}

func TestNoRowsAffectedDeleteThread(t *testing.T) {
	mockedDB, mockObj, err := sqlmock.New()
	db, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			Conn:                      mockedDB,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	threadRepo := CreateNewThreadRepository(db)

	defer mockedDB.Close()

	mockObj.ExpectBegin()
	mockObj.ExpectExec(regexp.QuoteMeta("UPDATE")).WithArgs(utils.AnyTime{}, uint(1)).WillReturnResult(sqlmock.NewResult(0, 0))
	mockObj.ExpectCommit()

	err = threadRepo.DeleteThread(uint(1))
	assert.Error(t, err)
}
