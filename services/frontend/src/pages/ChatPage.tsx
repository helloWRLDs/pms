import React, { useEffect, useRef, useState } from "react";

const ChatPage = () => {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const id = "123"; // You can dynamically set this later
    const socket = new WebSocket(`ws://localhost:8080/dashboard/${id}`);
    socketRef.current = socket;

    socket.onopen = () => {
      console.log("Connected to WebSocket server ✅");
    };

    socket.onmessage = (event) => {
      console.log("Received from server:", event.data);
      setMessages((prev) => [...prev, `Server: ${event.data}`]);
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onclose = () => {
      console.warn("WebSocket connection closed.");
    };

    return () => {
      socket.close();
    };
  }, []);

  const handleSendMessage = () => {
    if (
      message.trim() !== "" &&
      socketRef.current?.readyState === WebSocket.OPEN
    ) {
      socketRef.current.send(message);
      setMessages((prev) => [...prev, `You: ${message}`]);
      setMessage("");
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSendMessage();
    }
  };

  return (
    <div
      style={{
        maxWidth: "600px",
        margin: "50px auto",
        padding: "20px",
        border: "1px solid #ccc",
        borderRadius: "8px",
      }}
    >
      <h2>WebSocket Echo Chat</h2>
      <div
        style={{
          height: "300px",
          overflowY: "auto",
          border: "1px solid #ddd",
          padding: "10px",
          marginBottom: "10px",
          borderRadius: "4px",
        }}
      >
        {messages.map((msg, index) => (
          <div key={index} style={{ marginBottom: "8px" }}>
            {msg}
          </div>
        ))}
      </div>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder="Type your message..."
        style={{ width: "80%", padding: "8px", marginRight: "10px" }}
      />
      <button onClick={handleSendMessage} style={{ padding: "8px 16px" }}>
        Send
      </button>
    </div>
  );
};

export default ChatPage;
