import { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import useWs from "../../hooks/useWs";
import TiptapEditor from "../../components/text/TiptapEditor";
import { Document as DocType } from "../../lib/document/document";

const DocumentPage = () => {
  const documentID = useParams()["documentID"];

  const { val, send } = useWs(`ws://localhost:8080/ws/docs/${documentID}`);
  const [doc, setDoc] = useState<DocType | null>(null);
  const [content, setContent] = useState("");
  const [isInitialized, setIsInitialized] = useState(false);
  const lastReceivedContent = useRef("");

  useEffect(() => {
    console.log("WebSocket received:", val);
    if (val) {
      const receivedDoc: DocType =
        typeof val === "string" ? JSON.parse(val) : val;
      setDoc(receivedDoc);
      if (receivedDoc.body) {
        const decodedContent = atob(receivedDoc.body);
        if (decodedContent !== lastReceivedContent.current) {
          lastReceivedContent.current = decodedContent;
          setContent(decodedContent);
        }
      }
      setIsInitialized(true);
    }
  }, [val]);

  const handleContentChange = (newContent: string) => {
    setContent(newContent);
    if (doc) {
      const updatedDoc = { ...doc, body: btoa(newContent) };
      setDoc(updatedDoc);
    }
  };

  useEffect(() => {
    const interval = setInterval(() => {
      if (doc && content !== lastReceivedContent.current) {
        console.log("Sending document update via WebSocket");
        send(doc);
      }
    }, 500);
    return () => clearInterval(interval);
  }, [doc, content, send]);

  if (!isInitialized) {
    return (
      <div className="min-h-screen bg-primary-600 flex items-center justify-center">
        <div className="text-neutral-100 text-lg">Loading document...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-primary-600 text-neutral-100">
      <div className="max-w-4xl mx-auto px-6 py-8">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-accent-500 mb-2">
            {doc?.title || "Untitled Document"}
          </h1>
          <div className="flex items-center gap-4 text-sm text-neutral-400">
            <span>Document ID: {documentID}</span>
            <span>•</span>
            <span>Auto-save enabled</span>
          </div>
        </div>

        <div className="bg-secondary-200 rounded-lg p-6 shadow-lg">
          <TiptapEditor
            content={content}
            onChange={handleContentChange}
            placeholder="Start writing your document..."
            className="min-h-[500px]"
          />
        </div>
      </div>
    </div>
  );
};

export default DocumentPage;
