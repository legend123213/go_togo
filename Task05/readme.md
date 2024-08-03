Enhancing Task Management API with Persistent Data Storage using MongoDB and Mongo Go Driver
## Objective
The objective of this task is to extend the existing Task Management API with persistent data storage using MongoDB and the Mongo Go Driver. This enhancement will replace the in-memory database with MongoDB to provide data persistence across API restarts.

## Prerequisites
Before running the project, ensure you have the following installed:

- Go (version 1.16 or higher)
- MongoDB (either locally or a cloud instance)

## Steps to Run
1. Clone the Repository

```bash
git clone https://github.com/legend123213/go_togo.git
cd go_togo
cd Task05
```

2. Install Dependencies
Ensure that you have the necessary dependencies installed using go mod:

```bash
go mod tidy
```

3. Set Up MongoDB

- Set up a MongoDB instance either locally or using a cloud service provider.
- Ensure MongoDB is running and accessible.

4. Configure MongoDB Connection

- Update the MongoDB connection parameters in main.go or use environment variables/config files.

5. Build and Run

```bash
go build -o task_manager
./task_manager
```

This will build the executable and start the server. By default, the server will run on http://localhost:8080.

## Test API Endpoints
Use tools like Postman to test the API endpoints (e.g., create, read, update, delete tasks).

## Folder Structure
Follow the following folder structure for this task:

```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```
