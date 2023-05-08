# shawandpartners-backend-challenge

## Run
    Inside project root
        $ docker-compose up -d
    or
        $ docker pull mrpsousa/api-go:0.0.0
        $ docker run -d -p 8080:8080 mrpsousa/api-go:0.0.0

### Endpoints
    - GET - /api/users?since={number} (...:8080/api/users?since=2) 
    - GET - /api/users/:username/details (...:8080/api/users/octocat/details)
    - GET - /api/users/:username/repos (...:8080/api/users/diego-shawandpartners/repos)

### Tests
    Inside project root
        $ go test ./...
