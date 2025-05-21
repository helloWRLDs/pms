import { useEffect, useRef } from "react";

type DropDownListProps = React.HTMLAttributes<HTMLDivElement> & {
  visible: boolean;
  onClose?: () => void;
};
const DropDownList = ({
  children,
  className,
  visible = false,
  onClose,
  ...props
}: DropDownListProps) => {
  const ref = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (!visible) return;

    const handleClickOutside = (event: MouseEvent) => {
      if (ref.current && !ref.current.contains(event.target as Node)) {
        onClose?.();
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [visible, onClose]);

  if (!visible) return null;
  return (
    <div ref={ref} className={`absolute z-50 w-fit ${className}`} {...props}>
      {children}
    </div>
  );
};

type DropDownListElementProps = React.HTMLAttributes<HTMLDivElement> & {};
const DropDownListElement = ({
  children,
  className,
  ...props
}: DropDownListElementProps) => {
  return (
    <div className={`${className}`} {...props}>
      {children}
    </div>
  );
};

DropDownList.Element = DropDownListElement;

export default DropDownList;
