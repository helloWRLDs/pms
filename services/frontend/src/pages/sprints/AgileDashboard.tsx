import { FC, useEffect, useState } from "react";
import DashboardDndKit from "../../components/dashboard/DashboardDndKit";
import { usePageSettings } from "../../hooks/usePageSettings";
import { useNavigate, useParams } from "react-router-dom";
import { Modal } from "../../components/ui/Modal";

const AgileDashboard: FC = () => {
  usePageSettings({ requireAuth: true, title: "Agile Dashboard" });

  const navigate = useNavigate();

  const [manageTasksModal, setManageTasksModal] = useState(false);

  const sprintID = useParams().sprintID;
  console.log(sprintID);

  useEffect(() => {
    if (!sprintID) {
      navigate("/sprints");
      return;
    }
  }, [navigate, sprintID]);

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

      <section id="sprint-agile-dashboard">
        {!!sprintID && (
          <DashboardDndKit sprintID={sprintID} className="w-full" />
        )}
      </section>
    </div>
  );
};

export default AgileDashboard;
