# Fiber Go Template

This is the backend part of a full stack template, it uses Go and the Fiber framework.

[Frontend part repository](https://github.com/HamzaDLM/vue_ts_template)

## Usage

### Running locally 

To run the example locally, clone the project & use the helper makefile commands:
```bash
# clone the project
git clone https://github.com/HamzaDLM/fiber_go_template

# run the app with reload mode using Air
air -c .air.conf

# or using docker, build the image
docker build -t fiber_go_backend .

# and run it 
docker run -p [host_port]:[container_port] fiber_go_backend
```

Then you should be able to access the app using the following URLs:

- Web: <http://localhost:port/>
- API: <http://localhost:port/v1/>
- Doc: <http://localhost:port/v1/docs>

### Tests

To run the tests:
```bash
go test -v ./...
```

### Documentation

To update the api documentation use the following:
```bash
swag init
```

### Makefile

There is a makefile with helper commands to make life easier:

List available commands
```bash
make list
```

## Libraries

| Library   | Version    |
|--------------- | --------------- |
| fiber   | 2.51.0   |
| gorm   | 1.25.5   |
| zap   | 1.26.0   |
| swag  | 1.1.0  |
| testify | 1.8.4 | 

## Folder structure

```bash
.
├── config                  # Hold the configuration files of the project 
├── container               # Defines data and tools that are attached to most layers (env, logger...)
├── controller              # Defines API handlers 
├── database                # Database abstraction
├── docs                    # Auto-generated files used by swagger
├── logger                  # Provides zap logging functionalities 
├── middleware              # Defining middlewares 
├── public                  # Where the built frontend 
├── router                  # Provides routing
├── tests                   # Unit testing 
├── main.go                 # Entry point
├── makefile                # Helper commands
└── README.md               
```

## TODO

- [] CORS (not working)
- [] Deployment files (Docker, K8S, ...)
- [] Grafana dashboard template for monitoring
