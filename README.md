
#  Go Server Counter API
This is a simple Go server that uses the `http.ServeMux` router to handle HTTP requests.
## Run Locally

To run the application, use the following command:

```bash
  go run cmd/api/*.go
```
## Routes

This application has the following routes:

- **/counter**: Returns count the total number of requests that it has received during the previous 60 seconds.
## API Reference

#### Get Count

```http
  GET /counter
```

No parameters required. Returns the current scores for the models.

This route returns a JSON object with the total number of requests that it has received during the previous 60 seconds. The route continues to return the correct numbers after restarting it, by persisting data to a file (requests.glob).

Response:
- 200: JSON marshalled count
- 405: Method Not allowed
- 500: Internal server error


## Running Tests

To run tests, run the following command

```bash
  go test -v ./...
```

