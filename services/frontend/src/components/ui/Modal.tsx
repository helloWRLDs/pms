import { IoCloseSharp } from "react-icons/io5";
import { useOutsideClick } from "../../hooks/useModal";
import { useEffect } from "react";

type ModalSize = "sm" | "md" | "lg" | "xl" | "full";

type ModalProps = React.HTMLAttributes<HTMLDivElement> & {
  onClose: () => void;
  visible: boolean;
  title: string;
  size?: ModalSize;
  hideCloseButton?: boolean;
  preventClickOutside?: boolean;
  showBackdrop?: boolean;
  position?: "center" | "top";
  contentClassName?: string;
  headerClassName?: string;
  bodyClassName?: string;
  backdropClassName?: string;
};

const sizeClasses: Record<ModalSize, string> = {
  sm: "max-w-md",
  md: "max-w-lg",
  lg: "max-w-2xl",
  xl: "max-w-4xl",
  full: "max-w-full m-4",
};

export const Modal = ({
  visible,
  onClose,
  title,
  children,
  className = "",
  size = "md",
  hideCloseButton = false,
  preventClickOutside = false,
  showBackdrop = true,
  position = "center",
  contentClassName = "",
  headerClassName = "",
  bodyClassName = "",
  backdropClassName = "",
  ...props
}: ModalProps) => {
  const modalRef = useOutsideClick(() => {
    if (!preventClickOutside) {
      onClose();
    }
  });

  useEffect(() => {
    const handleEsc = (e: KeyboardEvent) => {
      if (e.key === "Escape" && !preventClickOutside) onClose();
    };
    document.addEventListener("keydown", handleEsc);
    return () => document.removeEventListener("keydown", handleEsc);
  }, [onClose, preventClickOutside]);

  if (!visible) return null;

  return (
    <div
      className={`fixed inset-0 z-50 ${
        position === "center" ? "flex items-center" : "flex items-start pt-20"
      } justify-center ${
        showBackdrop ? "bg-black/50 backdrop-blur-sm" : ""
      } ${backdropClassName}`}
      role="dialog"
      aria-modal="true"
      aria-labelledby="modal-title"
    >
      {/* Modal Animation Wrapper */}
      <div
        className={`
          w-full transform transition-all duration-300 ease-out
          ${visible ? "opacity-100 translate-y-0" : "opacity-0 -translate-y-4"}
          ${sizeClasses[size]}
          ${className}
        `}
      >
        {/* Modal Content */}
        <div
          ref={modalRef}
          className={`
            relative bg-primary-500/30 backdrop-blur-lg
            rounded-xl shadow-2xl border border-primary-400/30
            ${contentClassName}
          `}
          {...props}
        >
          {/* Header */}
          <div
            className={`
              flex justify-between items-center p-6
              border-b border-primary-400/30
              ${headerClassName}
            `}
          >
            <h3 id="modal-title" className="text-xl font-semibold text-white">
              {title}
            </h3>
            {!hideCloseButton && (
              <button
                onClick={onClose}
                className="text-white/60 hover:text-white transition-colors cursor-pointer rounded-lg p-1 hover:bg-white/10"
                aria-label="Close modal"
              >
                <IoCloseSharp size={24} />
              </button>
            )}
          </div>

          {/* Body */}
          <div className={`p-6 text-white/90 ${bodyClassName}`}>{children}</div>
        </div>
      </div>
    </div>
  );
};

// Helper component for Modal.Header if needed
export const ModalHeader: React.FC<{
  className?: string;
  children: React.ReactNode;
}> = ({ className = "", children }) => (
  <div
    className={`flex justify-between items-center p-6 border-b border-primary-400/30 ${className}`}
  >
    {children}
  </div>
);

// Helper component for Modal.Body if needed
export const ModalBody: React.FC<{
  className?: string;
  children: React.ReactNode;
}> = ({ className = "", children }) => (
  <div className={`p-6 text-white/90 ${className}`}>{children}</div>
);

// Helper component for Modal.Footer if needed
export const ModalFooter: React.FC<{
  className?: string;
  children: React.ReactNode;
}> = ({ className = "", children }) => (
  <div
    className={`flex justify-end gap-4 p-6 border-t border-primary-400/30 ${className}`}
  >
    {children}
  </div>
);
