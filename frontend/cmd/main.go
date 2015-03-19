package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/calavera/git-explorer/pb"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const serverAddr = "localhost:9090"

type server struct {
	gitConn *grpc.ClientConn
}

func startServer() {
	s := &server{}

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public")).ServeHTTP

	router.POST("/create", s.createRepo)
	router.POST("/r/:name/fetch", s.fetchRepo)
	router.POST("/r/:name/delete", s.deleteRepo)

	router.GET("/r/:name/tree/:ref", s.getTreeEntries)
	router.GET("/r/:name/tree/:ref/*path", s.getTreeEntries)
	router.GET("/r/:name/blob/:ref/*path", s.getBlob)

	log.Println("Frontend server started in :9091")
	log.Fatal(http.ListenAndServe(":9091", router))
}

func (s *server) closeConnection() {
	if s.gitConn != nil {
		s.gitConn.Close()
	}
}

func (s *server) repositoryClient() pb.RepositoryExplorerClient {
	if s.gitConn == nil {
		var err error
		s.gitConn, err = grpc.Dial(serverAddr)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
	}
	return pb.NewRepositoryExplorerClient(s.gitConn)
}

func (s *server) getBlob(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rep := repositoryFromParams(ps)
	blob, err := s.repositoryClient().GetBlobData(context.Background(), rep)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	rw.Write(blob.Data)
}

func (s *server) fetchRepo(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rep := &pb.Repository{Name: ps.ByName("name")}
	c, err := s.repositoryClient().FetchRepo(context.Background(), rep)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	u := fmt.Sprintf("/r/%s/tree/%s", rep.Name, c.Oid)
	http.Redirect(rw, r, u, 302)
}

func (s *server) createRepo(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rep := &pb.Repository{
		Url: r.FormValue("url"),
	}
	rr, err := s.repositoryClient().CreateRepo(context.Background(), rep)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	u := fmt.Sprintf("/r/%s/tree/%s", rr.Name, rr.Ref)
	http.Redirect(rw, r, u, 302)
}

func (s *server) deleteRepo(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rep := &pb.Repository{Name: ps.ByName("name")}
	_, err := s.repositoryClient().DeleteRepo(context.Background(), rep)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	http.Redirect(rw, r, "/", 302)
}

func (s *server) getTreeEntries(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Make sure that the writer supports flushing.
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/html")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	rep := repositoryFromParams(ps)
	stream, err := s.repositoryClient().ListTreeEntries(context.Background(), rep)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	writeButtons(rw, rep.Name)

	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		if entry.Type == pb.ObjectType_TreeType {
			fmt.Fprintf(rw, "<div><a href=\"/r/%s/tree/%s/%s\">%s</a></div>", rep.Name, rep.Ref, entry.Path, entry.Name)
		} else {
			fmt.Fprintf(rw, "<div><a href=\"/r/%s/blob/%s/%s\">%s</a></div>", rep.Name, entry.Oid, entry.Path, entry.Name)
		}
		flusher.Flush()
	}
}

func writeButtons(rw http.ResponseWriter, name string) {
	fetch := fmt.Sprintf("<form action=\"/r/%s/fetch\" method=\"post\"><button type\"submit\">Fetch repo!</button></form>", name)
	del := fmt.Sprintf("<form action=\"/r/%s/delete\" method=\"post\"><button type\"submit\">Delete repo!</button></form>", name)
	fmt.Fprintf(rw, "<div>%s%s</div>", fetch, del)
}

func repositoryFromParams(ps httprouter.Params) *pb.Repository {
	return &pb.Repository{
		Name: ps.ByName("name"),
		Ref:  ps.ByName("ref"),
		Base: strings.TrimLeft(ps.ByName("path"), "/"),
	}
}

func main() {
	startServer()
}
