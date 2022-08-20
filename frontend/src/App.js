// App.js
import React, { Component, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';

const App1 = () => {
  const [inputMsg, setInputMsg] = useState("");
  const [chatHistory, setChatHistory] = useState([]);
  //setChatHistory(arr)

  const handleMsg = (event) => {
    setInputMsg(event.target.value)
  }

  const addMessage = (msg) => {
    sendMsg(inputMsg)
    setChatHistory(chatHistory => [...chatHistory, inputMsg])
    setInputMsg("")
    document.getElementById("mainInput").value = "";
  }

  return (
    <div className="App-header">
      <Header />
      <ChatHistory chatHistoryMessages={chatHistory} />
      <input id="mainInput" onChange={(e) => {handleMsg(e)}}></input>
      <button onClick={addMessage}>Hit</button>
    </div>
  );
}

export default App1;