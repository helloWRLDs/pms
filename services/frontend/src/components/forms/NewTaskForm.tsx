import { useEffect, useState } from "react";
import { TaskCreation } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";
import { UserOptional } from "../../lib/user/user";
import { Sprint } from "../../lib/sprint/sprint";
import { Project } from "../../lib/project/project";

type Options = {
  className?: string;
  assignees?: UserOptional[];
  sprints?: Sprint[];
  project: Project;
  onSubmit: (task: TaskCreation) => void;
};

const PRIORITIES = [1, 2, 3, 4, 5];

const NewTaskForm = (opts: Options) => {
  const [task, setTask] = useState<TaskCreation>({
    title: "",
    body: "",
    status: getTaskStatuses[0],
    priority: 0,
    assignee_id: "",
    sprint_id: "",
    project_id: opts.project.id,
    due_date: {
      seconds: new Date().getTime(),
    },
  });

  useEffect(() => {
    console.log(task);
  }, [task]);

  const handleChange = (field: keyof TaskCreation, value: any) => {
    setTask({ ...task, [field]: value });
  };

  return (
    <>
      <div
        className={`mx-auto w-full p-6 space-y-6 bg-white rounded shadow ${opts.className}`}
      >
        <div className="space-y-2">
          <label htmlFor="task-title" className="block font-semibold">
            Title
          </label>
          <input
            type="text"
            id="task-title"
            placeholder="Task Title"
            className="w-full p-2 border rounded"
            value={task.title}
            onChange={(e) => handleChange("title", e.target.value)}
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="task-body" className="block font-semibold">
            Content
          </label>
          <textarea
            id="task-body"
            placeholder="Task Content"
            className="w-full p-2 border rounded"
            value={task.body}
            onChange={(e) => handleChange("body", e.target.value)}
          />
        </div>

        <div className="space-y-2">
          <label htmlFor="task-status" className="block font-semibold">
            Status
          </label>
          <select
            id="task-status"
            className="w-full p-2 border rounded"
            value={task.status}
            onChange={(e) => handleChange("status", e.target.value)}
          >
            {getTaskStatuses.map((status) => (
              <option key={status} value={status}>
                {capitalize(status.replace("_", " "))}
              </option>
            ))}
          </select>
        </div>

        <div className="space-y-2">
          <label htmlFor="task-priority" className="block font-semibold">
            Priority
          </label>
          <select
            id="task-priority"
            className="w-full p-2 border rounded"
            value={task.priority}
            onChange={(e) => handleChange("priority", parseInt(e.target.value))}
          >
            {PRIORITIES.map((priority) => (
              <option key={priority} value={priority}>
                {priority}
              </option>
            ))}
          </select>
        </div>

        <div className="space-y-2">
          <label htmlFor="task-assignee" className="block font-semibold">
            Assignee
          </label>
          <select
            id="task-assignee"
            className="w-full p-2 border rounded"
            value={task.assignee_id}
            onChange={(e) => handleChange("assignee_id", e.target.value)}
          >
            <option value="">Unassigned</option>
            {opts.assignees &&
              opts.assignees.map((assignee) => (
                <option key={assignee.id} value={assignee.id}>
                  {assignee.name}
                </option>
              ))}
          </select>
        </div>

        <div className="space-y-2">
          <label htmlFor="task-due-date" className="block font-semibold">
            Due Date
          </label>
          <input
            type="date"
            id="task-due-date"
            className="w-full p-2 border rounded"
            value={new Date(task.due_date.seconds).toISOString().split("T")[0]}
            onChange={(e) => {
              const date = new Date(e.target.value);
              setTask({ ...task, due_date: { seconds: date.getTime() } });
            }}
          />
        </div>

        {opts.sprints && opts.sprints.length > 0 && (
          <div className="space-y-2">
            <label htmlFor="task-sprint" className="block font-semibold">
              Sprint
            </label>
            <select
              id="task-sprint"
              className="w-full p-2 border rounded"
              value={task.sprint_id}
              onChange={(e) => handleChange("sprint_id", e.target.value)}
            >
              <option value="">No Sprint</option>
              {opts.sprints.map((sprint) => (
                <option key={sprint.id} value={sprint.id}>
                  {sprint.title}
                </option>
              ))}
            </select>
          </div>
        )}

        <div className="space-y-2">
          <label htmlFor="task-project" className="block font-semibold">
            Project
          </label>
          <input
            type="text"
            id="task-project"
            className="w-full p-2 border rounded bg-gray-100"
            disabled
            value={opts.project.title}
          />
        </div>

        <button
          type="submit"
          className="w-full bg-primary-400 text-white py-2 rounded hover:bg-accent-600 hover:text-primary-400 transition cursor-pointer"
          onClick={() => {
            task.due_date.seconds = Math.round(task.due_date.seconds / 1000);
            opts.onSubmit(task);
          }}
        >
          Create Task
        </button>
      </div>
    </>
  );
};

export default NewTaskForm;
