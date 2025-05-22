import {
  FC,
  useState,
  useLayoutEffect,
  ComponentProps,
  useEffect,
} from "react";
import {
  DragDropContext,
  Droppable,
  Draggable,
  DropResult,
} from "@hello-pangea/dnd";
import { Task } from "../../lib/task/task";
import TaskCard from "./TaskCard";
import { capitalize } from "../../lib/utils/string";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";
import useWs from "../../hooks/useWs";
import { getTaskStatuses } from "../../lib/task/status";
import { useProjectStore } from "../../store/selectedProjectStore";
import { useSprintStore } from "../../store/selectedSprintStore";

type TaskMap = Record<string, Task[]>; // { status: [tasks] }

const Dashboard: FC<ComponentProps<"table">> = (props) => {
  const [tasksByStatus, setTasksByStatus] = useState<TaskMap>({});
  const { project } = useProjectStore();
  const { sprint } = useSprintStore();
  useEffect(() => {
    console.log(`selected sprint: ${sprint?.id}`);
    console.log(`selected project: ${project?.id}`);
  }, [project, sprint]);

  const { val, send } = useWs(
    `ws://localhost:8080/ws/projects/${project?.id}/sprints/${sprint?.id}`
  );

  useLayoutEffect(() => {
    if (val) {
      const tasks: Task[] = JSON.parse(val);

      // Group tasks into a map by status
      const grouped: TaskMap = {};
      getTaskStatuses.forEach((status) => {
        grouped[status] = [];
      });

      if (tasks) {
        tasks.forEach((task) => {
          if (!grouped[task.status]) grouped[task.status] = [];
          grouped[task.status].push(task);
        });

        setTasksByStatus(grouped);
      }
    }
  }, [val]);

  const onDragEnd = (result: DropResult) => {
    const { source, destination } = result;
    if (!destination) return;

    const sourceStatus = source.droppableId;
    const destStatus = destination.droppableId;

    // Clone the current map
    const newTasksByStatus = { ...tasksByStatus };

    // Copy arrays to avoid mutating state directly
    const sourceTasks = [...(newTasksByStatus[sourceStatus] || [])];
    const destTasks = [...(newTasksByStatus[destStatus] || [])];

    // Remove task from source
    const [movedTask] = sourceTasks.splice(source.index, 1);

    // Update task status
    movedTask.status = destStatus;

    // Insert task into destination
    destTasks.splice(destination.index, 0, movedTask);

    // Update the map
    newTasksByStatus[sourceStatus] = sourceTasks;
    newTasksByStatus[destStatus] = destTasks;

    setTasksByStatus(newTasksByStatus);

    send(movedTask);

    toast.success(`Moved task "${movedTask.title}" to ${destStatus}`, {
      ...toastOpts,
      autoClose: 2000,
    });
  };

  return (
    <DragDropContext onDragEnd={onDragEnd}>
      <table
        className={`w-full h-[100vh] table-fixed border-collapse bg-primary-500 text-neutral-100 ${props.className}`}
      >
        <thead>
          <tr className="bg-primary-400 text-neutral-100">
            {getTaskStatuses.map((status) => (
              <th
                key={status}
                className="border border-primary-300 p-4 text-lg font-semibold"
              >
                {capitalize(status).replace(/_/g, " ")}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="text-white border-white">
          <tr>
            {getTaskStatuses.map((status) => (
              <td
                key={status}
                className="border border-primary-300 p-4 align-top w-1/4"
              >
                <Droppable droppableId={status}>
                  {(provided) => (
                    <div
                      ref={provided.innerRef}
                      {...provided.droppableProps}
                      className="min-h-[100px] p-2"
                    >
                      {(tasksByStatus[status] || []).map((task, index) => (
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
                              <TaskCard
                                onClick={() => {
                                  console.log(task);
                                }}
                                task={task}
                              />
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
