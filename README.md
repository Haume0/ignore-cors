
## Welcome to my project 'Ignore CORS'!
In cases where you are not authorized to access the api that gives CORS error,
your local development will be able to continue without interruption.
#### How to build:

> `go build .` 
> For building app for your operating system.

> `go run .`
> For starting app without building.

# How to use?
First, you must start executable file (icors.exe for windows) you can build anywhere with using source code.
then change fetch url to `http://localhost:8080/get?url=https://api.example.com/news`.
You can write any url you want after `?url=`.
and It's DONE!
