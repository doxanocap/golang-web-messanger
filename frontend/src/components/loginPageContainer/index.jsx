import React, {useState} from "react";
import {useNavigate} from "react-router-dom";

import "./index.css"

const LoginPageContainer = ({setUsername}) => {
    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const loginUser = async (e) => {
        e.preventDefault();
        const response = await fetch("http://localhost:8080/api/login", {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });
        const data = await response.json()
        navigate("/");
        window.location.reload()
        setUsername(data.username);
    }
    return (
        <div className="loginPageContainer">
            <form onSubmit={loginUser}>
                <h1>Sign Up page</h1>
                <input className="loginPageContainer-input1" type={"email"} id={"inputEmail"} placeholder={"Email"} onChange={e => setEmail(e.target.value)} />
                <input className="loginPageContainer-input2" type={"password"} id={"inputPassword"} placeholder={"Password"} onChange={e => setPassword(e.target.value)} />
                <button onClick={loginUser}>Login</button>
            </form>
        </div>
    )
};

export default LoginPageContainer;