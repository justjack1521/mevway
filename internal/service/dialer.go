package service

import (
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"google.golang.org/grpc"
	"mevway/internal/config"
)

func DialToAccessClient(config config.Application) (services.AccessServiceClient, error) {

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

	conn, err := grpc.Dial(config.AccessClient.ConnectionString(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return services.NewAccessServiceClient(conn), nil

}

func DialToLobbyClient(config config.Application) (services.MeviusMultiServiceClient, error) {
	conn, err := grpc.Dial(config.LobbyClient.ConnectionString(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return services.NewMeviusMultiServiceClient(conn), nil
}

func DialToSocialClient(config config.Application) (services.MeviusSocialServiceClient, error) {

	//cert, err := ioutil.ReadFile(config.PresenceClient.CertificatePath)
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

	conn, err := grpc.Dial(config.PresenceClient.ConnectionString(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return services.NewMeviusSocialServiceClient(conn), nil

}

func DialToGameClient(config config.Application) (services.MeviusGameServiceClient, error) {

	//cert, err := ioutil.ReadFile(config.GameClient.CertificatePath)
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

	conn, err := grpc.Dial(config.GameClient.ConnectionString(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return services.NewMeviusGameServiceClient(conn), nil

}

func DialToRankClient(config config.Application) (services.MeviusRankServiceClient, error) {

	//cert, err := ioutil.ReadFile(config.RankClient.CertificatePath)
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

	conn, err := grpc.Dial(config.RankClient.ConnectionString(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return services.NewMeviusRankServiceClient(conn), nil
}
