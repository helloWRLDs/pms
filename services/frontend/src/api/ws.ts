import { toast } from "react-toastify";
import { toastOpts } from "../utils/toast";

export interface WSConfig {
  baseURL: string;
  reconnectIntervalMs?: number;
}

export class WebSocketClient {
  private baseURL: string;
  private reconnectIntervalMs: number;
  private socket: WebSocket | null = null;
  private listeners: Map<string, (data: any) => void> = new Map();
  private shouldReconnect = true;

  constructor(config: WSConfig) {
    this.baseURL = config.baseURL;
    this.reconnectIntervalMs = config.reconnectIntervalMs || 3000;
  }

  connect() {
    this.socket = new WebSocket(this.baseURL);

    this.socket.onopen = () => {
      console.log("WebSocket connected ✅");
    };

    this.socket.onmessage = (event) => {
      try {
        const { event: eventName, data } = JSON.parse(event.data);
        const listener = this.listeners.get(eventName);
        if (listener) {
          listener(data);
        }
      } catch (err) {
        console.error("Invalid WebSocket message format:", event.data);
      }
    };

    this.socket.onclose = () => {
      console.warn("WebSocket closed! Attempting to reconnect...");
      if (this.shouldReconnect) {
        setTimeout(() => this.connect(), this.reconnectIntervalMs);
      }
    };

    this.socket.onerror = (error) => {
      console.error("WebSocket error:", error);
      toast.error("WebSocket connection error!", toastOpts);
      this.socket?.close();
    };
  }

  disconnect() {
    this.shouldReconnect = false;
    this.socket?.close();
  }

  send(event: string, data: any) {
    if (this.socket?.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify({ event, data }));
    } else {
      console.warn("WebSocket is not open. Cannot send message.");
      toast.warning("WebSocket not connected. Message not sent.", toastOpts);
    }
  }

  on(event: string, callback: (data: any) => void) {
    this.listeners.set(event, callback);
  }

  off(event: string) {
    this.listeners.delete(event);
  }
}
