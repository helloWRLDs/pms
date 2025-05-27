import { useEffect, useState } from "react";
import NewTaskForm from "../components/forms/NewTaskForm";
import { useProjectStore } from "../store/selectedProjectStore";
import { useNavigate } from "react-router-dom";
import { Modal } from "../components/ui/Modal";
import { taskAPI } from "../api/taskAPI";
import { infoToast } from "../lib/utils/toast";
import { Task, TaskCreation, TaskFilter } from "../lib/task/task";
import { usePageSettings } from "../hooks/usePageSettings";
import Table from "../components/ui/Table";
import { useQuery } from "@tanstack/react-query";
import { capitalize } from "../lib/utils/string";
import Badge from "../components/ui/Badge";
import { formatTime } from "../lib/utils/time";
import { Priority, PriorityFilterValues } from "../lib/task/priority";
import Paginator from "../components/ui/Paginator";
import { StatusFilterValues } from "../lib/task/status";
import { Switch } from "@headlessui/react";
import { useAuthStore } from "../store/authStore";
import TaskView from "../components/task/TaskView";
import { useCacheStore } from "../store/cacheStore";
import { BsFillPlusCircleFill } from "react-icons/bs";
import { ListItems } from "../lib/utils/list";
import { parseError } from "../lib/errors";
import authAPI from "../api/auth";
import { useCacheLoader } from "../hooks/useCacheLoader";

