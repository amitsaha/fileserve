### Start a HTTP server serving the files in a directory

Usage:

```
$ go run fileserve.go -d <directory> -p <port>
```

By default, the files in the current directory will be served on port 8080

```
$ go run fileserve.go 
```

To specify a different directory and port:

```
$ go run fileserve.go -d=/tmp -p=9090
```
