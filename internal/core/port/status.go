package port

import (
	"net"
)

type SystemStatusService interface {
	Status(address []net.IP) error
}
