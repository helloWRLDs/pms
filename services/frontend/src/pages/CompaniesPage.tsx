import { FC, useEffect, useState } from "react";
import { Company } from "../lib/company/company";
import { formatTime } from "../lib/utils/time";
import { usePageSettings } from "../hooks/usePageSettings";
import { authAPI } from "../api/authAPI";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { useAuthStore } from "../store/authStore";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { ListItems, Pagination } from "../lib/utils/list";
import DataTable from "../components/ui/DataTable";
import Paginator from "../components/ui/Paginator";

const CompaniesPage: FC = () => {
  usePageSettings({ requireAuth: true, title: "Companies" });

  const { isAuthenticated } = useAuthStore();
  const { selectCompany } = useCompanyStore();

  const [searchTerm, setSearchTerm] = useState("");
  const [selectedIndustry, setSelectedIndustry] = useState("all");
  const [isIndustryDropdownOpen, setIsIndustryDropdownOpen] = useState(false);
  const [pagination, setPagination] = useState<Pagination>({
    page: 1,
    per_page: 10,
  });

  const { data: companyList, isLoading: isCompanyListLoading } = useQuery<
    ListItems<Company>
  >({
    queryKey: ["companies", pagination.page, pagination.per_page],
    queryFn: () =>
      authAPI.listCompanies({
        page: pagination.page,
        per_page: pagination.per_page,
      }),
    enabled: isAuthenticated(),
  });

  useEffect(() => {
    console.log(companyList);
  }, [companyList]);

  const navigate = useNavigate();

  const handleSelectCompany = async (company: Company) => {
    try {
      const session = await authAPI.getSession();
      session.selected_company_id = company.id;
      await authAPI.updateSession(session);
      selectCompany(company);
      navigate(`/companies/${company.id}`);
    } catch (e) {
      console.error(e);
    }
  };

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
      <section>
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
                {isCompanyListLoading || !companyList ? (
                  <div className="p-4 text-center text-gray-500">
                    Loading companies...
                  </div>
                ) : companyList?.items?.length === 0 ? (
                  <div className="p-4 text-center text-gray-500">
                    No companies found.
                  </div>
                ) : (
                  <div>
                    <DataTable
                      heads={[
                        { label: "Name", key: "name" },
                        { label: "Codename", key: "codename" },
                        { label: "People", key: "people_count" },
                        { label: "Created", key: "created_at_formatted" },
                      ]}
                      data={companyList.items.map((entry) => ({
                        ...entry,
                        created_at_formatted: formatTime(
                          entry.created_at.seconds
                        ),
                      }))}
                      onRowClick={(item) => {
                        handleSelectCompany(item);
                      }}
                    />
                    <Paginator
                      page={companyList?.page ?? 0}
                      per_page={companyList?.per_page ?? 0}
                      total_items={companyList?.total_items ?? 0}
                      total_pages={companyList?.total_pages ?? 0}
                      onPageChange={(page) => {
                        setPagination({ ...pagination, page: page });
                      }}
                    />
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default CompaniesPage;
