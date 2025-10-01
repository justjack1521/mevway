package system

import (
	"errors"
	"fmt"
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

func (s *StatusService) Status(addresses []net.IP) error {

	if os.Getenv("MAINT_MODE") != "true" {
		return nil
	}

	var allowed = os.Getenv("MAINT_MODE_WHITELIST")

	if allowed == "" {
		return errServerMaintenance
	}

	var list = strings.Split(allowed, ",")
	if len(list) == 0 {
		return errServerMaintenance
	}

	var ips = make([]net.IP, 0)
	for _, value := range list {
		var ip = net.ParseIP(value)
		if ip != nil {
			ips = append(ips, ip)
		}
	}

	for _, value := range ips {
		for _, address := range addresses {
			fmt.Println(fmt.Sprintf("Check %s against %s", value, addresses))
			if value.Equal(address) {
				return nil
			}
		}
	}

	return errServerMaintenance

}
