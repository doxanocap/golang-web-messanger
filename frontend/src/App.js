import React, { Component, useState, useEffect } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';

const App = () => {
  const [chatHistory, setChatHistory] = useState([]);
  const [input, setInput] = useState("");
  
  const handleInput = (event) => {
    setInput(event.target.value)
  }

  const sendMessage = () => {
    sendMsg(input)
    var currentTime = new Date().toLocaleString();
    setChatHistory(chatHistory => [...chatHistory, {Time: currentTime, Username: "Doxa", Message: input}])
    setInput("")
    document.getElementById("mainInput").value = "";
  }


  async function ParsingChatHistory() {
    const response = await fetch("http://localhost:8080/put");
    const data = await response.json();
    const myArrStr = JSON.parse(data);
    setChatHistory(myArrStr)
  }

  useEffect(() => {
    ParsingChatHistory()
  }, [])

  return (
    <div className="App-header">
      <Header />
      <ChatHistory chatHistoryMessages={chatHistory} />
      <div className="ChatInput">
        <input id="mainInput" onChange={(e) => {handleInput(e)}} />
        <button onClick={() => {sendMessage()}}>Send</button>
      </div>
    </div>
  );
}

export default App;