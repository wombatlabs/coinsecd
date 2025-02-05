package server

import (
	"context"
	"github.com/wombatlabs/coinsecd/cmd/coinsecwallet/daemon/pb"
	"github.com/wombatlabs/coinsecd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
