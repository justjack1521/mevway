package dto

import "encoding/json"

type SocketClientRedis struct {
	SessionID string `redis:"SessionID" json:"SessionID"`
	UserID    string `redis:"UserID" json:"UserID"`
	PlayerID  string `redis:"PlayerID" json:"PlayerID"`
}

func (s SocketClientRedis) ToMapStringInterface() map[string]interface{} {
	return map[string]interface{}{
		"SessionID": s.SessionID,
		"UserID":    s.UserID,
		"PlayerID":  s.PlayerID,
	}
}

func (s SocketClientRedis) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
