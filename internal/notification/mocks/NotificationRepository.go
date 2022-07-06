package mocks

import (
	entity "macaiki/internal/notification/entity"

	"github.com/stretchr/testify/mock"
)

type NotificationRepository struct {
	mock.Mock
}

func (_m *NotificationRepository) StoreNotification(notification entity.Notification) error
func (_m *NotificationRepository) GetAllNotifications(userID uint) ([]entity.Notification, error)
func (_m *NotificationRepository) ReadAllNotifications(userID uint) error
func (_m *NotificationRepository) DeleleteAllNotifications(userID uint) error

type mockConstructorTestingTNewNotificationRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotificationRepository creates a new instance of NotificationRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotificationRepository(t mockConstructorTestingTNewNotificationRepository) *NotificationRepository {
	mock := &NotificationRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
