package rpc

import (
	"github.com/justjack1521/mevconn"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func DialToLobbyClient() (services.MeviusMultiServiceClient, error) {
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.MULTISERVICENAME)
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusMultiServiceClient(conn), nil
}

func DialToSocialClient() (services.MeviusSocialServiceClient, error) {
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.SOCIALSERVICENAME)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithChainUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusSocialServiceClient(conn), nil
}

func DialToGameClient() (services.MeviusGameServiceClient, error) {
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.GAMESERVICENAME)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusGameServiceClient(conn), nil
}

func DialToAdminClient() (services.MeviusAdminServiceClient, error) {
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.GAMESERVICENAME)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusAdminServiceClient(conn), nil
}

func DialToRankClient() (services.MeviusRankServiceClient, error) {
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.RANKSERVICENAME)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusRankServiceClient(conn), nil
}
