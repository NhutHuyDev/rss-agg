## RSS-AGG (Rich Site Summary Aggregation)
A simple [RSS](https://en.wikipedia.org/wiki/RSS) aggregator backend built by Go, that allows users to follow to and manage RSS feeds via [HTTP/1.1](https://pkg.go.dev/net/http).
- RESTful APIs using JSON format

- Handle background tasks for scraping blogs from RSS feeds via Goroutines

- PostgreSQL, [goose](https://github.com/pressly/goose), and [sqlc](https://github.com/sqlc-dev/sqlc)

- Manage routes via [Chi](https://github.com/go-chi/chi)

- [Go validator](https://github.com/go-playground/validator)

## How to Run
### Prerequisites
- Go 

- PostgreSQL

- ```sqlc``` installed for code generation

- ```goose``` insstalled for database migration

### Database Migration
**./sql/schema**
```
goose postgres <"database connection string"> up
```
```
goose postgres <"database connection string"> down
```

### Generate Code with ```sqlc```
```
sqlc generate
```

###  Run the Server
**./cmd/rssagg**
```
go run .
```

## API appendix

### User API

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/users`        | Get a user           | N/A                  | `{"id": "757a8984-da5d-4ef8-bf8e-ac855e8ecf47", "created_at": "2024-09-29T12:50:14.5237Z", "updated_at": "2024-09-29T12:50:14.5237Z", "name": "John\n", "api_key": "3c6b5bac369f498ace97215aa11284c4ee495fefb52d27c2bbb1b38e2cb4342f"}`
 | Yes             |
| POST    | `/v1/users`   | Create a specific user  | `{"name": "Huy"}` | `{"id": 1, "name": "John Doe", "email": "john@example.com"}` | No            |



### Notes:
- All responses are in JSON format as well.

- For POST, PUT, DELETE requests, you need to send the request body in JSON format.

- For endpoints marked with "Yes" in the Authentication column, a valid API key is required.

- The API key is sent using the `Authorization` header, formatted as follows: 
    ```
    Authorization: ApiKey <"API key">
    ```

## References
[1]. freeCodeCamp.org. (2023, May 11). Go Programming – Golang Course with Bonus Projects [Video]. YouTube. https://www.youtube.com/watch?v=un6ZyFkqFKo

[2]. Johnson, B. (2021, January 4). Standard Package Layout - Ben Johnson - Medium. Medium. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

[3]. Golang-Standards. (n.d.). GitHub - golang-standards/project-layout: Standard Go Project Layout. GitHub. https://github.com/golang-standards/project-layout

