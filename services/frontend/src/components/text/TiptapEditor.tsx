import { useEditor, EditorContent, JSONContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import Typography from "@tiptap/extension-typography";
import Link from "@tiptap/extension-link";
import Image from "@tiptap/extension-image";
import { useEffect, useRef, useCallback } from "react";

type TiptapEditorProps = {
  content: string;
  onChange: (content: string) => void;
  onUpdate?: (content: string) => void;
  placeholder?: string;
  editable?: boolean;
  className?: string;
};

const TiptapEditor = ({
  content,
  onChange,
  onUpdate,
  placeholder = "Start writing...",
  editable = true,
  className = "",
}: TiptapEditorProps) => {
  const isUpdatingFromProps = useRef(false);
  const lastContent = useRef(content);

  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        bulletList: {
          keepMarks: true,
          keepAttributes: false,
        },
        orderedList: {
          keepMarks: true,
          keepAttributes: false,
        },
      }),
      Placeholder.configure({
        placeholder,
      }),
      Typography,
      Link.configure({
        openOnClick: false,
        HTMLAttributes: {
          class:
            "text-accent-500 underline hover:text-accent-400 transition-colors",
        },
      }),
      Image.configure({
        HTMLAttributes: {
          class: "max-w-full h-auto rounded-lg",
        },
      }),
    ],
    content: content || "",
    editable,
    onUpdate: ({ editor }) => {
      if (!isUpdatingFromProps.current) {
        const html = editor.getHTML();
        lastContent.current = html;
        onChange(html);
        onUpdate?.(html);
      }
    },
    editorProps: {
      attributes: {
        class: `focus:outline-none min-h-[200px] p-4 text-neutral-100 ${className}`,
        spellcheck: "false",
      },
    },
  });

  // Handle content updates from props (websocket) while preserving cursor
  useEffect(() => {
    if (editor && content !== lastContent.current) {
      isUpdatingFromProps.current = true;

      // Store current cursor position
      const { from, to } = editor.state.selection;

      // Update content
      editor.commands.setContent(content || "", false);

      // Restore cursor position if possible
      const newDoc = editor.state.doc;
      const maxPos = newDoc.content.size;
      const safeFrom = Math.min(from, maxPos);
      const safeTo = Math.min(to, maxPos);

      // Use setTimeout to ensure the update is processed
      setTimeout(() => {
        try {
          editor.commands.setTextSelection({ from: safeFrom, to: safeTo });
        } catch (error) {
          // If cursor restoration fails, place cursor at the end
          editor.commands.focus("end");
        }
        isUpdatingFromProps.current = false;
      }, 0);
    }
  }, [content, editor]);

  // Handle editable state changes
  useEffect(() => {
    if (editor) {
      editor.setEditable(editable);
    }
  }, [editable, editor]);

  const addLink = useCallback(() => {
    if (!editor) return;

    const url = window.prompt("Enter URL:");
    if (url) {
      editor
        .chain()
        .focus()
        .extendMarkRange("link")
        .setLink({ href: url })
        .run();
    }
  }, [editor]);

  const addImage = useCallback(() => {
    if (!editor) return;

    const url = window.prompt("Enter image URL:");
    if (url) {
      editor.chain().focus().setImage({ src: url }).run();
    }
  }, [editor]);

  if (!editor) {
    return (
      <div className="bg-secondary-400/5 border border-secondary-200/20 rounded-lg p-4 animate-pulse">
        <div className="h-4 bg-secondary-100 rounded w-1/4 mb-4"></div>
        <div className="h-4 bg-secondary-100 rounded w-3/4 mb-2"></div>
        <div className="h-4 bg-secondary-100 rounded w-1/2"></div>
      </div>
    );
  }

  return (
    <div className="bg-secondary-400/5 border border-secondary-200/20 rounded-lg overflow-hidden">
      {/* Toolbar */}
      {editable && (
        <div className="border-b border-secondary-200/20 p-3 bg-secondary-300/10">
          <div className="flex flex-wrap items-center gap-2">
            <div className="flex items-center gap-1">
              <button
                onClick={() => editor.chain().focus().toggleBold().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("bold")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Bold
              </button>
              <button
                onClick={() => editor.chain().focus().toggleItalic().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("italic")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Italic
              </button>
              <button
                onClick={() => editor.chain().focus().toggleStrike().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("strike")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Strike
              </button>
            </div>

            <div className="w-px h-6 bg-secondary-100"></div>

            <div className="flex items-center gap-1">
              <button
                onClick={() =>
                  editor.chain().focus().toggleHeading({ level: 1 }).run()
                }
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("heading", { level: 1 })
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                H1
              </button>
              <button
                onClick={() =>
                  editor.chain().focus().toggleHeading({ level: 2 }).run()
                }
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("heading", { level: 2 })
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                H2
              </button>
              <button
                onClick={() =>
                  editor.chain().focus().toggleHeading({ level: 3 }).run()
                }
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("heading", { level: 3 })
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                H3
              </button>
            </div>

            <div className="w-px h-6 bg-secondary-100"></div>

            <div className="flex items-center gap-1">
              <button
                onClick={() => editor.chain().focus().toggleBulletList().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("bulletList")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Bullet List
              </button>
              <button
                onClick={() => editor.chain().focus().toggleOrderedList().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("orderedList")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Ordered List
              </button>
            </div>

            <div className="w-px h-6 bg-secondary-100"></div>

            <div className="flex items-center gap-1">
              <button
                onClick={addLink}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("link")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Link
              </button>
              <button
                onClick={addImage}
                className="px-3 py-1.5 text-xs rounded bg-secondary-100 text-neutral-200 hover:bg-secondary-50 transition-colors"
              >
                Image
              </button>
            </div>

            <div className="w-px h-6 bg-secondary-100"></div>

            <div className="flex items-center gap-1">
              <button
                onClick={() => editor.chain().focus().toggleBlockquote().run()}
                className={`px-3 py-1.5 text-xs rounded transition-colors ${
                  editor.isActive("blockquote")
                    ? "bg-accent-500 text-primary-700"
                    : "bg-secondary-100 text-neutral-200 hover:bg-secondary-50"
                }`}
              >
                Quote
              </button>
              <button
                onClick={() => editor.chain().focus().setHorizontalRule().run()}
                className="px-3 py-1.5 text-xs rounded bg-secondary-100 text-neutral-200 hover:bg-secondary-50 transition-colors"
              >
                Divider
              </button>
            </div>

            <div className="w-px h-6 bg-secondary-100"></div>

            <div className="flex items-center gap-1">
              <button
                onClick={() => editor.chain().focus().undo().run()}
                disabled={!editor.can().undo()}
                className="px-3 py-1.5 text-xs rounded bg-secondary-100 text-neutral-200 hover:bg-secondary-50 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Undo
              </button>
              <button
                onClick={() => editor.chain().focus().redo().run()}
                disabled={!editor.can().redo()}
                className="px-3 py-1.5 text-xs rounded bg-secondary-100 text-neutral-200 hover:bg-secondary-50 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Redo
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Editor Content */}
      <div className="min-h-[300px]">
        <style
          dangerouslySetInnerHTML={{
            __html: `
            .ProseMirror {
              color: #f5f5f5;
              line-height: 1.6;
            }
            .ProseMirror p {
              margin: 0.75em 0;
            }
            .ProseMirror strong {
              font-weight: 700;
              color: #fbbf24;
            }
            .ProseMirror em {
              font-style: italic;
              color: #a78bfa;
            }
            .ProseMirror s {
              text-decoration: line-through;
            }
            .ProseMirror h1 {
              font-size: 2em;
              font-weight: 700;
              margin: 1em 0 0.5em 0;
              color: #fbbf24;
            }
            .ProseMirror h2 {
              font-size: 1.5em;
              font-weight: 600;
              margin: 0.8em 0 0.4em 0;
              color: #fbbf24;
            }
            .ProseMirror h3 {
              font-size: 1.2em;
              font-weight: 600;
              margin: 0.6em 0 0.3em 0;
              color: #fbbf24;
            }
            .ProseMirror ul, .ProseMirror ol {
              padding-left: 1.5em;
              margin: 0.75em 0;
            }
            .ProseMirror ul {
              list-style-type: disc;
              list-style-position: outside;
            }
            .ProseMirror ol {
              list-style-type: decimal;
              list-style-position: outside;
            }
            .ProseMirror li {
              margin: 0.25em 0;
              display: list-item;
            }
            .ProseMirror blockquote {
              border-left: 4px solid #fbbf24;
              padding-left: 1em;
              margin: 1em 0;
              font-style: italic;
              color: #d1d5db;
            }
            .ProseMirror hr {
              border: none;
              border-top: 2px solid #4b5563;
              margin: 2em 0;
            }
            .ProseMirror a {
              color: #fbbf24;
              text-decoration: underline;
            }
            .ProseMirror a:hover {
              color: #f59e0b;
            }
            .ProseMirror img {
              max-width: 100%;
              height: auto;
              border-radius: 0.5rem;
              margin: 1em 0;
            }
            .ProseMirror .is-editor-empty:first-child::before {
              content: attr(data-placeholder);
              float: left;
              color: #9ca3af;
              pointer-events: none;
              height: 0;
            }
          `,
          }}
        />
        <EditorContent editor={editor} />
      </div>
    </div>
  );
};

export default TiptapEditor;
