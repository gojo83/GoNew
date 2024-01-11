Get all books:
curl http://localhost:8080/books

Create a new book:
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Book", "author":"New Author"}' http://localhost:8080/books

Update a book by ID:
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Title", "author":"Updated Author"}' http://localhost:8080/books/{id}

Delete a book by ID:
curl -X DELETE http://localhost:8080/books/{id}
