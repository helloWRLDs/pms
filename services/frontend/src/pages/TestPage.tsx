import { HTMLAttributes, useState, useRef, useEffect } from "react";
import useMetaCache from "../store/useMetaCache";
import { cn } from "../lib/utils/cn";
import { Button } from "../components/ui/Button";

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

  // Close dropdown when clicking outside
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

  // Close dropdown on escape key
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
      {/* Trigger */}
      <div onClick={handleTriggerClick} className="cursor-pointer">
        {children}
      </div>

      {/* Dropdown Menu */}
      {isOpen && (
        <div className="absolute top-full left-0 mt-1 bg-white border border-gray-200 rounded-md shadow-lg z-50 min-w-[160px]">
          {options.map((option, index) => (
            <button
              key={index}
              onClick={() => handleOptionClick(option)}
              className={cn(
                "w-full text-left px-4 py-2 text-sm hover:bg-gray-100 transition-colors duration-150",
                "first:rounded-t-md last:rounded-b-md",
                "cursor-pointer hover:bg-gray-300",
                option.isActive && "bg-blue-50 text-blue-600 font-medium"
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

const TestPage = () => {
  const metaCache = useMetaCache();

  const [selectedOption, setSelectedOption] = useState("bob");

  const dropdownOptions = [
    {
      label: "Bob",
      isActive: selectedOption === "bob",
      onClick: () => {
        setSelectedOption("bob");
        console.log("Selected: bob");
      },
    },
    {
      label: "Alice",
      isActive: selectedOption === "alice",
      onClick: () => {
        setSelectedOption("alice");
        console.log("Selected: alice");
      },
    },
    {
      label: "Charlie",
      isActive: selectedOption === "charlie",
      onClick: () => {
        setSelectedOption("charlie");
        console.log("Selected: charlie");
      },
    },
  ];

  return (
    <div className="px-8 py-5 bg-primary-500">
      <DropDownable options={dropdownOptions}>
        <Button>
          Select User: {selectedOption}
          <svg
            className="ml-2 h-4 w-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </Button>
      </DropDownable>
    </div>
  );
};

export default TestPage;
