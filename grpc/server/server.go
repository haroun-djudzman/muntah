package server

import (
	"context"
	"crypto/rand"
	"fmt"
	"net"

	pb "github.com/haroun-djudzman/muntah/proto"
	"google.golang.org/grpc"
)

type mulut struct {
	muntahSize [][]byte
}

type Handler struct {
	pb.UnimplementedMuntahServer
	mulut mulut
}

func New() *Handler {
	size := []int{10000, 25000, 50000, 75000, 100000, 200000}
	muntah := [6][]byte{}

	for i, v := range size {
		token := make([]byte, v)
		rand.Read(token)

		muntah[i] = token
	}

	mulut := mulut{
		muntahSize: muntah[:],
	}

	return &Handler{
		mulut: mulut,
	}
}

func (h *Handler) GetMuntah(ctx context.Context, in *pb.GetMuntahRequest) (*pb.GetMuntahResponse, error) {
	index := in.GetIndex()

	return &pb.GetMuntahResponse{
		Muntah: h.mulut.muntahSize[index-1],
	}, nil
}

func StartServer() error {
	muntahHandler := New()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 6969))
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterMuntahServer(server, muntahHandler)

	return server.Serve(lis)
}
