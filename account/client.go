package account

import (
	"context"

	pb "github.com/veetmoradiya3628/microservices-go-lang/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	service := pb.NewAccountServiceClient(conn)
	return &Client{conn: conn, service: service}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostAccount(ctx context.Context, name string) (*Account, error) {
	resp, err := c.service.PostAccount(
		ctx,
		&pb.PostAccountRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}
	return &Account{ID: resp.Account.Id, Name: resp.Account.Name}, nil
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	resp, err := c.service.GetAccount(
		ctx,
		&pb.GetAccountRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Account{ID: resp.Account.Id, Name: resp.Account.Name}, nil
}

func (c *Client) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	resp, err := c.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{Skip: skip, Take: take},
	)
	if err != nil {
		return nil, err
	}
	accounts := make([]Account, len(resp.Accounts))
	for i, acc := range resp.Accounts {
		accounts[i] = Account{ID: acc.Id, Name: acc.Name}
	}
	return accounts, nil
}
