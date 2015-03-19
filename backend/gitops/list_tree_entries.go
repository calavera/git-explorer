package gitops

import (
	"path/filepath"

	"github.com/calavera/git-explorer/pb"
	"github.com/libgit2/git2go"
)

func (s *server) ListTreeEntries(rep *pb.Repository, stream pb.RepositoryExplorer_ListTreeEntriesServer) error {
	r, err := s.openRepository(rep.Name)
	if err != nil {
		return err
	}
	defer r.Free()

	var t *git.Tree
	b, err := lookupTree(r, rep.Ref)
	if err != nil {
		return err
	}
	defer b.Free()

	if rep.Base != "" {
		e, err := b.EntryByPath(rep.Base)
		if err != nil {
			return err
		}

		if e.Type != git.ObjectTree {
			return err
		}

		t, err = r.LookupTree(e.Id)
		if err != nil {
			return err
		}
		defer t.Free()
	} else {
		t = b
	}

	singleTreeWalker := func(base string, entry *git.TreeEntry) int {
		path := filepath.Join(rep.Base, entry.Name)

		e := &pb.TreeEntry{
			Oid:  entry.Id.String(),
			Name: entry.Name,
			Path: path,
			Type: pb.ObjectType(int(entry.Type) + 2),
		}

		if err := stream.Send(e); err != nil {
			return -1
		}

		if entry.Type == git.ObjectTree {
			return 1
		}

		return 0
	}

	t.Walk(singleTreeWalker)

	return nil
}
