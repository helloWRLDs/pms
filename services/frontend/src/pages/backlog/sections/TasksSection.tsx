import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { BsKanbanFill, BsTrash, BsFilter } from "react-icons/bs";
import { FiCalendar, FiUser } from "react-icons/fi";
import { Switch } from "@headlessui/react";
import { Task, TaskFilter } from "../../../lib/task/task";
import { taskAPI } from "../../../api/taskAPI";
import { ListItems } from "../../../lib/utils/list";
import { parseError } from "../../../lib/errors";
import { useAuthStore } from "../../../store/authStore";
import { useAssigneeList, useSprintList } from "../../../hooks/useData";
import useMetaCache from "../../../store/useMetaCache";
import { Priority, PriorityFilterValues } from "../../../lib/task/priority";
import { StatusFilterValues } from "../../../lib/task/status";
import {
  getTaskTypeConfig,
  TaskType,
  TaskTypes,
} from "../../../lib/task/tasktype";
import { capitalize } from "../../../lib/utils/string";
import { formatTime } from "../../../lib/utils/time";
import { Button } from "../../../components/ui/Button";
import Table from "../../../components/ui/Table";
import Badge from "../../../components/ui/Badge";
import TaskCard from "../../../components/task/TaskCard";
import { ContextMenu } from "../../../components/ui/ContextMenu";
import Paginator from "../../../components/ui/Paginator";
import FilterButton from "../../../components/ui/button/FilterButton";
import { usePermission } from "../../../hooks/usePermission";
import { Permissions } from "../../../lib/permission";
import { infoToast, errorToast } from "../../../lib/utils/toast";
import { useEffect, useState } from "react";
import { Modal } from "../../../components/ui/Modal";
import TaskView from "../../../components/task/TaskView";

type ViewMode = "table" | "cards";

interface TasksSectionProps {
  filter: TaskFilter;
  viewMode: ViewMode;
  onFilterChange: (filter: TaskFilter) => void;
  onRefetch?: () => void;
  onTaskUpdate?: (updatedTask: Task) => void;
  currentTaskId?: string;
}

