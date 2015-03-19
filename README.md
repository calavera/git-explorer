# Git explorer

Experimental git browser using [GRPC](http://grpc.io) and [git2go](https://github.com/libgit2/git2go).

This repo includes 2 applications, `frontend` is the web application to access from your browser. `backend` is the git server that access the git data.

The git repositories are stored in `~/git-explorer-data` by default but you can modify that location by exporting `GIT_EXPLORER_STORAGE_PATH`.

Git explorer uses GRPC streaming to forward data to your browser on slow tree requests to avoid buffering on the frontend.


## Installing dependencies

```
$ make deps
```

## Running frontend and backend

```
$ make run
```

## Moking lacency on your localhost

```
$ make slugish-lo0
```

## License

MIT
