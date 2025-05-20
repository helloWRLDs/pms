import { FC, useEffect, useState } from "react";
import { Company, CompanyFilters } from "../lib/company/company";
import { formatTime } from "../lib/utils/time";
import { usePageSettings } from "../hooks/usePageSettings";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { useAuthStore } from "../store/authStore";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { ListItems, Pagination } from "../lib/utils/list";
import Paginator from "../components/ui/Paginator";
import companyAPI from "../api/company";
import Table from "../components/ui/Table";
import NewCompanyForm from "../components/forms/NewCompanyForm";
import { Modal } from "../components/ui/Modal";
import { useCacheStore } from "../store/cacheStore";
import { useCacheLoader } from "../hooks/useCacheLoader";

const CompaniesPage: FC = () => {
  usePageSettings({ requireAuth: true, title: "Companies" });

  const { companies, getCompanies, setCompanies } = useCacheStore();
  useEffect(() => {
    console.log("store companies");
    console.log(JSON.stringify(companies));
  }, [companies]);

  const { isAuthenticated, auth } = useAuthStore();
  const { selectCompany } = useCompanyStore();

  const [newCompanyModal, setNewCompanyModal] = useState<boolean>(false);

  const [filter, setFilter] = useState<CompanyFilters>({
    page: 1,
    per_page: 10,
  });

  const { data: companyList, isLoading: isCompanyListLoading } = useQuery<
    ListItems<Company>
  >({
    queryKey: ["companies", filter.page, filter.per_page, filter.company_name],
    queryFn: () =>
      companyAPI.list({
        page: filter.page,
        per_page: filter.per_page,
        user_id: auth?.user.id ?? "",
      }),
    enabled: isAuthenticated(),
  });

  useCacheLoader({ companyList: companyList });

  const navigate = useNavigate();

  const handleSelectCompany = async (company: Company) => {
    try {
      selectCompany(company);
      navigate(`/projects`);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="w-full px-5 py-10">
      <Modal
        title="Create new company"
        visible={newCompanyModal}
        onClose={() => setNewCompanyModal(false)}
        className="w-[50%] mx-auto bg-primary-300 text-white"
      >
        <NewCompanyForm
          onSubmit={(data) => {
            console.log(data);
            setNewCompanyModal(false);
          }}
        />
      </Modal>

      <section id="companies-header">
        <div className="flex justify-between items-center mb-4">
          <h2 className="font-bold text-2xl mb-4">Organizations</h2>
          <div>
            <input
              type="text"
              name="search"
              id=""
              placeholder="Search"
              className="border rounded-lg px-4 py-2 "
              onInput={(e) => {
                setFilter({ ...filter, company_name: e.currentTarget.value });
              }}
            />
          </div>
          <button
            className="px-4 py-2 border border-black rounded-md cursor-pointer"
            onClick={() => {
              setNewCompanyModal(true);
            }}
          >
            Create Company
          </button>
        </div>
      </section>

      <section id="companies-list">
        <div className="overflow-x-auto">
          {isCompanyListLoading || !companyList ? (
            <div className="p-4 text-center text-gray-500">
              Loading companies...
            </div>
          ) : companyList.items?.length === 0 ? (
            <div className="p-4 text-center text-gray-500">
              No companies found.
            </div>
          ) : (
            <div>
              <Table>
                <Table.Head>
                  <Table.Row>
                    <Table.HeadCell>Name</Table.HeadCell>
                    <Table.HeadCell>Codename</Table.HeadCell>
                    <Table.HeadCell>People</Table.HeadCell>
                    <Table.HeadCell>Created</Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                <Table.Body>
                  {companyList.items.map((company) => (
                    <Table.Row
                      key={company.id}
                      className="cursor-pointer"
                      onClick={() => handleSelectCompany(company)}
                    >
                      <Table.Cell>{company.name}</Table.Cell>
                      <Table.Cell>{company.codename}</Table.Cell>
                      <Table.Cell>{company.people_count}</Table.Cell>
                      <Table.Cell>
                        {formatTime(company.created_at.seconds)}
                      </Table.Cell>
                    </Table.Row>
                  ))}
                </Table.Body>
              </Table>
              <Paginator
                page={companyList?.page ?? 0}
                per_page={companyList?.per_page ?? 0}
                total_items={companyList?.total_items ?? 0}
                total_pages={companyList?.total_pages ?? 0}
                onPageChange={(page) => {
                  setFilter({ ...filter, page: page });
                }}
              />
            </div>
          )}
        </div>
      </section>
    </div>
  );
};

export default CompaniesPage;
