package port

import (
	"context"
	"net"
)

type SystemStatusService interface {
	Status(ctx context.Context, address []net.IP) error
}
