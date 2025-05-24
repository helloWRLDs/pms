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
import projectAPI from "../api/projectsAPI";
import Table from "../components/ui/Table";
import { UserFilter } from "../lib/user/user";
import AddParticipantForm from "../components/forms/AddParticipantForm";
import { infoToast } from "../lib/utils/toast";

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
  const [addUserModal, setAddUserModal] = useState(false);

  const {
    data: company,
    isLoading: isCompanyLoading,
    refetch: companyRefetch,
  } = useQuery({
    queryKey: ["company", selectedCompany?.id],
    queryFn: () => companyAPI.get(selectedCompany?.id ?? ""),
    enabled: !!selectedCompany?.id,
  });

  const [userFilter, setUserFilter] = useState<UserFilter>({
    page: 1,
    per_page: 10,
    company_id: company?.id ?? "",
  });

  const {
    data: users,
    isLoading: isUsersLoading,
    refetch: usersRefetch,
  } = useQuery({
    queryKey: [
      "users",
      userFilter.page,
      userFilter.per_page,
      userFilter.company_id,
    ],
    queryFn: () => authAPI.listUsers(userFilter),
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
            onFinish={async (data) => {
              try {
                await projectAPI.create(data);
              } catch (e) {
                console.error(e);
              } finally {
                setNewProjectModal(false);
                companyRefetch();
              }
            }}
          />
        </Modal>
        <Modal
          title="Add user"
          visible={addUserModal}
          onClose={() => {
            setAddUserModal(false);
          }}
        >
          <AddParticipantForm
            onFinish={async (userID) => {
              try {
                if (company?.id) {
                  await companyAPI.addParticipant(company.id, userID);
                  infoToast("user added");
                  usersRefetch();
                }
              } catch (e) {
                console.error(e);
              }
            }}
          />
        </Modal>
      </section>

      <section>
        <div className="container mx-auto mb-6 flex items-center justify-between">
          <h1 className="text-3xl font-bold">
            <span className="text-accent-500">
              {capitalize(company?.name ?? "")}
            </span>{" "}
            Company Dashboard
          </h1>
        </div>
      </section>

      <section className="mb-10">
        <div className="container mx-auto">
          <h2 className="text-2xl font-semibold mb-5">Projects</h2>
          <ProjectCardWrapper className="w-full flex-wrap">
            {isCompanyLoading ? (
              <p>Loading...</p>
            ) : company?.projects?.total_items === 0 ? (
              <p>No projects found.</p>
            ) : (
              company?.projects?.items?.map((project, i) => (
                <ProjectCardWrapper.Card
                  className="min-w-[30%] px-4 py-6 bg-secondary-200 hover:bg-secondary-100 text-neutral-200 cursor-pointer transition-all duration-300"
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
        </div>
      </section>

      <section>
        <div className="container mx-auto">
          <h2 className="text-2xl font-semibold mb-5">People</h2>
          <Table>
            <Table.Head>
              <Table.Row className="text-neutral-100 bg-primary-400">
                <Table.HeadCell className="w-[4rem]"></Table.HeadCell>
                <Table.HeadCell>№</Table.HeadCell>
                <Table.HeadCell>Name</Table.HeadCell>
                <Table.HeadCell>Email</Table.HeadCell>
              </Table.Row>
            </Table.Head>
            {isUsersLoading ? (
              <p>Loading...</p>
            ) : !users || !users.items || users.total_items === 0 ? (
              <Table.Body></Table.Body>
            ) : (
              <Table.Body>
                {users.items.map((user, i) => (
                  <Table.Row
                    key={user.id}
                    className="cursor-pointer bg-secondary-200 text-neutral-100 hover:bg-secondary-100 py-10"
                  >
                    <Table.Cell>
                      <div className="aspect-square w-[2rem]">
                        <img
                          src={`data:image/jpeg;base64,${user?.avatar_img}`}
                        />
                      </div>
                    </Table.Cell>
                    <Table.Cell>{i + 1}</Table.Cell>
                    <Table.Cell>{user.name}</Table.Cell>
                    <Table.Cell>{user.email}</Table.Cell>
                  </Table.Row>
                ))}
              </Table.Body>
            )}
          </Table>
          <button
            className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
            onClick={() => {
              setAddUserModal(true);
            }}
          >
            <BsFillPlusCircleFill
              size="30"
              className="mx-auto text-neutral-300 group-hover:text-accent-300 "
            />
          </button>

          {users && users.items && users.total_items > 0 && (
            <Paginator
              page={users.page ?? 0}
              per_page={users.per_page ?? 0}
              total_items={users.total_items ?? 0}
              total_pages={users.total_pages ?? 0}
              onPageChange={(page) => {
                setUserFilter({ ...userFilter, page: page });
              }}
            />
          )}
        </div>
      </section>
    </div>
  );
};

export default CompanyOverviewPage;
