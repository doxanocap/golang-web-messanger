import React from "react";
import "./index.css"

const MainPageContainer = ({CurrentUsersName}) => {
    if (typeof(CurrentUsersName) === "string" && CurrentUsersName.length > 0) {
        return (
            <div className="MainPageContainer">
                <div className="MainPageContainer-upperLarge">
                    <h2 className="MainPageContainer-upperLarge-h1">Hello dear user: {CurrentUsersName}!</h2>
                    <p className="MainPageContainer-upperLarge-p">This is my personal online Web Project. The reason I built is firstly: to connect React with Golang, work with websockets using gorillaMux/websocket, use one of the most powerfull router frameworks gonic/Gin, and to deploy my project to Docker/Heroku.</p>
                </div>
                <div className="MainPageContainer-parentOfSmall">
                    <div className="MainPageContainer-bottomDark">
                        <p className="MainPageContainer-bottomDark-p">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                    </div>
                    <div className="MainPageContainer-bottomLight">
                        <p className="MainPageContainer-bottomLight-p">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                    </div>
                </div>
            </div>
        )
    } else {
        return (
            <div className="MainPageContainer">
                <div className="MainPageContainer-upperLarge">
                    <h2 className="MainPageContainer-upperLarge-h1">Hello dear user!</h2>
                    <p className="MainPageContainer-upperLarge-p">This is my personal online Web Project. The reason I built is firstly: to connect React with Golang, work with websockets using gorillaMux/websocket, use one of the most powerfull router frameworks gonic/Gin, and to deploy my project to Docker/Heroku.</p>
                </div>
                <div className="MainPageContainer-parentOfSmall">
                    <div className="MainPageContainer-bottomDark">
                        <p className="MainPageContainer-bottomDark-p">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                    </div>
                    <div className="MainPageContainer-bottomLight">
                        <p className="MainPageContainer-bottomLight-p">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>
                    </div>
                </div>
            </div>
        )
    }
};

export default MainPageContainer;

