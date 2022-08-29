import React from "react";
import {useNavigate} from 'react-router-dom';

import "./index.css"

const MainPageHeader = () => {
    const navigate = useNavigate();

    const navigateToChat = () => {
        navigate('/webchat');
    };

    const navigateToSignUp = () => {
        navigate('/sing-up');
    };

    return (
        <div className="MainPageHeader">
            <h2>Realtime Chat App Header</h2>
            <button onClick={navigateToChat} >Go to the chat</button>
            <button onClick={navigateToSignUp}>Sing In</button>
        </div>
    )
};

export {MainPageHeader};