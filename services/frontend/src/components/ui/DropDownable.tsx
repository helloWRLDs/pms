import { useEffect, useRef, useState, HTMLAttributes } from "react";
import { cn } from "../../lib/utils/cn";

const DropDownable = ({
  children,
  className,
  options,
}: HTMLAttributes<HTMLDivElement> & {
  options: {
    label: string;
    isActive: boolean;
    onClick: () => void;
  }[];
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    if (isOpen) {
      document.addEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [isOpen]);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setIsOpen(false);
      }
    };
    if (isOpen) {
      document.addEventListener("keydown", handleKeyDown);
    }
    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [isOpen]);

  const handleTriggerClick = () => {
    setIsOpen(!isOpen);
  };

  const handleOptionClick = (option: (typeof options)[0]) => {
    option.onClick();
    setIsOpen(false);
  };

  return (
    <div className={cn("relative", className)} ref={dropdownRef}>
      <div onClick={handleTriggerClick} className="cursor-pointer">
        {children}
      </div>

      {/* Dropdown Menu */}
      {isOpen && (
        <div className="absolute top-full left-0 mt-1 bg-secondary-100 border border-primary-400/30 rounded-md shadow-lg z-50 min-w-[160px]">
          {options.map((option, index) => (
            <button
              key={index}
              onClick={() => handleOptionClick(option)}
              className={cn(
                "w-full text-left px-4 py-2 text-sm hover:bg-secondary-200 transition-colors duration-150",
                "first:rounded-t-md last:rounded-b-md",
                "cursor-pointer hover:bg-secondary-200",
                option.isActive && "bg-accent-50 text-accent-600 font-medium"
              )}
            >
              {option.label}
            </button>
          ))}
        </div>
      )}
    </div>
  );
};

export default DropDownable;
