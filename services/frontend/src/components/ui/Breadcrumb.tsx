import { FC, useState, useEffect } from "react";
import { cn } from "../../lib/utils/cn";

export interface BreadcrumbItem {
  id: string;
  label: string;
  href?: string;
  onClick?: () => void;
  isActive?: boolean;
  isClickable?: boolean;
}

interface BreadcrumbProps {
  items: BreadcrumbItem[];
  className?: string;
  maxItems?: number;
}

const Breadcrumb: FC<BreadcrumbProps> = ({
  items,
  className = "",
  maxItems,
}) => {
  const [screenSize, setScreenSize] = useState<"mobile" | "tablet" | "desktop">(
    "desktop"
  );

  // Responsive max items based on screen size
  useEffect(() => {
    const handleResize = () => {
      const width = window.innerWidth;
      if (width < 640) {
        setScreenSize("mobile");
      } else if (width < 1024) {
        setScreenSize("tablet");
      } else {
        setScreenSize("desktop");
      }
    };

    handleResize();
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  // Dynamic max items based on screen size
  const getMaxItems = () => {
    if (maxItems) return maxItems;

    switch (screenSize) {
      case "mobile":
        return 3;
      case "tablet":
        return 5;
      case "desktop":
        return 8;
      default:
        return 6;
    }
  };

  const dynamicMaxItems = getMaxItems();

  // Limit items if too many
  const displayItems =
    items.length > dynamicMaxItems
      ? [
          items[0], // Always show first item (Home)
          { id: "ellipsis", label: "...", isClickable: false },
          ...items.slice(-Math.max(dynamicMaxItems - 2, 1)), // Show last N items
        ]
      : items;

  // Z-index mapping for proper Tailwind classes
  const getZIndex = (index: number) => {
    const zLevel = Math.max(10 - index, 1);
    const zIndexMap: Record<number, string> = {
      10: "z-[10]",
      9: "z-[9]",
      8: "z-[8]",
      7: "z-[7]",
      6: "z-[6]",
      5: "z-[5]",
      4: "z-[4]",
      3: "z-[3]",
      2: "z-[2]",
      1: "z-[1]",
    };
    return zIndexMap[zLevel] || "z-[1]";
  };

  const getItemStyles = (
    index: number,
    isActive: boolean,
    isClickable: boolean
  ) => {
    const baseStyles = cn(
      "relative transition-all duration-200 rounded-r-3xl",
      // Responsive padding
      "px-3 py-1.5 sm:px-4 sm:py-2 lg:px-6 lg:py-2.5",
      getZIndex(index),
      index > 0 && "-ml-2 sm:-ml-3 lg:-ml-4"
    );

    let colorStyles = "";
    if (isActive) {
      colorStyles = "bg-accent-500 text-white shadow-lg";
    } else if (isClickable) {
      colorStyles = cn(
        "bg-secondary-100 text-accent-500 cursor-pointer",
        "hover:bg-secondary-200 active:bg-secondary-300",
        "transform hover:scale-105 active:scale-95"
      );
    } else {
      colorStyles = "bg-accent-100 text-accent-600";
    }

    return cn(baseStyles, colorStyles);
  };

  const handleItemClick = (item: BreadcrumbItem) => {
    if (!item.isClickable && !item.onClick) return;

    if (item.onClick) {
      item.onClick();
    } else if (item.href) {
      window.location.href = item.href;
    }
  };

  return (
    <nav
      className={cn(
        "flex items-center",
        // Responsive container
        "overflow-x-auto scrollbar-hide",
        "w-full max-w-full",
        className
      )}
      aria-label="Breadcrumb navigation"
    >
      <div className="flex items-center min-w-max">
        {displayItems.map((item, index) => {
          const isClickable = Boolean(
            item.isClickable !== false && (item.onClick || item.href)
          );
          const isActive = Boolean(item.isActive);

          return (
            <div
              key={item.id}
              className={cn(
                getItemStyles(index, isActive, isClickable),
                "breadcrumb-item"
              )}
              onClick={() => handleItemClick(item)}
              role={isClickable ? "button" : undefined}
              tabIndex={isClickable ? 0 : undefined}
              onKeyDown={(e) => {
                if (isClickable && (e.key === "Enter" || e.key === " ")) {
                  e.preventDefault();
                  handleItemClick(item);
                }
              }}
              aria-current={isActive ? "page" : undefined}
              aria-label={`${isActive ? "Current page: " : ""}${item.label}`}
            >
              <span
                className={cn(
                  "relative font-medium truncate",
                  // Responsive text sizes
                  "text-xs sm:text-sm",
                  // Responsive max widths
                  "max-w-16 sm:max-w-24 lg:max-w-32"
                )}
                title={item.label} // Tooltip for truncated text
              >
                {item.label}
              </span>

              {/* Hover effect overlay */}
              {isClickable && (
                <div className="absolute inset-0 bg-white/10 rounded-r-3xl opacity-0 hover:opacity-100 transition-opacity duration-200 pointer-events-none" />
              )}

              {/* Active item glow effect */}
              {isActive && (
                <div className="absolute inset-0 bg-gradient-to-r from-accent-400/20 to-accent-600/20 rounded-r-3xl pointer-events-none" />
              )}
            </div>
          );
        })}
      </div>

      {/* Fade out effect for overflow */}
      <div className="absolute right-0 top-0 bottom-0 w-8 bg-gradient-to-l from-[#0f0f0f] to-transparent pointer-events-none lg:hidden" />
    </nav>
  );
};

export default Breadcrumb;
