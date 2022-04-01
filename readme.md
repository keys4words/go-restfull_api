# pizza api (trash api :) )
1. data in slice
2. routes:
- localhost:8080/pizzas -> get json with all pizzas
- localhost:8080/pizza/{id} -> get json with this pizza ingredients if pizza is availabe or null
3. use response codes for errors & success requests
4. Collection of Tests in Thunder Client at pizzaApi.json

# books api (semi-trash api)
1. data in slice of structures
2. standard project structure
3. use godotenv
3. routes:
- base: localhost:8080/api/v1
- GET    base/books -> json with all books
- GET    base/books/1 -> json with book #1
- POST   base/book/ -> add new book
- PUT    base/books/1 -> update existing book
- DELETE base/books/1 -> delete existing book
4. Collection of Tests in Thunder Client at booksApi.json

# article API
1. use golang-standards/project-layout for project structuring
2. data in postgres DB
3. use configs in toml or .env formats
4. use golang-migrate for migrations -> migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up/down
5. add authentication by JWT for delete/post article
6. add logs by Logrus
7. Collection of tests articleApi
7. routes:
- base:   localhost:8080/api/v1
- GET     base/articles -> json with all articles
- GET     base/articles/1 -> json with article #1
- DELETE  base/articles/1 -> delete article #1, must be authorized
- POST    base/articles -> create new article, must be authorized
- POST    base/user/register -> register new user
- POST    base/user/auth -> authorization existing user


# REST API for solving quadratic equations
rest api to get number of roots of quadratic equation ax*x + bx + c = 0
1. routes:
- base: localhost:8080/api/v1
- POST  base/grab -> set equition params in body
{
  "A": 10,
  "B": 20,
  "C": 30
}
- GET   base/solve -> get number of roots

# Test First Development API example
factorial API
1. how can I check that it works?
2. define boundary conditions
3. start with tests - main_test.go
4. add local hash to reduce calculation
5. routes:
- base: localhost:8080
- GET base/factorial?num=7 -> request for calc factorial of "7" -> get factorial answer in response

# Gin & GORM API example
