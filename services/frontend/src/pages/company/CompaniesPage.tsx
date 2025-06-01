import { FC, useEffect, useState } from "react";
import { Company, CompanyFilters } from "../../lib/company/company";
import { formatTime } from "../../lib/utils/time";
import { usePageSettings } from "../../hooks/usePageSettings";
import { useAuthStore } from "../../store/authStore";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { ListItems} from "../../lib/utils/list";
import Paginator from "../../components/ui/Paginator";
import companyAPI from "../../api/company";
import Table, { TableColumn } from "../../components/ui/Table";
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
} from "react-icons/bs";
import { parseError } from "../../lib/errors";
import useMetaCache from "../../store/useMetaCache";
import { useAssigneeList } from "../../hooks/useSprintList";

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

  const companyColumns: TableColumn<Company>[] = [
    {
      header: "Name",
      accessor: (company) => (
        <div className="flex items-center gap-2">
          <BsBuilding className="text-accent-500" />
          <span className="font-medium">{company.name}</span>
        </div>
      ),
    },
    {
      header: "Codename",
      accessor: (company) => (
        <div className="px-2 py-1 bg-accent-500/10 text-accent-500 rounded-md w-fit">
          {company.codename}
        </div>
      ),
    },
    {
      header: "People",
      accessor: (company) => (
        <div className="flex items-center gap-2">
          <BsPeople />
          <span>{company.people_count ?? 0}</span>
        </div>
      ),
    },
    {
      header: "Projects",
      accessor: (company) => (
        <div className="flex items-center gap-2">
          <BsFolder />
          <span>{company.projects?.total_items ?? 0}</span>
        </div>
      ),
    },
    {
      header: "Created",
      accessor: (company) => (
        <div className="text-neutral-400">
          {formatTime(company.created_at.seconds)}
        </div>
      ),
    },
  ];

  return (
    <div className="min-h-[100lvh] bg-gradient-to-br from-primary-700 to-primary-600">
      <div className="container mx-auto px-6 py-8">
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

        {/* Header Section */}
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
                <h3 className="text-lg font-medium text-white">Total People</h3>
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
                  (acc, company) => acc + (company.projects?.total_items ?? 0),
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

        {/* Table Section */}
        <div className="bg-primary-500/50 backdrop-blur-sm rounded-xl border border-primary-400/30 overflow-hidden">
          <div className="h-[50vh]">
            <Table
              data={companyList?.items}
              columns={companyColumns}
              isLoading={isCompanyListLoading}
              emptyMessage="No companies found"
              onRowClick={handleSelectCompany}
            />
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
      </div>
    </div>
  );
};

export default CompaniesPage;
