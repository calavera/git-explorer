package gitops

import (
	"github.com/calavera/git-explorer/pb"
	"golang.org/x/net/context"
)

func (s *server) FetchRepo(ctx context.Context, rep *pb.Repository) (*pb.Commit, error) {
	r, err := s.openRepository(rep.Name)
	if err != nil {
		return nil, err
	}
	defer r.Free()

	rem, err := r.LookupRemote("origin")
	if err != nil {
		return nil, err
	}
	defer rem.Free()

	var refs []string
	err = rem.Fetch(refs, "")
	if err != nil {
		return nil, err
	}

	h, err := r.Head()
	if err != nil {
		return nil, err
	}

	return &pb.Commit{
		Oid: h.Target().String(),
	}, nil
}
