import {useNavigate} from 'react-router-dom';
import React from "react";
import "./index.css"


const PagesHeader = ({username,setUsername}) => {
    let menu;
    const navigate = useNavigate();
    const logout = async () => {
        await fetch("http://localhost:3000/api/logout", {
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
                <a  href="/" onClick={()=>{navigate('/')}}>Main Page</a>
                <a href="/login" onClick={()=>{navigate('/login')}}>Login</a>
                <a href="/register" onClick={()=>{navigate('/register')}}>Register</a>
            </nav>
        )
    } else {
        menu = (
            <nav>
                <a href="/webchat" onClick={()=>{navigate('/webchat')}}>My Chat</a>
                <a href="/" onClick={()=>{navigate('/')}}>Main Page</a>
                <a href="/" onClick={()=>{logout()}}>Logout</a>
                <a href='/'>User:{username}</a>
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

export default PagesHeader;