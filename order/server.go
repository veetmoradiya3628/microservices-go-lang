package order

import (
	"fmt"
	"net"

	"github.com/veetmoradiya3628/microservices-go-lang/account"
	"github.com/veetmoradiya3628/microservices-go-lang/catalog"
	"github.com/veetmoradiya3628/microservices-go-lang/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedOrderServiceServer
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewClient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &grpcServer{
		service: s,
		accountClient,
		catalogClient,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}
