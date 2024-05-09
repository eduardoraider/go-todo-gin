### Go TODO with GIN

This project implements a TODO API server in Go using Gin, a popular HTTP web framework.

#### Installation

```bash
go get -u github.com/gin-gonic/gin
```

#### Usage

To run the server, execute the following command:

```bash
go run cmd/api/main.go
```

#### Endpoints

- **GET /todos**: Retrieve all todo items.
- **GET /todos/:id**: Retrieve a todo item by its ID.
- **POST /todos**: Add a new todo item.
- **PUT /todos/:id**: Update an existing todo item.
- **DELETE /todos/:id**: Delete an existing todo item.

#### Using CLI (curl)

- **Get all todos:**
```bash
curl http://localhost:8080/todos
```

- **Get a todo by ID:**
```bash
curl http://localhost:8080/todos/1
```

- **Add a new todo:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "title": "New Todo", "done": false}' http://localhost:8080/todos
```

- **Update a todo:**
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Updated Todo", "done": true}' http://localhost:8080/todos/1
```

- **Delete a todo:**
```bash
curl -X DELETE http://localhost:8080/todos/1
```

#### Using Postman

- **Get all todos:**
    - **Method:** GET
    - **URL:** http://localhost:8080/todos

- **Get a todo by ID:**
    - **Method:** GET
    - **URL:** http://localhost:8080/todos/1
    - **Replace 1 with the ID of the todo you want to retrieve.**

- **Add a new todo:**
    - **Method:** POST
    - **URL:** http://localhost:8080/todos
    - **Headers:**
        - Key: Content-Type
        - Value: application/json
    - **Body (raw JSON):**
      ```json
      {
          "id": 1,
          "title": "New Todo",
          "done": false
      }
      ```

- **Update a todo:**
    - **Method:** PUT
    - **URL:** http://localhost:8080/todos/1
    - **Replace 1 with the ID of the todo you want to update.**
    - **Headers:**
        - Key: Content-Type
        - Value: application/json
    - **Body (raw JSON):**
      ```json
      {
          "id": 1,
          "title": "Updated Todo",
          "done": true
      }
      ```

- **Delete a todo:**
    - **Method:** DELETE
    - **URL:** http://localhost:8080/todos/1
    - **Replace 1 with the ID of the todo you want to delete.**


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer