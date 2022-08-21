# linkshortener

:construction: Work in progress :construction_worker::construction:

## Description
This is a link shortening API written in Go using the Gin web framework and is written to support a MongoDB backend. It's still in early development so is missing some features and documentation.

### Curent Functionality
User Logic - Create, Get, Edit, Delete, GetAll
Link Logic:
    - Create: create takes in a username and a URL and creates a shortened URL that redirects to the supplied URL
    - GetLink: getlink takes in a shortened URL parameter, queries the database and redirects the user to the saved URL
    - GetUserLinks: getuserlinks takes in a username, queries the link document collection for all links created by the user and returns the shortened and original URL

#### Why use these specific technologies?
- Go: I chose to write the project in Go since I'm interested in building backend infrastructure and Go is a simple but powerful language specifically designed for working on backend systems
- Gin: It's a popular web framework in active development and the built in middleware makes troubleshooting much easier
- MongoDB: I was originally using bboltdb but wanted to use a remote backend and learn more about NoSQL databases. 
- godotenv: It's helpful to have sensitive data in a seperate file that can be ignored when pushing code to github. 

## Project Goals
- #TODO


### Short term plans
- Complete and test the link controller logic
- Clean up code(inconsistent variable names, struct fields/tags, fix bugs)
- Add tests to existing packages that won't see major changes
- ~~Implement a timestamp instead of using unix time~~ - time is stored as a MongoDB Datetime primitive which returns as a formatted date/time
- Document the API and provide examples of API calls
- look into user authentication and protected routes

### Long term plans/wants
- Prep the project to be dockerized and allow important variables to be pulled from the .env file
- Implement an in memory counter that bulk updates DB documents every x hours





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
