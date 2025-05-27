import React from "react";

type ButtonProps = React.HTMLAttributes<HTMLButtonElement> & {};

export const Button = ({ children, className, ...props }: ButtonProps) => {
  return (
    <button
      className={`px-4 py-2 transition-all duration-300 cursor-pointer rounded-md outline-1 outline-accent-500 bg-secondary-100 text-accent-500 hover:text-secondary-100  hover:bg-accent-500 ${className}`}
      {...props}
    >
      {children}
    </button>
  );
};
