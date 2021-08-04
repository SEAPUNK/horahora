package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/SEAPUNK/horahora/archiver/internal/archiverequests"
	"github.com/SEAPUNK/horahora/archiver/internal/config"
	proto "github.com/SEAPUNK/horahora/archiver/protocol"
	"google.golang.org/grpc"
)

type archiverServer struct {
	proto.UnimplementedArchiverServer
	Cfg *config.Config
}

func NewGRPCServer(ctx context.Context, cfg *config.Config) error {
	archiverServer := initializeArchiverServer(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	proto.RegisterArchiverServer(serv, archiverServer)

	go func() {
		<-ctx.Done()
		serv.GracefulStop()
	}()

	return serv.Serve(lis)
}

func initializeArchiverServer(cfg *config.Config) archiverServer {
	return archiverServer{
		Cfg: cfg,
	}
}

func (s archiverServer) CreateArchiveRequest(ctx context.Context, req *proto.CreateArchiveRequestRequest) (*proto.CreateArchiveRequestResponse, error) {
	archiveId, err := archiverequests.CreateArchiveRequest(s.Cfg,
		&archiverequests.UserArchiveRequest{
			UserId: req.UserId,
			ArchiveRequest: archiverequests.ArchiveRequest{
				Query: req.Query,
			},
		})

	// we don't care if this fails, cuz it can get queued later on
	go func() {
		// TODO(ivan): we should still add some logging or something to this though
		archiverequests.QueueArchiveRequest(s.Cfg, archiveId)
	}()

	return &proto.CreateArchiveRequestResponse{
		ArchiveId: archiveId,
	}, err
}

func (s archiverServer) ListArchiveRequestsForUser(ctx context.Context, req *proto.ListArchiveRequestsForUserRequest) (*proto.ListArchiveRequestsForUserResponse, error) {
	var entries []*proto.UserArchiveRequest

	ars, err := archiverequests.ListArchiveRequestsForUser(s.Cfg, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, ar := range ars {
		entries = append(entries, &proto.UserArchiveRequest{
			ArchiveRequest: &proto.ArchiveRequest{
				Id:    ar.Id,
				Query: ar.Query,
				Error: ar.Error.String,
			},
		})
	}

	return &proto.ListArchiveRequestsForUserResponse{
		Entries: entries,
	}, nil
}
