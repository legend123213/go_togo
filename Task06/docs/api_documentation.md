
# 🚀 Task Management using Gin

This is a simple API documentation for the Task Manager RESTful API built using Gin (Go Framework).

## Authentication

All endpoints, except for login and signup, require authentication via Bearer Token.

- Type: Bearer Token
- Token: `<token>`

## Endpoints

### User

#### POST /api/v1/login

This endpoint is used for user login. The username and password should be provided in the request body as JSON data.

Request Body:

```json
{
   "username": "abel@gmail.com",
   "password": "abelabel"
}
```

Response:

- Status Code: 202
- Body:

```json
{
   "token": "your_jwt_token_here"
}
```

Access: Public

#### POST /api/v1/signup

This endpoint is used for user signup. The new user's username and password should be provided in the request body as JSON data. The new user will have the admin role.

Request Body:

```json
{
   "username": "abel wendmu",
   "password": "abelabel"
}
```

Response:

- Status Code: 202
- Body:

```json
{
   "token": "your_jwt_token_here"
}
```

Access: Public

#### GET /api/v1/user/:id

This endpoint is used to retrieve a specific user. You should include the user ID in the URL.

Response:

- Status Code: 200
- Body:

```json
{
   "id": "66adee6f1f41b59e0fe03a43",
   "username": "abel wendmu",
   "isAdmin": true
}
```

Access: Authenticated Users

#### GET /api/v1/users

This endpoint is used to retrieve all users in the system. Only admin users can perform this action.

Response:

- Status Code: 200
- Body:

```json
[
   {
      "id": "66adee6f1f41b59e0fe03a43",
      "username": "abel wendmu",
      "isAdmin": true
   },
   {
      "id": "66adee6f1f41b59e0fe03a44",
      "username": "john doe",
      "isAdmin": false
   }
]
```

Access: Admin Users

#### PATCH /api/v1/promote/:id

This endpoint is used to promote a user to admin. Only admin users can perform this action. You should include the user ID in the URL.

Response:

- Status Code: 202
- Body:

```json
{
   "message": "User promoted to admin"
}
```

Access: Admin Users

#### DELETE /api/v1/user/:id

This endpoint is used to delete a user. Only admin users can perform this action. You should include the user ID in the URL.

Response:

- Status Code: 204

Access: Admin Users

### Task

#### GET /api/v1/tasks

This endpoint is used to retrieve all tasks.

Response:

- Status Code: 200
- Body:

```json
[
   {
      "id": "66adee6f1f41b59e0fe03a45",
      "title": "Task Manager: Adding Endpoint handle Create",
      "description": "Building Endpoint that handles Creating Task",
      "due_date": "2006-01-02T15:04:05Z",
      "status": "Done"
   },
   {
      "id": "66adee6f1f41b59e0fe03a46",
      "title": "Task Manager: Adding Endpoint handle Update",
      "description": "Building Endpoint that handles Updating Task",
      "due_date": "2006-01-02T15:04:05Z",
      "status": "In Progress"
   }
]
```

Access: Authenticated Users

#### GET /api/v1/task/:id

This endpoint is used to retrieve a specific task. You should include the task ID in the URL.

Response:

- Status Code: 200
- Body:

```json
{
   "id": "66adee6f1f41b59e0fe03a45",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpoint that handles Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

Access: Authenticated Users

#### POST /api/v1/task

This endpoint is used to add a new task. The task details should be provided in the request body as JSON data.

Request Body:

```json
{
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpoint that handles Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

Response:

- Status Code: 202
- Body:

```json
{
   "id": "66adee6f1f41b59e0fe03a45",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpoint that handles Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

Access: Authenticated Users

#### PUT /api/v1/task/:id

This endpoint is used to update an existing task. You should include the task ID in the URL and provide the updated task details in the request body as JSON data.

Request Body:

```json
{
   "id": "66adee6f1f41b59e0fe03a45",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpoints that handle Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

Response:

- Status Code: 202
- Body:

```json
{
   "id": "66adee6f1f41b59e0fe03a45",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpoints that handle Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}
```

Access: Authenticated Users

#### DELETE /api/v1/task/:id

This endpoint is used to delete a task that was previously created. You should include the task ID in the URL.

Response:

- Status Code: 204

Access: Authenticated Users