import { FC } from "react";

type BadgeProps = React.HTMLAttributes<HTMLDivElement> & {
  className?: string;
};

const Badge: FC<BadgeProps> = ({ className, children, ...rest }) => {
  return (
    <div className={`rounded-xl text-sm py-1 px-3 w-fit ${className}`}>
      {children}
    </div>
  );
};

export default Badge;
