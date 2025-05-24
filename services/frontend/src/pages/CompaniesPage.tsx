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
import Input from "../components/ui/Input";
import { Button } from "../components/ui/Button";
import { BsFillPlusCircleFill } from "react-icons/bs";
import { parseError } from "../lib/errors";
import { useProjectStore } from "../store/selectedProjectStore";

const CompaniesPage: FC = () => {
  usePageSettings({ requireAuth: true, title: "Companies" });

  const { companies, getCompanies, setCompanies } = useCacheStore();
  useEffect(() => {
    console.log("store companies");
    console.log(JSON.stringify(companies));
  }, [companies]);

  const { isAuthenticated, auth, clearAuth } = useAuthStore();
  const { selectCompany } = useCompanyStore();
  const { selectProject } = useProjectStore();

  const [search, setSearch] = useState("");
  const [newCompanyModal, setNewCompanyModal] = useState<boolean>(false);

  const [filter, setFilter] = useState<CompanyFilters>({
    page: 1,
    per_page: 10,
  });

  const { data: companyList, isLoading: isCompanyListLoading } = useQuery<
    ListItems<Company>
  >({
    queryKey: ["companies", filter.page, filter.per_page, filter.company_name],
    queryFn: async () => {
      try {
        const res = await companyAPI.list({
          page: filter.page,
          per_page: filter.per_page,
          user_id: auth?.user.id ?? "",
        });
        return res;
      } catch (e) {
        if (parseError(e)?.status === 401) {
          clearAuth();
          navigate("/login");
        }
        return {} as ListItems<Company>;
      }
    },
    enabled: isAuthenticated(),
  });

  useCacheLoader({ companyList: companyList });

  const navigate = useNavigate();

  const handleSelectCompany = async (company: Company) => {
    try {
      selectCompany(company);
      selectProject(null);
      navigate(`/projects`);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-primary-600 text-neutral-100 ">
      <Modal
        title="Create new company"
        visible={newCompanyModal}
        onClose={() => setNewCompanyModal(false)}
        className="w-[50%] mx-auto bg-primary-300 text-white"
      >
        <NewCompanyForm
          onFinish={(data) => {
            companyAPI.create(data);
            setNewCompanyModal(false);
          }}
        />
      </Modal>

      <section id="companies-header">
        <div className="container mx-auto flex justify-between items-center mb-4">
          <h2 className="font-bold text-2xl mb-4">Organizations</h2>
          <div className="flex gap-4 items-baseline">
            <Input>
              <Input.Element
                type="text"
                label="Title"
                value={search}
                onInput={(e) => {
                  setSearch(e.currentTarget.value);
                }}
              />
            </Input>
            <Button
              onClick={() => {
                setFilter({ ...filter, company_codename: search });
              }}
            >
              Search
            </Button>
          </div>
        </div>
      </section>

      <section id="companies-list">
        <div className="overflow-x-auto container mx-auto shadow-xl">
          <div className="">
            <div className="h-[75lvh] w-full">
              <Table className="rounded-lg">
                <Table.Head>
                  <Table.Row className="text-neutral-100 bg-primary-400">
                    <Table.HeadCell>Name</Table.HeadCell>
                    <Table.HeadCell>Codename</Table.HeadCell>
                    <Table.HeadCell>People</Table.HeadCell>
                    <Table.HeadCell>Projects</Table.HeadCell>
                    <Table.HeadCell>Created</Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                {isCompanyListLoading ? (
                  <div className="p-4 text-center text-gray-500">
                    Loading companies...
                  </div>
                ) : !companyList ||
                  !companyList.items ||
                  companyList.items?.length === 0 ? (
                  <Table.Body></Table.Body>
                ) : (
                  <Table.Body>
                    {companyList.items.map((company) => (
                      <Table.Row
                        key={company.id}
                        className="cursor-pointer bg-secondary-200 text-neutral-100 hover:bg-secondary-100 transition-all duration-300"
                        onClick={() => handleSelectCompany(company)}
                      >
                        <Table.Cell className="py-4 px-2">
                          {company.name}
                        </Table.Cell>
                        <Table.Cell className="py-4 px-2">
                          {company.codename}
                        </Table.Cell>
                        <Table.Cell className="py-4 px-2">
                          {company.people_count}
                        </Table.Cell>
                        <Table.Cell className="py-4 px-2">
                          {company.projects?.total_items ?? 0}
                        </Table.Cell>
                        <Table.Cell className="py-4 px-2">
                          {formatTime(company.created_at.seconds)}
                        </Table.Cell>
                      </Table.Row>
                    ))}
                  </Table.Body>
                )}
              </Table>
              <button
                className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
                onClick={() => {
                  setNewCompanyModal(true);
                }}
              >
                <BsFillPlusCircleFill
                  size="30"
                  className="mx-auto text-neutral-300 group-hover:text-accent-300 "
                />
              </button>
            </div>
            {companyList && companyList.items && (
              <Paginator
                page={companyList.page ?? 0}
                per_page={companyList.per_page ?? 0}
                total_items={companyList.total_items ?? 0}
                total_pages={companyList.total_pages ?? 0}
                onPageChange={(page) => {
                  setFilter({ ...filter, page: page });
                }}
                className=""
              />
            )}
          </div>
        </div>
      </section>
    </div>
  );
};

export default CompaniesPage;
