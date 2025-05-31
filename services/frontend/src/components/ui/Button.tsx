import React from "react";

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: "default" | "outline" | "ghost";
};

export const Button = ({
  children,
  className,
  variant = "default",
  disabled,
  ...props
}: ButtonProps) => {
  const baseStyles =
    "px-4 py-2 transition-all duration-300 cursor-pointer rounded-lg outline-1 outline-accent-500";

  const variantStyles = {
    default: "bg-accent-500 text-white hover:bg-accent-600",
    outline:
      "border border-accent-500 text-accent-500 hover:bg-accent-500 hover:text-white",
    ghost: "text-accent-500 hover:bg-accent-500/10",
  };

  return (
    <button
      className={`
        ${baseStyles}
        ${variantStyles[variant]}
        ${disabled ? "opacity-50 cursor-not-allowed" : ""}
        ${className}
      `}
      disabled={disabled}
      {...props}
    >
      {children}
    </button>
  );
};
