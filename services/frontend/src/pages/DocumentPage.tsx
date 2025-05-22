import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useWs from "../hooks/useWs";
import Editor from "../components/text/Editor";

const DocumentPage = () => {
  const documentID = useParams()["documentID"];

  const { val, send } = useWs(`ws://localhost:8080/ws/docs/${documentID}`);
  const [doc, setDoc] = useState<Document | null>(null);
  const [content, setContent] = useState("");

  // Receive doc over WebSocket
  useEffect(() => {
    console.log(val);
    if (val) {
      const receivedDoc: Document =
        typeof val === "string" ? JSON.parse(val) : val;
      setDoc(receivedDoc);
      if (receivedDoc.body) {
        setContent(atob(receivedDoc.body));
      }
    }
  }, [val]);

  // Update body when content changes
  useEffect(() => {
    if (doc) {
      setDoc({ ...doc, body: btoa(content) });
    }
  }, [content]);

  // Send full doc every 3 seconds
  useEffect(() => {
    const interval = setInterval(() => {
      if (doc) {
        send(doc);
      }
    }, 50);
    return () => clearInterval(interval);
  });

  return (
    <div>
      <section></section>
      <section>
        <Editor
          content={content}
          onChange={(content) => {
            setContent(content);
          }}
        />
      </section>
    </div>
  );
};

export default DocumentPage;
