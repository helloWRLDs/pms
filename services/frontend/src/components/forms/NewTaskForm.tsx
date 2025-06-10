import { useEffect, useState } from "react";
import { TaskCreation } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";
import { Project } from "../../lib/project/project";
import Input from "../ui/Input";
import { Priority } from "../../lib/task/priority";
import { useAssigneeList, useSprintList } from "../../hooks/useData";
import { getTaskTypes } from "../../lib/task/tasktype";
import TiptapEditor from "../text/TiptapEditor";
import useMetaCache from "../../store/useMetaCache";

type NewTaskFormProps = {
  className?: string;
  project: Project;
  onSubmit: (data: TaskCreation) => void;
};

const PRIORITIES = [1, 2, 3, 4, 5];

const NewTaskForm = ({ onSubmit, className, ...props }: NewTaskFormProps) => {
  const metaCache = useMetaCache();
  const NULL_TASK: TaskCreation = {
    title: "",
    body: "",
    status: getTaskStatuses[0],
    priority: 1,
    assignee_id: "",
    sprint_id: "",
    project_id: props.project.id ?? "",
    due_date: {
      seconds: new Date().getTime(),
    },
    type: getTaskTypes[1],
  };
  const [newTask, setNewTask] = useState<TaskCreation>(NULL_TASK);
  const { assignees } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );
  const { sprints } = useSprintList(
    metaCache.metadata.selectedProject?.id ?? ""
  );

  useEffect(() => {
    console.log(newTask);
  }, [newTask]);

  const handleChange = (
    field: keyof TaskCreation,
    value: string | number | { seconds: number }
  ) => {
    setNewTask({ ...newTask, [field]: value });
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        const fixedTask: TaskCreation = {
          ...newTask,
          due_date: {
            seconds: Math.floor(newTask.due_date.seconds),
          },
        };
        onSubmit(fixedTask);
        setNewTask(NULL_TASK);
      }}
      className={`mx-auto space-y-4 ${className}`}
      {...props}
    >
      {/* Title - Full width */}
      <Input>
        <Input.Element
          type="text"
          label="Title"
          value={newTask.title}
          required={true}
          onChange={(e) => handleChange("title", e.currentTarget.value)}
        />
      </Input>

      {/* Body - Full width */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Body *
        </label>
        <div className="bg-white rounded-lg [&_.ProseMirror]:!text-gray-900 [&_.ProseMirror]:!bg-white [&_.ProseMirror_p]:!text-gray-900 [&_.ProseMirror_h1]:!text-gray-900 [&_.ProseMirror_h2]:!text-gray-900 [&_.ProseMirror_h3]:!text-gray-900 [&_.ProseMirror_li]:!text-gray-900 [&_.ProseMirror_ul]:!text-gray-900 [&_.ProseMirror_ol]:!text-gray-900">
          <TiptapEditor
            content={newTask.body}
            onChange={(content) => handleChange("body", content)}
            placeholder="Enter task description..."
            className="min-h-[150px] text-sm text-black "
          />
        </div>
      </div>

      {/* Type and Status - Flex row */}
      <div className="flex flex-col sm:flex-row gap-4">
        <div className="flex-1">
          <Input>
            <Input.Element
              type="select"
              label="Type"
              options={getTaskTypes.map((item) => {
                return {
                  label: capitalize(item.replace("_", " ")),
                  value: item,
                };
              })}
              value={newTask.type}
              onChange={(e) => handleChange("type", e.currentTarget.value)}
              required
            />
          </Input>
        </div>
        <div className="flex-1">
          <Input>
            <Input.Element
              type="select"
              label="Status"
              options={getTaskStatuses.map((item) => {
                return {
                  label: capitalize(item.replace("_", " ")),
                  value: item,
                };
              })}
              value={newTask.status}
              onChange={(e) => handleChange("status", e.currentTarget.value)}
              required
            />
          </Input>
        </div>
      </div>

      <div className="flex flex-col sm:flex-row gap-4">
        <div className="flex-1">
          <Input className="text-gray-700">
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
              required
            />
          </Input>
        </div>
        <div className="flex-1">
          <Input>
            <Input.Element
              type="select"
              label="Assignee"
              options={[
                { label: "Unassigned", value: "" },
                ...(assignees?.items?.map((item) => ({
                  label: `${item.first_name} ${item.last_name}`,
                  value: item.id ?? "",
                })) ?? []),
              ]}
              value={newTask.assignee_id}
              onChange={(e) =>
                handleChange("assignee_id", e.currentTarget.value)
              }
            />
          </Input>
        </div>
      </div>

      {/* Due date and Sprint - Flex row */}
      <div className="flex flex-col sm:flex-row gap-4">
        <div className="flex-1">
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
              required
            />
          </Input>
        </div>
        <div className="flex-1">
          {sprints && sprints.items && sprints.items.length > 0 ? (
            <Input>
              <Input.Element
                type="select"
                label="Sprint"
                options={[
                  { label: "No Sprint", value: "" },
                  ...(sprints && sprints.items && sprints.items.length > 0
                    ? sprints?.items?.map((item) => ({
                        label: item.title ?? "",
                        value: item.id ?? "",
                      }))
                    : []),
                ]}
                value={newTask.sprint_id}
                onChange={(e) =>
                  handleChange("sprint_id", e.currentTarget.value)
                }
              />
            </Input>
          ) : (
            <Input>
              <Input.Element
                type="text"
                label="Sprint"
                value="No sprints available"
                disabled
              />
            </Input>
          )}
        </div>
      </div>

      <Input>
        <Input.Element
          type="text"
          label="Project"
          value={props.project.code_name}
          disabled
        />
      </Input>

      <div className="flex justify-center pt-4">
        <input
          type="submit"
          className="w-full sm:w-auto px-8 py-3 bg-primary-400 text-white rounded-lg hover:bg-accent-600 hover:text-primary-400 transition-colors cursor-pointer font-medium"
          value="Create Task"
        />
      </div>
    </form>
  );
};

export default NewTaskForm;
