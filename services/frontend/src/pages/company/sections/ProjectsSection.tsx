import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { BsFillPlusCircleFill } from "react-icons/bs";
import ProjectCardWrapper from "../../../components/prioject/ProjectCard";
import { Modal } from "../../../components/ui/Modal";
import NewProjectForm from "../../../components/forms/NewProjectForm";
import projectAPI from "../../../api/projectsAPI";
import useMetaCache from "../../../store/useMetaCache";
import { usePermission } from "../../../hooks/usePermission";
import { Permissions } from "../../../lib/permission";

interface ProjectsSectionProps {
  companyId: string;
}

const ProjectsSection = ({ companyId }: ProjectsSectionProps) => {
  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const navigate = useNavigate();
  const [newProjectModal, setNewProjectModal] = useState<boolean>(false);

  const {
    data: projects,
    isLoading: isProjectsLoading,
    refetch: projectsRefetch,
  } = useQuery({
    queryKey: ["projects", companyId],
    queryFn: () =>
      projectAPI.list({
        page: 1,
        per_page: 1000,
        company_id: companyId,
      }),
    enabled: !!companyId && hasPermission(Permissions.PROJECT_READ_PERMISSION),
  });

  // Don't render if user doesn't have permission to read projects
  if (!hasPermission(Permissions.PROJECT_READ_PERMISSION)) {
    return null;
  }

  return (
    <section className="mb-10">
      <Modal
        title="Create new project"
        visible={newProjectModal}
        onClose={() => setNewProjectModal(false)}
        className="w-[50%] mx-auto bg-primary-300 text-white"
      >
        <NewProjectForm
          onFinish={async (data) => {
            try {
              await projectAPI.create(data);
            } catch (e) {
              console.error(e);
            } finally {
              setNewProjectModal(false);
              await projectsRefetch();
            }
          }}
        />
      </Modal>

      <div className="container mx-auto">
        <h2 className="text-2xl font-semibold mb-5">Projects</h2>
        <ProjectCardWrapper className="w-full flex-wrap">
          {isProjectsLoading ? (
            <p>Loading...</p>
          ) : (
            projects?.items?.map((project, i) => (
              <ProjectCardWrapper.Card
                className="min-w-[30%] px-4 py-6 bg-secondary-200 hover:bg-secondary-100 text-neutral-200 cursor-pointer transition-all duration-300"
                project={project}
                key={i}
                onClick={() => {
                  metaCache.setSelectedProject(project);
                  navigate(`/backlog`);
                }}
              />
            ))
          )}
          {hasPermission(Permissions.PROJECT_WRITE_PERMISSION) && (
            <ProjectCardWrapper.Card className="w-[30%] p-0">
              <div
                onClick={() => {
                  setNewProjectModal(true);
                }}
                className="bg-secondary-200 rounded-md w-full h-full flex justify-center group cursor-pointer hover:bg-secondary-100 py-4 transition-all duration-300"
              >
                <button className="cursor-pointer ">
                  <BsFillPlusCircleFill
                    size="30"
                    className="mx-auto  text-neutral-300 group-hover:text-accent-300 cursor-pointer"
                  />
                </button>
              </div>
            </ProjectCardWrapper.Card>
          )}
        </ProjectCardWrapper>
      </div>
    </section>
  );
};

export default ProjectsSection;
