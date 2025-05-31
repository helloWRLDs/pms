import { FC } from "react";
import {
  TaskType,
  TaskTypes,
  getTaskTypes,
  getTaskTypeConfig,
} from "../../lib/task/tasktype";

interface TaskTypeSelectProps {
  value?: TaskType;
  onChange: (type: TaskType) => void;
  placeholder?: string;
  className?: string;
  disabled?: boolean;
}

const TaskTypeSelect: FC<TaskTypeSelectProps> = ({
  value,
  onChange,
  placeholder = "Select task type",
  className = "",
  disabled = false,
}) => {
  const taskTypes = getTaskTypes;

  return (
    <select
      value={value || ""}
      onChange={(e) => onChange(e.target.value as TaskType)}
      className={`w-full px-3 py-2 bg-secondary-400/20 border border-secondary-400/30 rounded-lg text-neutral-100 focus:outline-none focus:ring-2 focus:ring-accent-500 focus:border-transparent ${className}`}
      disabled={disabled}
    >
      <option value="" disabled>
        {placeholder}
      </option>
      {taskTypes.map((type) => {
        const config = getTaskTypeConfig(type);
        return (
          <option
            key={type}
            value={type}
            className="bg-secondary-600 text-neutral-100"
          >
            {config.icon} {config.label}
          </option>
        );
      })}
    </select>
  );
};

export default TaskTypeSelect;
