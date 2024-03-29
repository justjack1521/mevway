package service

import (
	"github.com/justjack1521/mevconn"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func DialToAccessClient() (services.AccessServiceClient, error) {
	//cert, err := ioutil.ReadFile(config.AccessClient.CertificatePath)
	//if err != nil {
	//	return nil, err
	//}
	//
	//pool := x509.NewCertPool()
	//if !pool.AppendCertsFromPEM(cert) {
	//	return nil, fmt.Errorf("failed to add server CA's certificate")
	//}
	//
	//cred := credentials.NewTLS(&tls.Config{
	//	RootCAs: pool,
	//})
	config, err := mevconn.CreateGrpcServiceConfig(mevconn.AUTHSERVICENAME)
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewAccessServiceClient(conn), nil

}

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
	conn, err := grpc.Dial(config.ConnectionString(), grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor), grpc.WithInsecure())
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
