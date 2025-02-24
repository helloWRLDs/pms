import { useState, FC, ComponentProps } from "react";
import {
  DragDropContext,
  Droppable,
  Draggable,
  DropResult,
} from "@hello-pangea/dnd";
import { Task } from "../../lib/task";
import TaskCard from "./TaskCard";
import { capitalize } from "../../utils/string";
import { toast } from "react-toastify";
import { toastOpts } from "../../utils/toast";

// interface Props {
//   tasks: Task[];
// }

const testTasks = [
  {
    id: "1",
    title: "Fix Home Page on UI",
    text: "NavBar Footer Body DropDowns",
    executor: "John",
    backlog: "browser",
    status: "created",
  },
  {
    id: "2",
    title: "Layout according to sample",
    text: "Make initial layout for key pages.",
    executor: "John",
    backlog: "browser",
    status: "in progress",
  },
  {
    id: "3",
    title: "Proper CI/CD Pipelines",
    text: "Should include tests in GitHub Actions...",
    executor: "Bob",
    backlog: "docker",
    status: "created",
  },
  {
    id: "4",
    title: "Containerization",
    text: "Cover all the services with Dockerfiles.Test all of them before submitting",
    executor: "Alice",
    backlog: "docker",
    status: "done",
  },
];

interface Task {
  id: string;
  title: string;
  text: string;
  executor: string;
  backlog: string;
  status: string;
}

const Dashboard: FC<ComponentProps<"table">> = (props) => {
  const [tasks, setTasks] = useState<Task[]>(testTasks);

  const onDragEnd = (result: DropResult) => {
    const { source, destination } = result;

    if (!destination) return; // If dropped outside, do nothing
    if (destination.droppableId === source.droppableId) return; // If dropped in the same

    // Find the task being dragged
    const taskId = result.draggableId;
    const taskIndex = tasks.findIndex((task) => task.id === taskId);
    if (taskIndex === -1) return;

    // Update the task's status based on the new column
    const updatedTasks = [...tasks];
    updatedTasks[taskIndex] = {
      ...tasks[taskIndex],
      status: destination.droppableId,
    };

    setTasks(updatedTasks);
    toast.success(
      `Status for task with id = ${taskId} changed to ${destination.droppableId}`,
      { ...toastOpts, autoClose: 2000 }
    );
  };
  return (
    <DragDropContext onDragEnd={onDragEnd}>
      <table
        className={`w-full table-fixed border-collapse bg-primary-500 text-neutral-100 ${props.className}`}
      >
        <thead>
          <tr className="bg-primary-400 text-neutral-100">
            {Task.STATUSES.map((status) => (
              <th
                key={status}
                className="border border-primary-300 p-4 text-lg font-semibold"
              >
                {capitalize(Task.toString(status))}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="text-white border-white">
          <tr>
            {Task.STATUSES.map((status) => (
              <td
                key={status}
                className="border border-primary-300 p-4 align-top w-1/4"
              >
                <Droppable droppableId={Task.toString(status)}>
                  {(provided) => (
                    <div
                      ref={provided.innerRef}
                      {...provided.droppableProps}
                      className="min-h-[100px] p-2"
                    >
                      {/* Filter and map tasks based on status */}
                      {tasks
                        .filter(
                          (task) => Task.fromString(task.status) === status
                        )
                        .map((task, index) => (
                          <Draggable
                            key={task.id}
                            draggableId={task.id}
                            index={index}
                          >
                            {(provided) => (
                              <div
                                ref={provided.innerRef}
                                {...provided.draggableProps}
                                {...provided.dragHandleProps}
                                className="mb-2 p-2 rounded shadow cursor-grab"
                              >
                                <TaskCard {...task} />
                              </div>
                            )}
                          </Draggable>
                        ))}
                      {provided.placeholder}
                    </div>
                  )}
                </Droppable>
              </td>
            ))}
          </tr>
        </tbody>
      </table>
    </DragDropContext>
  );
};

export default Dashboard;
