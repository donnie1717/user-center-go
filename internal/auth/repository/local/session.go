package local

type Session struct {
	OnlineUser map[string]string
}

// NewSession init method
func NewSession() *Session {
	onlineUser := make(map[string]string, 32)
	return &Session{OnlineUser: onlineUser}
}
