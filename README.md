# shawandpartners-backend-challenge

## Run
    Inside project root
        $ docker-compose up -d
    or
        $ docker pull mrpsousa/api-go:0.0.3
        $ docker run -d -p 8080:8080 mrpsousa/api-go:0.0.3

### Endpoints
    - GET - /api/users?since={number} (example: localhost:8080/api/users?since=2) (returning 50 as maximum)
    - GET - /api/users/:username/details (example: localhost:8080/api/users/octocat/details)
    - GET - /api/users/:username/repos (example: localhost:8080/api/users/diego-shawandpartners/repos)

### Run tests
    Inside project root
        $ go test ./...

### Tests Coverage
    - GET - http://localhost:8080/api/coverage
    (I couldn't finish in time)

### Documentation
    - GET - http://localhost:8080/api/docs
    (I couldn't finish in time)

