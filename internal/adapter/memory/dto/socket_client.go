package dto

type SocketClientRedis struct {
	SessionID string `redis:"session_id"`
	UserID    string `redis:"user_id"`
	PlayerID  string `redis:"player_id"`
}
