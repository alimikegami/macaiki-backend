// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "macaiki/internal/user/dto"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	threaddto "macaiki/internal/thread/dto"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// BanComment provides a mock function with given fields: userRole, commentReportID
func (_m *UserUsecase) BanComment(userRole string, commentReportID uint) error {
	ret := _m.Called(userRole, commentReportID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(userRole, commentReportID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BanCommunity provides a mock function with given fields: userRole, communityReportID
func (_m *UserUsecase) BanCommunity(userRole string, communityReportID uint) error {
	ret := _m.Called(userRole, communityReportID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(userRole, communityReportID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BanThread provides a mock function with given fields: userRole, threadReportID
func (_m *UserUsecase) BanThread(userRole string, threadReportID uint) error {
	ret := _m.Called(userRole, threadReportID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(userRole, threadReportID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BanUser provides a mock function with given fields: userRole, userReportID
func (_m *UserUsecase) BanUser(userRole string, userReportID uint) error {
	ret := _m.Called(userRole, userReportID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(userRole, userReportID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeEmail provides a mock function with given fields: id, info
func (_m *UserUsecase) ChangeEmail(id uint, info dto.UserLoginRequest) (string, error) {
	ret := _m.Called(id, info)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, dto.UserLoginRequest) string); ok {
		r0 = rf(id, info)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, dto.UserLoginRequest) error); ok {
		r1 = rf(id, info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangePassword provides a mock function with given fields: id, passwordInfo
func (_m *UserUsecase) ChangePassword(id uint, passwordInfo dto.UserChangePasswordRequest) error {
	ret := _m.Called(id, passwordInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, dto.UserChangePasswordRequest) error); ok {
		r0 = rf(id, passwordInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id, curentUserID, curentUser
func (_m *UserUsecase) Delete(id uint, curentUserID uint, curentUser string) error {
	ret := _m.Called(id, curentUserID, curentUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(id, curentUserID, curentUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Follow provides a mock function with given fields: userID, userFollowerID
func (_m *UserUsecase) Follow(userID uint, userFollowerID uint) error {
	ret := _m.Called(userID, userFollowerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, userFollowerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id, tokenUserID
func (_m *UserUsecase) Get(id uint, tokenUserID uint) (dto.UserDetailResponse, error) {
	ret := _m.Called(id, tokenUserID)

	var r0 dto.UserDetailResponse
	if rf, ok := ret.Get(0).(func(uint, uint) dto.UserDetailResponse); ok {
		r0 = rf(id, tokenUserID)
	} else {
		r0 = ret.Get(0).(dto.UserDetailResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(id, tokenUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: userID, search
func (_m *UserUsecase) GetAll(userID uint, search string) ([]dto.UserResponse, error) {
	ret := _m.Called(userID, search)

	var r0 []dto.UserResponse
	if rf, ok := ret.Get(0).(func(uint, string) []dto.UserResponse); ok {
		r0 = rf(userID, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardAnalytics provides a mock function with given fields: userRole
func (_m *UserUsecase) GetDashboardAnalytics(userRole string) (dto.AdminDashboardAnalytics, error) {
	ret := _m.Called(userRole)

	var r0 dto.AdminDashboardAnalytics
	if rf, ok := ret.Get(0).(func(string) dto.AdminDashboardAnalytics); ok {
		r0 = rf(userRole)
	} else {
		r0 = ret.Get(0).(dto.AdminDashboardAnalytics)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportedComment provides a mock function with given fields: userRole, commentReportID
func (_m *UserUsecase) GetReportedComment(userRole string, commentReportID uint) (dto.ReportedCommentResponse, error) {
	ret := _m.Called(userRole, commentReportID)

	var r0 dto.ReportedCommentResponse
	if rf, ok := ret.Get(0).(func(string, uint) dto.ReportedCommentResponse); ok {
		r0 = rf(userRole, commentReportID)
	} else {
		r0 = ret.Get(0).(dto.ReportedCommentResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(userRole, commentReportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportedCommunity provides a mock function with given fields: userRole, communityReportID
func (_m *UserUsecase) GetReportedCommunity(userRole string, communityReportID uint) (dto.ReportedCommunityResponse, error) {
	ret := _m.Called(userRole, communityReportID)

	var r0 dto.ReportedCommunityResponse
	if rf, ok := ret.Get(0).(func(string, uint) dto.ReportedCommunityResponse); ok {
		r0 = rf(userRole, communityReportID)
	} else {
		r0 = ret.Get(0).(dto.ReportedCommunityResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(userRole, communityReportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportedThread provides a mock function with given fields: userRole, threadReportID
func (_m *UserUsecase) GetReportedThread(userRole string, threadReportID uint) (dto.ReportedThreadResponse, error) {
	ret := _m.Called(userRole, threadReportID)

	var r0 dto.ReportedThreadResponse
	if rf, ok := ret.Get(0).(func(string, uint) dto.ReportedThreadResponse); ok {
		r0 = rf(userRole, threadReportID)
	} else {
		r0 = ret.Get(0).(dto.ReportedThreadResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(userRole, threadReportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportedUser provides a mock function with given fields: userRole, userReportID
func (_m *UserUsecase) GetReportedUser(userRole string, userReportID uint) (dto.ReportedUserResponse, error) {
	ret := _m.Called(userRole, userReportID)

	var r0 dto.ReportedUserResponse
	if rf, ok := ret.Get(0).(func(string, uint) dto.ReportedUserResponse); ok {
		r0 = rf(userRole, userReportID)
	} else {
		r0 = ret.Get(0).(dto.ReportedUserResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint) error); ok {
		r1 = rf(userRole, userReportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReports provides a mock function with given fields: curentUserRole
func (_m *UserUsecase) GetReports(curentUserRole string) ([]dto.BriefReportResponse, error) {
	ret := _m.Called(curentUserRole)

	var r0 []dto.BriefReportResponse
	if rf, ok := ret.Get(0).(func(string) []dto.BriefReportResponse); ok {
		r0 = rf(curentUserRole)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.BriefReportResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(curentUserRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadByToken provides a mock function with given fields: userID, tokenUserID
func (_m *UserUsecase) GetThreadByToken(userID uint, tokenUserID uint) ([]threaddto.DetailedThreadResponse, error) {
	ret := _m.Called(userID, tokenUserID)

	var r0 []threaddto.DetailedThreadResponse
	if rf, ok := ret.Get(0).(func(uint, uint) []threaddto.DetailedThreadResponse); ok {
		r0 = rf(userID, tokenUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]threaddto.DetailedThreadResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, tokenUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFollowers provides a mock function with given fields: tokenUserID, getFollowingUserID
func (_m *UserUsecase) GetUserFollowers(tokenUserID uint, getFollowingUserID uint) ([]dto.UserResponse, error) {
	ret := _m.Called(tokenUserID, getFollowingUserID)

	var r0 []dto.UserResponse
	if rf, ok := ret.Get(0).(func(uint, uint) []dto.UserResponse); ok {
		r0 = rf(tokenUserID, getFollowingUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(tokenUserID, getFollowingUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFollowing provides a mock function with given fields: tokenUserID, getFollowingUserID
func (_m *UserUsecase) GetUserFollowing(tokenUserID uint, getFollowingUserID uint) ([]dto.UserResponse, error) {
	ret := _m.Called(tokenUserID, getFollowingUserID)

	var r0 []dto.UserResponse
	if rf, ok := ret.Get(0).(func(uint, uint) []dto.UserResponse); ok {
		r0 = rf(tokenUserID, getFollowingUserID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(tokenUserID, getFollowingUserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: loginInfo
func (_m *UserUsecase) Login(loginInfo dto.UserLoginRequest) (dto.LoginResponse, error) {
	ret := _m.Called(loginInfo)

	var r0 dto.LoginResponse
	if rf, ok := ret.Get(0).(func(dto.UserLoginRequest) dto.LoginResponse); ok {
		r0 = rf(loginInfo)
	} else {
		r0 = ret.Get(0).(dto.LoginResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.UserLoginRequest) error); ok {
		r1 = rf(loginInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *UserUsecase) Register(_a0 dto.UserRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.UserRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Report provides a mock function with given fields: userID, userReportedID, ReportCategoryID
func (_m *UserUsecase) Report(userID uint, userReportedID uint, ReportCategoryID uint) error {
	ret := _m.Called(userID, userReportedID, ReportCategoryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) error); ok {
		r0 = rf(userID, userReportedID, ReportCategoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendOTP provides a mock function with given fields: email
func (_m *UserUsecase) SendOTP(email dto.SendOTPRequest) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.SendOTPRequest) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBackgroundImage provides a mock function with given fields: id, img
func (_m *UserUsecase) SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error) {
	ret := _m.Called(id, img)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, *multipart.FileHeader) string); ok {
		r0 = rf(id, img)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *multipart.FileHeader) error); ok {
		r1 = rf(id, img)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetProfileImage provides a mock function with given fields: id, img
func (_m *UserUsecase) SetProfileImage(id uint, img *multipart.FileHeader) (string, error) {
	ret := _m.Called(id, img)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, *multipart.FileHeader) string); ok {
		r0 = rf(id, img)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, *multipart.FileHeader) error); ok {
		r1 = rf(id, img)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unfollow provides a mock function with given fields: userID, userFollowerID
func (_m *UserUsecase) Unfollow(userID uint, userFollowerID uint) error {
	ret := _m.Called(userID, userFollowerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, userFollowerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: userUpdate, id
func (_m *UserUsecase) Update(userUpdate dto.UserUpdateRequest, id uint) (dto.UserUpdateResponse, error) {
	ret := _m.Called(userUpdate, id)

	var r0 dto.UserUpdateResponse
	if rf, ok := ret.Get(0).(func(dto.UserUpdateRequest, uint) dto.UserUpdateResponse); ok {
		r0 = rf(userUpdate, id)
	} else {
		r0 = ret.Get(0).(dto.UserUpdateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.UserUpdateRequest, uint) error); ok {
		r1 = rf(userUpdate, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyOTP provides a mock function with given fields: email, OTPCode
func (_m *UserUsecase) VerifyOTP(email string, OTPCode string) error {
	ret := _m.Called(email, OTPCode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, OTPCode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
