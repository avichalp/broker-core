package service

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	golog "github.com/ipfs/go-log/v2"
	"github.com/textileio/broker-core/broker"
	"github.com/textileio/broker-core/cmd/piecerd/piecer"
	"github.com/textileio/broker-core/dshelper/txndswrap"
	"github.com/textileio/broker-core/finalizer"
	pb "github.com/textileio/broker-core/gen/broker/piecer/v1"
	"github.com/textileio/broker-core/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var log = golog.Logger("piecer/service")

// Config defines params for Service configuration.
type Config struct {
	Listener net.Listener

	IpfsClient *httpapi.HttpApi
	Broker     broker.Broker
	Datastore  txndswrap.TxnDatastore

	DaemonFrequency time.Duration
	RetryDelay      time.Duration
}

// Service is a gRPC service wrapper around a piecer.
type Service struct {
	pb.UnimplementedAPIServiceServer

	server    *grpc.Server
	piecer    *piecer.Piecer
	finalizer *finalizer.Finalizer
}

var _ pb.APIServiceServer = (*Service)(nil)

// New returns a new Service.
func New(conf Config) (*Service, error) {
	fin := finalizer.NewFinalizer()

	lib, err := piecer.New(conf.Datastore, conf.IpfsClient, conf.Broker, conf.DaemonFrequency, conf.RetryDelay)
	if err != nil {
		return nil, fin.Cleanupf("creating piecer: %v", err)
	}
	fin.Add(lib)

	s := &Service{
		server:    grpc.NewServer(),
		piecer:    lib,
		finalizer: fin,
	}

	go func() {
		pb.RegisterAPIServiceServer(s.server, s)
		if err := s.server.Serve(conf.Listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Errorf("server error: %v", err)
		}
	}()

	return s, nil
}

// ReadyToPrepare indicates that a batch is ready to be prepared.
func (s *Service) ReadyToPrepare(ctx context.Context, r *pb.ReadyToPrepareRequest) (*pb.ReadyToPrepareResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if r.StorageDealId == "" {
		return nil, status.Error(codes.InvalidArgument, "storage deal id is empty")
	}
	dataCid, err := cid.Decode(r.DataCid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "decoding data cid: %s", err)
	}
	if !dataCid.Defined() {
		return nil, status.Error(codes.InvalidArgument, "data cid is undefined")
	}

	if err := s.piecer.ReadyToPrepare(ctx, broker.StorageDealID(r.StorageDealId), dataCid); err != nil {
		return nil, status.Errorf(codes.Internal, "queuing data-cid to be prepared: %s", err)
	}

	return &pb.ReadyToPrepareResponse{}, nil
}

// Close the service.
func (s *Service) Close() error {
	rpc.StopServer(s.server)
	log.Info("service was shutdown")
	return s.finalizer.Cleanup(nil)
}