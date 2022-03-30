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
5. add authentication by JWT for all operations
6. routes:
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	// a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("Delete")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")

	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")

# REST API for solving quadratic equations
rest api to get number of roots of quadratic equation
1. routes:
- base: localhost:8080/api/v1
- POST  base/grab -> set equition params in body
{
  "A": 10,
  "B": 20,
  "C": 30
}
- GET   base/solve -> get number of roots

