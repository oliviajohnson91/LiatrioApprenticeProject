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

## API Endpoint

The API endpoint expects a `name` value in the request body and returns a JSON object containing the provided name and the current timestamp.

### Request Body

```json
{
  "name": "Olivia"
}
```

### Example Response

```json
{
  "message": "My name is Olivia",
  "timestamp": 1791232163
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

Using `curl`:

```bash
curl -X GET http://localhost:3000/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Olivia"}'
```

Alternatively, create a `GET` request to `http://localhost:3000/` using Postman or another API client. Be sure to include the `name` value in the JSON request body.
