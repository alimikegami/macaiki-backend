package mysql

import (
	"macaiki/internal/thread/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	threadReport = entity.ThreadReport{
		UserID:           uint(1),
		ThreadID:         uint(1),
		ReportCategoryID: uint(2),
	}
)

func TestSuccessfullCreateThreadReport(t *testing.T) {
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

	mockThreadRepo := CreateNewThreadRepository(db)
	mockObj.ExpectBegin()
	mockObj.ExpectExec("INSERT INTO `thread_reports` (`created_at`,`updated_at`,`deleted_at`,`user_id`,`thread_id`,`report_category_id`) VALUES (?, ?, NULL, ?, ?, ?)").WithArgs(time.Now(), time.Now(), 1, 1, 1)
	mockObj.ExpectCommit()

	err = mockThreadRepo.CreateThreadReport(threadReport)

	defer mockedDB.Close()

	assert.Error(t, err)
}
