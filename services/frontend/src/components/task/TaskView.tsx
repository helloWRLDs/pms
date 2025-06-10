import { useQuery } from "@tanstack/react-query";
import { taskAPI } from "../../api/taskAPI";
import { TaskCommentCreation, TaskCommentFilter } from "../../lib/task/comment";
import { useEffect, useState } from "react";
import {
  MdOutlineSend,
  MdAssignment,
  MdCode,
  MdTimer,
  MdFolder,
  MdEdit,
  MdCategory,
} from "react-icons/md";
import { formatTime } from "../../lib/utils/time";
import { useAuthStore } from "../../store/authStore";
import Paginator from "../ui/Paginator";
import { infoToast } from "../../lib/utils/toast";
import useMetaCache from "../../store/useMetaCache";
import { getTaskTypeConfig } from "../../lib/task/tasktype";
import Editor from "../text/Editor";
import HTMLReactParser from "html-react-parser";
import {
  useAssigneeList,
  useProjectList,
  useSprintList,
} from "../../hooks/useData";
import DropDownable from "../ui/DropDownable";

type TaskViewProps = React.HTMLAttributes<HTMLDivElement> & {
  task_id: string;
  refetchTasks?: () => void;
};

const MetadataItem: React.FC<{
  icon: React.ReactNode;
  label: string;
  value: string | React.ReactNode;
  onClick?: () => void;
  copyable?: boolean;
  children?: React.ReactNode;
}> = ({ icon, label, value, onClick, copyable, children }) => (
  <div className="relative">
    <div className="flex items-center gap-1.5 px-2 py-1.5 bg-secondary-400/5 rounded text-xs min-w-0">
      <span className="text-accent-500 flex-shrink-0">{icon}</span>
      <span className="text-neutral-400 font-medium flex-shrink-0">
        {label}:
      </span>
      <span
        className={`text-neutral-200 truncate min-w-0 ${
          onClick || copyable
            ? "cursor-pointer hover:text-accent-500 transition-colors"
            : ""
        }`}
        onClick={onClick}
        title={typeof value === "string" ? value : undefined}
      >
        {value || "none"}
      </span>
    </div>
    {children}
  </div>
);

