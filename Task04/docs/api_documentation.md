markdown
# 🚀 Task Management API Documentation

This is a simple API documentation for the Task Manager RESTful API built using the Gin framework (Go).

## Base URL

`http://localhost:8000/api/v1`

## Endpoints

### Tasks

#### Get Tasks

- **URL:** `/task`
- **Method:** `GET`
- **Description:** This endpoint is used to get tasks. A successful GET response will have a `200 OK` status and should include some kind of response body (e.g., JSON data).

##### Request
- **Headers:** None
- **Path Variables:** None

##### Response
- **Status Code:** `200 OK`
- **Body:** JSON data representing the tasks.

##### Example
```sh
curl -X GET "http://localhost:8000/api/v1/task"
```

#### Add Task

- **URL:** `/task`
- **Method:** `POST`
- **Description:** This endpoint is used to submit a task to the API via the request body. This request submits JSON data, and the data is reflected in the response.

##### Request
- **Headers:** None
- **Body:**
```json
{
   "Title": "Task Manager: Adding Endpoint handle Create",
   "Description": "Building Endpoint that handle Creating Task",
   "Due_date": "01-02-2024",
   "Status": "Done"
}
```

##### Response
- **Status Code:** `200 OK` or `201 Created`
- **Body:** JSON data representing the created task.

##### Example
```sh
curl -X POST "http://localhost:8000/api/v1/task" -H "Content-Type: application/json" -d '{
   "Title": "Task Manager: Adding Endpoint handle Create",
   "Description": "Building Endpoint that handle Creating Task",
   "Due_date": "01-02-2024",
   "Status": "Done"
}'
```

#### Delete Task

- **URL:** `/task/:id`
- **Method:** `DELETE`
- **Description:** This endpoint is used to delete a task that was previously created via a POST request. You typically identify the entity being updated by including an identifier in the URL.

##### Request
- **Headers:** None
- **Path Variables:**
   - `id` (integer): The ID of the task to delete.

##### Response
- **Status Code:** `200 OK`, `202 Accepted`, or `204 No Content`
- **Body:** None

##### Example
```sh
curl -X DELETE "http://localhost:8000/api/v1/task/2"
```

#### Update Task

- **URL:** `/task/:id`
- **Method:** `PUT`
- **Description:** This endpoint is used to overwrite an existing piece of data. For instance, after you create an entity with a POST request, you may want to modify that later. You can do that using a PUT request. You typically identify the entity being updated by including an identifier in the URL.

##### Request
- **Headers:** None
- **Path Variables:**
   - `id` (integer): The ID of the task to update.
- **Body:**
```json
{
   "ID": 5,
   "Title": "Task Manager RESTful API",
   "Description": "Building simple CRUD operations to handle task management",
   "Due_date": "01-01-2024",
   "Status": "Done"
}
```

##### Response
- **Status Code:** `200 OK`, `201 Created`, or `204 No Content`
- **Body:** JSON data representing the updated task.

##### Example
```sh
curl -X PUT "http://localhost:8000/api/v1/task/5" -H "Content-Type: application/json" -d '{
   "ID": 5,
   "Title": "Task Manager RESTful API",
   "Description": "Building simple CRUD operations to handle task management",
   "Due_date": "01-01-2024",
   "Status": "Done"
}'
```

### Variables

- `base_url`: The base URL for the API. Default is `http://localhost:8000/api/v1`.

### Testing

#### Pre-request Script

None specified.

#### Tests

##### GetTasks:

```javascript
pm.test("Status code is 200", function () {
      pm.response.to.have.status(200);
});
```

##### AddTask:

```javascript
pm.test("Successful POST request", function () {
      pm.expect(pm.response.code).to.be.oneOf([200, 201]);
});
```

##### DeleteTask:

```javascript
pm.test("Successful DELETE request", function () {
      pm.expect(pm.response.code).to.be.oneOf([200, 202, 204]);
});
```

##### UpdateTask:

```javascript
pm.test("Successful PUT request", function () {
      pm.expect(pm.response.code).to.be.oneOf([200, 201, 204]);
});
```

```
```
