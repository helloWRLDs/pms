import { FiUser } from "react-icons/fi";

import { FiCalendar } from "react-icons/fi";
import { formatTime } from "../../lib/utils/time";
import Badge from "../ui/Badge";
import { capitalize } from "../../lib/utils/string";
import { Priority } from "../../lib/task/priority";
import { useSprintList } from "../../hooks/useData";
import useMetaCache from "../../store/useMetaCache";
import { useAssigneeList } from "../../hooks/useData";
import { Task } from "../../lib/task/task";

const TaskCard = ({ task, onClick }: { task: Task; onClick: () => void }) => {
  const metaCache = useMetaCache();
  const { getAssigneeName } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );
  const { getSprintName } = useSprintList(
    metaCache.metadata.selectedProject?.id ?? ""
  );

  return (
    <div
      onClick={onClick}
      className="bg-secondary-200 rounded-lg p-4 cursor-pointer hover:bg-secondary-100 transition-all duration-200 group"
    >
      <div className="flex justify-between items-start mb-3">
        <span className="text-sm font-mono text-neutral-400">{task.code}</span>
        <Badge className={`text-${new Priority(task.priority).getColor()}-500`}>
          {new Priority(task.priority).toString()}
        </Badge>
      </div>

      <h3 className="text-lg font-semibold mb-2 group-hover:text-accent-400 transition-colors">
        {task.title}
      </h3>

      <div className="flex items-center gap-2 mb-3">
        <Badge className="text-white bg-primary-400">
          {capitalize(task.status.replace("_", " "))}
        </Badge>
        {task.sprint_id && (
          <Badge className="bg-accent-500/10 text-accent-400">
            {getSprintName(task.sprint_id)}
          </Badge>
        )}
      </div>

      <div className="grid grid-cols-2 gap-3 text-sm text-neutral-300">
        <div className="flex items-center gap-2">
          <FiUser className="text-accent-500" />
          <span>{getAssigneeName(task.assignee_id)}</span>
        </div>
        <div className="flex items-center gap-2">
          <FiCalendar className="text-accent-500" />
          <span>{formatTime(task.due_date.seconds)}</span>
        </div>
      </div>
    </div>
  );
};

export default TaskCard;
