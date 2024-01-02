# Run
run `go mod tidy` to fix dependencies.

run `go mod vendor` to bring dependenciy packages.

run `go run ./cmd/.` to start app.

### Run with custom config
This application use environment for config that there is in [`config/.env`](github.com/maadiii/helli/config/.env) to run in dev mode, so you can change that envs or alternate set `MODE` env in os envs to `product` to use envs of os.

### Run test
run `go test ./... -v` to run tests.

# Description of project structure
`internal` directory is main directory that consist of `application` for web handlers, `domain` for handling business logic, `data_access` for handling Data Layer, and `entity` that contains entities of business.
