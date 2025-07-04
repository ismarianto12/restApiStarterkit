# REST API Documentation

## Overview
This application is a REST API built with Go (Golang). It provides endpoints for managing resources and supports standard HTTP methods such as GET, POST, PUT, and DELETE.

## Requirements
- Go 1.18 or higher
- Any REST client (e.g., Postman, curl)

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/ismarianto/golangRest.git
    ```
2. Navigate to the project directory:
    ```bash
    cd golangRest
    ```
3. Build and run the application:
    ```bash
    go run main.go
    ```

## Endpoints

### GET /resources
Retrieve a list of all resources.

#### Response
- **200 OK**: Returns an array of resources.
- **500 Internal Server Error**: If there is an issue with the server.

---

### POST /resources
Create a new resource.

#### Request Body
```json
{
  "name": "Resource Name",
  "description": "Resource Description"
}
```

#### Response
- **201 Created**: Returns the created resource.
- **400 Bad Request**: If the request body is invalid.

---

### PUT /resources/{id}
Update an existing resource.

#### Request Body
```json
{
  "name": "Updated Name",
  "description": "Updated Description"
}
```

#### Response
- **200 OK**: Returns the updated resource.
- **404 Not Found**: If the resource does not exist.

---

### DELETE /resources/{id}
Delete a resource.

#### Response
- **200 OK**: If the resource is successfully deleted.
- **404 Not Found**: If the resource does not exist.

## Error Handling
All errors are returned in the following format:
```json
{
  "error": "Error message"
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contact
For questions or support, please contact [ismarianto@example.com](mailto:ismarianto@example.com).