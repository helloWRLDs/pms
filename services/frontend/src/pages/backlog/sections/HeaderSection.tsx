import {
  BsKanbanFill,
  BsFillPlusCircleFill,
  BsTable,
  BsGrid,
} from "react-icons/bs";
import { Button } from "../../../components/ui/Button";
import useMetaCache from "../../../store/useMetaCache";
import { usePermission } from "../../../hooks/usePermission";
import { Permissions } from "../../../lib/permission";

type ViewMode = "table" | "cards";

interface ViewToggleProps {
  view: ViewMode;
  onViewChange: (view: ViewMode) => void;
}

const ViewToggle = ({ view, onViewChange }: ViewToggleProps) => (
  <div className="flex items-center gap-2 bg-secondary-200 rounded-lg p-1">
    <button
      onClick={() => onViewChange("table")}
      className={`p-2 rounded-md transition-colors ${
        view === "table"
          ? "bg-accent-500 text-primary-700"
          : "text-neutral-300 hover:text-accent-400"
      }`}
    >
      <BsTable size={20} />
    </button>
    <button
      onClick={() => onViewChange("cards")}
      className={`p-2 rounded-md transition-colors ${
        view === "cards"
          ? "bg-accent-500 text-primary-700"
          : "text-neutral-300 hover:text-accent-400"
      }`}
    >
      <BsGrid size={20} />
    </button>
  </div>
);

interface HeaderSectionProps {
  viewMode: ViewMode;
  onViewChange: (view: ViewMode) => void;
  onCreateTask: () => void;
}

const HeaderSection = ({
  viewMode,
  onViewChange,
  onCreateTask,
}: HeaderSectionProps) => {
  const metaCache = useMetaCache();
  const { hasPermission } = usePermission();

  // Don't render if no project is selected
  if (!metaCache.metadata.selectedProject) {
    return null;
  }

  return (
    <div className="px-6 py-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
          <div>
            <h1 className="text-3xl font-bold flex items-center gap-3">
              <BsKanbanFill className="text-accent-500" />
              <span className="text-accent-500">
                {metaCache.metadata.selectedProject.title}
              </span>
              <span>Backlog</span>
            </h1>
            <p className="text-neutral-300 mt-2">
              Manage and track your project tasks
            </p>
          </div>

          <div className="flex items-center gap-4">
            <ViewToggle view={viewMode} onViewChange={onViewChange} />
            {hasPermission(Permissions.TASK_WRITE_PERMISSION) && (
              <Button
                onClick={onCreateTask}
                className="flex items-center gap-2 bg-accent-500 text-white hover:bg-accent-400"
              >
                <BsFillPlusCircleFill />
                Create Task
              </Button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default HeaderSection;
