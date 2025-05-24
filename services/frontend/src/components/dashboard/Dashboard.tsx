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
import { Modal } from "../ui/Modal";
import TaskView from "../task/TaskView";
import { useCacheStore } from "../../store/cacheStore";

type TaskMap = Record<string, Task[]>; // { status: [tasks] }

const Dashboard: FC<ComponentProps<"table">> = (props) => {
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [taskViewModal, setTaskViewModal] = useState(false);
  const [tasksByStatus, setTasksByStatus] = useState<TaskMap>({});
  const { project } = useProjectStore();
  const { sprint } = useSprintStore();

  const { val, send } = useWs(
    `ws://localhost:8080/ws/projects/${project?.id}/sprints/${sprint?.id}`
  );

  const { assignees } = useCacheStore();

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
    if (source.droppableId === destination.droppableId) return;

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
    <div>
      <section id="modals">
        <Modal
          title={capitalize(selectedTask?.title ?? "Task View")}
          visible={taskViewModal}
          onClose={() => {
            setTaskViewModal(false);
          }}
          className="bg-secondary-300"
        >
          <TaskView
            assignees={{
              total_items: 0,
              total_pages: 0,
              page: 0,
              per_page: 0,
              items: [],
              // items: Object.values(assignees),
            }}
            task={selectedTask ?? ({} as Task)}
          />
        </Modal>
      </section>
      <section>
        <DragDropContext onDragEnd={onDragEnd}>
          <table
            className={`w-full h-[100vh] table-fixed border-collapse bg-primary-500 text-neutral-100 ${props.className}`}
          >
            <thead>
              <tr className="bg-primary-400 text-neutral-100 h-[3rem]">
                {getTaskStatuses.map((status) => (
                  <th
                    key={status}
                    className="border border-primary-300 p-4 text-lg font-semibold"
                  >
                    {capitalize(status).replace("_", " ")}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody className="text-white border-white">
              <tr>
                {getTaskStatuses.map((status) => (
                  <td
                    key={status}
                    className="border border-primary-300 p-4 align-top"
                  >
                    <Droppable droppableId={status}>
                      {(provided) => (
                        <div
                          ref={provided.innerRef}
                          {...provided.droppableProps}
                          className="p-2"
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
                                      setSelectedTask(task);
                                      setTaskViewModal(true);
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
      </section>
    </div>
  );
};

export default Dashboard;
