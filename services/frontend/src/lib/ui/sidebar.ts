import { IconType } from "react-icons";

export type SideBarItem = {
  className?: string;
  isEnabled: boolean;
  label: string;
  icon: IconType;
  onClick?: () => void;
  badge?: JSX.Element;
};
