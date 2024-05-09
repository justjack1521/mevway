package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	uuid "github.com/satori/go.uuid"
)

type CustomerPlayerIDRepository struct {
	client services.AccessServiceClient
}

func NewCustomerPlayerIDRepository(client services.AccessServiceClient) *CustomerPlayerIDRepository {
	return &CustomerPlayerIDRepository{client: client}
}

func (r *CustomerPlayerIDRepository) Get(ctx context.Context, customer string) (uuid.UUID, error) {
	result, err := r.client.CustomerSearch(ctx, &protoaccess.CustomerSearchRequest{CustomerId: customer})
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromString(result.PlayerId)
}
