import { useQuery } from "@tanstack/react-query";
import { useCompanyStore } from "../store/selectedCompanyStore";
import companyAPI from "../api/company";
import { useEffect, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { capitalize } from "../lib/utils/string";
import ProjectCardWrapper from "../components/prioject/ProjectCard";
import { Modal } from "../components/ui/Modal";
import NewProjectForm from "../components/forms/NewProjectForm";
import { useProjectStore } from "../store/selectedProjectStore";
import { useNavigate } from "react-router-dom";
import { usePageSettings } from "../hooks/usePageSettings";
import { Layouts } from "../lib/layout/layout";
import { useCacheLoader } from "../hooks/useCacheLoader";
import { useCacheStore } from "../store/cacheStore";

const CompanyOverviewPage = () => {
  usePageSettings({
    title: "Dashboard",
    requireAuth: true,
    layout: Layouts.Companies,
  });
  const { selectedCompany } = useCompanyStore();
  const { selectProject } = useProjectStore();

  const navigate = useNavigate();

  const [newProjectModal, setNewProjectModal] = useState<boolean>(false);

  const { data: company, isLoading: isCompanyLoading } = useQuery({
    queryKey: ["company", selectedCompany?.id],
    queryFn: () => companyAPI.get(selectedCompany?.id ?? ""),
    enabled: !!selectedCompany?.id,
  });

  const { projects } = useCacheStore();
  useCacheLoader({ projectList: company?.projects });

  useEffect(() => {
    if (!selectedCompany) {
      navigate("/companies");
    }
  }, [selectedCompany]);

  return (
    <div className="w-full px-5 py-10">
      <section>
        <Modal
          title="Create new project"
          visible={newProjectModal}
          onClose={() => setNewProjectModal(false)}
          className="w-[50%] mx-auto bg-primary-300 text-white"
        >
          <NewProjectForm
            onSubmit={(data) => {
              console.log(data);
            }}
          />
        </Modal>
      </section>

      <section className="mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <div className="mb-6 flex items-center justify-between">
          <h1 className="text-2xl font-bold text-gray-900">
            {capitalize(company?.name ?? "")} Company Dashboard
          </h1>
          <button
            onClick={() => {
              setNewProjectModal(true);
            }}
            className="shadow-2xl group px-4 py-2 bg-secondary-100 text-white rounded-md hover:bg-accent-300 hover:text-secondary-100 cursor-pointer flex items-center"
          >
            <IoMdAdd className="mr-2 transition-transform duration-300 group-hover:rotate-90" />
            New Project
          </button>
        </div>
      </section>
      <section>
        <ProjectCardWrapper>
          {isCompanyLoading ? (
            <p>Loading...</p>
          ) : company?.projects?.total_items === 0 ? (
            <p>No projects found.</p>
          ) : (
            company?.projects?.items?.map((project, i) => (
              <ProjectCardWrapper.Card
                project={project}
                key={i}
                onClick={() => {
                  selectProject(project);
                  navigate(`/backlog`);
                }}
              />
            ))
          )}
        </ProjectCardWrapper>
      </section>
    </div>
  );
};

export default CompanyOverviewPage;
