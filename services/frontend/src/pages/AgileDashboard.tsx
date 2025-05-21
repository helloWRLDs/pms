import { FC, useEffect, useState } from "react";
import Dashboard from "../components/dashboard/Dashboard";
import { usePageSettings } from "../hooks/usePageSettings";
import { useSprintStore } from "../store/selectedSprintStore";
import { useNavigate } from "react-router-dom";
import { Modal } from "../components/ui/Modal";
import { useQuery } from "@tanstack/react-query";
import { SprintFilter } from "../lib/sprint/sprint";
import { useProjectStore } from "../store/selectedProjectStore";
import { TaskFilter } from "../lib/task/task";
import { taskAPI } from "../api/taskAPI";

const AgileDashboard: FC = () => {
  usePageSettings({ requireAuth: false, title: "Agile Dashboard" });

  const navigate = useNavigate();

  const { sprint } = useSprintStore();
  const { project } = useProjectStore();

  const [manageTasksModal, setManageTasksModal] = useState(false);

  const [filter, setFilter] = useState<TaskFilter>({
    page: 1,
    per_page: 10,
    sprint_id: sprint?.id,
    sprint_name: sprint?.title,
    priority: 0,
    assignee_id: "",
    status: "",
  });

  const [tasksToAddFilter, settasksToAddFilter] = useState<TaskFilter>({
    page: 1,
    per_page: 10,
    sprint_id: "",
    sprint_name: "",
    priority: 0,
    assignee_id: "",
    status: "",
    project_id: project?.id,
  });

  const { data: tasksToAddList, isLoading: isTasksToAddListLoading } = useQuery(
    {
      queryKey: [
        "task-to-add",
        tasksToAddFilter.page,
        tasksToAddFilter.per_page,
        tasksToAddFilter.title,
      ],
      queryFn: () => taskAPI.list(tasksToAddFilter),
      enabled: manageTasksModal,
    }
  );

  const {
    data: sprintTaskList,
    isLoading: isSprintTaskListLoading,
    refetch,
  } = useQuery({
    queryKey: ["tasks", filter.page, filter.per_page, filter.title],
    queryFn: () => taskAPI.list(filter),
    enabled: manageTasksModal,
  });

  useEffect(() => {
    console.log(sprintTaskList);
  }, [sprintTaskList]);

  useEffect(() => {
    if (!sprint) {
      navigate("/sprints");
    }
  }, [sprint]);

  return (
    <div className="w-full flex flex-col p-4 bg-primary-600 text-neutral-100">
      <section id="modals">
        <Modal
          visible={manageTasksModal}
          title="Manage Tasks"
          onClose={() => setManageTasksModal(false)}
          className="w-[80%]"
        >
          <div>
            <div className="flex ">
              <input
                type="text"
                name=""
                id=""
                className="bg-white text-black"
              />
            </div>
            <button className="px-4 py-2 rounded-md bg-accent-500 text-black cursor-pointer">
              Add Task
            </button>
          </div>
        </Modal>
      </section>

      <section id="sprint-data">
        <div className="sprint-data-header">
          <h2 className="text-3xl font-bold mb-4">{sprint?.title} Dashboard</h2>
        </div>
        <div className="sprint-data-options">
          <button
            className="px-4 py-2 rounded-md bg-accent-500 text-primary-500"
            onClick={() => {
              setManageTasksModal(true);
            }}
          >
            Manage Tasks
          </button>
        </div>
      </section>

      <section id="sprint-agile-dashboard">
        <Dashboard className="w-full" />
      </section>
    </div>
  );
};

export default AgileDashboard;
