import { FC, useEffect, useState } from "react";
import { Company, CompanyFilters } from "../../lib/company/company";
import { formatTime } from "../../lib/utils/time";
import { usePageSettings } from "../../hooks/usePageSettings";
import { useAuthStore } from "../../store/authStore";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { ListItems } from "../../lib/utils/list";
import Paginator from "../../components/ui/Paginator";
import companyAPI from "../../api/company";
import Table from "../../components/ui/Table";
import NewCompanyForm from "../../components/forms/NewCompanyForm";
import { Modal } from "../../components/ui/Modal";
import Input from "../../components/ui/Input";
import { Button } from "../../components/ui/Button";
import {
  BsFillPlusCircleFill,
  BsSearch,
  BsBuilding,
  BsPeople,
  BsFolder,
  BsThreeDotsVertical,
} from "react-icons/bs";
import { parseError } from "../../lib/errors";
import useMetaCache from "../../store/useMetaCache";
import { useAssigneeList } from "../../hooks/useSprintList";
import { ContextMenu } from "../../components/ui/ContextMenu";

const CompaniesPage: FC = () => {
  usePageSettings({ requireAuth: true, title: "Companies" });
  const metaCache = useMetaCache();
  const { isAuthenticated, auth, clearAuth } = useAuthStore();

  const [search, setSearch] = useState("");
  const [newCompanyModal, setNewCompanyModal] = useState<boolean>(false);

  const [filter, setFilter] = useState<CompanyFilters>({
    page: 1,
    per_page: 10,
  });

  const {
    data: companyList,
    isLoading: isCompanyListLoading,
    refetch: refetchCompanyList,
  } = useQuery<ListItems<Company>>({
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

  const navigate = useNavigate();

  const handleSelectCompany = async (company: Company) => {
    try {
      metaCache.setSelectedCompany(company);
      navigate(`/projects`);
    } catch (e) {
      console.error(e);
    }
  };

  const { refetchAssignees } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );

  useEffect(() => {
    if (metaCache.metadata.selectedCompany) {
      refetchAssignees();
    }
  }, [metaCache.metadata.selectedCompany]);

  const LoadingRow = () => (
    <Table.Row>
      <Table.Cell>
        <div className="animate-pulse">
          <div className="h-4 bg-primary-300/50 rounded w-3/4"></div>
        </div>
      </Table.Cell>
      <Table.Cell>
        <div className="animate-pulse">
          <div className="h-4 bg-primary-300/50 rounded w-1/2"></div>
        </div>
      </Table.Cell>
      <Table.Cell>
        <div className="animate-pulse">
          <div className="h-4 bg-primary-300/50 rounded w-1/3"></div>
        </div>
      </Table.Cell>
      <Table.Cell>
        <div className="animate-pulse">
          <div className="h-4 bg-primary-300/50 rounded w-1/3"></div>
        </div>
      </Table.Cell>
      <Table.Cell>
        <div className="animate-pulse">
          <div className="h-4 bg-primary-300/50 rounded w-1/2"></div>
        </div>
      </Table.Cell>
    </Table.Row>
  );

  return (
    <div className="min-h-[100lvh] bg-gradient-to-br from-primary-700 to-primary-600">
      <div className="container mx-auto px-6 py-8">
        <section id="modal-section">
          <Modal
            title="Create new company"
            visible={newCompanyModal}
            onClose={() => setNewCompanyModal(false)}
            className="w-[50%] mx-auto bg-primary-300 text-white"
          >
            <NewCompanyForm
              onFinish={async (data) => {
                await companyAPI.create(data);
                await refetchCompanyList();
                setNewCompanyModal(false);
              }}
            />
          </Modal>
        </section>

        <section id="header-section">
          <div className="mb-8">
            <div className="flex justify-between items-center mb-6">
              <div>
                <h1 className="text-3xl font-bold text-white mb-2">
                  Organizations
                </h1>
                <p className="text-neutral-300">
                  Manage and organize your companies
                </p>
              </div>
              <Button
                onClick={() => setNewCompanyModal(true)}
                className="flex items-center gap-2 bg-accent-500 text-white hover:bg-accent-600"
              >
                <BsFillPlusCircleFill />
                <span>New Organization</span>
              </Button>
            </div>

            {/* Stats Section */}
            <div className="grid grid-cols-3 gap-6 mb-8">
              <div className="bg-primary-500/50 backdrop-blur-sm rounded-xl p-6 border border-primary-400/30">
                <div className="flex items-center gap-3 mb-2">
                  <BsBuilding className="text-2xl text-accent-500" />
                  <h3 className="text-lg font-medium text-white">
                    Total Companies
                  </h3>
                </div>
                <p className="text-3xl font-bold text-white">
                  {companyList?.total_items ?? 0}
                </p>
              </div>
              <div className="bg-primary-500/50 backdrop-blur-sm rounded-xl p-6 border border-primary-400/30">
                <div className="flex items-center gap-3 mb-2">
                  <BsPeople className="text-2xl text-accent-500" />
                  <h3 className="text-lg font-medium text-white">
                    Total People
                  </h3>
                </div>
                <p className="text-3xl font-bold text-white">
                  {companyList?.items?.reduce(
                    (acc, company) => acc + (company.people_count ?? 0),
                    0
                  ) ?? 0}
                </p>
              </div>
              <div className="bg-primary-500/50 backdrop-blur-sm rounded-xl p-6 border border-primary-400/30">
                <div className="flex items-center gap-3 mb-2">
                  <BsFolder className="text-2xl text-accent-500" />
                  <h3 className="text-lg font-medium text-white">
                    Total Projects
                  </h3>
                </div>
                <p className="text-3xl font-bold text-white">
                  {companyList?.items?.reduce(
                    (acc, company) =>
                      acc + (company.projects?.total_items ?? 0),
                    0
                  ) ?? 0}
                </p>
              </div>
            </div>

            {/* Search Section */}
            <div className="flex gap-4 items-center bg-primary-500/50 backdrop-blur-sm p-4 rounded-xl border border-primary-400/30">
              <div className="flex-1">
                <Input>
                  <Input.Element
                    type="text"
                    label="Search organizations"
                    value={search}
                    onInput={(e) => setSearch(e.currentTarget.value)}
                  />
                </Input>
              </div>
              <Button
                onClick={() => setFilter({ ...filter, company_name: search })}
                className="flex items-center gap-2"
              >
                <BsSearch />
                <span>Search</span>
              </Button>
            </div>
          </div>
        </section>

        <section id="table-section">
          <div className="bg-primary-500/50 backdrop-blur-sm rounded-xl border border-primary-400/30 overflow-hidden">
            <div className="h-[50vh] overflow-x-auto">
              <table className="w-full border-collapse">
                <Table.Head className="bg-primary-400/70 text-white sticky top-0">
                  <Table.Row>
                    <Table.HeadCell>Name</Table.HeadCell>
                    <Table.HeadCell>Codename</Table.HeadCell>
                    <Table.HeadCell>Address</Table.HeadCell>
                    <Table.HeadCell>BIN</Table.HeadCell>
                    <Table.HeadCell>People</Table.HeadCell>
                    <Table.HeadCell>Projects</Table.HeadCell>
                    <Table.HeadCell>Created</Table.HeadCell>
                    <Table.HeadCell></Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                <Table.Body className="divide-y divide-primary-400/20">
                  {isCompanyListLoading ? (
                    Array(5)
                      .fill(0)
                      .map((_, index) => <LoadingRow key={index} />)
                  ) : !companyList?.items || companyList.items.length === 0 ? (
                    <Table.Row>
                      <td colSpan={5} className="px-4 py-12 text-center">
                        <div className="text-white/70">
                          <BsBuilding className="mx-auto text-4xl mb-4 opacity-50" />
                          <p className="text-lg font-medium mb-2">
                            No companies found
                          </p>
                          <p className="text-sm opacity-75">
                            Create your first organization to get started
                          </p>
                        </div>
                      </td>
                    </Table.Row>
                  ) : (
                    companyList.items.map((company) => (
                      <Table.Row
                        key={company.id}
                        onClick={() => handleSelectCompany(company)}
                        className="bg-primary-600/30 hover:bg-primary-500/40 text-white cursor-pointer transition-all duration-200 hover:shadow-lg"
                      >
                        <Table.Cell>
                          <div className="flex items-center gap-3">
                            <div className="p-2 bg-accent-500/20 rounded-lg">
                              <BsBuilding className="text-accent-400" />
                            </div>
                            <div>
                              <span className="font-medium text-white">
                                {company.name}
                              </span>
                            </div>
                          </div>
                        </Table.Cell>
                        <Table.Cell>
                          <div className="inline-flex px-3 py-1 bg-accent-500/20 text-accent-400 rounded-full text-sm font-medium border border-accent-500/30">
                            {company.codename}
                          </div>
                        </Table.Cell>
                        <Table.Cell>{company.address}</Table.Cell>
                        <Table.Cell>{company.bin}</Table.Cell>
                        <Table.Cell>
                          <div className="flex items-center gap-2 text-white/80">
                            <BsPeople className="text-accent-400" />
                            <span className="font-medium">
                              {company.people_count ?? 0}
                            </span>
                          </div>
                        </Table.Cell>
                        <Table.Cell>
                          <div className="flex items-center gap-2 text-white/80">
                            <BsFolder className="text-accent-400" />
                            <span className="font-medium">
                              {company.projects?.total_items ?? 0}
                            </span>
                          </div>
                        </Table.Cell>
                        <Table.Cell>
                          <div className="text-white/60 text-sm">
                            {formatTime(company.created_at.seconds)}
                          </div>
                        </Table.Cell>
                        <Table.Cell>
                          <ContextMenu
                            placement="left"
                            trigger={<BsThreeDotsVertical />}
                            items={[
                              { label: "Edit", onClick: () => {} },
                              { label: "Delete", onClick: () => {} },
                            ]}
                          />
                        </Table.Cell>
                      </Table.Row>
                    ))
                  )}
                </Table.Body>
              </table>
            </div>
            {companyList && companyList.items && (
              <div className="p-4 border-t border-primary-400/30">
                <Paginator
                  page={Number(companyList.page)}
                  per_page={Number(companyList.per_page)}
                  total_items={Number(companyList.total_items)}
                  total_pages={Number(companyList.total_pages)}
                  onPageChange={(page) =>
                    setFilter({ ...filter, page: Number(page) })
                  }
                />
              </div>
            )}
          </div>
        </section>
      </div>
    </div>
  );
};

export default CompaniesPage;
