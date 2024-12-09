package port

import "context"

type ContactRepository interface {
	Create(ctx context.Context, email string, content string) error
}
