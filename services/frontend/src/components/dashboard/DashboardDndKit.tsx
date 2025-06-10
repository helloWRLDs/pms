import {
  FC,
  useState,
  useLayoutEffect,
  ComponentProps,
  useEffect,
} from "react";
import {
  DndContext,
  DragOverlay,
  MouseSensor,
  TouchSensor,
  useSensor,
  useSensors,
  DragEndEvent,
  UniqueIdentifier,
  DragStartEvent,
  DragOverEvent,
  useDroppable,
  pointerWithin,
} from "@dnd-kit/core";
import {
  SortableContext,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { Task } from "../../lib/task/task";
import TaskCard from "./TaskCard";
import { capitalize } from "../../lib/utils/string";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";
import useWs from "../../hooks/useWs";
import { getTaskStatuses } from "../../lib/task/status";
import { Modal } from "../ui/Modal";
import TaskView from "../task/TaskView";
import SortableTaskCard from "./SortableTaskCard";

type TaskMap = Record<string, Task[]>;

type DashboardDndKitProps = ComponentProps<"div"> & {
  sprintID: string;
};

interface DroppableContainerProps {
  id: string;
  items: Task[];
  onTaskClick: (task: Task) => void;
}

const DroppableContainer: FC<DroppableContainerProps> = ({
  id,
  items,
  onTaskClick,
}) => {
  const { setNodeRef, isOver } = useDroppable({
    id,
    data: {
      type: "container",
      accepts: ["task"],
    },
  });

  return (
    <div
      ref={setNodeRef}
      className={`flex-1 p-4 min-h-[calc(100vh-12rem)] rounded-lg transition-colors duration-200 ${
        isOver
          ? "bg-primary-400/50 outline-2 outline-accent-500"
          : "bg-primary-500/50"
      }`}
      data-status={id}
    >
      <div className="text-lg font-semibold mb-4">
        {capitalize(id).replace("_", " ")}
      </div>
      <SortableContext
        items={items.map((t) => t.id)}
        strategy={verticalListSortingStrategy}
      >
        <div className="space-y-2">
          {items.map((task) => (
            <SortableTaskCard
              key={task.id}
              task={task}
              onClick={() => onTaskClick(task)}
            />
          ))}
        </div>
      </SortableContext>
    </div>
  );
};

const StatusToggle: FC<{
  status: string;
  enabled: boolean;
  onToggle: (status: string) => void;
}> = ({ status, enabled, onToggle }) => (
  <div
    className="flex items-center gap-2 p-2 cursor-pointer hover:bg-primary-400/20 rounded transition-colors"
    onClick={() => onToggle(status)}
  >
    <input
      type="checkbox"
      checked={enabled}
      onChange={() => onToggle(status)}
      className="form-checkbox h-4 w-4 text-accent-500 rounded"
    />
    <span className="text-neutral-100">
      {capitalize(status).replace("_", " ")}
    </span>
  </div>
);

const DashboardDndKit: FC<DashboardDndKitProps> = ({
  sprintID,
  ...props
}: DashboardDndKitProps) => {
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [taskViewModal, setTaskViewModal] = useState(false);
  const [tasksByStatus, setTasksByStatus] = useState<TaskMap>({});
  const [activeId, setActiveId] = useState<UniqueIdentifier | null>(null);
  const [activeContainer, setActiveContainer] = useState<string | null>(null);
  const [dragOverStatus, setDragOverStatus] = useState<string | null>(null);
  const [lastDroppedContainer, setLastDroppedContainer] = useState<
    string | null
  >(null);
  const [statusFilterModal, setStatusFilterModal] = useState(false);
  const [enabledStatuses, setEnabledStatuses] = useState<Set<string>>(
    new Set(getTaskStatuses)
  );

  console.log(activeContainer);
  console.log(dragOverStatus);

  // Configure sensors for better drag detection
  const sensors = useSensors(
    useSensor(MouseSensor, {
      activationConstraint: {
        distance: 5,
      },
    }),
    useSensor(TouchSensor, {
      activationConstraint: {
        delay: 250,
        tolerance: 5,
      },
    })
  );

  const { val, send, close } = useWs(
    `ws://localhost:8080/ws/sprints/${sprintID}`
  );

  // Cleanup WebSocket connection when component unmounts
  useEffect(() => {
    return () => {
      console.log("Cleaning up WebSocket connection");
      close();
    };
  }, [close]);

  useLayoutEffect(() => {
    if (!val) return;

    try {
      const tasks: Task[] = JSON.parse(val);

      // Group tasks into a map by status
      const grouped: TaskMap = {};
      getTaskStatuses.forEach((status) => {
        grouped[status] = [];
      });

      tasks?.forEach((task) => {
        if (!grouped[task.status]) grouped[task.status] = [];
        grouped[task.status].push(task);
      });

      setTasksByStatus(grouped);
    } catch (error) {
      console.error("Error parsing WebSocket data:", error);
      toast.error("Error updating task board", toastOpts);
    }
  }, [val]);

  const findContainer = (id: UniqueIdentifier): string | undefined => {
    const stringId = id.toString();
    if (stringId in tasksByStatus) return stringId;

    return Object.keys(tasksByStatus).find((key) =>
      tasksByStatus[key].some((task) => task.id === stringId)
    );
  };

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const activeContainer = findContainer(active.id);
    setActiveId(active.id);
    setActiveContainer(activeContainer || null);
    setDragOverStatus(null);
    setLastDroppedContainer(null);
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    if (!over) {
      setDragOverStatus(null);
      return;
    }

    // Only allow dropping in containers
    if (!over.data.current?.type || over.data.current.type !== "container") {
      return;
    }

    const activeContainer = findContainer(active.id);
    const overContainer = over.id.toString();

    if (!activeContainer || activeContainer === overContainer) {
      return;
    }

    setDragOverStatus(overContainer);
  };

  const handleDragEnd = async (event: DragEndEvent) => {
    const { active, over } = event;

    setActiveId(null);
    setActiveContainer(null);
    setDragOverStatus(null);

    // If no valid drop target or not a container, do nothing
    if (
      !over ||
      !over.data.current?.type ||
      over.data.current.type !== "container"
    ) {
      return;
    }

    const activeContainer = findContainer(active.id);
    const overContainer = over.id.toString();

    // Prevent duplicate drops in the same container
    if (
      !activeContainer ||
      activeContainer === overContainer ||
      overContainer === lastDroppedContainer
    ) {
      return;
    }

    setLastDroppedContainer(overContainer);

    // Find the task that was dragged
    const task = Object.values(tasksByStatus)
      .flat()
      .find((t) => t.id === active.id.toString());

    if (!task) return;

    // Update task status
    const updatedTask = { ...task, status: overContainer };

    // Update local state
    setTasksByStatus((prev) => {
      // First, ensure we're not duplicating the task
      const sourceItems = prev[activeContainer].filter(
        (item) => item.id !== active.id.toString()
      );

      // Then update both source and destination
      return {
        ...prev,
        [activeContainer]: sourceItems,
        [overContainer]: [updatedTask, ...(prev[overContainer] || [])],
      };
    });

    // Send updated task through websocket
    send(updatedTask);

    toast.success(`Moved task "${task.title}" to ${overContainer}`, {
      ...toastOpts,
      autoClose: 2000,
    });
  };

  const handleDragCancel = () => {
    setActiveId(null);
    setActiveContainer(null);
    setDragOverStatus(null);
    setLastDroppedContainer(null);
  };

  const activeTask = activeId
    ? Object.values(tasksByStatus)
        .flat()
        .find((task) => task.id === activeId)
    : null;

  const toggleStatus = (status: string) => {
    setEnabledStatuses((prev) => {
      const next = new Set(prev);
      if (next.has(status)) {
        // Don't allow disabling the last status
        if (next.size > 1) {
          next.delete(status);
        }
      } else {
        next.add(status);
      }
      return next;
    });
  };

  return (
    <div className={props.className}>
      <section id="modals">
        <Modal
          title={capitalize(selectedTask?.title ?? "Task View")}
          visible={taskViewModal}
          onClose={() => {
            setTaskViewModal(false);
          }}
          className="bg-secondary-300"
        >
          <TaskView task_id={selectedTask?.id ?? ""} />
        </Modal>

        <Modal
          visible={statusFilterModal}
          title="Filter Status Columns"
          onClose={() => setStatusFilterModal(false)}
          className="bg-secondary-300"
        >
          <div className="space-y-2">
            {getTaskStatuses.map((status) => (
              <StatusToggle
                key={status}
                status={status}
                enabled={enabledStatuses.has(status)}
                onToggle={toggleStatus}
              />
            ))}
          </div>
        </Modal>
      </section>

      <DndContext
        sensors={sensors}
        onDragStart={handleDragStart}
        onDragOver={handleDragOver}
        onDragEnd={handleDragEnd}
        onDragCancel={handleDragCancel}
        collisionDetection={pointerWithin}
      >
        <div className="flex flex-col h-screen bg-primary-600 p-4">
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-3xl font-bold text-neutral-100">
              Sprint Dashboard
            </h2>
            <button
              onClick={() => setStatusFilterModal(true)}
              className="px-4 py-2 bg-primary-400 hover:bg-primary-300 transition-colors rounded-lg text-neutral-100 flex items-center gap-2"
            >
              <span>Filter Columns</span>
              <span className="text-sm bg-primary-300 px-2 py-0.5 rounded">
                {enabledStatuses.size}/{getTaskStatuses.length}
              </span>
            </button>
          </div>

          <div className="flex gap-4 flex-1 overflow-x-auto pb-4">
            {getTaskStatuses
              .filter((status) => enabledStatuses.has(status))
              .map((status) => (
                <DroppableContainer
                  key={status}
                  id={status}
                  items={tasksByStatus[status] || []}
                  onTaskClick={(task) => {
                    setSelectedTask(task);
                    setTaskViewModal(true);
                  }}
                />
              ))}
          </div>
        </div>

        <DragOverlay>
          {activeTask && <TaskCard task={activeTask} onClick={() => {}} />}
        </DragOverlay>
      </DndContext>
    </div>
  );
};

export default DashboardDndKit;
