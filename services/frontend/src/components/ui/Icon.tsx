import { FC } from "react";
import { IconType } from "react-icons";
import * as FaIcons from "react-icons/fa";
import * as GoIcons from "react-icons/go";

const projectIcons: Record<string, IconType> = {
  android: FaIcons.FaAndroid,
  apple: FaIcons.FaApple,
  server: FaIcons.FaServer,
  desktop: FaIcons.FaDesktop,
  browser: GoIcons.GoBrowser,
  docker: FaIcons.FaDocker,
};

export const getAllProjectIcons = (): string[] => Object.keys(projectIcons);

interface Props {
  name: string;
  size?: number;
  color?: string;
  className?: string;
}

export const Icon: FC<Props> = ({
  className,
  name,
  size = 24,
  color = "black",
}) => {
  const IconComponent = projectIcons[name];
  if (!IconComponent) return <span>❓</span>;
  return <IconComponent size={size} color={color} className={className} />;
};
