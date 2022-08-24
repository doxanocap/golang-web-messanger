import React, { useEffect, useState } from "react";
import "./ChatHistory.css";

const ChatHistory = ({chatHistoryMessages}) => {
    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {chatHistoryMessages.map((item) => (
          <p key={item.Time}>{item.Message}</p>
        ))}
      </div>
    );
}

export default ChatHistory;