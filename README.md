# Fiber Go Template

This is the backend part of a full stack template, it uses Go and the Fiber framework.

![Frontend part repository](https://github.com/HamzaDLM/vue_ts_template)

## Usage

### Running locally 

To run the example locally, clone the project & use the helper makefile commands:
```
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
```
go test -v ./...
```

### Documentation

To update the api documentation use the following:
```
swag init
```

### Makefile

There is a makefile with helper commands to make life easier:

List available commands
```
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


## TODO

- [] CORS (not working)
- [] Deployment files (Docker, K8S, ...)
- [] Grafana dashboard template for monitoring
