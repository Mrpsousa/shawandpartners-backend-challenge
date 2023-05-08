# shawandpartners-backend-challenge

## Run
    Inside project root
        $ docker-compose up -d

### Endpoints
    - GET - /api/users?since={number} (...:8080/api/users?since=2) 
    - GET - /api/users/:username/details (...:8080/api/users/octocat/details)
    - GET - /api/users/:username/repos (...:8080/api/users/octocat/repos)

### Tests
    Inside project root
        $ go test ./...
