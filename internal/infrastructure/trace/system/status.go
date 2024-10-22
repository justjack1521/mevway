package system

import (
	"errors"
	"net"
	"os"
	"strings"
)

var (
	errServerMaintenance = errors.New("service is undergoing maintenance")
)

type StatusService struct {
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Status(address net.IP) error {

	if os.Getenv("MAINTENANCE_MODE") != "true" {
		return nil
	}

	var allowed = os.Getenv("MAINTENANCE_MODE_WHITELIST")

	if allowed == "" {
		return errServerMaintenance
	}

	var list = strings.Split(allowed, ",")
	if len(list) == 0 {
		return errServerMaintenance
	}

	var ips = make([]net.IP, len(list))

	for index, value := range list {
		ips[index] = net.ParseIP(value)
	}

	for _, value := range ips {
		if value.Equal(address) {
			return nil
		}
	}

	return errServerMaintenance

}
