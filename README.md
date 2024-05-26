
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

No parameters are required. Returns the current scores for the models.

This route returns a JSON object with the total number of requests that it has received during the previous 60 seconds. The route continues to return the correct numbers after restarting it, by persisting data to a file (requests.glob).

Response:
- 200: JSON marshalled count
- 405: Method Not allowed
- 500: Internal server error


Example:

![Screenshot 2024-05-26 at 15 19 41](https://github.com/takahiromitsui/go-server-counter/assets/78789212/b54bced5-1c35-4f7b-ab88-24ab26cc7a9c)



## Running Tests

To run tests, run the following command

```bash
  go test -v ./...
```

## Total amount of time I spent on this project

I spent roughly 4 hours finishing this project, including testing.
![Screenshot 2024-05-26 at 14 59 17](https://github.com/takahiromitsui/go-server-counter/assets/78789212/9d3012bf-9e9f-499b-926e-28be1f791ab5)
