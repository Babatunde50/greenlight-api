package session

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
	IsSessionExpired(maxlifetime int64) bool
}
