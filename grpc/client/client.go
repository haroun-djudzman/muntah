package client

import (
	"context"

	muntahProto "github.com/haroun-djudzman/muntah/proto"

	"google.golang.org/grpc"
)

type MuntahGRPC struct {
	muntahProto.MuntahClient
	conn *grpc.ClientConn
}

func NewClientConn(address string) (*MuntahGRPC, error) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, address,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &MuntahGRPC{
		MuntahClient: muntahProto.NewMuntahClient(conn),
		conn:         conn,
	}, nil
}