const TaskView = ({ task_id, refetchTasks, ...props }: TaskViewProps) => {
  const { auth } = useAuthStore();
  const metadata = useMetaCache();
  const { getSprintName, sprints } = useSprintList(
    metadata.metadata.selectedProject?.id ?? ""
  );
  const { getProjectName } = useProjectList(
    metadata.metadata.selectedCompany?.id ?? ""
  );
  const { getAssigneeName, assignees } = useAssigneeList(
    metadata.metadata.selectedCompany?.id ?? ""
  );

  const { data: task, refetch: refetchTask } = useQuery({
    queryKey: ["task", task_id],
    queryFn: () => taskAPI.get(task_id),
    enabled: !!task_id,
  });

  useEffect(() => {
    if (task) {
      console.log(JSON.stringify(task, null, 2));
    }
  }, [task]);

  const [showCommentEditor, setShowCommentEditor] = useState(false);

  const [newComment, setNewComment] = useState<TaskCommentCreation>({
    body: "",
    task_id: task_id,
    user_id: auth?.user.id || "",
  });

  const [filter, setFilter] = useState<TaskCommentFilter>({
    page: 1,
    per_page: 5,
    task_id: "",
    user_id: "",
  });

  const {
    data: commentData,
    isLoading: isCommentDataLoading,
    refetch,
  } = useQuery({
    queryKey: [
      "task-comment",
      task_id,
      filter.page,
      filter.per_page,
      filter.user_id,
    ],
    queryFn: () => taskAPI.listComments(task_id, filter),
    enabled: true,
  });

  return (
    <div className={`space-y-3 ${props.className}`} {...props}>
      {task ? (
        <>
          {/* Task Description */}
          <div className="bg-secondary-400/5 border border-secondary-200/20 rounded-lg">
            <div className="flex items-center justify-between mb-3">
              <h3 className="text-base font-semibold text-neutral-200 flex items-center gap-2">
                <MdEdit className="text-accent-500" />
                Body
              </h3>
            </div>
            <div className="prose prose-invert prose-sm max-w-none text-neutral-200 leading-relaxed text-md">
              {task.body ? (
                HTMLReactParser(task.body)
              ) : (
                <p className="text-neutral-400 italic">
                  No description provided.
                </p>
              )}
            </div>
          </div>

          {/* Compact Metadata Bar */}
          <div className="bg-secondary-300/20 rounded-lg p-2 border border-secondary-200/30">
            <div className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-1.5">
              <MetadataItem
                icon={<MdFolder className="text-sm" />}
                label="Project"
                value={getProjectName(task.project_id ?? "") ?? "none"}
              />
              <DropDownable
                options={[
                  {
                    label: "unassigned",
                    isActive: !task.assignee_id,
                    onClick: async () => {
                      await taskAPI.unassign(task_id, task.assignee_id ?? "");
                      refetchTask();
                    },
                  },
                  ...(assignees?.items && assignees?.items.length > 0
                    ? assignees?.items.map((assignee) => ({
                        label: `${assignee.first_name} ${assignee.last_name}`,
                        isActive: assignee.id === task.assignee_id,
                        onClick: async () => {
                          await taskAPI.assign(task_id, assignee.id);
                          refetchTask();
                          refetchTasks?.();
                        },
                      }))
                    : []),
                ]}
              >
                <MetadataItem
                  icon={<MdAssignment className="text-sm" />}
                  label="Assignee"
                  value={getAssigneeName(task.assignee_id)}
                />
              </DropDownable>
              <MetadataItem
                icon={<MdCode className="text-sm" />}
                label="Code"
                value={task.code || "none"}
                copyable
                onClick={() => {
                  if (task.code) {
                    navigator.clipboard.writeText(task.code);
                    infoToast("Copied to clipboard!");
                  }
                }}
              />
              <MetadataItem
                icon={<MdCategory className="text-sm" />}
                label="Type"
                value={
                  task.type ? (
                    <div className="flex items-center gap-1">
                      <span className="text-sm">
                        {getTaskTypeConfig(task.type as any).icon}
                      </span>
                      <span>{getTaskTypeConfig(task.type as any).label}</span>
                    </div>
                  ) : (
                    "none"
                  )
                }
              />
              <DropDownable
                options={[
                  {
                    label: "none",
                    isActive: !task.sprint_id,
                    onClick: async () => {
                      await taskAPI.update(task_id, { ...task, sprint_id: "" });
                      refetchTask();
                      refetchTasks && refetchTasks();
                    },
                  },
                  ...(sprints?.items && sprints?.items.length > 0
                    ? sprints?.items.map((sprint) => ({
                        label: sprint.title,
                        isActive: sprint.id === task.sprint_id,
                        onClick: async () => {
                          await taskAPI.update(task_id, {
                            ...task,
                            sprint_id: sprint.id,
                          });
                          refetchTask();
                          refetchTasks && refetchTasks();
                        },
                      }))
                    : []),
                ]}
              >
                <MetadataItem
                  icon={<MdTimer className="text-sm" />}
                  label="Sprint"
                  value={getSprintName(task.sprint_id)}
                />
              </DropDownable>
            </div>
          </div>

          {/* Comments Section */}
          <div className="space-y-3">
            <div className="flex items-center justify-between border-b border-secondary-200/20 pb-2">
              <h3 className="text-base font-semibold text-neutral-200">
                Comments
              </h3>
              <button
                onClick={() => setShowCommentEditor(!showCommentEditor)}
                className="px-3 py-1.5 text-xs rounded-lg bg-accent-500 text-secondary-100 hover:bg-accent-400 transition-colors flex items-center gap-1.5"
              >
                <MdEdit className="text-sm" />
                {showCommentEditor ? "Cancel" : "Add Comment"}
              </button>
            </div>

            {/* New Comment Editor (Collapsible) */}
            {showCommentEditor && (
              <div className="space-y-2 bg-secondary-300/10 border border-secondary-200/30 rounded-lg p-3">
                <div className="bg-secondary-400/5 border border-secondary-300/50 rounded-lg overflow-hidden">
                  <Editor
                    content={newComment.body}
                    onChange={(content) =>
                      setNewComment({ ...newComment, body: content })
                    }
                  />
                </div>
                <div className="flex justify-end gap-2">
                  <button
                    onClick={() => {
                      setShowCommentEditor(false);
                      setNewComment({ ...newComment, body: "" });
                    }}
                    className="px-4 py-1.5 text-xs rounded-lg bg-secondary-100 text-neutral-200 hover:bg-secondary-50 transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    className="px-4 py-1.5 text-xs rounded-lg bg-accent-500 text-secondary-100 hover:bg-accent-400 transition-colors flex items-center gap-1.5"
                    onClick={async () => {
                      if (!newComment.body.trim()) return;
                      await taskAPI.createComment(newComment);
                      setNewComment({ ...newComment, body: "" });
                      setShowCommentEditor(false);
                      refetch();
                    }}
                  >
                    <MdOutlineSend className="text-sm" />
                    Post Comment
                  </button>
                </div>
              </div>
            )}

            {/* Comments List */}
            <div className="space-y-2">
              {isCommentDataLoading ? (
                <div className="text-center py-6 text-neutral-400">
                  <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-accent-500 mx-auto mb-2"></div>
                  Loading comments...
                </div>
              ) : !commentData?.items?.length ? (
                <div className="text-center py-6 text-neutral-400 bg-secondary-400/5 rounded-lg">
                  <div className="text-3xl mb-1">💬</div>
                  <div className="text-sm">No comments yet</div>
                </div>
              ) : (
                <div className="space-y-2">
                  {commentData.items
                    .sort((a, b) => b.created_at.seconds - a.created_at.seconds)
                    .map((comment, i) => (
                      <div
                        key={i}
                        className="bg-secondary-400/5 border border-secondary-200/20 rounded-lg p-3 hover:bg-secondary-400/10 transition-colors"
                      >
                        <div className="flex items-center gap-2 mb-2">
                          {comment.user.avatar_url ? (
                            <img
                              src={comment.user.avatar_url}
                              alt=""
                              className="w-8 h-8 rounded-full bg-neutral-300 ring-2 ring-secondary-200/30"
                            />
                          ) : comment.user.avatar_img ? (
                            <img
                              src={`data:image/jpeg;base64,${comment.user.avatar_img}`}
                              alt={`${comment.user.first_name}'s avatar`}
                              className="w-8 h-8 rounded-full bg-neutral-300 ring-2 ring-secondary-200/30"
                            />
                          ) : (
                            <div className="w-8 h-8 rounded-full bg-neutral-300 ring-2 ring-secondary-200/30"></div>
                          )}
                          <div className="flex-1 min-w-0">
                            <div className="text-xs font-semibold text-neutral-200 truncate">
                              {comment.user.first_name} {comment.user.last_name}
                            </div>
                            <div className="text-xs text-neutral-400">
                              {formatTime(comment.created_at.seconds)}
                            </div>
                          </div>
                        </div>
                        <div className="text-sm text-neutral-200 leading-relaxed ml-10">
                          {HTMLReactParser(comment.body)}
                        </div>
                      </div>
                    ))}
                </div>
              )}

              {commentData &&
                commentData.items &&
                commentData.items.length > 0 && (
                  <div className="flex justify-center pt-2">
                    <Paginator
                      page={filter.page}
                      per_page={filter.per_page}
                      total_items={commentData.total_items}
                      total_pages={commentData.total_pages}
                      onPageChange={(page) => {
                        setFilter({ ...filter, page });
                        refetch();
                      }}
                    />
                  </div>
                )}
            </div>
          </div>
        </>
      ) : (
        <div className="text-center py-6 text-neutral-400">
          <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-accent-500 mx-auto mb-2"></div>
          Loading task...
        </div>
      )}
    </div>
  );
};

export default TaskView;