const TasksSection = ({
  filter,
  viewMode,
  onFilterChange,
  onRefetch,
  onTaskUpdate,
  currentTaskId,
}: TasksSectionProps) => {
  const [taskViewModal, setTaskViewModal] = useState(false);
  const [task, setTask] = useState<Task | null>(null);

  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const navigate = useNavigate();
  const { auth, clearAuth } = useAuthStore();

  const { getAssigneeName } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );
  const { sprints, getSprintName } = useSprintList(
    metaCache.metadata.selectedProject?.id ?? ""
  );

  const {
    data: taskList,
    isLoading,
    refetch: refetchTasks,
    error,
  } = useQuery<ListItems<Task>>({
    queryKey: [
      "tasks",
      filter.page,
      filter.per_page,
      filter.status,
      filter.priority,
      filter.assignee_id,
      filter.sprint_id,
      filter.type,
    ],
    queryFn: async () => {
      try {
        const response = await taskAPI.list(filter);
        return response;
      } catch (e) {
        if (parseError(e)?.status === 401) {
          clearAuth();
          navigate("/login");
        }
        throw e;
      }
    },
    enabled:
      !!metaCache.metadata.selectedProject?.id &&
      hasPermission(Permissions.TASK_READ_PERMISSION),
  });

  useEffect(() => {
    if (taskList) {
      console.log(JSON.stringify(taskList, null, 2));

      // If there's a current task being viewed in a modal, update it with fresh data
      if (currentTaskId && onTaskUpdate && taskList.items) {
        const updatedTask = taskList.items.find(
          (task) => task.id === currentTaskId
        );
        if (updatedTask) {
          onTaskUpdate(updatedTask);
        }
      }
    }
  }, [taskList, currentTaskId, onTaskUpdate]);

  const handleDeleteTask = async (taskId: string) => {
    if (!hasPermission(Permissions.TASK_DELETE_PERMISSION)) return;

    try {
      await taskAPI.delete(taskId);
      infoToast("Task deleted successfully");
      refetchTasks();
      onRefetch?.();
    } catch (error) {
      errorToast("Failed to delete task");
      console.error(error);
    }
  };

  // Don't render if user doesn't have permission to read tasks
  if (!hasPermission(Permissions.TASK_READ_PERMISSION)) {
    return null;
  }

  return (
    <div className="px-6 pb-8">
      <div className="max-w-7xl mx-auto">
        {/* Task Filters */}
        <div className="bg-secondary-200/50 rounded-lg p-4 mb-8">
          <div className="flex flex-wrap items-center gap-6">
            <div className="flex items-center gap-3">
              <BsFilter className="text-accent-500" size={20} />
              <span className="text-neutral-300">Filters:</span>
            </div>

            <div className="flex flex-wrap items-center gap-4">
              <label className="flex items-center gap-2">
                <span>Priority</span>
                <FilterButton
                  label="Priority"
                  value={filter.priority?.toString()}
                  options={PriorityFilterValues}
                  onChange={(value) =>
                    onFilterChange({
                      ...filter,
                      priority: parseInt(value, 10),
                      page: 1,
                    })
                  }
                />
              </label>

              <label className="flex items-center gap-2">
                <span>Status</span>
                <FilterButton
                  label="Status"
                  value={filter.status}
                  options={StatusFilterValues}
                  onChange={(value) =>
                    onFilterChange({ ...filter, status: value, page: 1 })
                  }
                />
              </label>

              <label className="flex items-center gap-2">
                <span>Sprint</span>
                <FilterButton
                  label="Sprint"
                  value={filter.sprint_id}
                  options={[
                    { value: "", label: "All" },
                    ...((sprints?.items &&
                      sprints?.items.map((sprint) => ({
                        label: sprint.title,
                        value: sprint.id,
                      }))) ??
                      []),
                  ]}
                  onChange={(value) =>
                    onFilterChange({ ...filter, sprint_id: value, page: 1 })
                  }
                />
              </label>

              <label className="flex items-center gap-2">
                <span>Type</span>
                <FilterButton
                  label="Type"
                  value={filter.type}
                  options={[
                    { value: "", label: "All" },
                    ...Object.values(TaskTypes).map((type) => ({
                      label: getTaskTypeConfig(type).label,
                      value: type,
                    })),
                  ]}
                  onChange={(value) =>
                    onFilterChange({
                      ...filter,
                      type: value as TaskType,
                      page: 1,
                    })
                  }
                />
              </label>

              <label className="flex items-center gap-2">
                <span>My Tasks</span>
                <Switch
                  checked={!!filter.assignee_id}
                  onChange={() => {
                    onFilterChange({
                      ...filter,
                      assignee_id: filter.assignee_id
                        ? ""
                        : auth?.user.id ?? "",
                      page: 1,
                    });
                  }}
                  className={`${
                    filter.assignee_id ? "bg-accent-500" : "bg-secondary-100"
                  } relative inline-flex h-6 w-11 items-center rounded-full transition-colors`}
                >
                  <span
                    className={`${
                      filter.assignee_id ? "translate-x-6" : "translate-x-1"
                    } inline-block h-4 w-4 transform rounded-full bg-white transition-transform`}
                  />
                </Switch>
              </label>
            </div>
          </div>
        </div>

        {/* Task Display */}
        {error ? (
          <div className="text-center py-12">
            <p className="text-red-500 mb-4">Failed to load tasks</p>
            <Button onClick={() => refetchTasks()}>Retry</Button>
          </div>
        ) : viewMode === "cards" ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            {isLoading
              ? Array(6)
                  .fill(0)
                  .map((_, i) => (
                    <div
                      key={i}
                      className="bg-secondary-200 rounded-lg p-4 animate-pulse"
                    >
                      <div className="h-4 bg-secondary-100 rounded w-1/4 mb-4"></div>
                      <div className="h-6 bg-secondary-100 rounded w-3/4 mb-4"></div>
                      <div className="h-8 bg-secondary-100 rounded w-1/2 mb-4"></div>
                      <div className="grid grid-cols-2 gap-4">
                        <div className="h-4 bg-secondary-100 rounded"></div>
                        <div className="h-4 bg-secondary-100 rounded"></div>
                      </div>
                    </div>
                  ))
              : taskList?.items?.map((task) => (
                  <TaskCard
                    key={task.id}
                    task={task}
                    onClick={() => {
                      setTask(task);
                      setTaskViewModal(true);
                    }}
                  />
                ))}
          </div>
        ) : (
          <div className="mb-8 bg-secondary-200 rounded-lg overflow-hidden shadow-lg">
            <div className="min-h-[50vh] overflow-x-auto">
              <table className="w-full border-collapse rounded-lg overflow-hidden">
                <Table.Head className="bg-primary-400 text-neutral-100 sticky top-0 z-10">
                  <Table.Row>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[8%] border-r border-primary-300">
                      ID
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[25%] border-r border-primary-300">
                      Title
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[12%] border-r border-primary-300">
                      Type
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[10%] border-r border-primary-300">
                      Status
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[10%] border-r border-primary-300">
                      Priority
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[15%] border-r border-primary-300">
                      Assignee
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[13%] border-r border-primary-300">
                      Sprint
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[15%]">
                      Due Date
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[15%]">
                      Actions
                    </Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                <Table.Body className="divide-y divide-secondary-100">
                  {isLoading ? (
                    Array(10)
                      .fill(0)
                      .map((_, index) => (
                        <Table.Row
                          key={index}
                          className="bg-secondary-200 hover:bg-secondary-100 transition-all duration-200 animate-pulse"
                        >
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-16"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-3/4"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-6 bg-secondary-100 rounded w-20"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-6 bg-secondary-100 rounded w-20"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-6 bg-secondary-100 rounded w-16"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-24"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-20"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4">
                            <div className="h-4 bg-secondary-100 rounded w-28"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4"></Table.Cell>
                        </Table.Row>
                      ))
                  ) : !taskList?.items || taskList.items.length === 0 ? (
                    <Table.Row className="bg-secondary-200">
                      <Table.Cell
                        className="px-6 py-12 text-center text-neutral-400 italic text-lg"
                        {...({ colSpan: 9 } as any)}
                      >
                        <div className="flex flex-col items-center gap-3">
                          <BsKanbanFill className="text-4xl text-neutral-500" />
                          <span>No tasks found</span>
                          <span className="text-sm text-neutral-500">
                            {filter.status ||
                            filter.priority ||
                            filter.assignee_id ||
                            filter.sprint_id ||
                            filter.type
                              ? "Try adjusting your filters or create a new task"
                              : "Create your first task to get started"}
                          </span>
                        </div>
                      </Table.Cell>
                    </Table.Row>
                  ) : (
                    taskList.items.map((task) => {
                      const priority = Priority.fromNumber(task.priority);

                      return (
                        <Table.Row
                          key={task.id}
                          onClick={() => {
                            setTask(task);
                            setTaskViewModal(true);
                          }}
                          className="bg-secondary-200 hover:bg-secondary-100 transition-all duration-200 cursor-pointer group border-l-4 border-l-transparent hover:border-l-accent-500"
                        >
                          <Table.Cell className="px-6 py-4 font-mono text-neutral-400 text-sm border-r border-secondary-100 group-hover:text-accent-400 transition-colors">
                            {task.code}
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 font-medium border-r border-secondary-100 group-hover:text-accent-400 transition-colors">
                            <div
                              className="truncate max-w-xs"
                              title={task.title}
                            >
                              {task.title}
                            </div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            {task.type ? (
                              <div className="flex items-center gap-2">
                                <span className="text-lg">
                                  {getTaskTypeConfig(task.type as any).icon}
                                </span>
                                <Badge
                                  className="px-3 py-1 rounded-full text-xs font-semibold uppercase tracking-wide"
                                  style={{
                                    backgroundColor: `${
                                      getTaskTypeConfig(task.type as any).color
                                    }20`,
                                    color: getTaskTypeConfig(task.type as any)
                                      .color,
                                  }}
                                >
                                  {getTaskTypeConfig(task.type as any).label}
                                </Badge>
                              </div>
                            ) : (
                              <span className="text-neutral-500 text-sm italic">
                                No Type
                              </span>
                            )}
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <Badge
                              className={`bg-${task.status.toLowerCase()}-500/20 text-${task.status.toLowerCase()}-500 px-3 py-1 rounded-full text-xs font-semibold uppercase tracking-wide`}
                            >
                              {capitalize(task.status.replace("_", " "))}
                            </Badge>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <Badge
                              className={`bg-${priority.getColor()}-500/20 text-${priority.getColor()}-500 px-3 py-1 rounded-full text-xs font-semibold uppercase tracking-wide`}
                            >
                              {priority.toString()}
                            </Badge>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="flex items-center gap-2">
                              <FiUser
                                className="text-accent-500 flex-shrink-0"
                                size={16}
                              />
                              <span className="truncate">
                                {getAssigneeName(task.assignee_id)}
                              </span>
                            </div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="flex items-center gap-2">
                              <div
                                className={`w-2 h-2 rounded-full flex-shrink-0 ${
                                  task.sprint_id
                                    ? "bg-accent-500"
                                    : "bg-neutral-500"
                                }`}
                              ></div>
                              <span className="truncate text-sm">
                                {getSprintName(task.sprint_id)}
                              </span>
                            </div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4">
                            <div className="flex items-center gap-2 text-sm">
                              <FiCalendar
                                className="text-accent-500 flex-shrink-0"
                                size={16}
                              />
                              <span>{formatTime(task.due_date.seconds)}</span>
                            </div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4">
                            {hasPermission(
                              Permissions.TASK_DELETE_PERMISSION
                            ) && (
                              <ContextMenu
                                items={[
                                  {
                                    label: "Delete",
                                    icon: <BsTrash />,
                                    onClick: () => handleDeleteTask(task.id),
                                  },
                                ]}
                              />
                            )}
                          </Table.Cell>
                        </Table.Row>
                      );
                    })
                  )}
                </Table.Body>
              </table>
            </div>
          </div>
        )}

        {/* Tasks Pagination */}
        {taskList && taskList.total_items > 0 && (
          <div className="flex justify-center">
            <Paginator
              page={taskList.page}
              per_page={taskList.per_page}
              total_items={taskList.total_items}
              total_pages={taskList.total_pages}
              onPageChange={(page) => onFilterChange({ ...filter, page })}
            />
          </div>
        )}
        {/*Modals */}
        <Modal
          title={task?.title ?? "Task Details"}
          visible={taskViewModal}
          onClose={() => setTaskViewModal(false)}
          className="bg-secondary-300"
          size="2xl"
        >
          {task && <TaskView task_id={task.id} refetchTasks={refetchTasks} />}
        </Modal>
      </div>
    </div>
  );
};

export default TasksSection;
