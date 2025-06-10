import { useEffect, useState } from "react";
import NewTaskForm from "../../components/forms/NewTaskForm";
import { useNavigate } from "react-router-dom";
import { Modal } from "../../components/ui/Modal";
import { taskAPI } from "../../api/taskAPI";
import { infoToast, errorToast } from "../../lib/utils/toast";
import { TaskCreation, TaskFilter } from "../../lib/task/task";
import { usePageSettings } from "../../hooks/usePageSettings";
import useMetaCache from "../../store/useMetaCache";
import { usePermission } from "../../hooks/usePermission";
import { Permissions } from "../../lib/permission";

// Section Components
import { HeaderSection, TasksSection, SprintsSection } from "./sections";

type ViewMode = "table" | "cards";

const BacklogPage = () => {
  usePageSettings({ requireAuth: true, title: "Backlog" });

  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const navigate = useNavigate();
  const [viewMode, setViewMode] = useState<ViewMode>("table");

  const [taskCreationModal, setTaskCreationModal] = useState(false);
  const [taskFilter, setTaskFilter] = useState<TaskFilter>({
    page: 1,
    per_page: 5,
    project_id: metaCache.metadata.selectedProject?.id,
  });

  useEffect(() => {
    if (!metaCache.metadata.selectedProject?.id) {
      navigate("/projects");
    }
  }, [metaCache.metadata.selectedProject, navigate]);

  useEffect(() => {
    setTaskFilter((prev) => ({
      ...prev,
      project_id: metaCache.metadata.selectedProject?.id,
    }));
  }, [metaCache.metadata.selectedProject?.id]);

  const handleCreateTask = async (data: TaskCreation) => {
    try {
      await taskAPI.create(data);
      infoToast("Task created successfully");
      setTaskCreationModal(false);
    } catch (error) {
      errorToast("Failed to create task");
      console.error(error);
    }
  };

  const handleViewChange = (newView: ViewMode) => {
    setViewMode(newView);
    setTaskFilter((prev) => ({
      ...prev,
      per_page: newView === "cards" ? 12 : 10,
      page: 1,
    }));
  };

  // Don't render anything if no project is selected
  if (!metaCache.metadata.selectedProject?.id) {
    return (
      <div className="h-[100lvh] bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100 flex items-center justify-center overflow-auto">
        <div className="text-center">
          <h2 className="text-2xl font-semibold mb-4">Loading Project...</h2>
          <p className="text-neutral-200">
            Please wait while we load your project data.
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="h-[100lvh] bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100 overflow-auto">
      {/* Header Section */}
      <HeaderSection
        viewMode={viewMode}
        onViewChange={handleViewChange}
        onCreateTask={() => setTaskCreationModal(true)}
      />

      {/* Tasks Section with integrated filters */}
      {hasPermission(Permissions.TASK_READ_PERMISSION) && (
        <TasksSection
          filter={taskFilter}
          viewMode={viewMode}
          onFilterChange={setTaskFilter}
        />
      )}

      {/* Sprints Section */}
      {hasPermission(Permissions.SPRINT_READ_PERMISSION) && (
        <SprintsSection projectId={metaCache.metadata.selectedProject.id} />
      )}

      {/* No Permissions Message */}
      {!hasPermission(Permissions.TASK_READ_PERMISSION) &&
        !hasPermission(Permissions.SPRINT_READ_PERMISSION) && (
          <div className="px-6 pb-8">
            <div className="max-w-7xl mx-auto text-center py-20">
              <div className="bg-secondary-200/50 backdrop-blur-sm rounded-lg border border-secondary-300/30 p-12">
                <h2 className="text-2xl font-semibold mb-4 text-white/80">
                  Access Restricted
                </h2>
                <p className="text-neutral-200">
                  You don't have permission to view any sections of this project
                  backlog. Please contact your administrator for access.
                </p>
              </div>
            </div>
          </div>
        )}

      {/* Modals */}
      <Modal
        title="Create New Task"
        visible={taskCreationModal}
        onClose={() => setTaskCreationModal(false)}
        className="w-full max-w-2xl mx-auto bg-secondary-300 max-h-[90vh] overflow-y-auto"
        size="lg"
      >
        {metaCache.metadata.selectedProject && (
          <div className="p-1">
            <NewTaskForm
              project={metaCache.metadata.selectedProject}
              onSubmit={handleCreateTask}
            />
          </div>
        )}
      </Modal>
    </div>
  );
};

export default BacklogPage;
