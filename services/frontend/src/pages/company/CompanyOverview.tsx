import { useQuery } from "@tanstack/react-query";
import companyAPI from "../../api/company";
import { useEffect, useState } from "react";
import { capitalize } from "../../lib/utils/string";
import ProjectCardWrapper from "../../components/prioject/ProjectCard";
import { Modal } from "../../components/ui/Modal";
import NewProjectForm from "../../components/forms/NewProjectForm";
import { useNavigate } from "react-router-dom";
import { usePageSettings } from "../../hooks/usePageSettings";
import { Layouts } from "../../lib/layout/layout";
import authAPI from "../../api/authAPI";
import Paginator from "../../components/ui/Paginator";
import { BsFillPlusCircleFill, BsThreeDotsVertical } from "react-icons/bs";
import projectAPI from "../../api/projectsAPI";
import Table from "../../components/ui/Table";
import { UserFilter, User } from "../../lib/user/user";
import AddParticipantForm from "../../components/forms/AddParticipantForm";
import { infoToast } from "../../lib/utils/toast";
import useMetaCache from "../../store/useMetaCache";
import { Profile } from "../../components/profile/Profile";
import { useAssigneeList } from "../../hooks/useSprintList";
import { formatTime } from "../../lib/utils/time";
import { ContextMenu } from "../../components/ui/ContextMenu";
import { TbTrash } from "react-icons/tb";
import { usePermission } from "../../hooks/usePermission";
import { Permissions } from "../../lib/permission";

const CompanyOverviewPage = () => {
  usePageSettings({
    title: "Dashboard",
    requireAuth: true,
    layout: Layouts.Companies,
  });

  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const navigate = useNavigate();

  const [newProjectModal, setNewProjectModal] = useState<boolean>(false);
  const [addUserModal, setAddUserModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [showProfileModal, setShowProfileModal] = useState(false);

  const { data: company } = useQuery({
    queryKey: ["company", metaCache.metadata.selectedCompany?.id],
    queryFn: () => companyAPI.get(metaCache.metadata.selectedCompany?.id ?? ""),
    enabled: !!metaCache.metadata.selectedCompany?.id,
  });

  const {
    data: projects,
    isLoading: isProjectsLoading,
    refetch: projectsRefetch,
  } = useQuery({
    queryKey: ["projects", metaCache.metadata.selectedCompany?.id],
    queryFn: () =>
      projectAPI.list({
        page: 1,
        per_page: 1000,
        company_id: metaCache.metadata.selectedCompany?.id ?? "",
      }),
    enabled: !!metaCache.metadata.selectedCompany?.id,
  });

  const [userFilter, setUserFilter] = useState<UserFilter>({
    page: 1,
    per_page: 10,
    company_id: company?.id ?? "",
  });
  console.log("selected company", company);

  const { data: users, refetch: usersRefetch } = useQuery({
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

  const { assignees } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100 ">
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
                await projectsRefetch();
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
                }
              } catch (e) {
                console.error(e);
              } finally {
                await usersRefetch();
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
            {isProjectsLoading ? (
              <p>Loading...</p>
            ) : (
              projects?.items?.map((project, i) => (
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
          <table className="w-full">
            <Table.Head>
              <Table.HeadCell></Table.HeadCell>
              <Table.HeadCell>№</Table.HeadCell>
              <Table.HeadCell>Name</Table.HeadCell>
              <Table.HeadCell>Email</Table.HeadCell>
              <Table.HeadCell>Joined</Table.HeadCell>
              <Table.HeadCell></Table.HeadCell>
            </Table.Head>
            <Table.Body>
              {assignees?.items?.map((assignee, index) => (
                <Table.Row key={assignee.id}>
                  <Table.Cell>
                    <div className="aspect-square w-[2rem]">
                      {assignee.avatar_url ? (
                        <img
                          src={assignee.avatar_url}
                          alt={`${assignee.first_name}'s avatar`}
                          className="rounded-full"
                        />
                      ) : assignee.avatar_img ? (
                        <img
                          src={`data:image/jpeg;base64,${assignee.avatar_img}`}
                          alt={`${assignee.first_name}'s avatar`}
                        />
                      ) : (
                        <div className="aspect-square w-[2rem] bg-secondary-200 rounded-full"></div>
                      )}
                    </div>
                  </Table.Cell>
                  <Table.Cell>{index + 1}</Table.Cell>
                  <Table.Cell>
                    {assignee.first_name} {assignee.last_name}
                  </Table.Cell>

                  <Table.Cell>{assignee.email}</Table.Cell>
                  <Table.Cell>
                    {formatTime(assignee.created_at.seconds)}
                  </Table.Cell>
                  <Table.Cell>
                    <ContextMenu
                      placement="left"
                      items={[
                        {
                          icon: <TbTrash />,
                          label: "Delete",
                          onClick: async () => {
                            try {
                              await companyAPI.removeParticipant(
                                company?.id ?? "",
                                assignee.id
                              );
                              infoToast("user removed");
                            } catch (e) {
                              console.error(e);
                            } finally {
                              await usersRefetch();
                            }
                          },
                        },
                      ]}
                    />
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </table>
          <button
            className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
            disabled={!hasPermission(Permissions.USER_ADD_PERMISSION)}
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
