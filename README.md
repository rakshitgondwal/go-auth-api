Authorization and Authentication API in Golang
This is a backend API service in Golang that handles authorization and authentication for a web app where users in an organization can sign in and list all other users in their organization. The API follows REST API conventions and covers the following functionalities:

User Login
User Logout
Admin User adds a new User account (by providing the username & password)
Admin User deletes an existing User account from their organization
List all Users in their organization
Setup
Prerequisites
Golang should be installed on your system. You can download and install it from the official Golang website.
MongoDB should be installed on your system. You can download and install it from the official MongoDB website.
Installation
Clone the repository: git clone https://github.com/your-username/auth-api.git
Navigate to the project directory: cd auth-api
Install the dependencies: go mod download
Configuration
The configuration file is located in the configs directory. The following configuration options are available:

server.port - The port on which the server should run (default: 8080)
db.uri - The URI of the MongoDB database (default: mongodb://localhost:27017)
db.name - The name of the MongoDB database (default: auth)
jwt.secret - The secret used to sign the JWT tokens (default: my-secret)
jwt.access_token_expiry - The expiry time of the access token in minutes (default: 60)
jwt.refresh_token_expiry - The expiry time of the refresh token in minutes (default: 1440)

Installation
Prerequisites
Go 1.16 or later
MongoDB 4.4 or later
Postman or any API development environment
Clone the repository
shell
Copy code
$ git clone https://github.com/<username>/<repository>.git
$ cd <repository>
Install dependencies
go
Copy code
$ go mod tidy
Setup environment variables
Create a .env file in the root directory and add the following environment variables:

makefile
Copy code
MONGODB_URI=<mongodb_uri>
MONGODB_DB=<mongodb_database>
JWT_SECRET=<jwt_secret_key>
MONGODB_URI: the MongoDB URI to connect to the database.
MONGODB_DB: the name of the MongoDB database.
JWT_SECRET: the secret key used to sign the JWT tokens.
Run the API
go
Copy code
$ go run main.go
The API will be available at http://localhost:8080.


Design Decisions
Framework
I chose to use the Echo framework for this project due to its simplicity, performance, and ease of use. Echo is a lightweight web framework that provides a minimalistic approach to building web applications and APIs.

Database
I chose MongoDB as the database for this project due to its flexibility, scalability, and ease of use. MongoDB is a document-oriented NoSQL database that allows for easy data modeling and flexible data structures.

ORM
I decided not to use an ORM for this project due to the simplicity of the data model and the limited number of database operations required by the API. Instead, I used the official MongoDB driver for Go to interact with the database.

JWT
For JWT token generation, I used the golang-jwt library, which provides a simple and easy-to-use interface for generating and verifying JWT tokens.

API Design
I followed REST API conventions for designing the API endpoints and used HTTP methods and status codes to represent the different actions and outcomes of the API requests. I also used middleware to handle authentication and authorization and to enforce input validation and error handling.