import { FC } from "react";
import { Icon } from "../ui/Icon";
import { Task } from "../../lib/task/task";
import { useCacheStore } from "../../store/cacheStore";
import { capitalize } from "../../lib/utils/string";

interface Props {
  task: Task;
  onClick?: () => void;
}

const TaskCard: FC<Props> = (props: Props) => {
  const { getAssignee } = useCacheStore();
  return (
    <div
      className="bg-primary-100 text-white rounded-md p-2 flex flex-wrap justify-between flex-col min-h-[100px] cursor-grab hover:bg-accent-700 transition-all ease-in-out duration-300 group"
      onClick={props.onClick}
      style={{
        backgroundImage:
          "repeating-linear-gradient(315deg, oklab(1 0 5.96046e-8 / 0.1) 0px, oklab(1 0 5.96046e-8 / 0.1) 1px, rgba(0, 0, 0, 0) 0px, rgba(0, 0, 0, 0) 50%)",
      }}
    >
      <h4 className="font-bold text-xl text-muted-700 group-hover:text-secondary-500 transition-all duration-300">
        {props.task.title}
      </h4>
      <div className="flex flex-row justify-between float-end text-xs">
        <p className="font-normal text-muted-700 group-hover:text-secondary-500 transition-colors">
          {props.task.assignee_id
            ? getAssignee(props.task.assignee_id)?.name
            : "None"}
        </p>
        <p className="font-normal text-muted-700 group-hover:text-secondary-500 transition-colors">
          status: {capitalize(props.task.status).replace("_", " ")}
        </p>
        <div className="flex items-center">
          <Icon
            name={props.task.project_id ?? ""}
            size={20}
            color="white"
            className="mr-1"
          />
          {/* <p>{props.backlog}</p> */}
        </div>
      </div>
    </div>
  );
};

export default TaskCard;
