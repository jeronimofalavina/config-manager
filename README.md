# Configuration Management Service
The Configuration Management Service is a simple HTTP service that allows you to store and retrieve configurations based on certain conditions. It provides basic CRUD (Create, Read, Update, Delete) operations for managing configurations.

## Project Structure
The project is organized with the following file structure:
```
.
├── Dockerfile
├── README.md
├── TODO
├── cmd
│   └── api
│       ├── handler
│       │   └── configs.go
│       ├── main.go
│       ├── models
│       │   └── config.go
│       └── routes
│           └── routes.go
├── go.mod
├── go.sum
├── kube
│   ├── app.yaml
│   └── kind.yaml
└── test
    └── main_test.go
```
### How to Run
Before running the service, make sure you have the following tools installed:
* Docker
* kind (Kubernetes in Docker)
* kubectl (Kubernetes command-line tool)
* Go (if you want to build and run the service locally)

### Creating and Deleting the Cluster
Your can run the following commands to deploy
```shell
make create_cluster
```

```shell
make delete_cluster
```
#### To deploy the code in the new kind cluster
```shell
make deploy
```
#### To test the code
```shell
make test
```

#### Configuration API endpoints

* GET /configs: List all configurations.
* POST /configs: Create a new configuration.
* GET /configs/<name>: Get a configuration by name.
* PUT /configs/<name>: Update a configuration by name.
* PATCH /configs/<name>: Partially update a configuration by name.
* DELETE /configs/<name>: Delete a configuration by name.
* GET /search: Query configurations based on metadata.

### API Operations
You can interact with the service using the provided API operations in the Makefile:

* make list: List all configurations.
* make create: Create a new configuration (example provided in examples/create.json).
* make get PARAM=<config-name>: Get a configuration by name (replace <config-name> with the desired name).
* make updatePUT PARAM=<config-name>: Update a configuration using the HTTP PUT method (replace <config-name> with the desired name).
* make updatePATCH PARAM=<config-name>: Update a configuration using the HTTP PATCH method (replace <config-name> with the desired name).
* make delete PARAM=<config-name>: Delete a configuration by name (replace <config-name> with the desired name).
* make query: Perform a sample query operation.
