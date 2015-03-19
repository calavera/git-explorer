package gitops

import (
	"net/url"
	"strings"

	"github.com/calavera/git-explorer/pb"
	"github.com/libgit2/git2go"
	"golang.org/x/net/context"
)

func (s *server) CreateRepo(ctx context.Context, rep *pb.Repository) (*pb.Repository, error) {
	opts := &git.CloneOptions{Bare: true}
	u := rep.Url

	x, _ := url.Parse(u)
	parts := strings.Split(x.Path, "/")
	name := parts[len(parts)-1]

	r, err := git.Clone(u, repositoryPath(s.storagePath, name), opts)
	if err != nil {
		return nil, err
	}
	defer r.Free()

	return &pb.Repository{
		Name: name,
		Ref:  "master",
	}, nil
}
