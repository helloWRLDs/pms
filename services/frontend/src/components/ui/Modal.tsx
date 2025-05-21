import { IoCloseSharp } from "react-icons/io5";
import { useOutsideClick } from "../../hooks/useModal";
import { useEffect } from "react";

type ModalProps = React.HTMLAttributes<HTMLDivElement> & {
  onClose: () => void;
  visible: boolean;
  title: string;
};

export const Modal = ({
  visible,
  onClose,
  title,
  children,
  className,
}: ModalProps) => {
  const modalRef = useOutsideClick(onClose);

  useEffect(() => {
    const handleEsc = (e: KeyboardEvent) => {
      if (e.key === "Escape") onClose();
    };
    document.addEventListener("keydown", handleEsc);
    return () => document.removeEventListener("keydown", handleEsc);
  }, [onClose]);

  if (!visible) return null;

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm"
      role="dialog"
      aria-modal="true"
      aria-labelledby="modal-title"
    >
      <div
        ref={modalRef}
        className={`rounded-xl shadow-2xl max-w-lg w-full relative p-6 ${className}`}
      >
        <div className="flex justify-between items-center border-b-1 border-gray-200 pb-3">
          <h3 id="modal-title" className="text-xl font-semibold">
            {title}
          </h3>
          <button
            onClick={onClose}
            className="text-gray-600 hover:text-red-600 transition-colors cursor-pointer"
            aria-label="Close modal"
          >
            <IoCloseSharp size={28} />
          </button>
        </div>
        <div className="mt-4">{children}</div>
      </div>
    </div>
  );
};
