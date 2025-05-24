import {
  QueryObserverBaseResult,
  RefetchOptions,
  useQuery,
} from "@tanstack/react-query";
import { Task } from "../../lib/task/task";
import { taskAPI } from "../../api/taskAPI";
import { TaskCommentCreation, TaskCommentFilter } from "../../lib/task/comment";
import { useEffect, useState } from "react";
import { MdOutlineSend } from "react-icons/md";
import { formatTime } from "../../lib/utils/time";
import { useAuthStore } from "../../store/authStore";
import Paginator from "../ui/Paginator";
import DropDownList from "../ui/DropDown";
import { useCacheStore } from "../../store/cacheStore";
import { infoToast } from "../../lib/utils/toast";
import { ListItems } from "../../lib/utils/list";
import { User } from "../../lib/user/user";

type TaskViewProps = React.HTMLAttributes<HTMLDivElement> & {
  task: Task;
  assignees: ListItems<User> | undefined;
  refetchTasks?: (
    options?: RefetchOptions | undefined
  ) => Promise<QueryObserverBaseResult<ListItems<Task>, Error>>;
};

const TaskView = ({ task, assignees, ...props }: TaskViewProps) => {
  const { auth } = useAuthStore();

  const { getAssignee } = useCacheStore();
  const [reassignDropDown, setReassignDropDown] = useState(false);

  const [newComment, setNewComment] = useState<TaskCommentCreation>({
    body: "",
    task_id: task.id,
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
      task.id,
      filter.page,
      filter.per_page,
      filter.user_id,
    ],
    queryFn: () => taskAPI.listComments(task.id, filter),
    enabled: true,
  });

  const [addSprintDropDown, setAddSprintDropDown] = useState(false);
  const { getSprint, getProject, sprints } = useCacheStore();
  return (
    <div className={` ${props.className}`} {...props}>
      <div id="body" className="text-lg w-full">
        <div id="body-main" className="min-h-[5rem]">
          {task.body}
        </div>
        <div
          id="body-secondary"
          className="text-sm flex gap-4 text-neutral-400 flex-wrap"
        >
          <div id="body-secondary-project">
            Project:{" "}
            {task.project_id
              ? getProject(task.project_id)?.title ?? ""
              : "none"}
          </div>
          <div id="body-secondary-assignee">
            <div>
              Assignee:{" "}
              <span
                className="cursor-pointer hover:underline"
                onClick={() => {
                  setReassignDropDown(true);
                }}
              >
                {task.assignee_id
                  ? getAssignee(task.assignee_id)?.name
                  : "none"}
              </span>
            </div>
            <DropDownList
              visible={reassignDropDown}
              onClose={() => setReassignDropDown(false)}
            >
              {assignees &&
                assignees.items.length > 0 &&
                assignees.items.map((assignee) => (
                  <DropDownList.Element
                    className="w-[5rem] px-4 py-2 bg-secondary-100 cursor-pointer text-neutral-200 hover:bg-accent-500 hover:text-secondary-100"
                    onClick={async () => {
                      try {
                        if (task.assignee_id) {
                          await taskAPI.unassign(task.id, task.assignee_id);
                        }
                        await taskAPI.assign(task.id, assignee.id);
                      } catch (e) {
                        console.error(e);
                      } finally {
                        setReassignDropDown(false);
                        props.refetchTasks && props.refetchTasks();
                      }
                    }}
                  >
                    {assignee.name}
                  </DropDownList.Element>
                ))}
            </DropDownList>
          </div>
          <div id="body-secondary-code">
            Code:{" "}
            <span
              className={`${task.code && "cursor-pointer hover:underline"}`}
              onClick={() => {
                if (task.code) {
                  navigator.clipboard.writeText(task.code);
                  infoToast("Copied to clipboard!");
                }
              }}
            >
              {task.code ?? "none"}
            </span>
          </div>
          <div id="body-secondary-sprint">
            Sprint:{" "}
            <span
              className="hover:underline cursor-pointer absolute"
              onClick={() => {
                setAddSprintDropDown(true);
              }}
            >
              {task.sprint_id ? getSprint(task.sprint_id)?.title : "none"}
            </span>
            <DropDownList
              visible={addSprintDropDown}
              onClose={() => {
                setAddSprintDropDown(false);
              }}
            >
              {sprints &&
                Object.entries(sprints).map(([string, sprint]) => (
                  <DropDownList.Element
                    key={string}
                    onClick={() => {
                      console.log(sprint);
                      const updatedTask: Task = JSON.parse(
                        JSON.stringify(task)
                      );
                      updatedTask.sprint_id = sprint.id;
                      console.log(JSON.stringify(updatedTask));
                      taskAPI.update(updatedTask.id, updatedTask);
                      refetch();
                      props.refetchTasks && props.refetchTasks();
                      setAddSprintDropDown(false);
                    }}
                    className="w-[10rem] px-4 py-2 bg-secondary-100 cursor-pointer text-neutral-200 hover:bg-accent-500 hover:text-secondary-100"
                  >
                    {sprint.title}
                  </DropDownList.Element>
                ))}
            </DropDownList>
          </div>
        </div>
      </div>
      <div id="comment" className="mt-8">
        <div id="comment-new" className="flex items-start gap-4 mb-4">
          <textarea
            placeholder="Add a comment"
            value={newComment.body}
            onChange={(e) => {
              setNewComment({ ...newComment, body: e.currentTarget.value });
            }}
            name="new-comment"
            id="new-comment"
            className="px-2 py-1 rounded-md w-full h-[4rem] border border-black"
          ></textarea>
          <button className="px-2 py-2 rounded-md group cursor-pointer bg-primary-500 text-accent-500 hover:bg-accent-500 group:transition-all duration-200">
            <MdOutlineSend
              className="group-hover:text-primary-500 text-accent-500"
              onClick={async () => {
                await taskAPI.createComment(newComment);
                setNewComment({ ...newComment, body: "" });
                refetch();
              }}
            />
          </button>
        </div>
        <div>
          {isCommentDataLoading ? (
            <p>Loading...</p>
          ) : commentData?.items?.length === 0 ||
            !commentData ||
            !commentData.items ? (
            <p>No comments found</p>
          ) : (
            commentData &&
            commentData.items
              .sort((a, b) => b.created_at.seconds - a.created_at.seconds)
              .map((comment, i) => (
                <div key={i} className="">
                  <div className="text-xs flex items-center gap-2  w-fit px-2 py-1 rounded-tr-md rounded-tl-md">
                    <img
                      src={`data:image/jpeg;base64,${comment.user.avatar_img}`}
                      className="w-5 h-5 rounded-full bg-neutral-300"
                    />
                    <div>{comment.user.name}</div>
                  </div>
                  <div className="text-sm px-2 py-1   rounded-md rounded-tl-none relative mb-4">
                    <div className="max-w-[80%]">
                      <span className="">{comment.body}</span>
                    </div>
                    <div className="absolute top-[0.4rem] right-[0.5rem] text-sm text-gray-500">
                      {formatTime(comment.created_at.seconds)}
                    </div>
                  </div>
                </div>
              ))
          )}
          {commentData && commentData.items && commentData.items.length > 0 && (
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
          )}
        </div>
      </div>
    </div>
  );
};

export default TaskView;
