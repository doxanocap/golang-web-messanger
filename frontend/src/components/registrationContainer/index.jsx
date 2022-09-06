import React, {useState} from "react";
import {useNavigate} from 'react-router-dom';
import "./index.css"

const RegistrationPageContainer = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const regNewUser = async (e) => {
        e.preventDefault();

        await fetch("http://localhost:8080/api/register", {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                username,
                email,
                password
            })
        });
        navigate('/login');
        window.location.reload();
    }

    return (
        <div className="RegistrationPageContainer">
            <form onSubmit={regNewUser}>
                <h2>Sign Up page</h2>
                <input className="RegistrationPageContainer-input1" type={"username"} id={"inputUsername"} placeholder={"Username"} onChange={e => setUsername(e.target.value)} />
                <input className="RegistrationPageContainer-input2" type={"email"} id={"inputEmail"} placeholder={"Email"} onChange={e => setEmail(e.target.value)} />
                <input className="RegistrationPageContainer-input3" type={"password"} id={"inputPassword"} placeholder={"Password"} onChange={e => setPassword(e.target.value)} />
                <button onClick={regNewUser}>Register</button>
            </form>
        </div>
    )
};

export default RegistrationPageContainer;