import React, { Component, useState, useEffect } from "react";
import "./App.css";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
const socket = new WebSocket('ws://localhost:8080/ws');


const App = () => {
  const [chatHistory, setChatHistory] = useState("");
  const [input, setInput] = useState("");
  socket.onmessage = (msg) => {
    var currentTime = new Date().toLocaleString();
    setChatHistory(chatHistory => [...chatHistory, {Time: currentTime, Username: JSON.parse(msg.data).Username, Message: JSON.parse(msg.data).Message}])
  }

  socket.onopen = () => {
    console.log("Successfully Connected");
    ParsingChatHistory();
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

  async function ParsingChatHistory() {
    const response = await fetch("http://localhost:8080/put");
    const data = await response.json();
    const myArrStr = JSON.parse(data);
    setChatHistory(myArrStr)
  } 

  return (
    <div className="App-header">
      <Header />
      <ChatHistory chatHistoryMessages={chatHistory} /> 
      <div className="ChatInput">
        <input id="mainInput" onChange={(e) => {handleInput(e)}} onKeyDown={(e) => { if (e.key === 'Enter') { sendMessage(e) }}}/>
        <button onClick={() => {sendMessage(input)}}>Send</button>
      </div>
    </div>
  );
}

export default App;