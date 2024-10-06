package socket

import uuid "github.com/satori/go.uuid"

type Message struct {
	UserID    uuid.UUID
	PlayerID  uuid.UUID
	SessionID uuid.UUID
	CommandID uuid.UUID
	Service   ServiceIdentifier
	Operation OperationIdentifier
	Data      []byte
}

type ServiceIdentifier struct {
	ID int
}

type OperationIdentifier struct {
	ID int
}

type Response interface {
	MarshallBinary() ([]byte, error)
}
