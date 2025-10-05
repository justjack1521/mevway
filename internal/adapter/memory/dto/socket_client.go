package dto

import "encoding/json"

type SocketClientRedis struct {
	SessionID string `redis:"SessionID" json:"SessionID"`
	UserID    string `redis:"UserID" json:"UserID"`
	PlayerID  string `redis:"PlayerID" json:"PlayerID"`
	PatchID   string `redis:"PatchID" json:"PatchID"`
}

func (s SocketClientRedis) ToMapStringInterface() map[string]interface{} {
	return map[string]interface{}{
		"SessionID": s.SessionID,
		"UserID":    s.UserID,
		"PlayerID":  s.PlayerID,
		"PatchID":   s.PatchID,
	}
}

func (s SocketClientRedis) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
