# CatAPI Project

## Introduction
This project is a RESTful API for managing spy cats and missions.

## Prerequisites
- Go 1.16+
- MySQL (or any compatible database)
- Postman (optional, for testing endpoints)

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd CatApi
```

2. Set up environment variables:
Create a .env file in the root of the project or alongside the executable with the following content:
```bash
DATABASE_URL="host=localhost user=user password=password dbname=spycats port=5432 sslmode=disable"
```

3. Build and start the application:
```bash
go build cmd/catapi/main.go
```
This will start the application on localhost:8080 (default).

4. Explore the API documentation:
Open your browser and go to [Swagger UI](http://localhost:8080/swagger/index.html#/) to explore and test the API endpoints.

5. Testing with Postman:
Import the provided Postman collection (CatAPI.postman_collection.json) to test the endpoints directly.

## Additional Notes
- Auto Migration: The application uses auto migration to manage database schema changes. Ensure the .env file is correctly configured for database access.
- Swagger Documentation: Swagger files (swagger.json and swagger.yaml) are generated automatically. You can access the Swagger UI at [Swagger UI](http://localhost:8080/swagger/index.html#/) after starting the application.