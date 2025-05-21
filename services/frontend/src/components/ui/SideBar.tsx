import { useState } from "react";

type SideBarElementProps = React.HTMLAttributes<HTMLDivElement> & {
  isActive?: boolean;
};

const SideBarElement = ({
  className,
  isActive = false,
  ...props
}: SideBarElementProps) => {
  return (
    <div
      className={`flex items-center gap-2 p-2 rounded  cursor-pointer group ${className}`}
      {...props}
    >
      {props.children}
    </div>
  );
};

type SideBarProps = React.HTMLAttributes<HTMLDivElement> & {
  logo?: {
    href: string;
    imgSrc: string;
    label: string;
  };
};

const SideBar = ({
  children,
  logo,
  className = "",
  ...props
}: SideBarProps) => {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="fixed top-4 left-4 z-50 p-2 bg-white shadow-md rounded lg:hidden"
      >
        {isOpen ? (
          <span className="w-5 h-5" onClick={() => setIsOpen(false)}>
            Close
          </span>
        ) : (
          <span className="w-5 h-5" onClick={() => setIsOpen(true)}>
            Open
          </span>
        )}
      </button>
      <aside
        className={`fixed top-0 left-0 z-40 h-screen w-64 transform transition-transform duration-300 ease-in-out bg-gray-50 dark:bg-gray-800 ${
          isOpen ? "translate-x-0" : "w-0"
        } md:translate-x-0 md:w-64`}
        {...props}
      >
        <div className="px-3 py-4 overflow-y-none bg-primary-500 relative h-full">
          <div className="space-y-2 font-medium">
            {logo && (
              <div id="aside-logo">
                <a href={logo.href} className="flex items-center ps-2.5 mb-5">
                  <img
                    src={logo.imgSrc}
                    className="h-6 me-3 sm:h-7"
                    alt={`${logo.label} Logo`}
                  />
                  <span className="self-center text-xl font-semibold whitespace-nowrap text-accent-500">
                    {logo.label}
                  </span>
                </a>
              </div>
            )}
            {children}
          </div>
        </div>
      </aside>
    </>
  );
};

SideBar.Element = SideBarElement;

export default SideBar;
