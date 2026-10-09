# Liatrio Apprenticeship Project

## Project Description

A simple Go/Fiber web application that exposes an HTTP API endpoint. The application is containerized with Docker and will be deployed to a cloud platform using a CI/CD pipeline.

## Running Locally

After cloning the repository, run the following command in your terminal:

```bash
go run main.go
```

The application will run on `http://localhost:3000`.

### Testing the API

You can test the API endpoint using `curl`, Postman, or a similar API client.

Using `curl`:

```bash
curl -X GET http://localhost:3000/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Olivia"}'
```

Alternatively, create a `GET` request to `http://localhost:3000/` using Postman or another API client. Be sure to include the `name` value in the JSON request body.

## API Endpoints

### GET Endpoint
The API endpoint does not accept a request body and always returns a JSON object containing a default name, timestamp, and formatted time.

#### Example Response

```json
{
  "message": "My name is Olivia",
  "timestamp": 1791572962232,
  "formatted_time": "19:09:22 10-09-2026"
}
```

### POST Endpoint

The API endpoint expects a `name` value in the request body and returns a JSON object containing the provided name and the current timestamp.

#### Request Body

```json
{
  "name": "Olivia"
}
```

#### Example Response

```json
{
  "message": "My name is Olivia",
  "timestamp": 1791572962232,
  "formatted_time": "19:09:22 10-09-2026"
}
```

## Running with Docker

Build the Docker image:

```bash
docker build -t server .
```

Run the Docker container:

```bash
docker run -p 3000:3000 server
```

The application will now be available at `http://localhost:3000`.

### Testing the Dockerized API

#### GET
Using `curl`:

```bash
curl -X GET http://localhost:3000/
```

Alternatively, create a `GET` request to `http://localhost:3000/` using Postman or another API client.

#### POST
Using `curl`:

```bash
curl -X POST http://localhost:3000/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Olivia"}'
```

Alternatively, create a `POST` request to `http://localhost:3000/` using Postman or another API client. Be sure to include the `name` value in the JSON request body.


## GitHub Actions
The repository has a GitHub Actions Workflow that
  - Builds the application's Docker image
  - Verifies the application functionality using Liatrio's GitHub [apprentice-action/https://github.com/liatrio/github-actions/tree/master/apprentice-action]
  - On successful testing, pushes the image to Docker Hub
  - On successful push, deploys the image to Google Cloud Platform

### Testing the Deployed Application
Send a GET request to https://go-api-979593700395.us-central1.run.app/ without a request body as shown in previous instructions.

Send a POST request to https://go-api-979593700395.us-central1.run.app/ with the name in the body as shown in previous instructions.