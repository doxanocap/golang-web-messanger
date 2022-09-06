import React from "react";
import "./index.css"

const MainPageContainer = ({CurrentUsersName}) => {
    let text;
    if (typeof(CurrentUsersName) === "string" && CurrentUsersName.length > 0) {
        text = (
            <h2 className="MainPageContainer-upperLarge-h1">Hello dear user: {CurrentUsersName}!</h2>
        )
    } else {
        text = (
            <h2 className="MainPageContainer-upperLarge-h1">Hello dear user!</h2>
        )
    }
    return (
        <div className="MainPageContainer">
            <div className="MainPageContainer-upperLarge">
                {text}
                
                <p className="MainPageContainer-upperLarge-p">This is my personal online Web Project. The reason I built is firstly: to connect React with Golang, work with websockets using gorillaMux/websocket, use one of the most powerfull router frameworks gonic/Gin, and to deploy my project to Docker/Heroku.</p>
            </div>
            <div className="MainPageContainer-parentOfSmall">
                <div className="MainPageContainer-bottomDark">
                    <p className="MainPageContainer-bottomDark-p">Backend:</p>
                    <p className="MainPageContainer-bottomDark-p">Language: Golang</p>
                    <p className="MainPageContainer-bottomDark-p">Frameworks: Gin Web Framework, Gorilla WebSocket, JWT v4, bcrypt, lib/pq, net/http</p> 
                    <p className="MainPageContainer-bottomDark-p">Brief Description: Developed RESTful API to create online WebChat application using WebSockets and Goroutine channels to pass messages and information about user. Activated CORS to estabilish connection between frontend and backend. Also, connected DataBase to store all messages, registered user information, and information about all online users. Applied autentification system using cookies and JWT, bcrypt to hash password of the user.</p>
                </div>
                <div className="MainPageContainer-bottomLight">
                    <p className="MainPageContainer-bottomLight-p">Frontend:</p>
                    <p className="MainPageContainer-bottomLight-p">Languages: JavaScript, HTML, CSS</p>
                    <p className="MainPageContainer-bottomLight-p">Frameworks: React</p>
                    <p className="MainPageContainer-bottomLight-p">Frameworks: Handled GET, POST requests from the backend using Fetch and async/await funtions, handled messages from Websocket connection. Designed all pages by myself without using bootstrap, MUI, and others.</p>
                </div>
            </div>
        </div>
    )
};

export default MainPageContainer;

