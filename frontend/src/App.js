import {BrowserRouter, Route, Routes } from 'react-router-dom';
import {React, useState, useEffect} from "react";
import Register from './pages/register';
import MainPage from './pages/mainpage';
import Webchat from './pages/webchat';
import Login from './pages/login';
import PagesHeader from  './components/pagesHeader'

const App = () => {
    const [username, setUsername] = useState("")

    useEffect(()=> {
        (
            async () => {
                const response = await fetch("https://webchat-doxa.herokuapp.com/api/user", {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const data = await response.json()
                console.log(data);
                setUsername(data.username)
            }
        )();
    });
    return (
        <BrowserRouter>
            <PagesHeader setUsername={setUsername} username={username}/>
            <Routes>
                <Route path={"/"} element={<MainPage CurrentUsersName={username} />} />
                <Route path={"/webchat"} element={<Webchat Username={username}/>}/>
                <Route path={"/login"} element={<Login setUsername={setUsername}/>}/>
                <Route path={"/register"} element={<Register />}/>
            </Routes>
        </BrowserRouter>
    );
}

export default App;