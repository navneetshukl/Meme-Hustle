import React, { useEffect, useState } from "react";

const Memes = () => {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080");

    ws.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.onmessage = (event) => {
      console.log(`Received: ${event.data}`);
      setMessages((prev) => [...prev, `Server: ${event.data}`]);
    };

    ws.onclose = () => {
      console.log("WebSocket closed");
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  const sendMessage = () => {
    if (socket && input.trim() !== "") {
      socket.send(input);
      setMessages((prev) => [...prev, `You: ${input}`]);
      setInput("");
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>WebSocket Chat</h2>
      <div style={{ marginBottom: 10 }}>
        {messages.map((msg, index) => (
          <div key={index}>{msg}</div>
        ))}
      </div>
      <input
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type your message"
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default Memes;
