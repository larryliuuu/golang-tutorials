# Go notes

## Setup

#### Enable dependency tracking

`go mod init [module-path]`

#### Dependency management

`go mod tidy` updates and clears module dependencies.

`go mod download` downloads modules from the internet and stores them locally.

`go mod vendor`  create a "vendor" directory within the root of a Go project's source code, which contains copies of all of the modules that are required by the project, including their dependencies. This allows a project to have a self-contained copy of its dependencies, rather than relying on those dependencies to be present in the global cache or in a shared location.

Vendoring is useful in situations where an application needs to be deployed to an environment where the Go toolchain and the modules required by the application may not be available.

When using vendor, go build, go run, go test and other go command will prioritize the vendor folder modules than normal modules.

It is important to notice that once a project vendors its dependencies, it can't use go mod command anymore and will have to manage the dependencies on its own by updating them inside the vendor folder.

#### Exportables

Upper-case 

