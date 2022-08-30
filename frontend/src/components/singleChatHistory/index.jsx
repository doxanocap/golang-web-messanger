import React, { useEffect, useState } from "react";
import "./index.css";

const SingleChatHistory = ({ChatHistoryMessages}) => {
  if (typeof(ChatHistoryMessages) === "object") {
    return (
      <div className="SingleChatHistory">
        <h2>Chat History</h2>
        {ChatHistoryMessages.map((item, i) => (
          <p key={item.message + item.time}>{}{item.username} : {item.message}</p>
        ))}
      </div>
    );
  }
}

export default SingleChatHistory; 