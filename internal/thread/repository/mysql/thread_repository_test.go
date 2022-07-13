package mysql

// func TestSuccessfullCreateThreadReport(t *testing.T) {
// 	mockedDB, mockObj, err := sqlmock.New()
// 	db, err := gorm.Open(mysql.Dialector{
// 		&mysql.Config{
// 			Conn:                      mockedDB,
// 			SkipInitializeWithVersion: true,
// 		},
// 	}, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	mockThreadRepo := CreateNewThreadRepository(db)
// 	mockObj.ExpectBegin()

// 	mockObj.ExpectCommit()

// 	defer mockedDB.Close()
// }
