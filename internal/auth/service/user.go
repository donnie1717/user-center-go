package service

import (
	"github.com/google/uuid"
	"sync"
	"user-center/internal/auth/repository"
)

type UserService interface {
	PasswordLogin(username, password string) (bool, string)
	GetBySessionId(sessionId string) string
}

type UserServiceImpl struct {
	repository repository.SessionRepository
}

var userService UserService
var once sync.Once

// GetUserService 获取单例登录服务
func GetUserService() UserService {
	once.Do(func() {
		sessionRepository := repository.NewSimpleSessionRepository()
		userService = &UserServiceImpl{repository: sessionRepository}
	})
	return userService
}

func (s *UserServiceImpl) PasswordLogin(username, password string) (bool, string)  {
	// 后续校验需要放在数据库
	if username == "13055653558" && password == "123456" {
		session := s.repository.GetSession()

		sessionId := uuid.New().String()
		session.OnlineUser[sessionId] = username
		return true, sessionId
	}

	return false, ""
}

func (s *UserServiceImpl) GetBySessionId(sessionId string) string {
	username, exists := s.repository.GetSession().OnlineUser[sessionId]
	if exists {
		return username
	}
	return ""
}





