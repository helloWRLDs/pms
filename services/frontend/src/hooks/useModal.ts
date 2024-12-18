import React, { useRef, useState } from "react";

// returns: isVisible, toggle, close, open
export const useModal = (initialState = false) => {
  const [visible, setVisible] = useState(initialState);

  const open = () => {
    setVisible(true);
  };

  const close = () => {
    setVisible(false);
  };

  const toggle = () => {
    setVisible((prev) => !prev);
  };
  return [visible, toggle, close, open] as const;
};

export const useOutsideClick = (callback: Function) => {
  const ref = useRef<HTMLDivElement | null>(null);

  React.useEffect(() => {
    const handleClick = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        callback();
      }
    };

    document.addEventListener("click", handleClick, true);

    return () => {
      document.removeEventListener("click", handleClick, true);
    };
  }, []);

  return ref;
};
