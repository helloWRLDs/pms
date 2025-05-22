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
import authAPI from "../api/auth";
import Paginator from "../components/ui/Paginator";
import { BsFillPlusCircleFill } from "react-icons/bs";

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

  const { data: users } = useQuery({
    queryKey: ["users", company?.id],
    queryFn: () =>
      authAPI.listUsers({
        page: 1,
        per_page: 10000,
        company_id: company?.id ?? "",
      }),
    enabled: !!company?.id,
  });

  useCacheLoader({ projectList: company?.projects });
  useCacheLoader({ userList: users });

  useEffect(() => {
    if (!selectedCompany) {
      navigate("/companies");
    }
  }, [selectedCompany]);

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-primary-600 text-neutral-100 ">
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

      <section>
        <div className="container mx-auto mb-6 flex items-center justify-between">
          <h1 className="text-2xl font-bold">
            {capitalize(company?.name ?? "")} Company Dashboard
          </h1>
        </div>
      </section>
      <section>
        <ProjectCardWrapper className="container mx-auto">
          {isCompanyLoading ? (
            <p>Loading...</p>
          ) : company?.projects?.total_items === 0 ? (
            <p>No projects found.</p>
          ) : (
            company?.projects?.items?.map((project, i) => (
              <ProjectCardWrapper.Card
                className="w-[30%] px-4 py-6 bg-secondary-200 hover:bg-secondary-100 text-neutral-200 cursor-pointer transition-all duration-300"
                project={project}
                key={i}
                onClick={() => {
                  selectProject(project);
                  navigate(`/backlog`);
                }}
              />
            ))
          )}
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
        </ProjectCardWrapper>
      </section>
    </div>
  );
};

export default CompanyOverviewPage;
