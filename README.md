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
- Implement an timestamp instead of using unix time
- Document the API and provide examples of API calls

### Long term plans/wants
- Prep the project to be dockerized and allow important variables to be pulled from the .env file
- Implement an in memory counter that bulk updates DB documents every x hours
- look into user authentication and protected routes