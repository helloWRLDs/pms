import { FC } from "react";
import { Icon } from "../ui/Icon";
import { Task } from "../../lib/task/task";
import { useCacheStore } from "../../store/cacheStore";
import { capitalize } from "../../lib/utils/string";
import { formatTime } from "../../lib/utils/time";
import { MdAccessTime, MdPerson, MdLabel } from "react-icons/md";
import { Priority } from "../../lib/task/priority";
import Badge from "../ui/Badge";

interface Props {
  task: Task;
  onClick?: () => void;
  isDragging?: boolean;
}

const TaskCard: FC<Props> = ({ task, onClick, isDragging }) => {
  const { getAssignee } = useCacheStore();
  const assignee = task.assignee_id ? getAssignee(task.assignee_id) : null;
  const priority = new Priority(task.priority);

  return (
    <div
      className={`
        bg-secondary-100 rounded-lg p-4 
        flex flex-col gap-3
        cursor-grab active:cursor-grabbing
        hover:bg-secondary-50 
        transition-all duration-200
        border border-secondary-200
        ${isDragging ? "shadow-lg scale-105" : "hover:shadow-md"}
      `}
      onClick={onClick}
    >
      {/* Task Code and Priority */}
      <div className="flex justify-between items-start">
        <span className="text-xs font-mono text-neutral-400">{task.code}</span>
        <Badge className={`text-${priority.getColor()}-500`}>
          {priority.toString()}
        </Badge>
      </div>

      {/* Title */}
      <h3 className="text-base font-medium text-neutral-100 line-clamp-2">
        {task.title}
      </h3>

      {/* Task Info */}
      <div className="grid grid-cols-2 gap-2 text-xs text-neutral-400">
        {/* Assignee */}
        <div className="flex items-center gap-1.5">
          <MdPerson className="text-accent-500" />
          <span className="truncate">
            {assignee
              ? `${assignee.first_name} ${assignee.last_name}`
              : "Unassigned"}
          </span>
        </div>

        {/* Due Date */}
        <div className="flex items-center gap-1.5">
          <MdAccessTime className="text-accent-500" />
          <span className="truncate">{formatTime(task.due_date.seconds)}</span>
        </div>

        {/* Status */}
        <div className="flex items-center gap-1.5">
          <MdLabel className="text-accent-500" />
          <span className="truncate">
            {capitalize(task.status.replace("_", " "))}
          </span>
        </div>

        {/* Project Icon */}
        {task.project_id && (
          <div className="flex items-center gap-1.5">
            <Icon
              name={task.project_id}
              size={16}
              className="text-accent-500"
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default TaskCard;
