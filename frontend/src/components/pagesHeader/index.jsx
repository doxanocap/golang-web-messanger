import {useNavigate} from 'react-router-dom';
import React, {useEffect, useState} from "react";
import "./index.css"


const PagesHeader = ({username,setUsername}) => {
    let menu;
    const navigate = useNavigate();
    const logout = async () => {
        await fetch("http://localhost:8080/api/logout", {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });
        navigate('/')
        setUsername("")
        window.location.reload()
    };
    if (username === "" || typeof(username)==="undefined") {
        menu = (
            <nav>
                <a href='/'>Main Page</a>
                <a href='/login'>Login</a>
                <a href='/register'>Register</a>
            </nav>
        )
    } else {
        menu = (
            <nav>
                <a href='/webchat'>My Chat</a>
                <a href='/'>Main Page</a>
                <a onClick={()=>{logout()}}>Logout</a>
                <a>User:{username}</a>
            </nav>
        )
    }
    return (
        <div className="pagesHeader">
            <h2>WebChat</h2>
            {menu}
        </div>
    )
};
//
// const checkUser = async () => {
//     const response = await fetch("http://localhost:8080/api/user", {
//         method: "GET",
//         headers: {'Content-Type': 'application/json'},
//         credentials: 'include',
//     });
//     const data = await response.json()
//     if (data !== {"message":"unauthenticated"}) {
//         setUsername(data.username)
//     }
// }

export default PagesHeader;