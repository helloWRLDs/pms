import { useQuery } from "@tanstack/react-query";
import { useCompanyStore } from "../../store/selectedCompanyStore";
import companyAPI from "../../api/company";
import { useEffect, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { capitalize } from "../../lib/utils/string";
import ProjectCardWrapper from "../../components/prioject/ProjectCard";
import { Modal } from "../../components/ui/Modal";
import NewProjectForm from "../../components/forms/NewProjectForm";
import { useProjectStore } from "../../store/selectedProjectStore";
import { useNavigate } from "react-router-dom";
import { usePageSettings } from "../../hooks/usePageSettings";
import { Layouts } from "../../lib/layout/layout";
import authAPI from "../../api/authAPI";
import Paginator from "../../components/ui/Paginator";
import { BsFillPlusCircleFill } from "react-icons/bs";
import projectAPI from "../../api/projectsAPI";
import Table, { TableColumn } from "../../components/ui/Table";
import { UserFilter, User } from "../../lib/user/user";
import AddParticipantForm from "../../components/forms/AddParticipantForm";
import { infoToast } from "../../lib/utils/toast";
import useMetaCache from "../../store/useMetaCache";
import { Profile } from "../../components/profile/Profile";

const CompanyOverviewPage = () => {
  usePageSettings({
    title: "Dashboard",
    requireAuth: true,
    layout: Layouts.Companies,
  });

  const metaCache = useMetaCache();
  const navigate = useNavigate();

  const [newProjectModal, setNewProjectModal] = useState<boolean>(false);
  const [addUserModal, setAddUserModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [showProfileModal, setShowProfileModal] = useState(false);

  const {
    data: company,
    isLoading: isCompanyLoading,
    refetch: companyRefetch,
  } = useQuery({
    queryKey: ["company", metaCache.metadata.selectedCompany?.id],
    queryFn: () => companyAPI.get(metaCache.metadata.selectedCompany?.id ?? ""),
    enabled: !!metaCache.metadata.selectedCompany?.id,
  });

  const [userFilter, setUserFilter] = useState<UserFilter>({
    page: 1,
    per_page: 10,
    company_id: company?.id ?? "",
  });
  console.log("selected company", company);

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

  useEffect(() => {
    if (!metaCache.metadata.selectedCompany?.id) {
      navigate("/companies");
    }
  }, [metaCache.metadata.selectedCompany?.id]);

  const userColumns: TableColumn<User>[] = [
    {
      header: "",
      accessor: (user) => (
        <div className="aspect-square w-[2rem]">
          <img
            src={`data:image/jpeg;base64,${user?.avatar_img}`}
            alt={`${user.first_name}'s avatar`}
          />
        </div>
      ),
      className: "w-[4rem]",
    },
    {
      header: "№",
      accessor: (user: User, index?: number) => (index ?? 0) + 1,
    },
    {
      header: "Name",
      accessor: (user) => `${user.first_name} ${user.last_name}`,
    },
    {
      header: "Email",
      accessor: "email",
    },
  ];

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
          className="bg-secondary-100"
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
        <Modal
          title="User Profile"
          visible={showProfileModal}
          onClose={() => {
            setShowProfileModal(false);
            setSelectedUser(null);
          }}
          className="w-[80%] max-w-4xl mx-auto bg-primary-300 text-white"
        >
          {selectedUser && (
            <Profile
              user={selectedUser}
              variant="modal"
              isEditable={false}
              onClose={() => {
                setShowProfileModal(false);
                setSelectedUser(null);
              }}
            />
          )}
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
                    metaCache.setSelectedProject(project);
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
          <Table
            data={users?.items}
            columns={userColumns}
            isLoading={isUsersLoading}
            emptyMessage="No users found"
            className="rounded-lg overflow-hidden"
            onRowClick={(user) => {
              setSelectedUser(user);
              setShowProfileModal(true);
            }}
          />
          <button
            className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
            onClick={() => {
              setAddUserModal(true);
            }}
          >
            <BsFillPlusCircleFill
              size="30"
              className="mx-auto text-neutral-300 group-hover:text-accent-300"
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
