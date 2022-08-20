import React, { Component } from "react";
import "./ChatHistory.css";

const ChatHistory = ({chatHistoryMessages}) => {
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