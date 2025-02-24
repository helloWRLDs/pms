import { FC } from "react";
// import GroupIcon from "../../assets/icons/group_icon.svg?react";
import { Icon } from "../ui/Icon";

interface Props {
  title?: string;
  text: string;
  executor?: string;
  backlog?: string;
  status?: string;
}

const TaskCard: FC<Props> = (props: Props) => {
  return (
    <div
      className="bg-neutral-500 text-white rounded-md p-2 flex flex-wrap justify-between flex-col h-52 cursor-grab"
      style={{
        backgroundImage:
          "repeating-linear-gradient(315deg, oklab(1 0 5.96046e-8 / 0.1) 0px, oklab(1 0 5.96046e-8 / 0.1) 1px, rgba(0, 0, 0, 0) 0px, rgba(0, 0, 0, 0) 50%)",
      }}
    >
      <h4 className="font-bold text-xl">{props.title}</h4>
      <p className="font-normal">{props.text}</p>
      <div className="flex flex-row justify-between float-end text-xs">
        <p>{props.executor}</p>
        <p>status: {props.status}</p>
        <div className="flex items-center">
          <Icon
            name={props.backlog ?? ""}
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
