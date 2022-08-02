# linkshortener

:warning: Very early WIP :warning:

The goal is to have a basic, but functional URL shortener with a built-in Database/keystore(currently bboltDB) and a basic frontend that allows user creation, log-in/logout and retrieval of a users created URLs.



## TODOs
- test different formats of storing data. 
- Revise/restructure DB 
- encode/decode functions for storing information as []byte and retrieving information and turning it into something returnable to the user
- user sign-up service - super basic, just username/password(good chance to learn about password hashing)
- website with basic functionality to present a GUI - using html templates
- create routes and route handlers
- setup http services
- write tests


## Finished
- First DB draft

### stretch goals
- https server with self-signed cert?
