package gitops

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/libgit2/git2go"
)

type server struct {
	storagePath string
}

func NewServer() *server {
	return &server{
		storagePath: os.Getenv("GIT_EXPLORER_STORAGE_PATH"),
	}
}

func (s *server) openRepository(name string) (*git.Repository, error) {
	return git.OpenRepository(repositoryPath(s.storagePath, name))
}

func repositoryPath(storage, name string) string {
	return filepath.Join(storage, md5Hash(name))
}

func md5Hash(name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name))
	return hex.EncodeToString(hasher.Sum(nil))
}

func lookupTree(r *git.Repository, ref string) (*git.Tree, error) {
	if ref == "" {
		ref = "master"
	}

	o, err := r.RevparseSingle(ref)
	commitId := o.Id()
	o.Free()

	commit, err := r.LookupCommit(commitId)
	if err != nil {
		return nil, err
	}
	defer commit.Free()

	return commit.Tree()
}
