// api/index.js
var socket = new WebSocket("ws://localhost:8080/ws");

const connect = (cb) => {
  console.log("connecting");
  socket.addEventListener('message', (event) => {
    console.log('Message from server ', event.data);
  });

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

const sendMsg = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };