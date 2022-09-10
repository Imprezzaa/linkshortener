# linkshortener

:construction: Work in progress :construction_worker::construction:

## Description
The intent of the project was to create a basic link shortening service written in Go but the project kept growing in scope as I learned more about different technologies and went from learning about one concept to another. The code ended up being messy and hard to follow since it was a mixture of reading documentation and following tutorials by different authors. I decided to take a step back and focus on learning about the concepts I thought were interesting and could fit within the project. 

I ended up reading both "Let's Go" and "Let's Go Further" by Alex Edwards and they provide a great blueprint for writing web applications from the ground up and hits on a ton of important topics in a cohesive way. The bulk of the code will be from the Greenlight project from the latter with modifications to make it work for the purposes of this project. Figure I could learn from the book(s) while also tinkering around with PostgreSQL, HTTPS, self-signed certs, TLS configs etc. 


I decided to restart the project and re-focus efforts into building a solid project with a defined scope. 

### Current Functionality
In it's current state it is just a basic HTTPS server that can connect to a PosgreSQL via enironment variable or CLI flag. 

#### Why use these specific technologies?
- Go: I chose to write the project in Go since I'm interested in building backend infrastructure and Go is a simple but powerful language specifically designed for working on backend systems
- httprouter: I wanted a lighter HTTP router that extended the base net/http package without needing to pull in a lot of extra code.

## Project Goals
- Build a functional link shortening service backend with the following components
    - HTTPS server (based on a self-signed certificate for now)
    - Route handlers providing multiple endpoints and method based routing that allow user creation, login/logout, shortened link creation, patching, deletion and viewing
        - This will be provided by an underlying SQL database and handlers will use SQL queries to provide it's functionality
    - User verification
        - Using bcrypt to hash and verify users have sent the correct password
    - Routing middleware
        - Authorization middleware
        - Leveled logging middleware that sends logs to stdout
    - User and Link models used to pass data to a SQL server.


### Short term plans
- HTTPS server w/ routes
- 

### Long term plans/wants
- Add a front-end with usable forms based on http templates





TODO: Update with more resources as the project progresses
### Resources Used:
- The original idea for the project was based on Gophercises
    - https://www.calhoun.io/
    - I went off the rails a bit, but I learned quite a bit from the original video and got some experience with BoltDB
- The basic layout of this version of the project was based on this article and it's a good resource to learn about some basics of the Mongo DB Go driver and gives examples of using different MongoDB functions
    - https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m
- Along with the MongoDB docs themselves 
    - https://www.mongodb.com/docs/drivers/go/current/usage-examples/#std-label-golang-usage-examples
    - https://www.mongodb.com/docs/drivers/go/current/quick-start/#std-label-golang-quickstart
- And the excellent mongoDB blog posts with guides about using the driver
    -   https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
- For learning about bcrypt and jwt authentication 
    - https://codewithmukesh.com/blog/jwt-authentication-in-golang/#Testing_Golang_API_with_VSCode_REST_Client
