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
  project: Project;
};
const ProjectCard = ({
  className,
  children,
  project,
  ...props
}: ProjectCardProps) => {
  return (
    <div
      className="bg-white px-3 py-6 w-full rounded-lg shadow border border-black backdrop-opacity-100"
      {...props}
    >
      <div id="card-header" className="flex gap-2 items-center">
        <h4>{project.title}</h4>
        <Badge className="inline bg-accent-600 text-black">
          {project.code_name}
        </Badge>
      </div>
      <div>
        <Badge className="bg-blue-500 text-white">{project.status}</Badge>
      </div>
      <div
        id="card-progress"
        className="flex items-center justify-between gap-5"
      >
        <div
          className="bg-[rgb(41,43,41)] rounded-full h-2 transition-all duration-300"
          style={{
            width: `${
              (project.done_tasks / project.total_tasks) * 100
                ? (project.done_tasks / project.total_tasks) * 100
                : 1
            }%`,
          }}
        ></div>
        {project.done_tasks && project.total_tasks ? (
          <p>{(project.done_tasks / project.total_tasks) * 100}%</p>
        ) : (
          <p>No Tasks</p>
        )}
      </div>
      <div>{formatTime(project.created_at.seconds)}</div>
    </div>
  );
};

ProjectCardWrapper.Card = ProjectCard;

export default ProjectCardWrapper;
