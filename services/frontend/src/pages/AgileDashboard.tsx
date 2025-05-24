import { FC, useEffect, useState } from "react";
import Dashboard from "../components/dashboard/Dashboard";
import { usePageSettings } from "../hooks/usePageSettings";
import { useSprintStore } from "../store/selectedSprintStore";
import { useNavigate } from "react-router-dom";
import { Modal } from "../components/ui/Modal";
import { useProjectStore } from "../store/selectedProjectStore";

const AgileDashboard: FC = () => {
  usePageSettings({ requireAuth: true, title: "Agile Dashboard" });

  const navigate = useNavigate();

  const { sprint } = useSprintStore();
  const { project } = useProjectStore();

  const [manageTasksModal, setManageTasksModal] = useState(false);

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
          className="w-[80%] bg-secondary-200"
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
        <div className="sprint-data-header mb-5">
          <h2 className="text-3xl font-bold">
            <span className="text-accent-500">{sprint?.title}</span> Dashboard
          </h2>
        </div>
        <div className="sprint-data-options"></div>
      </section>

      <section id="sprint-agile-dashboard">
        <Dashboard className="w-full" />
      </section>
    </div>
  );
};

export default AgileDashboard;
