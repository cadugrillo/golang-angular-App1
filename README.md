<h1>What is Golang-Angular-App1?</h1>

This Web App is part of a series of small projects being created during my quest to learn new skills in the modern software development world. It is written mostly using:

- Golang (or Go) programming language. You can find more info at https://go.dev/.
- Angular framework. You can find more info at https://angular.io/.


<h3>This app works in frontend / backend architecture where:</h3>

**golang-app1 container**   - contains an HTTP API endpoint written in Go using Gin framework acting as backend of a Todo List.
(the backend app can be found at https://hub.docker.com/repository/docker/cadugrillo/golang-app1).  

**angular-app1 container** - contains the webpage based on Angular framework acting as the frontend of a Todo List.  
(the frontend app can be found at https://hub.docker.com/repository/docker/cadugrillo/angular-app1).


<h3>You can access a running version of the app at:</h3>

**http://app1.cadugrillo.com**


<h1>Where can I find the source code?</h1>

You can fork this git repository.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/cadugrillo/golang-angular-App1.git)

<h1>How to deploy ?</h1>

The easiest way to deploy is using docker-compose.yml file found in this git repository.

1. Copy the file to your desired folder
2.  From the root of the folder run in a terminal "docker-compose up -d"
3. From your local browser navigate to http://localhost:80.

<h1>Description of the HTTP API endpoint</h1>

The service listens to port 4300 and handles all methods (**GET**, **POST**, **PUT**, and **DELETE**) available at the API endpoint (**/todo**). 

**GET -** returns a list with all tasks.  
**POST -** adds a new task to the list and returns all tasks.  
**DELETE -** deletes a task based on its id.  
**PUT -** sets the task flag "complete" and returns all tasks.
