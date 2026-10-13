package account

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	// pb.
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, req *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccount(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{Name: account.Name, Id: account.ID}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: &pb.Account{Name: account.Name, Id: account.ID}}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, req *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	accounts, err := s.service.GetAccounts(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	accountProtos := make([]*pb.Account, len(accounts))
	for i, acc := range accounts {
		accountProtos[i] = &pb.Account{Name: acc.Name, Id: acc.ID}
	}
	return &pb.GetAccountsResponse{Accounts: accountProtos}, nil
}
