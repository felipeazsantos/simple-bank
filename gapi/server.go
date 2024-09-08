package gapi

import (
	"fmt"

	db "github.com/felipeazsantos/simple_bank/db/sqlc"
	"github.com/felipeazsantos/simple_bank/pb"
	"github.com/felipeazsantos/simple_bank/token"
	"github.com/felipeazsantos/simple_bank/util"
)

// Server serves gRPC requests for our baking service
type Server struct {
	pb.SimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
