import { useEffect, useState } from "react";
import NewTaskForm from "../components/forms/NewTaskForm";
import { useProjectStore } from "../store/selectedProjectStore";
import { useNavigate } from "react-router-dom";
import { Modal } from "../components/ui/Modal";
import { taskAPI } from "../api/taskAPI";
import { infoToast } from "../lib/utils/toast";

type Options = {};

const BacklogPage = (opts: Options) => {
  const { project } = useProjectStore();
  const navigate = useNavigate();

  const [taskCreationModal, setTaskCreationModal] = useState(true);

  useEffect(() => {
    console.log(project);
    if (!project) {
      navigate("/projects");
    }
  }, []);

  return (
    <>
      <section>
        <button onClick={() => setTaskCreationModal(true)}>Create Task</button>
        <Modal
          title="Create New Task"
          visible={taskCreationModal}
          onClose={() => {
            setTaskCreationModal(false);
          }}
          className="w-1/2 mx-auto text-sm"
        >
          {project && (
            <NewTaskForm
              project={project}
              onSubmit={(taskCreation) => {
                console.log(taskCreation);
                taskAPI
                  .createTask(taskCreation)
                  .then((res) => {
                    console.log(res);
                    infoToast("Task created successfully");
                    setTaskCreationModal(false);
                  })
                  .catch((error) => {
                    console.error(error);
                  });
              }}
            />
          )}
        </Modal>
      </section>

      {/* <section id="new-task">
        <h2>Create Task</h2>
        {project && (
          <NewTaskForm
            project={project}
            onSubmit={(taskCreation) => {
              console.log(taskCreation);
            }}
          />
        )}
      </section> */}
    </>
  );
};

export default BacklogPage;
