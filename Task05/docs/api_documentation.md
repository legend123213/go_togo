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
> Response Examples

> Success

```json
{
  "tasks": [
    {
      "Title": "Task Manager: Adding Endpoint handle Create",
      "Description": "Building Endpoint that handle Creating Task",
      "Due_date": "01-02-2024",
      "Status": "Done"
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» tasks|[object]|true|none||none|
|»» Title|string|false|none||none|
|»» Description|string|false|none||none|
|»» Due_date|string|false|none||none|
|»» Status|string|false|none||none|

### POST
**AddTask**
- Endpoint: `http://localhost:8000/api/v1/task`
- Description: This is a POST request used to add a new task to the API. The request body should contain JSON data representing the task.
- Response: A successful POST request typically returns a 200 OK or 201 Created response code.
> Body Parameters

```json
{
  "Title": "Task Manager: Adding Endpoint handle Create",
  "Description": "Building Endpoint that handle Creating Task",
  "Due_date": "01-02-2024",
  "Status": "Done"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Success

```json
{
  "task": {
    "Title": "Task Manager: Adding Endpoint handle Create",
    "Description": "Building Endpoint that handle Creating Task",
    "Due_date": "01-02-2024",
    "Status": "Done"
  },
  "message": "Task added successfully"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» task|object|true|none||none|
|»» Title|string|true|none||none|
|»» Description|string|true|none||none|
|»» Due_date|string|true|none||none|
|»» Status|string|true|none||none|
|» message|string|true|none||none|

### DELETE
**DeleteTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a DELETE request used to delete a task that was previously created via a POST request. You should include the task ID in the URL.
- Response: A successful DELETE request typically returns a 200 OK, 202 Accepted, or 204 No Content response code.
### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

### PUT
**UpdateTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a PUT request used to update an existing task. You should include the task ID in the URL.
- Response: A successful PUT request typically returns a 200 OK, 201 Created, or 204 No Content response code.
> Body Parameters

```
|-
{
   "_id":"66adee6f1f41b59e0fe03a43",
   "title": "Task Manager: Adding Endpoint handle Create",
   "description": "Building Endpointss that handle Creating Task",
   "due_date": "2006-01-02T15:04:05Z",
   "status": "Done"
}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|
|body|body|string| no |none|

> Response Examples

> Success

```json
{
  "edited_task": {
    "_id":"66adee6f1f41b59e0fe03a43",
    "Title": "Task Manager: Adding Endpoint handle Create",
    "Description": "Building Endpoint that handle Creating Task",
    "Due_date": "01-02-2024",
    "Status": "Done"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» edited_task|object|true|none||none|
|»» Title|string|true|none||none|
|»» Description|string|true|none||none|
|»» Due_date|string|true|none||none|
|»» Status|string|true|none||none|


### GET
**GetTask**
- Endpoint: `http://localhost:8000/api/v1/task/:id`
- Description: This is a GET request used to retrieve a specific task by its ID. You should include the task ID in the URL.
- Response: A successful GET request will have a 200 OK status and may include HTML web content or JSON data.

GET /api/v1/task/{id}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|id|path|string| yes |none|

> Response Examples

> Success

```json
{
  "task": {
    "Title": "Task Manager: Adding Endpoint handle Create",
    "Description": "Building Endpoint that handle Creating Task",
    "Due_date": "01-02-2024",
    "Status": "Done"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» task|object|true|none||none|
|»» Title|string|true|none||none|
|»» Description|string|true|none||none|
|»» Due_date|string|true|none||none|
|»» Status|string|true|none||none|
