import { FC } from "react";
import { Icon } from "../ui/Icon";
import { Task } from "../../lib/task";

interface Props {
  task: Task;
  onClick?: () => void;
}

const TaskCard: FC<Props> = (props: Props) => {
  return (
    <div
      className="bg-primary-100 text-white rounded-md p-2 flex flex-wrap justify-between flex-col h-52 cursor-grab hover:bg-accent-700 transition-all ease-in-out duration-300 group"
      onClick={props.onClick}
      style={{
        backgroundImage:
          "repeating-linear-gradient(315deg, oklab(1 0 5.96046e-8 / 0.1) 0px, oklab(1 0 5.96046e-8 / 0.1) 1px, rgba(0, 0, 0, 0) 0px, rgba(0, 0, 0, 0) 50%)",
      }}
    >
      <h4 className="font-bold text-xl text-muted-700 group-hover:text-secondary-500 transition-">
        {props.task.title}
      </h4>
      <p className="font-normal text-muted-700 group-hover:text-secondary-500 transition-colors">
        {props.task.body}
      </p>
      <div className="flex flex-row justify-between float-end text-xs">
        <p className="font-normal text-muted-700 group-hover:text-secondary-500 transition-colors">
          {props.task.executor}
        </p>
        <p className="font-normal text-muted-700 group-hover:text-secondary-500 transition-colors">
          status: {props.task.status}
        </p>
        <div className="flex items-center">
          <Icon
            name={props.task.backlog_id ?? ""}
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
