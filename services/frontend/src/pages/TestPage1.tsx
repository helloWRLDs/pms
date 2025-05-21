import { useState, useRef, useEffect } from "react";
import JoditEditor from "jodit-react";
import HTMLReactParser from "html-react-parser";

const TestPage1 = () => {
  const editor = useRef(null);
  const [content, setContent] = useState("");

  useEffect(() => {
    console.log(HTMLReactParser(content));
  }, [content]);

  document.querySelector(`a[href="https://xdsoft.net/jodit/"]`)?.remove();

  return (
    <div className="container">
      <JoditEditor
        ref={editor}
        value={content}
        config={{
          toolbarAdaptive: true,
        }}
        onChange={(newContent) => setContent(newContent)}
      />

      <div>{content}</div>
    </div>
  );
};

export default TestPage1;
