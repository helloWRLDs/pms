import { useState, useRef, useEffect } from "react";
import JoditEditor from "jodit-react";
import HTMLReactParser from "html-react-parser";

const initialContent = `<p>some text</p>`;

const TestPage1 = () => {
  const editor = useRef(null);
  const [content, setContent] = useState(initialContent);

  useEffect(() => {
    console.log(content);
  }, [content]);

  document.querySelector(`a[href="https://xdsoft.net/jodit/"]`)?.remove();

  return (
    <div className="container">
      <JoditEditor
        ref={editor}
        value={content}
        onChange={(newContent) => setContent(newContent)}
      />

      <div>{HTMLReactParser(content)}</div>
    </div>
  );
};

export default TestPage1;
