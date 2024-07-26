package server

import (
	"context"
	"github.com/coinsec/coinsecd/cmd/coinsecwallet/daemon/pb"
	"github.com/coinsec/coinsecd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
