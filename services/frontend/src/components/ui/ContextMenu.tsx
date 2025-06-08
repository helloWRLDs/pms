import { Menu, MenuButton, MenuItem, MenuItems } from "@headlessui/react";
import { cn } from "../../lib/utils/cn";
import { ReactNode, useRef, useState } from "react";
import { BsThreeDotsVertical } from "react-icons/bs";

export interface ContextMenuItemProps {
  /** Icon to display next to the label */
  icon?: ReactNode;
  /** Text label for the menu item */
  label: string;
  /** Click handler for the menu item */
  onClick: (e: React.MouseEvent) => void;
  /** Additional CSS classes */
  className?: string;
  /** Whether the item is disabled */
  disabled?: boolean;
  /** Visual variant for the menu item */
  variant?: "default" | "danger" | "warning";
}

interface ContextMenuProps {
  trigger?: ReactNode;
  items: ContextMenuItemProps[];
  className?: string;
  menuClassName?: string;
  placement?: "left" | "right";
}

export const ContextMenu = ({
  trigger,
  items,
  className,
  menuClassName,
  placement = "right",
}: ContextMenuProps) => {
  const buttonRef = useRef<HTMLButtonElement>(null);
  const [buttonRect, setButtonRect] = useState<DOMRect | null>(null);

  const getVariantStyles = (
    variant: ContextMenuItemProps["variant"] = "default"
  ) => {
    switch (variant) {
      case "danger":
        return "text-red-400 hover:text-red-300 hover:bg-red-500/20 focus:bg-red-500/30 active:bg-red-500/40";
      case "warning":
        return "text-yellow-400 hover:text-yellow-300 hover:bg-yellow-500/20 focus:bg-yellow-500/30 active:bg-yellow-500/40";
      default:
        return "text-neutral-100 hover:text-white hover:bg-secondary-200 focus:bg-secondary-200 active:bg-secondary-300";
    }
  };

  trigger = trigger ?? <BsThreeDotsVertical />;

  const handleMenuToggle = (e: React.MouseEvent) => {
    e.stopPropagation();
    if (buttonRef.current) {
      setButtonRect(buttonRef.current.getBoundingClientRect());
    }
  };

  const getMenuStyle = () => {
    if (!buttonRect) return {};

    const viewportWidth = window.innerWidth;
    const menuWidth = 192;

    let left = buttonRect.left;
    let top = buttonRect.bottom + 4;

    if (left + menuWidth > viewportWidth) {
      left = buttonRect.right - menuWidth;
    }

    if (left < 8) {
      left = 8;
    }

    return {
      position: "fixed" as const,
      left: `${left}px`,
      top: `${top}px`,
      zIndex: 9999,
    };
  };

  return (
    <Menu as="div" className={cn("relative inline-block", className)}>
      {trigger && (
        <MenuButton
          ref={buttonRef}
          className={cn(
            "text-neutral-300 hover:text-accent-400 hover:bg-secondary-100 focus:outline-none focus:ring-2 focus:ring-accent-500/50 focus:ring-offset-2 focus:ring-offset-secondary-200",
            "group p-2 rounded-md transition-all duration-200 ease-in-out hover:scale-105 active:scale-95 cursor-pointer"
          )}
          onClick={handleMenuToggle}
        >
          {({ open }) => (
            <span
              className={cn(
                "block transition-transform duration-200 ease-in-out ",
                open ? "rotate-90" : "rotate-0"
              )}
            >
              {trigger}
            </span>
          )}
        </MenuButton>
      )}

      <MenuItems
        style={getMenuStyle()}
        className={cn(
          "w-48 bg-secondary-100 rounded-lg shadow-xl border border-secondary-50 py-1 focus:outline-none",
          "transform transition-all duration-200 ease-out",
          "data-[closed]:scale-95 data-[closed]:opacity-0 data-[open]:scale-100 data-[open]:opacity-100",
          "hover:shadow-2xl hover:border-secondary-100",
          menuClassName
        )}
      >
        {items.map((item, index) => (
          <MenuItem key={index} disabled={item.disabled}>
            {({ active, focus }) => (
              <button
                className={cn(
                  "flex items-center w-full px-4 py-2 text-sm gap-3 transition-all duration-200 ease-in-out",
                  "group relative overflow-hidden cursor-pointer",
                  active || focus
                    ? "bg-secondary-200 transform translate-x-1"
                    : "",
                  getVariantStyles(item.variant),
                  item.disabled &&
                    "opacity-50 cursor-not-allowed hover:bg-transparent hover:text-neutral-400",
                  item.className
                )}
                onClick={(e) => {
                  e.stopPropagation();
                  if (!item.disabled) {
                    item.onClick(e);
                  }
                }}
                disabled={item.disabled}
              >
                {/* Hover background effect */}
                <span
                  className={cn(
                    "absolute inset-0 bg-gradient-to-r from-accent-500/10 to-transparent",
                    "transform transition-transform duration-300 ease-out",
                    active || focus ? "translate-x-0" : "-translate-x-full"
                  )}
                />

                {item.icon && (
                  <span
                    className={cn(
                      "flex-shrink-0 text-accent-500 transition-all duration-200",
                      active || focus ? "text-accent-400 scale-110" : "",
                      item.disabled && "text-neutral-500"
                    )}
                  >
                    {item.icon}
                  </span>
                )}

                <span
                  className={cn(
                    "truncate relative z-10 transition-all duration-200",
                    active || focus ? "font-medium" : ""
                  )}
                >
                  {item.label}
                </span>

                {/* Ripple effect on hover */}
                <span
                  className={cn(
                    "absolute right-0 top-1/2 w-2 h-2 bg-accent-500/30 rounded-full transform -translate-y-1/2 transition-all duration-300 ease-out",
                    active || focus
                      ? "scale-50 opacity-100"
                      : "scale-0 opacity-0"
                  )}
                />
              </button>
            )}
          </MenuItem>
        ))}

        {/* Menu container hover glow effect */}
        <div className="absolute inset-0 rounded-lg bg-gradient-to-br from-accent-500/5 to-transparent opacity-0 hover:opacity-100 transition-opacity duration-300 pointer-events-none" />
      </MenuItems>
    </Menu>
  );
};

// Export individual components for more granular control if needed
export const ContextMenuButton = MenuButton;
export const ContextMenuItems = MenuItems;
export const ContextMenuItem = MenuItem;