const BacklogPage = () => {
  usePageSettings({ requireAuth: true, title: "Backlog" });

  const { project } = useProjectStore();
  const { auth, clearAuth } = useAuthStore();
  const { getSprint } = useCacheStore();

  const navigate = useNavigate();

  const [task, setTask] = useState<Task>();
  const [taskViewModal, setTaskViewModal] = useState(false);
  const [taskCreationModal, setTaskCreationModal] = useState(false);
  const [filter, setFilter] = useState<TaskFilter>({
    page: 1,
    per_page: 10,
  });

  const {
    data: taskList,
    isLoading: isTaskListLoading,
    refetch: refetachTasks,
    error,
    isLoadingError,
  } = useQuery<ListItems<Task>>({
    queryKey: [
      "tasks",
      filter.page,
      filter.per_page,
      filter.status,
      filter.priority,
      filter.assignee_id,
      filter.sprint_id,
    ],
    queryFn: async () => {
      try {
        const res = await taskAPI.list(filter);
        return res;
      } catch (e) {
        if (parseError(e)?.status == 401) {
          clearAuth();
          navigate("/login");
        }
      }
      return {} as ListItems<Task>;
    },
    enabled: !!project?.id,
  });

  const { data: assignees } = useQuery({
    queryKey: ["assignees", project?.company_id],
    queryFn: async () => {
      try {
        const res = await authAPI.listUsers({
          page: 1,
          per_page: 100,
          company_id: project?.company_id ?? "",
        });
        return res;
      } catch (e) {
        if (parseError(e)?.status === 401) {
          clearAuth();
          navigate("/login");
        }
      }
    },
    enabled: !!project?.company_id,
  });

  useCacheLoader({ userList: assignees });

  useEffect(() => {
    console.log(`check error: `, error, isLoadingError);
  }, [error]);

  useEffect(() => {
    if (!project) {
      navigate("/projects");
    }
  }, []);

  const { getAssignee } = useCacheStore();

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-primary-600 text-neutral-100">
      <section id="modals">
        <Modal
          title={`${task?.title ?? "Task Preview"}`}
          visible={taskViewModal}
          onClose={() => setTaskViewModal(false)}
          className="bg-secondary-300"
        >
          {task && (
            <TaskView
              refetchTasks={refetachTasks}
              assignees={assignees}
              task={task}
              className=""
            />
          )}
        </Modal>
        <Modal
          title="Create New Task"
          visible={taskCreationModal}
          onClose={() => {
            setTaskCreationModal(false);
          }}
          className="w-[50%] mx-auto bg-secondary-300"
        >
          {project && (
            <NewTaskForm
              project={project}
              onSubmit={(data: TaskCreation) => {
                taskAPI
                  .create(data)
                  .then((res) => {
                    console.log(res);
                    infoToast("Task created successfully");
                    setTaskCreationModal(false);
                    refetachTasks();
                  })
                  .catch((error) => {
                    console.error(error);
                  });
              }}
              className=""
            />
          )}
        </Modal>
      </section>

      <section id="task-header">
        <div className="container mx-auto ">
          <div className="mb-5">
            <h1 className="font-bold text-3xl">
              <span className="text-accent-500">
                {project && project.title}
              </span>{" "}
              Backlog
            </h1>
          </div>

          <div className="flex items-center mb-4">
            <div className="flex gap-4 items-center mb-4">
              <label
                htmlFor="filter-priority"
                className="flex gap-2 items-center"
              >
                <span>Priority</span>
                <select
                  name=""
                  id=""
                  className="px-4 py-2 border rounded-lg text-neutral-100 bg-primary-400"
                >
                  {PriorityFilterValues.map((priority) => (
                    <option
                      value={priority.value}
                      onClick={() => {
                        setFilter({ ...filter, priority: priority.value });
                      }}
                    >
                      {priority.label}
                    </option>
                  ))}
                </select>
              </label>

              <label
                htmlFor="filter-status"
                className="flex gap-2 items-center"
              >
                <span>Status</span>
                <select
                  name=""
                  id=""
                  className="px-4 py-2 border rounded-lg text-neutral-100 bg-primary-400 cursor-pointer"
                >
                  {StatusFilterValues.map((status) => (
                    <option
                      value={status.value}
                      onClick={() => {
                        setFilter({ ...filter, status: status.value });
                      }}
                    >
                      {status.label}
                    </option>
                  ))}
                </select>
              </label>

              <label
                htmlFor="filter-status"
                className="flex gap-2 items-center"
              >
                <span>Show mine</span>
                <Switch
                  checked={!!filter.assignee_id}
                  onClick={() => {
                    if (filter.assignee_id !== "") {
                      setFilter({ ...filter, assignee_id: "" });
                    } else {
                      setFilter({
                        ...filter,
                        assignee_id: auth?.user.id ?? "",
                      });
                    }
                    console.log(filter);
                  }}
                  className="cursor-pointer bg-secondary-100 group inline-flex h-6 w-11 items-center rounded-full  transition data-checked:bg-accent-600"
                >
                  <span className="size-4 translate-x-1 rounded-full bg-white transition group-data-checked:translate-x-6" />
                </Switch>
              </label>
            </div>
          </div>
        </div>
      </section>

      <section id="task-list">
        <div className="overflow-x-auto container mx-auto shadow-xl">
          <div>
            <div className="h-[75lvh] w-full">
              <Table className="rounded-lg">
                <Table.Head>
                  <Table.Row className="text-neutral-100 bg-primary-400">
                    <Table.HeadCell>ID</Table.HeadCell>
                    <Table.HeadCell>Title</Table.HeadCell>
                    <Table.HeadCell>Status</Table.HeadCell>
                    <Table.HeadCell>Priority</Table.HeadCell>
                    <Table.HeadCell>Sprint</Table.HeadCell>
                    <Table.HeadCell>Assignee</Table.HeadCell>
                    <Table.HeadCell>Due to</Table.HeadCell>
                    <Table.HeadCell>Created</Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                {isTaskListLoading ? (
                  <p>Loading...</p>
                ) : !taskList ||
                  !taskList.items ||
                  taskList.items.length === 0 ? (
                  <Table.Body></Table.Body>
                ) : (
                  <Table.Body>
                    {taskList.items.map((task, i) => (
                      <Table.Row
                        key={i}
                        onClick={() => {
                          setTask(task);
                          setTaskViewModal(true);
                        }}
                        className="cursor-pointer bg-secondary-200 text-neutral-100 hover:bg-secondary-100"
                      >
                        <Table.Cell className="py-5">{task.code}</Table.Cell>
                        <Table.Cell>{task.title}</Table.Cell>
                        <Table.Cell>
                          <Badge className="text-white bg-primary-400">
                            {capitalize(task.status.replace("_", ""))}
                          </Badge>
                        </Table.Cell>
                        <Table.Cell>
                          <Badge
                            className={`text-${new Priority(
                              task.priority
                            ).getColor()}-500`}
                          >
                            {new Priority(task.priority).toString()}
                          </Badge>
                        </Table.Cell>
                        <Table.Cell>
                          {task.sprint_id
                            ? getSprint(task.sprint_id)?.title
                            : "none"}
                        </Table.Cell>
                        <Table.Cell>
                          {task.assignee_id
                            ? getAssignee(task.assignee_id)?.name ?? "none "
                            : "none"}
                        </Table.Cell>
                        <Table.Cell>
                          {formatTime(task.due_date.seconds)}
                        </Table.Cell>
                        <Table.Cell>
                          {formatTime(task.created_at.seconds)}
                        </Table.Cell>
                      </Table.Row>
                    ))}
                  </Table.Body>
                )}
              </Table>
              <button
                className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
                onClick={() => {
                  setTaskCreationModal(true);
                }}
              >
                <BsFillPlusCircleFill
                  size="30"
                  className="mx-auto text-neutral-300 group-hover:text-accent-300 "
                />
              </button>
            </div>
            {taskList && taskList.items && (
              <Paginator
                page={taskList.page ?? 0}
                per_page={taskList.per_page ?? 0}
                total_items={taskList.total_items ?? 0}
                total_pages={taskList.total_pages ?? 0}
                onPageChange={(page) => {
                  setFilter({ ...filter, page: page });
                }}
              />
            )}
          </div>
        </div>
      </section>
    </div>
  );
};

export default BacklogPage;
