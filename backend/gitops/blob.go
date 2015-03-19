package gitops

import (
	"github.com/calavera/git-explorer/pb"
	"github.com/libgit2/git2go"
	"golang.org/x/net/context"
)

func (s *server) GetBlobData(ctx context.Context, rep *pb.Repository) (*pb.Blob, error) {
	r, err := s.openRepository(rep.Name)
	if err != nil {
		return nil, err
	}
	defer r.Free()

	id, err := git.NewOid(rep.Ref)
	if err != nil {
		return nil, err
	}

	b, err := r.LookupBlob(id)
	if err != nil {
		return nil, err
	}
	defer b.Free()

	return &pb.Blob{
		Oid:  b.Id().String(),
		Data: b.Contents(),
	}, nil
}
