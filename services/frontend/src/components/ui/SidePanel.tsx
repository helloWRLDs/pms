import { FC, ReactNode, useState } from "react";
import { GiHamburgerMenu } from "react-icons/gi";

interface Props {
  isOpen?: boolean;
  closedContent: ReactNode;
  openContent: ReactNode;
  className?: string;
}

const SidePanel: FC<Props> = (props) => {
  const [isOpen, setIsOpen] = useState(props.isOpen ?? false);

  return (
    <div
      className={`transition-all duration-300 ease-in-out bg-primary-500 shadow-lg text-neutral-100 min-h-dvh flex flex-col overflow-y-hidden ${
        isOpen ? "w-1/5" : "w-[5%]"
      } ${props.className}`}
    >
      {/* Content */}
      <div className="flex flex-col w-full overflow-auto min-h-[85svh]">
        {isOpen ? props.openContent : props.closedContent}
      </div>

      {/* Fixed Burger Button */}
      <div className="flex justify-center">
        <GiHamburgerMenu
          size={30}
          className="cursor-pointer"
          onClick={() => setIsOpen(!isOpen)}
          color="white"
        />
      </div>
    </div>
  );
};

export default SidePanel;
