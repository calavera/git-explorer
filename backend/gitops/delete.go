package gitops

import (
	"os"

	"github.com/calavera/git-explorer/pb"
	"golang.org/x/net/context"
)

func (s *server) DeleteRepo(ctx context.Context, r *pb.Repository) (*pb.Empty, error) {
	repo := repositoryPath(s.storagePath, r.Name)
	err := os.RemoveAll(repo)
	return &pb.Empty{}, err
}
