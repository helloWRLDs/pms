import { FC, useState } from "react";
import TiptapEditor from "./TiptapEditor";

interface TextEditorProps {
  initialContent?: string;
  onSave?: (content: string) => void;
  className?: string;
}

const TextEditor: FC<TextEditorProps> = ({
  initialContent = "",
  onSave,
  className = "",
}) => {
  const [content, setContent] = useState(initialContent);

  const handleChange = (newContent: string) => {
    setContent(newContent);
    onSave?.(newContent);
  };

  return (
    <div className={className}>
      <TiptapEditor
        content={content}
        onChange={handleChange}
        placeholder="Write something amazing..."
        className="min-h-[200px] p-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </div>
  );
};

export default TextEditor;
