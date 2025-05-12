import { FC, useEffect } from "react";
import { IoMdAdd } from "react-icons/io";
import { usePageSettings } from "../hooks/usePageSettings";
import { formatTime } from "../lib/utils/time";
import { authAPI } from "../api/authAPI";
import { useQuery } from "@tanstack/react-query";
import { useCompanyStore } from "../store/selectedCompanyStore";
import DataTable from "../components/ui/DataTable";
import { Project } from "../lib/project/project";
import { useProjectStore } from "../store/selectedProjectStore";
import { useNavigate } from "react-router-dom";

const ProjectOverviewPage: FC = () => {
  usePageSettings({ title: "Project overview", requireAuth: true });

  const { selectedCompany } = useCompanyStore();
  const { selectProject } = useProjectStore();
  const navigate = useNavigate();

  const { data: company, isLoading: isCompanyLoading } = useQuery({
    queryKey: ["company", selectedCompany?.id],
    queryFn: () => authAPI.getCompany(selectedCompany?.id ?? ""),
    enabled: !!selectedCompany?.id,
  });

  useEffect(() => {
    console.log(company);
  }, [company]);

  const handleSelectProject = (project: Project) => {
    selectProject(project);
    navigate(`/projects/${project.id}/backlog`);
  };

  return (
    <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <div className="mb-6 flex items-center justify-between">
        <div className="flex items-center">
          <button
            // onClick={() => setSelectedOrg(null)}
            className="mr-4 text-gray-600 hover:text-gray-900 !rounded-button"
          >
            <i className="fas fa-arrow-left"></i>
          </button>
          <h1 className="text-2xl font-bold text-gray-900">
            {company?.name} Dashboard
          </h1>
        </div>
        <button className="shadow-2xl group px-4 py-2 bg-secondary-100 text-white rounded-md hover:bg-accent-300 hover:text-secondary-100 cursor-pointer flex items-center">
          <IoMdAdd className="mr-2 transition-transform duration-300 group-hover:rotate-90" />
          New Project
        </button>
      </div>

      {isCompanyLoading ? (
        <p>Loading...</p>
      ) : (
        <div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            <div className="bg-white p-6 rounded-lg shadow">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-medium text-gray-900">
                  Total Projects
                </h3>
                <i className="fas fa-project-diagram text-[rgb(41,43,41)] text-xl"></i>
              </div>
              <p className="text-3xl font-bold text-gray-900">
                {company?.projects?.total_items}
              </p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-medium text-gray-900">
                  Active Projects
                </h3>
                <i className="fas fa-clock text-[rgb(41,43,41)] text-xl"></i>
              </div>
              <p className="text-3xl font-bold text-gray-900">
                {company &&
                company?.projects?.items &&
                company?.projects?.items?.length > 0
                  ? company?.projects?.items.filter(
                      (project) => project.status === "ACTIVE"
                    ).length
                  : 0}
              </p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-medium text-gray-900">
                  Team Members
                </h3>
                <i className="fas fa-users text-[rgb(41,43,41)] text-xl"></i>
              </div>
              <p className="text-3xl font-bold text-gray-900">
                {company?.people_count}
              </p>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow">
            <div className="px-6 py-4 border-b border-gray-200">
              <h2 className="text-xl font-semibold text-gray-900">Projects</h2>
            </div>
            <div className="overflow-x-auto">
              {isCompanyLoading ? (
                <p>Loading...</p>
              ) : company?.projects?.total_items === 0 ? (
                <p>No projects found.</p>
              ) : isCompanyLoading ? (
                <p>Loading...</p>
              ) : company &&
                company.projects.items &&
                company?.projects.items.length === 0 ? (
                <p>No projects found.</p>
              ) : (
                <DataTable
                  heads={[
                    { label: "Project Name", key: "title" },
                    { label: "Status", key: "status" },
                    { label: "Progress", key: "progress" },
                    { label: "Timeline", key: "timeline" },
                    { label: "Team Size", key: "teamSize" },
                  ]}
                  data={company?.projects.items?.map((entry) => ({
                    ...entry,
                    timeline: `${formatTime(
                      entry.created_at.seconds
                    )} - present`,
                    teamSize: 10,
                    progress: 50,
                  }))}
                  onRowClick={handleSelectProject}
                />
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ProjectOverviewPage;
