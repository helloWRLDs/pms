import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { BsFillPlusCircleFill } from "react-icons/bs";
import { TbCalendarEvent } from "react-icons/tb";
import { Modal } from "../../../components/ui/Modal";
import Table from "../../../components/ui/Table";
import { Button } from "../../../components/ui/Button";
import Paginator from "../../../components/ui/Paginator";
import sprintAPI from "../../../api/sprintAPI";
import { usePermission } from "../../../hooks/usePermission";
import { Permissions } from "../../../lib/permission";
import {
  Sprint,
  SprintCreation,
  SprintFilter,
} from "../../../lib/sprint/sprint";
import { formatTime } from "../../../lib/utils/time";
import { infoToast, errorToast } from "../../../lib/utils/toast";
import NewSprintForm from "../../../components/forms/NewSprintForm";
import { useNavigate } from "react-router-dom";
interface SprintsSectionProps {
  projectId: string;
}

const SprintsSection = ({ projectId }: SprintsSectionProps) => {
  const { hasPermission } = usePermission();
  const navigate = useNavigate();
  const [newSprintModal, setNewSprintModal] = useState(false);
  const [editSprintModal, setEditSprintModal] = useState(false);
  const [selectedSprint, setSelectedSprint] = useState<Sprint | null>(null);

  const [sprintFilter, setSprintFilter] = useState<SprintFilter>({
    project_id: projectId,
    page: 1,
    per_page: 10,
  });

  const {
    data: sprints,
    refetch: refetchSprints,
    isLoading,
  } = useQuery({
    queryKey: [
      "sprints",
      sprintFilter.project_id,
      sprintFilter.page,
      sprintFilter.per_page,
    ],
    queryFn: () => sprintAPI.list(sprintFilter),
    enabled: !!projectId && hasPermission(Permissions.SPRINT_READ_PERMISSION),
  });

  const handleCreateSprint = async (sprintData: SprintCreation) => {
    try {
      await sprintAPI.create(sprintData);
      infoToast("Sprint created successfully");
      setNewSprintModal(false);
      await refetchSprints();
    } catch (error) {
      errorToast("Failed to create sprint");
      console.error("Failed to create sprint:", error);
    }
  };

  const handleUpdateSprint = async (sprintData: Sprint) => {
    try {
      if (selectedSprint) {
        await sprintAPI.update(selectedSprint.id, sprintData);
        infoToast("Sprint updated successfully");
        setEditSprintModal(false);
        setSelectedSprint(null);
        await refetchSprints();
      }
    } catch (error) {
      errorToast("Failed to update sprint");
      console.error("Failed to update sprint:", error);
    }
  };

  const getSprintStatusBadge = (sprint: Sprint) => {
    const now = new Date();
    const startDate = new Date(sprint.start_date.seconds * 1000);
    const endDate = new Date(sprint.end_date.seconds * 1000);

    if (now < startDate) {
      return (
        <span className="bg-blue-500/20 text-blue-400 px-2 py-1 rounded-full text-xs font-semibold uppercase tracking-wide border border-blue-500/30">
          Planned
        </span>
      );
    } else if (now >= startDate && now <= endDate) {
      return (
        <span className="bg-green-500/20 text-green-400 px-2 py-1 rounded-full text-xs font-semibold uppercase tracking-wide border border-green-500/30">
          Active
        </span>
      );
    } else {
      return (
        <span className="bg-gray-500/20 text-gray-400 px-2 py-1 rounded-full text-xs font-semibold uppercase tracking-wide border border-gray-500/30">
          Completed
        </span>
      );
    }
  };

  // Don't render if user doesn't have permission to read sprints
  if (!hasPermission(Permissions.SPRINT_READ_PERMISSION)) {
    return null;
  }

  return (
    <section className="mb-10">
      <Modal
        title="Create New Sprint"
        visible={newSprintModal}
        onClose={() => setNewSprintModal(false)}
        className="w-[60%] mx-auto bg-primary-300 text-white"
      >
        <NewSprintForm
          onFinish={(data) => {
            handleCreateSprint(data);
            setNewSprintModal(false);
          }}
        />
      </Modal>

      <Modal
        title="Edit Sprint"
        visible={editSprintModal}
        onClose={() => {
          setEditSprintModal(false);
          setSelectedSprint(null);
        }}
        className="w-[60%] mx-auto bg-primary-300 text-white"
      >
        {selectedSprint && (
          <SprintForm
            initialSprint={selectedSprint}
            onSubmit={handleUpdateSprint}
            onCancel={() => {
              setEditSprintModal(false);
              setSelectedSprint(null);
            }}
            projectId={projectId}
            isEditing
          />
        )}
      </Modal>

      <div className="px-6">
        <div className="max-w-7xl mx-auto">
          <div className="flex items-center justify-between mb-5">
            <h2 className="text-2xl font-semibold text-white flex items-center gap-3">
              <TbCalendarEvent className="text-accent-500" />
              Sprints
            </h2>
            {hasPermission(Permissions.SPRINT_WRITE_PERMISSION) && (
              <Button
                onClick={() => setNewSprintModal(true)}
                className="bg-accent-500 hover:bg-accent-600 text-white px-4 py-2 rounded-md flex items-center gap-2"
              >
                <BsFillPlusCircleFill size={16} />
                Create Sprint
              </Button>
            )}
          </div>

          <div className="bg-secondary-200/50 backdrop-blur-sm rounded-lg border border-secondary-300/30 overflow-visible">
            <table className="w-full">
              <Table.Head className="bg-primary-400/70 text-white">
                <Table.HeadCell>Sprint Name</Table.HeadCell>
                <Table.HeadCell>Status</Table.HeadCell>
                <Table.HeadCell>Start Date</Table.HeadCell>
                <Table.HeadCell>End Date</Table.HeadCell>
                <Table.HeadCell>Duration</Table.HeadCell>
                <Table.HeadCell>Actions</Table.HeadCell>
              </Table.Head>
              <Table.Body className="divide-y divide-secondary-300/20">
                {isLoading
                  ? Array(3)
                      .fill(0)
                      .map((_, index) => (
                        <Table.Row
                          key={index}
                          className="bg-secondary-200/30 animate-pulse"
                        >
                          <Table.Cell>
                            <div className="h-4 bg-secondary-100 rounded w-3/4"></div>
                          </Table.Cell>
                          <Table.Cell>
                            <div className="h-6 bg-secondary-100 rounded w-20"></div>
                          </Table.Cell>
                          <Table.Cell>
                            <div className="h-4 bg-secondary-100 rounded w-24"></div>
                          </Table.Cell>
                          <Table.Cell>
                            <div className="h-4 bg-secondary-100 rounded w-24"></div>
                          </Table.Cell>
                          <Table.Cell>
                            <div className="h-4 bg-secondary-100 rounded w-16"></div>
                          </Table.Cell>
                          <Table.Cell>
                            <div className="h-4 bg-secondary-100 rounded w-8"></div>
                          </Table.Cell>
                        </Table.Row>
                      ))
                  : sprints?.items?.map((sprint) => {
                      const duration = Math.ceil(
                        (sprint.end_date.seconds - sprint.start_date.seconds) /
                          (24 * 60 * 60)
                      );

                      return (
                        <Table.Row
                          key={sprint.id}
                          className="bg-secondary-200/30 hover:bg-secondary-100/40 text-white transition-colors"
                          onClick={() => {
                            navigate(`/sprints/${sprint.id}`);
                          }}
                        >
                          <Table.Cell>
                            <div className="flex items-center gap-3">
                              <div className="p-2 bg-accent-500/20 rounded-full">
                                <TbCalendarEvent
                                  className="text-accent-400"
                                  size={16}
                                />
                              </div>
                              <div>
                                <div className="font-semibold text-white">
                                  {sprint.title}
                                </div>
                                {sprint.description && (
                                  <div className="text-sm text-white/60 truncate max-w-xs">
                                    {sprint.description}
                                  </div>
                                )}
                              </div>
                            </div>
                          </Table.Cell>
                          <Table.Cell>
                            {getSprintStatusBadge(sprint)}
                          </Table.Cell>
                          <Table.Cell className="text-white/80">
                            {formatTime(sprint.start_date.seconds)}
                          </Table.Cell>
                          <Table.Cell className="text-white/80">
                            {formatTime(sprint.end_date.seconds)}
                          </Table.Cell>
                          <Table.Cell className="text-white/80">
                            {duration} days
                          </Table.Cell>
                          <Table.Cell>
                            {/* {(hasPermission(
                              Permissions.SPRINT_WRITE_PERMISSION
                            ) ||
                              hasPermission(
                                Permissions.SPRINT_DELETE_PERMISSION
                              )) && (
                              <ContextMenu
                                placement="left"
                                trigger={<BsThreeDotsVertical />}
                                items={[
                                  ...(hasPermission(
                                    Permissions.SPRINT_WRITE_PERMISSION
                                  )
                                    ? [
                                        {
                                          icon: <TbEdit />,
                                          label: "Edit Sprint",
                                          onClick: () => {
                                            setSelectedSprint(sprint);
                                            setEditSprintModal(true);
                                          },
                                        },
                                      ]
                                    : []),
                                ]}
                              />
                            )} */}
                          </Table.Cell>
                        </Table.Row>
                      );
                    })}
              </Table.Body>
            </table>

            {hasPermission(Permissions.SPRINT_WRITE_PERMISSION) && (
              <button
                className="w-full cursor-pointer group hover:bg-secondary-100/40 py-4 transition-all duration-300"
                onClick={() => setNewSprintModal(true)}
              >
                <BsFillPlusCircleFill
                  size="30"
                  className="mx-auto text-white/60 group-hover:text-accent-400"
                />
              </button>
            )}

            {(!sprints?.items || sprints.items.length === 0) && (
              <div className="text-center py-8 text-white/70">
                <TbCalendarEvent
                  size={48}
                  className="mx-auto mb-4 opacity-50"
                />
                <p>
                  No sprints found. Create your first sprint to get started.
                </p>
              </div>
            )}
          </div>

          {sprints && sprints.items && sprints.total_items > 0 && (
            <div className="mt-4">
              <Paginator
                page={sprints.page ?? 0}
                per_page={sprints.per_page ?? 0}
                total_items={sprints.total_items ?? 0}
                total_pages={sprints.total_pages ?? 0}
                onPageChange={(page) => {
                  setSprintFilter({ ...sprintFilter, page: page });
                }}
              />
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

// Simple Sprint Form Component (we'll need to create a proper form)
const SprintForm = ({
  initialSprint,
  onSubmit,
  onCancel,
  projectId,
  isEditing = false,
}: {
  initialSprint?: Sprint;
  onSubmit: (data: any) => void;
  onCancel: () => void;
  projectId: string;
  isEditing?: boolean;
}) => {
  const [formData, setFormData] = useState({
    title: initialSprint?.title || "",
    description: initialSprint?.description || "",
    start_date: initialSprint
      ? new Date(initialSprint.start_date.seconds * 1000)
          .toISOString()
          .split("T")[0]
      : "",
    end_date: initialSprint
      ? new Date(initialSprint.end_date.seconds * 1000)
          .toISOString()
          .split("T")[0]
      : "",
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const sprintData = {
      ...formData,
      project_id: projectId,
      start_date: {
        seconds: Math.floor(new Date(formData.start_date).getTime() / 1000),
        nanos: 0,
      },
      end_date: {
        seconds: Math.floor(new Date(formData.end_date).getTime() / 1000),
        nanos: 0,
      },
    };

    if (isEditing && initialSprint) {
      onSubmit({ ...initialSprint, ...sprintData });
    } else {
      onSubmit(sprintData);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label className="block text-sm font-medium mb-2">Sprint Title</label>
        <input
          type="text"
          value={formData.title}
          onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          className="w-full px-3 py-2 bg-secondary-100 border border-secondary-200 rounded-md text-neutral-800"
          required
        />
      </div>

      <div>
        <label className="block text-sm font-medium mb-2">Description</label>
        <textarea
          value={formData.description}
          onChange={(e) =>
            setFormData({ ...formData, description: e.target.value })
          }
          className="w-full px-3 py-2 bg-secondary-100 border border-secondary-200 rounded-md text-neutral-800"
          rows={3}
        />
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div>
          <label className="block text-sm font-medium mb-2">Start Date</label>
          <input
            type="date"
            value={formData.start_date}
            onChange={(e) =>
              setFormData({ ...formData, start_date: e.target.value })
            }
            className="w-full px-3 py-2 bg-secondary-100 border border-secondary-200 rounded-md text-neutral-800"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">End Date</label>
          <input
            type="date"
            value={formData.end_date}
            onChange={(e) =>
              setFormData({ ...formData, end_date: e.target.value })
            }
            className="w-full px-3 py-2 bg-secondary-100 border border-secondary-200 rounded-md text-neutral-800"
            required
          />
        </div>
      </div>

      <div className="flex justify-end gap-3 pt-4">
        <Button
          type="button"
          onClick={onCancel}
          className="px-4 py-2 bg-secondary-200 text-neutral-700 hover:bg-secondary-300"
        >
          Cancel
        </Button>
        <Button
          type="submit"
          className="px-4 py-2 bg-accent-500 text-white hover:bg-accent-600"
        >
          {isEditing ? "Update Sprint" : "Create Sprint"}
        </Button>
      </div>
    </form>
  );
};

export default SprintsSection;
