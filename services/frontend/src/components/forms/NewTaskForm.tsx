import { useEffect, useState } from "react";
import { TaskCreation } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";
import { UserOptional } from "../../lib/user/user";
import { Sprint } from "../../lib/sprint/sprint";
import { Project } from "../../lib/project/project";

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
      <div className="relative z-0 mb-4 mt-8">
        <input
          type="text"
          value={newTask.title}
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) => handleChange("title", e.currentTarget.value)}
        />
        <label
          htmlFor="new-project-description"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Title
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <textarea
          value={newTask.body}
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewTask({ ...newTask, body: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-task-body"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Body
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <select
          id="new-task-status"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          value={newTask.status}
          onChange={(e) => handleChange("status", e.currentTarget.value)}
        >
          {getTaskStatuses.map((status) => (
            <option key={status} value={status} className="text-black">
              {capitalize(status.replace("_", " "))}
            </option>
          ))}
        </select>
        <label
          htmlFor="new-task-status"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Status
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <select
          id="new-task-priority"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          value={newTask.priority}
          onChange={(e) =>
            handleChange("priority", parseInt(e.currentTarget.value))
          }
        >
          {PRIORITIES.map((priority) => (
            <option key={priority} value={priority}>
              {priority}
            </option>
          ))}
        </select>
        <label
          htmlFor="new-task-priority"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Priority
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <select
          id="new-task-assignee"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          value={newTask.assignee_id}
          onChange={(e) => handleChange("assignee_id", e.currentTarget.value)}
        >
          <option value="">Unassigned</option>
          {props.assignees &&
            props.assignees.map((assignee) => (
              <option key={assignee.id} value={assignee.id}>
                {assignee.name}
              </option>
            ))}
        </select>
        <label
          htmlFor="new-task-status"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Assignee
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <input
          type="date"
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          value={new Date(newTask.due_date.seconds).toISOString().split("T")[0]}
          onChange={(e) => {
            handleChange("due_date", {
              seconds: new Date(e.currentTarget.value).getTime() / 1000,
            });
          }}
        />
        <label
          htmlFor="new-project-description"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Due Date
        </label>
      </div>

      {props.sprints && props.sprints.length > 0 && (
        <div className="relative z-0 mb-4 mt-8">
          <select
            id="new-task-sprint"
            className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
            value={newTask.sprint_id}
            onChange={(e) => handleChange("sprint_id", e.currentTarget.value)}
          >
            <option value="">No Sprint</option>
            {props.sprints.map((sprint) => (
              <option key={sprint.id} value={sprint.id}>
                {sprint.title}
              </option>
            ))}
          </select>
          <label
            htmlFor="new-task-sprint"
            className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
          >
            Sprint
          </label>
        </div>
      )}

      <div className="relative z-0 mb-4 mt-8">
        <input
          type="text"
          value={props.project.code_name}
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          disabled
        />
        <label
          htmlFor="new-project-description"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Project
        </label>
      </div>

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
