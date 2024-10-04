package system

import (
	"errors"
	"os"
)

var (
	errServerMaintenance = errors.New("service is undergoing maintenance")
)

type StatusService struct {
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Status() error {
	if os.Getenv("MAINTENANCE_MODE") == "true" {
		return errServerMaintenance
	}
	return nil
}
