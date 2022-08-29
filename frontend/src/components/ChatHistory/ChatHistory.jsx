import React, { useEffect, useState } from "react";
import "./ChatHistory.css";

const ChatHistory = ({chatHistoryMessages}) => {
  if (typeof(chatHistoryMessages) !== "object") {
    return null
  } 
  return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {chatHistoryMessages.map((item, i) => (
          <p key={item.Time}>{item.Username} : {item.Message}</p>
        ))}
      </div>
    );
}

export default ChatHistory; 