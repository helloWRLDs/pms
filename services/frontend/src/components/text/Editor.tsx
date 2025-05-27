import JoditEditor from "jodit-react";
import { useRef } from "react";

type EditorProps = {
  content: string;
  onChange: (updatedContent: string) => void;
};

const Editor = ({ content, onChange }: EditorProps) => {
  const editor = useRef(null);
  document.querySelector(`a[href="https://xdsoft.net/jodit/"]`)?.remove();

  return (
    <div>
      <JoditEditor
        ref={editor}
        value={content}
        onChange={(newContent) => onChange(newContent)}
      />
    </div>
  );
};

export default Editor;
