import { useCallback, useEffect, useRef, useState } from "react";

const useWs = (url: string, reconnectInterval = 5000) => {
  const [isReady, setIsReady] = useState(false);
  const [val, setVal] = useState<string | null>(null);
  const ws = useRef<WebSocket | null>(null);
  const reconnectTimer = useRef<NodeJS.Timeout | null>(null);
  const shouldReconnect = useRef<boolean>(true);

  const connect = useCallback(() => {
    if (ws.current) {
      ws.current.close(); // Close previous if exists
    }

    const socket = new WebSocket(url);

    socket.onopen = () => {
      console.log("WebSocket connected");
      setIsReady(true);
    };

    socket.onclose = () => {
      console.warn("WebSocket disconnected, will retry...");
      setIsReady(false);

      if (shouldReconnect.current) {
        reconnectTimer.current = setTimeout(() => {
          connect();
        }, reconnectInterval);
      }
    };

    socket.onmessage = (event) => {
      setVal(event.data);
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
      socket.close();
    };

    ws.current = socket;
  }, [url, reconnectInterval]);

  const close = useCallback(() => {
    console.log("Closing WebSocket connection");
    shouldReconnect.current = false;
    if (reconnectTimer.current) {
      clearTimeout(reconnectTimer.current);
      reconnectTimer.current = null;
    }
    if (ws.current) {
      ws.current.close();
      ws.current = null;
    }
  }, []);

  useEffect(() => {
    connect();

    return () => {
      close();
    };
  }, [connect, close]);

  const send = useCallback((data: any) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(data));
    } else {
      console.warn("WebSocket not connected. Cannot send:", data);
    }
  }, []);

  return { isReady, val, send, close };
};

export default useWs;
