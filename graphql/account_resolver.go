package main

import "context"

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, pagination *PaginationInput, obj *Account) ([]*Order, error) {

}
