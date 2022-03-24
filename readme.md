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
- base: localhost:8080/api/v1/
- GET    base/books -> json with all books
- GET    base/books/1 -> json with book #1
- POST   base/book/ -> add new book
- PUT    base/books/1 -> update existing book
- DELETE base/books/1 -> delete existing book
4. Collection of Tests in Thunder Client at booksApi.json