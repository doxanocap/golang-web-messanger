# https://webchat-doxa.herokuapp.com/
# Online Web Messanger

This is my personal online Web Project. The reason I built is firstly: to connect React with Golang, work with websockets using gorillaMux/websocket, use one of the most powerfull router frameworks gonic/Gin, and to deploy my project to Docker/Heroku.

# Backend:

## Language: Golang

### Frameworks: Gin Web Framework, Gorilla WebSocket, JWT v4, bcrypt, lib/pq, net/http

Brief Description: Developed RESTful API to create online WebChat application using WebSockets and Goroutine channels to pass messages and information about user. Activated CORS to estabilish connection between frontend and backend. Also, connected DataBase to store all messages, registered user information, and information about all online users. Applied autentification system using cookies and JWT, bcrypt to hash password of the user.

# Frontend:

## Languages: JavaScript, HTML, CSS

### Frameworks: React

Brief Description: Handled GET, POST requests from the backend using Fetch and async/await funtions, handled messages from Websocket connection. Designed all pages by myself without using bootstrap, MUI, and others.