import { useState, ReactNode, FC } from "react";
import { IoClose, IoMenu } from "react-icons/io5";
import { FiChevronDown, FiChevronRight } from "react-icons/fi";
import { SideBarItem } from "../../lib/ui/sidebar";

interface SideBarProps {
  logo?: {
    href: string;
    imgSrc: string;
    label: string;
  };
  children?: ReactNode;
  className?: string;
}

interface SideBarElementProps {
  children: ReactNode;
  icon?: ReactNode;
  onClick?: () => void;
  className?: string;
  isActive?: boolean;
  level?: number;
}

interface SideBarDropdownProps extends SideBarElementProps {
  isOpen?: boolean;
  defaultOpen?: boolean;
  badge?: ReactNode;
  label: string;
}

const SideBarDropdown: FC<SideBarDropdownProps> = ({
  children,
  icon,
  onClick,
  className = "",
  isActive,
  isOpen,
  defaultOpen = false,
  badge,
  label,
  level = 0,
}) => {
  const [isExpanded, setIsExpanded] = useState(defaultOpen);

  const handleClick = () => {
    setIsExpanded(!isExpanded);
    onClick?.();
  };

  return (
    <div className="relative">
      <button
        onClick={handleClick}
        className={`
          w-full flex items-center justify-between px-4 py-2.5
          text-white/80 hover:text-white
          bg-[var(--color-primary-400)] hover:bg-[var(--color-primary-300)] active:bg-[var(--color-primary-200)]
          transition-all duration-200 rounded-lg
          ${isActive ? "bg-[var(--color-primary-200)] text-white" : ""}
          ${className}
        `}
      >
        <div className="flex items-center gap-3">
          {icon}
          <span className="text-sm font-medium">{label}</span>
        </div>
        <div className="flex items-center gap-2">
          {badge}
          {isExpanded ? (
            <FiChevronDown className="text-white/60" />
          ) : (
            <FiChevronRight className="text-white/60" />
          )}
        </div>
      </button>
      {isExpanded && (
        <div className={`mt-1 ${level === 0 ? "ml-3" : "ml-4"}`}>
          <div
            className={`
            ${level === 0 ? "pl-4" : "pl-3"} 
            ${level === 0 ? "border-l border-[var(--color-primary-300)]" : ""}
          `}
          >
            {children}
          </div>
        </div>
      )}
    </div>
  );
};

const SideBarElement: FC<SideBarElementProps> = ({
  children,
  icon,
  onClick,
  className = "",
  isActive,
  level = 0,
}) => {
  return (
    <button
      onClick={onClick}
      className={`
        w-full flex items-center gap-3 px-4 py-2.5
        text-white/80 hover:text-white
        bg-[var(--color-primary-400)] hover:bg-[var(--color-primary-300)] active:bg-[var(--color-primary-200)]
        transition-all duration-200 rounded-lg
        ${isActive ? "bg-[var(--color-primary-200)] text-white" : ""}
        ${level > 0 ? "ml-4" : ""}
        ${className}
      `}
    >
      {icon}
      <span className="text-sm font-medium">{children}</span>
    </button>
  );
};

const SideBar: FC<SideBarProps> & {
  Element: typeof SideBarElement;
  Dropdown: typeof SideBarDropdown;
} = ({ logo, children, className = "" }) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="fixed top-4 left-4 z-50 p-2.5 bg-[var(--color-accent-400)] text-white shadow-lg rounded-lg lg:hidden hover:bg-[var(--color-accent-300)] transition-colors duration-200"
        aria-label="Toggle Menu"
      >
        {isOpen ? (
          <IoClose className="w-6 h-6" />
        ) : (
          <IoMenu className="w-6 h-6" />
        )}
      </button>

      {/* Overlay */}
      <div
        className={`fixed inset-0 bg-black/50 backdrop-blur-sm z-30 transition-opacity duration-300 lg:hidden ${
          isOpen ? "opacity-100" : "opacity-0 pointer-events-none"
        }`}
        onClick={() => setIsOpen(false)}
      />

      <aside
        className={`
          fixed top-0 left-0 z-40 h-screen 
          transform transition-all duration-300 ease-in-out 
          bg-[var(--color-primary-500)] border-r border-[var(--color-primary-300)]
          ${
            isOpen
              ? "translate-x-0 w-72"
              : "-translate-x-full w-0 lg:w-72 lg:translate-x-0"
          }
          ${className}
        `}
      >
        {logo && (
          <div className="flex items-center gap-3 px-6 py-4 border-b border-[var(--color-primary-300)]">
            <a href={logo.href} className="flex items-center gap-3">
              <img src={logo.imgSrc} className="h-8" alt="Logo" />
              <span className="text-xl font-semibold text-white">
                {logo.label}
              </span>
            </a>
          </div>
        )}
        <nav className="flex flex-col h-[calc(100vh-4rem)] overflow-y-auto p-3 space-y-1">
          {children}
        </nav>
      </aside>
    </>
  );
};

SideBar.Element = SideBarElement;
SideBar.Dropdown = SideBarDropdown;

export default SideBar;
