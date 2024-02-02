# Go CRUD Server with Gorilla Mux

This is a simple Go server implementing CRUD (Create, Read, Update, Delete) operations using the Gorilla Mux router and the `net/http` package.

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/tyler1836/crud-go.git
   ```
## Change into the project directory:

```bash
Copy code
cd your-repo
Build and run the server:
```

- The server should now be running on http://localhost:8080.

## Usage
### Endpoints
- ``` GET /movies: Get a list of all movies.```
- ``` GET /movies/{id}: Get details of a specific item.```
- ``` POST /movies: Create a new item.```
- ``` PUT /movies/{id}: Update an existing item.```
- ``` DELETE /movies/{id}: Delete an item.```
## Example Requests
- Get all movies
```curl http://localhost:8080/movies```
- Get a specific item
```curl http://localhost:8080/movies/{id}```
- Create a new item
```curl -X POST -H "Content-Type: application/json" -d '{"name":"New Item","description":"Description of the new item"}' http://localhost:8080/movies```
- Update an item
```curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Item","description":"Updated description"}' http://localhost:8080/movies/{id}```
- Delete an item
```curl -X DELETE http://localhost:8080/movies/{id}```
## Dependencies
* Gorilla Mux: A powerful URL router and dispatcher for Go.
## Contributing
* Feel free to contribute by opening issues or submitting pull requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.