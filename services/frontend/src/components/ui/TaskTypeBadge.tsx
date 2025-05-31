import { FC } from "react";
import { TaskType, getTaskTypeConfig } from "../../lib/task/tasktype";

interface TaskTypeBadgeProps {
  type: TaskType;
  size?: "sm" | "md" | "lg";
  showIcon?: boolean;
  showLabel?: boolean;
  className?: string;
}

const TaskTypeBadge: FC<TaskTypeBadgeProps> = ({
  type,
  size = "md",
  showIcon = true,
  showLabel = true,
  className = "",
}) => {
  const config = getTaskTypeConfig(type);

  const sizeClasses = {
    sm: "px-2 py-1 text-xs",
    md: "px-3 py-1 text-sm",
    lg: "px-4 py-2 text-base",
  };

  return (
    <span
      className={`inline-flex items-center gap-1 rounded-full font-medium text-white ${sizeClasses[size]} ${className}`}
      style={{ backgroundColor: config.color }}
      title={config.description}
    >
      {showIcon && <span className="text-xs">{config.icon}</span>}
      {showLabel && <span>{config.label}</span>}
    </span>
  );
};

export default TaskTypeBadge;
