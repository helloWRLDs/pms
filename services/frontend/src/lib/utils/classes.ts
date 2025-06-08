type Size = "sm" | "md" | "lg" | "xl" | "2xl" | "full";

export const getSizeClass = (size: Size) => {
  switch (size) {
    case "sm":
      return "text-sm";
    case "md":
      return "text-md";
    case "lg":
      return "text-lg";
    case "xl":
      return "text-xl";
    case "2xl":
      return "text-2xl";
    case "full":
      return "max-w-full";
  }
};
