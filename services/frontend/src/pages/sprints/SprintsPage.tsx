import { useEffect, useState } from "react";
import { usePageSettings } from "../../hooks/usePageSettings";
import { Modal } from "../../components/ui/Modal";
import NewSprintForm from "../../components/forms/NewSprintForm";
import sprintAPI from "../../api/sprintAPI";
import { useQuery } from "@tanstack/react-query";
import { Sprint, SprintFilter } from "../../lib/sprint/sprint";
import { formatTime } from "../../lib/utils/time";
import Paginator from "../../components/ui/Paginator";
import { useNavigate } from "react-router-dom";
import { BsFillPlusCircleFill, BsCalendarEvent } from "react-icons/bs";
import { FiSearch } from "react-icons/fi";
import { ListItems } from "../../lib/utils/list";
import { parseError } from "../../lib/errors";
import { useAuthStore } from "../../store/authStore";
import { Button } from "../../components/ui/Button";
import useMetaCache from "../../store/useMetaCache";
import { Layouts } from "../../lib/layout/layout";
import { toast } from "react-toastify";
import { toastOpts } from "../../lib/utils/toast";
import { AiOutlineClockCircle } from "react-icons/ai";
import { MdOutlineTaskAlt } from "react-icons/md";

const SprintsPage = () => {
  usePageSettings({
    title: "Sprints",
    requireAuth: true,
    layout: Layouts.Projects,
  });

  const { clearAuth } = useAuthStore();
  const navigate = useNavigate();
  const metaCache = useMetaCache();

  const [sprintCreationModal, setSprintCreationModal] = useState(false);
  const [search, setSearch] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const [filter, setFilter] = useState<SprintFilter>({
    page: 1,
    per_page: 12,
    project_id: metaCache.metadata.selectedProject?.id,
  });

  const {
    data: sprintList,
    isLoading,
    refetch: refetchSprints,
    error,
  } = useQuery<ListItems<Sprint>>({
    queryKey: [
      "sprints",
      filter.page,
      filter.per_page,
      filter.project_id,
      filter.title,
    ],
    queryFn: async () => {
      console.log("Sprint list query running with filter:", filter);
      try {
        const response = await sprintAPI.list(filter);
        console.log("Sprint list API response:", response);
        return response;
      } catch (e) {
        console.error("Sprint list API error:", e);
        if (parseError(e)?.status === 401) {
          clearAuth();
          navigate("/login");
        }
        throw e;
      }
    },
    enabled: !!metaCache.metadata.selectedProject?.id,
  });

  useEffect(() => {
    console.log("Selected project:", metaCache.metadata.selectedProject);
    if (!metaCache.metadata.selectedProject?.id) {
      navigate("/projects");
      toast.error("Please select a project first", toastOpts);
    }
  }, [metaCache.metadata.selectedProject?.id, navigate]);

  const handleCreateSprint = async (data: any) => {
    try {
      setIsSubmitting(true);
      await sprintAPI.create(data);
      await refetchSprints();
      setSprintCreationModal(false);
      toast.success("Sprint created successfully", toastOpts);
    } catch (error) {
      toast.error("Failed to create sprint", toastOpts);
      console.error("Error creating sprint:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="w-full h-[100vh] overflow-y-auto bg-gradient-to-br from-primary-700 to-primary-600">
      <div className="flex flex-col gap-6 p-6">
        {/* Header Section */}
        <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4 bg-primary-400 p-4 rounded-lg">
          <div className="flex flex-col md:flex-row items-start md:items-center gap-4 w-full md:w-auto">
            <h1 className="text-2xl font-semibold text-neutral-100 flex items-center gap-2">
              <BsCalendarEvent className="text-accent-500" />
              {metaCache.metadata.selectedProject?.title} Sprints
            </h1>
            <div className="relative w-full md:w-64">
              <input
                type="text"
                placeholder="Search sprints..."
                value={search}
                onChange={(e) => {
                  setSearch(e.target.value);
                  setFilter((prev) => ({
                    ...prev,
                    title: e.target.value || undefined,
                    page: 1,
                  }));
                }}
                className="block w-full px-4 py-2 text-sm text-neutral-100 bg-primary-300 border border-primary-300 rounded-lg focus:ring-accent-500 focus:border-accent-500"
              />
              <FiSearch className="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400" />
            </div>
          </div>
          <Button
            onClick={() => setSprintCreationModal(true)}
            className="flex items-center gap-2 bg-accent-500 hover:bg-accent-600 text-white px-4 py-2 rounded-lg transition-colors duration-200 w-full md:w-auto justify-center"
            disabled={isSubmitting}
          >
            <BsFillPlusCircleFill />
            New Sprint
          </Button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4 gap-4">
          {isLoading ? (
            Array.from({ length: 12 }).map((_, index) => (
              <div
                key={`skeleton-${index}`}
                className="bg-primary-400 rounded-lg p-4 animate-pulse"
              >
                <div className="h-6 bg-primary-300 rounded w-3/4 mb-4"></div>
                <div className="h-4 bg-primary-300 rounded w-1/2 mb-2"></div>
                <div className="h-4 bg-primary-300 rounded w-2/3"></div>
              </div>
            ))
          ) : error ? (
            <div className="col-span-full text-center text-red-500 bg-primary-400 rounded-lg p-4">
              Failed to load sprints. Please try again.
            </div>
          ) : !sprintList ||
            !sprintList.items ||
            sprintList.items.length === 0 ? (
            <div className="col-span-full text-center text-neutral-400 bg-primary-400 rounded-lg p-8">
              <div className="flex flex-col items-center gap-4">
                <BsCalendarEvent className="text-4xl" />
                <p>
                  No sprints found. Create your first sprint to get started!
                </p>
                <Button
                  onClick={() => setSprintCreationModal(true)}
                  className="flex items-center gap-2 bg-accent-500 hover:bg-accent-600"
                >
                  <BsFillPlusCircleFill />
                  Create Sprint
                </Button>
              </div>
            </div>
          ) : (
            sprintList?.items.map((sprint) => (
              <div
                key={sprint.id}
                onClick={() => {
                  navigate(`/sprints/${sprint.id}`);
                }}
                className="bg-primary-400 rounded-lg p-4 cursor-pointer hover:bg-primary-300 transition-all duration-200 transform hover:scale-[1.02] border border-transparent hover:border-accent-500"
              >
                <h3 className="text-lg font-semibold text-neutral-100 mb-2 flex items-center gap-2">
                  <BsCalendarEvent className="text-accent-500" />
                  {sprint.title}
                </h3>
                <p className="text-sm text-neutral-400 mb-4 line-clamp-2">
                  {sprint.description || "No description provided"}
                </p>
                <div className="flex items-center gap-4 text-xs text-neutral-300">
                  <span className="flex items-center gap-1">
                    <AiOutlineClockCircle className="text-accent-500" />
                    {formatTime(sprint.start_date.seconds)} -{" "}
                    {formatTime(sprint.end_date.seconds)}
                  </span>
                  <span className="flex items-center gap-1">
                    <MdOutlineTaskAlt className="text-accent-500" />
                    {sprint.tasks?.length || 0} tasks
                  </span>
                </div>
              </div>
            ))
          )}
        </div>

        {sprintList && sprintList.total_pages > 1 && (
          <div className="flex justify-center mt-4">
            <Paginator
              page={sprintList.page}
              per_page={sprintList.per_page}
              total_items={sprintList.total_items}
              total_pages={sprintList.total_pages}
              onPageChange={(page) => setFilter((prev) => ({ ...prev, page }))}
            />
          </div>
        )}
      </div>

      <Modal
        visible={sprintCreationModal}
        onClose={() => !isSubmitting && setSprintCreationModal(false)}
        title="Create New Sprint"
        className="bg-primary-400"
      >
        <NewSprintForm onFinish={handleCreateSprint} />
      </Modal>
    </div>
  );
};

export default SprintsPage;
