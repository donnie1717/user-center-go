package repository

import (
	"sync"
	"user-center/internal/auth/repository/local"
)

// SessionRepository 持久层接口
type SessionRepository interface {
	GetSession() *local.Session
}

// SimpleSessionRepository 持久层接口实现
type SimpleSessionRepository struct {
	session *local.Session
	once sync.Once
}

// GetSession 实现接口
func (r *SimpleSessionRepository) GetSession() *local.Session {
	r.once.Do(func() {
		r.session = local.NewSession()
	})
	return r.session
}

// NewSimpleSessionRepository 构造方法
func NewSimpleSessionRepository() SessionRepository {
	return &SimpleSessionRepository{}
}

