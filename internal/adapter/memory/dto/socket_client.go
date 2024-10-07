package dto

import "encoding/json"

type SocketClientRedis struct {
	SessionID string `redis:"session_id" json:"session_id"`
	UserID    string `redis:"user_id" json:"user_id"`
	PlayerID  string `redis:"player_id" json:"player_id"`
}

func (s SocketClientRedis) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
