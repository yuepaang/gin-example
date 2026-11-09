# data-microservices
Web services for data preprocessing

## Build and security checks

Use Go 1.27.3 or newer, including the latest security patches. Run the Gin
example explicitly, since the repository root also contains two independent
example programs (`staticserver.go` and `test.go`):

```sh
go run main.go
```

The security workflow builds `main.go` and checks its imported packages against
the Go vulnerability database weekly and on pull requests. To run the same
checks locally:

```sh
go mod verify
go build -mod=readonly -o gin-example main.go
go run golang.org/x/vuln/cmd/govulncheck@v1.9.0 -mode=binary -scan=package gin-example
```

This check covers the Gin executable, not every example or unused dependency.
`go test ./...` currently fails on existing duplicate entry points, the legacy
import in `models`, and undefined settings variables. These require separate
cleanup before a repository-wide source scan can run.
