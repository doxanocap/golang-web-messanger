import React, { useEffect, useState } from "react";
import "./ChatHistory.css";
const axios = require('axios');

const ChatHistory = ({chatHistoryMessages}) => {
  const url = "http://localhost:8080/12";
    var headers = {}
    
    fetch(url, {
        method : "GET",
        mode: 'cors',
        headers: headers
    })
    .then((response) => {
        if (!response.ok) {
            throw new Error(response.error)
        }
        console.log(response);
        return response.json();
    })
    .then(data => {
        console.log(document.getElementsByTagName('pre').value);
        document.getElementsByTagName('pre').value = data.messages
    })
    .catch(function(error) {
        document.getElementsByTagName('pre').value = error;
    });
console.log(headers);
    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {chatHistoryMessages.map((message, index) => (
            <p key={index}>{message}</p>
          ))
        }
      </div>
    );
}

export default ChatHistory;