import { FC, useEffect, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { usePageSettings } from "../hooks/usePageSettings";
import authAPI from "../api/auth";
import useAuth from "../hooks/useAuth";
import { Company } from "../lib/company/company";
import { formatTime } from "../lib/utils/formatTime";

const DashboardPage: FC = () => {
  usePageSettings({ title: "Dashboard", requireAuth: true });

  const [org, setOrg] = useState<Company | null>(null);
  const { access_token } = useAuth();

  useEffect(() => {
    authAPI(access_token)
      .getCompany("8f557202-0853-4672-aafb-a0b6cae7067a")
      .then((res) => {
        console.log(res);
        setOrg(res);
      })
      .catch((e) => {
        console.error(e);
      });
  }, []);

  const ORG = {
    id: 2,
    companyName: "Global Finance Corp",
    codeName: "GFINCORP",
    createdAt: "2024-02-20",
    numberOfPeople: 300,
    industry: "Finance",
    status: "active",
    projects: [
      {
        id: 3,
        name: "Digital Banking App",
        status: "completed",
        progress: 100,
        startDate: "2024-01-10",
        endDate: "2024-04-10",
        teamSize: 15,
      },
    ],
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
            {org?.name} Dashboard
          </h1>
        </div>
        <button className="shadow-2xl group px-4 py-2 bg-secondary-100 text-white rounded-md hover:bg-accent-300 hover:text-secondary-100 cursor-pointer flex items-center">
          <IoMdAdd className="mr-2 transition-transform duration-300 group-hover:rotate-90" />
          New Project
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <div className="bg-white p-6 rounded-lg shadow">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-lg font-medium text-gray-900">
              Total Projects
            </h3>
            <i className="fas fa-project-diagram text-[rgb(41,43,41)] text-xl"></i>
          </div>
          <p className="text-3xl font-bold text-gray-900">
            {org?.projects?.total_items}
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
            {
              org?.projects?.items.filter(
                (project) => project.status === "ACTIVE"
              ).length
            }
          </p>
        </div>
        <div className="bg-white p-6 rounded-lg shadow">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-lg font-medium text-gray-900">Team Members</h3>
            <i className="fas fa-users text-[rgb(41,43,41)] text-xl"></i>
          </div>
          <p className="text-3xl font-bold text-gray-900">
            {org?.people_count}
          </p>
        </div>
      </div>

      <div className="bg-white rounded-lg shadow">
        <div className="px-6 py-4 border-b border-gray-200">
          <h2 className="text-xl font-semibold text-gray-900">Projects</h2>
        </div>
        <div className="overflow-x-auto">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Project Name
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Status
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Progress
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Timeline
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Team Size
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {org?.projects?.items.map((project) => (
                <tr key={project.id}>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="text-sm font-medium text-gray-900">
                      {project.title}
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span
                      className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${
                        project.status === "active"
                          ? "bg-green-100 text-green-800"
                          : project.status === "completed"
                          ? "bg-blue-100 text-blue-800"
                          : "bg-yellow-100 text-yellow-800"
                      }`}
                    >
                      {project.status.replace("_", " ")}
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="w-full bg-gray-200 rounded-full h-2.5">
                      <div
                        className="bg-[rgb(41,43,41)] h-2.5 rounded-full"
                        style={{ width: `${project.progress ?? 0}%` }}
                      ></div>
                    </div>
                    <span className="text-sm text-gray-500 mt-1">
                      {project.progress ?? 0}%
                    </span>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <div className="text-sm text-gray-500">
                      {formatTime(project.created_at.seconds)} - present
                    </div>
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {/* {project.teamSize} members */}
                    20 members
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <button className="text-[rgb(41,43,41)] hover:text-[rgb(31,33,31)] mr-4 !rounded-button">
                      Edit
                    </button>
                    <button className="text-red-600 hover:text-red-900 !rounded-button">
                      Remove
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default DashboardPage;
