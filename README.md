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

- ```goose``` installed for database migration

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

**./cmd/rssagg/.env**
```
PORT=<"port">
DB_CONNECTION_STR=<"database connection string">
```
**./cmd/rssagg**
```
go run .
```

## Database Diagram
![rssagg](https://github.com/user-attachments/assets/51f5dccf-be1d-41d3-b7fd-8e23d868e9ba)

## API appendix

### User APIs

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/users`        | Get a user           | N/A                  | `{"id": "757a8984-da5d-4ef8-bf8e-ac855e8ecf47", "created_at": "2024-09-29T12:50:14.5237Z", "updated_at": "2024-09-29T12:50:14.5237Z", "name": "John", "api_key": "3c6b5bac369f498ace97215aa11284c4ee495fefb52d27c2bbb1b38e2cb4342f"}`| Yes             |
| POST    | `/v1/users`   | Create a specific user  | `{"name": "John"}` |  `{"id": "757a8984-da5d-4ef8-bf8e-ac855e8ecf47", "created_at": "2024-09-29T12:50:14.5237Z", "updated_at": "2024-09-29T12:50:14.5237Z", "name": "John", "api_key": "3c6b5bac369f498ace97215aa11284c4ee495fefb52d27c2bbb1b38e2cb4342f"}` | No            |

### Feed APIs

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/feeds?limit=10&page=1`        | Get all feeds           | N/A                  | `{"feeds": [{"id": "74b363cc-69ba-43b7-ac11-21f18a58f2c6", "created_at": "2024-10-03T07:22:44.796866Z", "updated_at": "2024-10-03T14:23:15.561915Z", "name": "Wagslane blog", "url": "https://www.wagslane.dev/index.xml"}], "total": 5, "total_page": 2, "current_page": 1, "next_page": 2}`| No             |
| POST    | `/v1/feeds`   | Create a specific feed  | `{"name": "Wagslane blog", "url": "https://www.wagslane.dev/index.xml"}` | `{"id": "74b363cc-69ba-43b7-ac11-21f18a58f2c6", "created_at": "2024-10-03T07:22:44.796866Z", "updated_at": "2024-10-03T07:22:44.796866Z", "name": "Wagslane blog", "url": "https://www.wagslane.dev/index.xml"}` | Yes            |

### Feed Follow APIs

| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/feed_follows`        | Get all feeds that a specific user followed    | N/A                  | `[{"id": "57af56ce-bed1-4d0e-879c-073555835a9e", "created_at": "2024-10-01T09:11:05.593611Z", "updated_at": "2024-10-01T09:11:05.593611Z", "feed_id": "7d9a2d5d-1193-49fb-a841-2cc5bda80c66", "user_id": "757a8984-da5d-4ef8-bf8e-ac855e8ecf47"}]`| Yes             |
| POST    | `/v1/feed_follows`   | Follow a specific feed  | `{"feed_id": "31b23de0-37cf-4479-b4c8-8576ff294603"}` | `{"id": "cce009a0-7155-4145-ada3-e91cba2aea7a", "created_at": "2024-10-03T07:30:06.166133Z", "updated_at": "2024-10-03T07:30:06.166133Z", "feed_id": "31b23de0-37cf-4479-b4c8-8576ff294603", "user_id": "757a8984-da5d-4ef8-bf8e-ac855e8ecf47"}` | Yes            |
| DELETE    | `/v1/feed_follows/:feed_folow_id`   | Unfollow a specific feed  | N/A | `{}` | Yes            |

### Post APIs
| Method | Endpoint       | Description                     | Request Body Example         | Response Body Example                                       | Authentication |
|--------|----------------|----------------------------------|----------------------|-------------------------------------------------------------|----------------|
| GET    | `/v1/posts?limit=10&page=1`        | Get all posts from a RSS feed that a specific user followed | N/A                  | `{"posts": [{"id": "31882c8b-053c-4901-be7b-e0da2b192b57", "created_at": "2024-10-03T05:29:46.457209Z", "updated_at": "2024-10-03T05:29:46.457209Z", "title": "The Boot.dev Beat. April 2024", "description": "Pythogoras returned in our second community-wide boss battle. He was vanquished, and there was much rejoicing.", "published_at": "2024-04-03T00:00:00Z", "url": "https://blog.boot.dev/news/bootdev-beat-2024-04/", "feed_id": "7d9a2d5d-1193-49fb-a841-2cc5bda80c66"}], "total": 384, "total_page": 384, "current_page": 10, "next_page": 11}` | Yes             |

### Notes:
- All responses are in JSON format as well.

- For endpoints marked with "Yes" in the Authentication column, a valid API key is required.

- The API key is sent using the `Authorization` header, formatted as follows: 
    ```
    Authorization: ApiKey <"API key">
    ```

- Can use [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension to test these APIs

## References
[1]. freeCodeCamp.org. (2023, May 11). Go Programming – Golang Course with Bonus Projects [Video]. YouTube. https://www.youtube.com/watch?v=un6ZyFkqFKo

[2]. Johnson, B. (2021, January 4). Standard Package Layout - Ben Johnson - Medium. Medium. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

[3]. Golang-Standards. (n.d.). GitHub - golang-standards/project-layout: Standard Go Project Layout. GitHub. https://github.com/golang-standards/project-layout

