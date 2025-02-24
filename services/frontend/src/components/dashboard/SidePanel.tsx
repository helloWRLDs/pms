import { ComponentProps, FC } from "react";

const SidePanel: FC<ComponentProps<"div">> = (props) => {
  return (
    <div
      {...props}
      className={`transition-all duration-300 ease-in-out bg-primary-500 shadow-lg text-neutral-100 min-h-dvh top-1/9 ${
        props.className || ""
      }`}
    >
      {props.children}
    </div>
  );
};

export default SidePanel;
