import { FC, useEffect, useState } from "react";
import { List } from "../lib/utils";
import { Company } from "../lib/company/company";
import useAuth from "../hooks/useAuth";
import { formatTime } from "../lib/utils/formatTime";
import { DataTable } from "../components/ui/DataTable";
import { usePageSettings } from "../hooks/usePageSettings";
import authAPI from "../api/auth";

const CompaniesPage: FC = () => {
  usePageSettings({ requireAuth: true, title: "Companies" });

  const { access_token, isAuthenticated, updateField } = useAuth();
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedIndustry, setSelectedIndustry] = useState("all");
  const [isIndustryDropdownOpen, setIsIndustryDropdownOpen] = useState(false);
  const [companyList, setCompanyList] = useState<List<Company>>();
  const [pagination, setPagination] = useState({ page: 1, per_page: 10 });

  const loadCompanies = async () => {
    try {
      const res = await authAPI(access_token).listCompanies({
        page: 1,
        per_page: 10,
      });
      setCompanyList(res);
    } catch (e) {
      console.error(e);
    }
  };

  const handleSelectCompany = async (company: Company) => {
    try {
      const session = await authAPI(access_token).getSession();
      session.selected_company_id = company.id;
      await authAPI(access_token).updateSession(session);

      updateField({ selected_company_id: company.id });
    } catch (e) {
      console.error(e);
    }
  };

  useEffect(() => {
    if (isAuthenticated) {
      loadCompanies();
    }
  }, []);
  const COMPANIES = [
    {
      id: 1,
      companyName: "Tech Innovators Ltd",
      codeName: "TECHINNO",
      createdAt: "2024-01-15",
      numberOfPeople: 150,
      industry: "Technology",
      status: "active",
      projects: [
        {
          id: 1,
          name: "AI Platform Development",
          status: "active",
          progress: 75,
          startDate: "2024-02-01",
          endDate: "2024-06-30",
          teamSize: 12,
        },
        {
          id: 2,
          name: "Cloud Migration",
          status: "in_progress",
          progress: 45,
          startDate: "2024-03-15",
          endDate: "2024-08-30",
          teamSize: 8,
        },
      ],
    },
    {
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
    },
    {
      id: 3,
      companyName: "Healthcare Solutions",
      codeName: "HCSOL",
      createdAt: "2024-03-10",
      numberOfPeople: 200,
      industry: "Healthcare",
      status: "inactive",
      projects: [
        {
          id: 4,
          name: "Patient Management System",
          status: "in_progress",
          progress: 60,
          startDate: "2024-02-15",
          endDate: "2024-07-30",
          teamSize: 10,
        },
      ],
    },
    {
      id: 4,
      companyName: "Green Energy Systems",
      codeName: "GRENES",
      createdAt: "2024-03-25",
      numberOfPeople: 120,
      industry: "Energy",
      status: "active",
      projects: [
        {
          id: 5,
          name: "Solar Panel Monitoring",
          status: "active",
          progress: 85,
          startDate: "2024-03-01",
          endDate: "2024-05-30",
          teamSize: 6,
        },
      ],
    },
    {
      id: 5,
      companyName: "Digital Marketing Pro",
      codeName: "DIGMPRO",
      createdAt: "2024-04-01",
      numberOfPeople: 80,
      industry: "Marketing",
      status: "active",
      projects: [
        {
          id: 6,
          name: "Campaign Analytics Tool",
          status: "in_progress",
          progress: 30,
          startDate: "2024-04-01",
          endDate: "2024-09-30",
          teamSize: 7,
        },
      ],
    },
  ];
  const INDUSTRIES = [
    "all",
    "Technology",
    "Finance",
    "Healthcare",
    "Energy",
    "Marketing",
  ];
  return (
    <>
      <div className="min-h-screen bg-gray-50">
        <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
          <div className="bg-white rounded-lg shadow">
            <div className="px-4 py-5 border-b border-gray-200 sm:px-6">
              <div className="flex flex-col md:flex-row md:items-center md:justify-between">
                <h2 className="text-2xl font-bold text-gray-900">
                  Organizations
                </h2>
                <div className="mt-4 md:mt-0 flex flex-col md:flex-row md:space-x-4 space-y-4 md:space-y-0">
                  <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <i className="fas fa-search text-gray-400"></i>
                    </div>
                    <input
                      type="text"
                      className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:ring-1 focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm"
                      placeholder="Search organizations..."
                      value={searchTerm}
                      onChange={(e) => setSearchTerm(e.target.value)}
                    />
                  </div>
                  <div className="relative">
                    <button
                      type="button"
                      className="inline-flex justify-between w-48 rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-1 focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] !rounded-button"
                      onClick={() =>
                        setIsIndustryDropdownOpen(!isIndustryDropdownOpen)
                      }
                    >
                      {selectedIndustry.charAt(0).toUpperCase() +
                        selectedIndustry.slice(1)}
                      <i className="fas fa-chevron-down ml-2"></i>
                    </button>
                    {isIndustryDropdownOpen && (
                      <div className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-10">
                        <div
                          className="py-1"
                          role="menu"
                          aria-orientation="vertical"
                        >
                          {INDUSTRIES.map((industry) => (
                            <button
                              key={industry}
                              className={`${
                                selectedIndustry === industry
                                  ? "bg-gray-100 text-gray-900"
                                  : "text-gray-700"
                              } block px-4 py-2 text-sm w-full text-left hover:bg-gray-100`}
                              onClick={() => {
                                setSelectedIndustry(industry);
                                setIsIndustryDropdownOpen(false);
                              }}
                            >
                              {industry.charAt(0).toUpperCase() +
                                industry.slice(1)}
                            </button>
                          ))}
                        </div>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            </div>
            <div className="overflow-x-auto">
              <DataTable
                data={companyList?.items || []}
                page={pagination.page}
                perPage={pagination.per_page}
                total={companyList?.total || 0}
                onPageChange={(page) =>
                  setPagination((prev) => ({ ...prev, page }))
                }
                columns={[
                  { header: "Company Name", accessor: (c) => c.name },
                  { header: "Code Name", accessor: (c) => c.codename },
                  {
                    header: "Created At",
                    accessor: (c) => formatTime(c.created_at.seconds),
                  },
                  { header: "People", accessor: (c) => c.people_count },
                  { header: "Industry", accessor: () => "Test" },
                  { header: "Status", accessor: () => "Active" },
                ]}
                renderActions={(comp) => (
                  <>
                    <button
                      className="cursor-pointer"
                      onClick={() => {
                        handleSelectCompany(comp);
                      }}
                    >
                      Select
                    </button>
                    <button className="text-[rgb(41,43,41)] hover:text-[rgb(31,33,31)] mr-4 !rounded-button">
                      Edit
                    </button>
                    <button className="text-red-600 hover:text-red-900 !rounded-button">
                      Delete
                    </button>
                  </>
                )}
              />
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default CompaniesPage;
