import { useState } from "react";
import { SprintCreation } from "../../lib/sprint/sprint";
import { useProjectStore } from "../../store/selectedProjectStore";

type NewSprintFormProps = {
  className?: string;
  onSubmit: (data: SprintCreation) => void;
};

const TASKS = [
  {
    id: "d961d6bf-b298-4f86-9e5a-33d7f4f797ac",
    title: "registration module",
    body: "make functional registration module",
    status: "PENDING",
    project_id: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
    priority: 3,
    created_at: { seconds: 1747588448, nanos: 475379000 },
    updated_at: { seconds: 1747570448, nanos: 475272000 },
    due_date: { seconds: 1747570377600 },
    code: "front-1",
  },
  {
    id: "04aa3fdd-0a1d-4019-bb5c-2e285a5b670b",
    title: "Login Page",
    body: "make responsive nice login page",
    status: "CREATED",
    project_id: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
    priority: 1,
    created_at: { seconds: 1747588221, nanos: 604845000 },
    updated_at: { seconds: 1747570221, nanos: 604250000 },
    due_date: { seconds: 1747570204800 },
    code: "front-4",
  },
  {
    id: "1295e7fd-8738-424f-af4e-f7acaee8fdba",
    title: "nginx server",
    body: "configure nginx server",
    status: "CREATED",
    project_id: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
    priority: 2,
    created_at: { seconds: 1747258074, nanos: 925439000 },
    updated_at: { seconds: 1747240074, nanos: 924797000 },
    due_date: { seconds: 1747180800 },
    code: "front-3",
  },
  {
    id: "91700b42-9bf0-43da-9ff4-75d61e1c809c",
    title: "vite configuration",
    body: "Setup vite configuration and tailwind styles",
    status: "DONE",
    project_id: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
    priority: 3,
    created_at: { seconds: 1747196829, nanos: 22786000 },
    updated_at: { seconds: 1747196829, nanos: 22786000 },
    due_date: { seconds: -62135596800 },
    code: "front-2",
  },
];

const NewSprintForm = ({
  onSubmit,
  className,
  ...props
}: NewSprintFormProps) => {
  const { project: selectedProject } = useProjectStore();

  const NULL_SPRINT: SprintCreation = {
    title: "",
    description: "",
    start_date: {
      seconds: new Date().getTime() / 1000,
    },
    end_date: {
      seconds: new Date().getTime() / 1000,
    },
    project_id: selectedProject?.id ?? "",
    tasks: [],
  };

  const [newSprint, setNewSprint] = useState<SprintCreation>(NULL_SPRINT);
  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        onSubmit(newSprint);
      }}
      className="text-black text-lg"
      {...props}
    >
      <div className="relative z-0 mb-4 mt-8">
        <input
          type="text"
          value={newSprint.title}
          id="new-sprint-title"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none  dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewSprint({ ...newSprint, title: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-sprint-title"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Title
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <textarea
          value={newSprint.description}
          id="new-sprint-title"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none  dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewSprint({ ...newSprint, description: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-sprint-title"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Description
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <input
          type="date"
          id="new-sprint-start-date"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none  dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          value={
            new Date(newSprint.start_date.seconds * 1000)
              .toISOString()
              .split("T")[0]
          }
          onChange={(e) => {
            setNewSprint({
              ...newSprint,
              start_date: {
                seconds: new Date(e.currentTarget.value).getTime() / 1000,
              },
            });
          }}
        />
        <label
          htmlFor="new-sprint-start-date"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Start Date
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <input
          type="date"
          id="new-sprint-start-date"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none  dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          value={
            new Date(newSprint.end_date.seconds * 1000)
              .toISOString()
              .split("T")[0]
          }
          onChange={(e) => {
            setNewSprint({
              ...newSprint,
              end_date: {
                seconds: new Date(e.currentTarget.value).getTime() / 1000,
              },
            });
          }}
        />
        <label
          htmlFor="new-sprint-start-date"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          End date
        </label>
      </div>

      <div className="relative z-0 mb-4 mt-8">
        <select
          id="new-task-assignee"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none  dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          value={newSprint.tasks.length}
          //   onChange={(e) => handleChange("assignee_id", e.currentTarget.value)}
        >
          <option value="">No tasks</option>
          {TASKS.map((task) => (
            <option key={task.id} value={task.id}>
              {`${task.code} - ${task.title}`}
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

      <div className="mx-auto w-fit">
        <input
          type="submit"
          value="Create"
          className="px-4 py-2 bg-primary-500 text-accent-400 rounded-md cursor-pointer hover:bg-accent-500 hover:text-primary-500"
        />
      </div>
    </form>
  );
};

export default NewSprintForm;
