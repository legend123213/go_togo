RESTful API For Tasks
# Task Management using Gin

This is a simple API documentation for a Task Manager RESTful API built using Gin (Go Framework).

## Tasks
CRUD Operations

### GET
**GetTasks**
- Endpoint: `http://localhost:8000/api/v1/tasks`
- Description: This is a GET request used to retrieve tasks from the API. There is no request body for a GET request, but you can use path variables to specify the resource you want data on (e.g., `/4` to get task with ID 4).
- Response: A successful GET response will have a 200 OK status and may include HTML web content or JSON data.

### POST
**AddTask**
- Endpoint: `http://localhost:8000/api/v1/task`
- Description: This is a POST request used to add a new task to the API. The request body should contain JSON data representing the task.
- Response: A successful POST request typically returns a 200 OK or 201 Created response code.

**Request Body Example:**
```json
{
   "Title":"Task Manager: Adding Endpoint handle Create",
   "Description":"Building Endpoint that handle Creating Task",
   "Due_date":"01-02-2024",
   "Status":"Done"
}
```

### DELETE
**DeleteTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a DELETE request used to delete a task that was previously created via a POST request. You should include the task ID in the URL.
- Response: A successful DELETE request typically returns a 200 OK, 202 Accepted, or 204 No Content response code.

### PUT
**UpdateTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a PUT request used to update an existing task. You should include the task ID in the URL.
- Response: A successful PUT request typically returns a 200 OK, 201 Created, or 204 No Content response code.

**Request Body Example:**
```json
{
   "_id":"66adee6f1f41b59e0fe03a43",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpointss that handle Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

### GET
**GetTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a GET request used to retrieve a specific task by its ID. You should include the task ID in the URL.
- Response: A successful GET request will have a 200 OK status and may include HTML web content or JSON data.

