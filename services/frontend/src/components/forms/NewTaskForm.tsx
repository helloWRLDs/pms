import { useEffect, useState } from "react";
import { TaskCreation } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";
import { UserOptional } from "../../lib/user/user";
import { Sprint } from "../../lib/sprint/sprint";
import { Project } from "../../lib/project/project";
import Input from "../ui/Input";
import { Priority } from "../../lib/task/priority";

type NewTaskFormProps = {
  className?: string;
  assignees?: UserOptional[];
  sprints?: Sprint[];
  project: Project;
  onSubmit?: (data: TaskCreation) => void;
};

const PRIORITIES = [1, 2, 3, 4, 5];

const NewTaskForm = ({ onSubmit, className, ...props }: NewTaskFormProps) => {
  const NULL_TASK: TaskCreation = {
    title: "",
    body: "",
    status: getTaskStatuses[0],
    priority: 0,
    assignee_id:
      props.assignees && props.assignees.length > 0
        ? props.assignees[0].id
        : "",
    sprint_id:
      props.sprints && props.sprints?.length > 0 ? props.sprints[0].id : "",
    project_id: props.project.id ?? "",
    due_date: {
      seconds: new Date().getTime(),
    },
  };
  const [newTask, setNewTask] = useState<TaskCreation>(NULL_TASK);

  useEffect(() => {
    console.log(newTask);
  }, [newTask]);

  const handleChange = (field: keyof TaskCreation, value: any) => {
    setNewTask({ ...newTask, [field]: value });
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        onSubmit && onSubmit(newTask);
        setNewTask(NULL_TASK);
      }}
      className={`mx-auto ${className}`}
      {...props}
    >
      <Input>
        <Input.Element
          type="text"
          label="Title"
          value={newTask.title}
          required={true}
          onChange={(e) => handleChange("title", e.currentTarget.value)}
        />
      </Input>

      <Input>
        <Input.Element
          type="textarea"
          label="Body"
          value={newTask.body}
          onChange={(e) => handleChange("body", e.currentTarget.value)}
        />
      </Input>

      <Input>
        <Input.Element
          type="select"
          label="Status"
          options={getTaskStatuses.map((item) => {
            return { label: capitalize(item.replace("_", " ")), value: item };
          })}
        />
      </Input>

      <Input>
        <Input.Element
          type="select"
          label="Priority"
          options={PRIORITIES.map((item) => {
            return {
              label: new Priority(item).toString(),
              value: item,
            };
          })}
          value={newTask.priority}
          onChange={(e) => {
            handleChange("priority", parseInt(e.currentTarget.value));
          }}
        />
      </Input>

      <Input>
        <Input.Element
          type="select"
          label="Assignee"
          options={[
            { label: "Unassigned", value: "" },
            ...(props.assignees?.map((item) => ({
              label: item.name ?? "",
              value: item.id ?? "",
            })) ?? []),
          ]}
        />
      </Input>

      <Input>
        <Input.Element
          label="Due date"
          type="date"
          value={
            new Date(newTask.due_date.seconds * 1000)
              .toISOString()
              .split("T")[0]
          }
          onChange={(e) =>
            handleChange("due_date", {
              seconds: new Date(e.currentTarget.value).getTime() / 1000,
            })
          }
        />
      </Input>

      {props.sprints && props.sprints.length > 0 && (
        <Input>
          <Input.Element
            type="select"
            label="Sprint"
            options={[
              { label: "No Sprint", value: "" },
              ...(props.sprints?.map((item) => ({
                label: item.title ?? "",
                value: item.id ?? "",
              })) ?? []),
            ]}
          />
        </Input>
      )}

      <Input>
        <Input.Element
          type="text"
          label="Project"
          value={props.project.code_name}
          disabled
        />
      </Input>

      <div className="w-1/2 mx-auto mt-8">
        <input
          type="submit"
          className="w-full px-4 py-2 bg-primary-400 text-white rounded hover:bg-accent-600 hover:text-primary-400 transition cursor-pointer"
          value="Create Task"
        />
      </div>
    </form>
  );
};

export default NewTaskForm;
