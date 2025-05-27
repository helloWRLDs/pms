import { Project } from "../../lib/project/project";
import { formatTime } from "../../lib/utils/time";
import Badge from "../ui/Badge";

type ProjectCardWrapperProps = React.HTMLAttributes<HTMLDivElement> & {};
const ProjectCardWrapper = ({
  className,
  children,
  ...props
}: ProjectCardWrapperProps) => {
  return (
    <div className={`flex w-full gap-4 ${className}`} {...props}>
      {children}
    </div>
  );
};

type ProjectCardProps = React.HTMLAttributes<HTMLDivElement> & {
  project?: Project;
};
const ProjectCard = ({
  className,
  children,
  project,
  ...props
}: ProjectCardProps) => {
  return (
    <div
      className={`rounded-lg shadow  backdrop-opacity-100 ${className}`}
      {...props}
    >
      {project ? (
        <div>
          <div id="card-header" className="flex gap-2 items-center">
            <h4>{project.title}</h4>
            <Badge className="inline bg-accent-600 text-black">
              {project.code_name}
            </Badge>
          </div>
          <div>
            <Badge className="bg-blue-500 text-neutral-100">
              {project.status}
            </Badge>
          </div>
          <div
            id="card-progress"
            className="flex items-center justify-between gap-5"
          >
            <div
              className="bg-green-500 rounded-full h-2 transition-all duration-300"
              style={{
                width: `${
                  (project.done_tasks / project.total_tasks) * 100
                    ? (project.done_tasks / project.total_tasks) * 100
                    : 1
                }%`,
              }}
            ></div>
            {project.done_tasks && project.total_tasks ? (
              <p>
                {Math.round((project.done_tasks / project.total_tasks) * 100)}%
              </p>
            ) : (
              <p>No Tasks</p>
            )}
          </div>
          <div>{formatTime(project.created_at.seconds)}</div>
        </div>
      ) : (
        <>{children}</>
      )}
    </div>
  );
};

ProjectCardWrapper.Card = ProjectCard;

export default ProjectCardWrapper;
