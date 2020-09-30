# Overview

The project is developed by Golang based on the echo framework, and the database uses SQLite GORM. It is a pure back-end project that integrates swagger and provides APIs.

# Usage
**Operating environment requirements**
 - Mac/Linux
 - Docker

**Build project**

    git clone git@github.com:bobyeoh/todo-backend.git
    cd todo-backend
    sh build.sh
    
default url: http://localhost:8888/api/swagger/index.html

**Test**
When the container is running

    sh test.sh

**Preview**

https://todo.001.gs/api/swagger/index.html


# Dependency

 1. Echo framework
 2. GORM
 3. go-playground/validator/v10
 4. Swagger
 5. Gomail.v2
 6. SQLite
 7. Testify

# Design Considerations

 1. For the scalability of the system and future deployment to the database cluster, I split the database operation into repositories (read) and services (write), so that the database can be easily separated from reading and writing.
 2. Whether it is user information, columns or tasks, all are stored in database tables. Based on the existing system, more functions can be easily implement: for example, you can add, edit and delete columns.
 3. Originally, I implemented the JWT system through echo middleware, cookie and Redis, but because of the need to simplify the deployment process, I refactored the code of the JWT part to enable echo middleware, cookie and sqlite to achieve the authentication function, based on this In the future, a multi-role authority system can be implemented very easily.
 4. In order to prevent hackers from cracking the login password through exhaustive methods, I added a logic: once the password is entered incorrectly 3 times, the account will be locked for 10 minutes. This can effectively reduce the exhaustive efficiency.
 
# Project structure

    ├── README.md
	├── app
	│   ├── controllers // Business logic and APIs unit testing
	│   ├── database // Database configuration
	│   ├── models // ORM models
	│   ├── permission // JWT middleware
	│   ├── repositories // Read database operations
	│   ├── requests // Incoming parameter model definition and data validation
	│   ├── responses // Response model definition for APIs
	│   ├── routes // router
	│   ├── server.go // Create instance
	│   ├── services // Write database operations
	│   ├── utils // Static functions
	│   └── validation // Request model data verification
	├── build.sh // Build project
	├── dev.env // Environment variable for development
	├── dockerfile // docker file
	├── docs // swagger folder
	│   ├── docs.go
	│   ├── swagger.json
	│   └── swagger.yaml
	├── go.mod
	├── go.sum
	├── main.go // main file
	├── start.sh // Start container
	├── stop.sh // Stop container
	└── test.sh // unit test

