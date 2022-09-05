import React, { Component, useState, useEffect } from "react";
import PagesHeader from '../pagesHeader/index';
import "./index.css";

const socket = new WebSocket('ws://localhost:8080/api/websocket');

const SingleChatContainer = ({Username}) => {
  const [usersList, setUsersList] = useState();
  const [onlineUsersList, setOnlineUsersList] = useState();
  const [chatHistory, setChatHistory] = useState();
  const [input, setInput] = useState("");
  socket.onmessage = (msg) => {
    var currentTime = new Date().toLocaleString();
    setChatHistory(chatHistory => [...chatHistory, {time: currentTime, username: JSON.parse(msg.data).username, message: JSON.parse(msg.data).message}])
  }
  socket.onopen = () => {
    console.log("Successfully Connected");
    ParsingChatHistory();
    ParseAllUsers();
    ParseOnlineUsers();
  };
  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };
  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };

  const sendMessage = () => {
    socket.send(input);
    setInput("")
    document.getElementById("mainInput").value = "";
  } 
  
  const handleInput = (event) => {
    setInput(event.target.value)
  }

  const ParsingChatHistory = async () => {
    const response = await fetch("http://localhost:8080/api/fetch");
    const data = await response.json();
    const myArrStr = JSON.parse(data);
    setChatHistory(myArrStr)
  }

  const ParseAllUsers = async () => {
    const response = await fetch("http://localhost:8080/api/all-users");
    const data = await response.json();
    const myArrStr = JSON.parse(data);
    setUsersList(myArrStr);
  }

  const ParseOnlineUsers = async () => {
    const response = await fetch("http://localhost:8080/api/online-users");
    const data = await response.json();
    const myArrStr = JSON.parse(data);
    setOnlineUsersList(myArrStr);
  }

  return (
      <div className="SingleChatContainer">
        <div className="left-panel">
          {typeof(onlineUsersList) === "object" ?
            onlineUsersList.map((item, i) => {
              return (
                <div className="users-boxes">
                  <p className="username" key={item.token}>{item.username}</p>
                  <p className="message" key={item.token+item.username+item.id}>{item.email}</p>
                </div>
              )
              }) : <p>Nulllll</p>}
        </div>
        <div>
          <div className="chatHistory">
            <ul key={"Qwe"}>
              {typeof(chatHistory) === "object" ? 
                chatHistory.map((item, i) => {
                  if (item.username === Username) {
                    return (
                      <li key={item.username+item.time+item.message}  className="right-sided">
                        <div className="text-blocks">
                          <p className="message">{item.message}</p>
                          <p className="time">{item.time.substr(12,5)}</p>
                        </div>
                      </li>
                    )
                  } else {
                    return (
                      <li key={item.username+item.time+item.message}  className="left-sided">
                        <div key={item.username+item.time+item.message} className="text-blocks">
                          <p className="username">{item.username}</p>
                          <p className="message">{item.message}</p>
                          <p className="time">{item.time.substr(12,5)}</p>
                        </div> 
                      </li>
                    )
                  }
                }) : null}
            </ul>
          </div>
          <div className="messageInput">
            <input id="mainInput" placeholder="Write your message!" onChange={(e) => {handleInput(e)}} onKeyDown={(e) => { if (e.key === 'Enter') { sendMessage(e) }}}/>
            <button onClick={() => {sendMessage(input)}}>Send</button>
          </div>
        </div>
      </div>
  );
}

export default SingleChatContainer;