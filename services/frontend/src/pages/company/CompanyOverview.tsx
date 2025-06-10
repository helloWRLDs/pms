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
import { useAssigneeList } from "../../hooks/useData";
import { formatTime } from "../../lib/utils/time";
import { ContextMenu } from "../../components/ui/ContextMenu";
import { TbTrash, TbEdit, TbUserCog } from "react-icons/tb";
import { usePermission } from "../../hooks/usePermission";
import { Permissions, Permission } from "../../lib/permission";
import { Role, RoleFilter } from "../../lib/roles";
import { Button } from "../../components/ui/Button";
import { cn } from "../../lib/utils/cn";
import RoleForm from "../../components/forms/RoleForm";

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
  const [newRoleModal, setNewRoleModal] = useState(false);
  const [editRoleModal, setEditRoleModal] = useState(false);
  const [selectedRole, setSelectedRole] = useState<Role | null>(null);

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

  const [roleFilter, setRoleFilter] = useState<RoleFilter>({
    company_id: company?.id ?? "",
    page: 1,
    per_page: 10,
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

  const { data: roles, refetch: refetchRoles } = useQuery({
    queryKey: [
      "roles",
      roleFilter.company_id,
      roleFilter.page,
      roleFilter.per_page,
    ],
    queryFn: () => authAPI.listRoles(roleFilter),
    enabled: !!company?.id,
  });

  useEffect(() => {
    if (!metaCache.metadata.selectedCompany?.id) {
      navigate("/companies");
    }
  }, [metaCache.metadata.selectedCompany?.id]);

  useEffect(() => {
    if (company?.id) {
      setRoleFilter((prev) => ({ ...prev, company_id: company.id }));
    }
  }, [company?.id]);

  const { assignees } = useAssigneeList(
    metaCache.metadata.selectedCompany?.id ?? ""
  );

  const handleCreateRole = async (roleData: Omit<Role, "created_at">) => {
    try {
      await authAPI.createRole({
        ...roleData,
        company_id: company?.id,
        created_at: {
          seconds: Math.floor(Date.now() / 1000),
          nanos: 0,
        },
      });
      infoToast("Role created successfully");
      setNewRoleModal(false);
      await refetchRoles();
    } catch (error) {
      console.error("Failed to create role:", error);
    }
  };

  const handleUpdateRole = async (roleData: Role) => {
    try {
      if (selectedRole) {
        await authAPI.updateRole(selectedRole.name, roleData);
        infoToast("Role updated successfully");
        setEditRoleModal(false);
        setSelectedRole(null);
        await refetchRoles();
      }
    } catch (error) {
      console.error("Failed to update role:", error);
    }
  };

  const handleDeleteRole = async (roleName: string) => {
    try {
      await authAPI.deleteRole(roleName);
      infoToast("Role deleted successfully");
      await refetchRoles();
    } catch (error) {
      console.error("Failed to delete role:", error);
    }
  };

  return (
    <div className="min-h-[100lvh] px-5 py-10 bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100">
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
            onFinish={async (userID, role) => {
              try {
                if (company?.id) {
                  await companyAPI.addParticipant(company.id, userID, role);
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
        <Modal
          title="Create New Role"
          visible={newRoleModal}
          onClose={() => setNewRoleModal(false)}
          className="w-[60%] mx-auto bg-primary-300 text-white"
        >
          <RoleForm
            onSubmit={handleCreateRole}
            onCancel={() => setNewRoleModal(false)}
          />
        </Modal>
        <Modal
          title="Edit Role"
          visible={editRoleModal}
          onClose={() => {
            setEditRoleModal(false);
            setSelectedRole(null);
          }}
          className="w-[60%] mx-auto bg-primary-300 text-white"
        >
          {selectedRole && (
            <RoleForm
              initialRole={selectedRole}
              onSubmit={handleUpdateRole}
              onCancel={() => {
                setEditRoleModal(false);
                setSelectedRole(null);
              }}
              isEditing
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

      {/* Roles Section */}
      <section className="mb-10">
        <div className="container mx-auto">
          <div className="flex items-center justify-between mb-5">
            <h2 className="text-2xl font-semibold">Roles & Permissions</h2>
            {hasPermission(Permissions.ROLE_WRITE_PERMISSION) && (
              <Button
                onClick={() => setNewRoleModal(true)}
                className="bg-accent-500 hover:bg-accent-600 text-white px-4 py-2 rounded-md flex items-center gap-2"
              >
                <BsFillPlusCircleFill size={16} />
                Create Role
              </Button>
            )}
          </div>

          <div className="bg-primary-500/50 backdrop-blur-sm rounded-lg border border-primary-400/30 overflow-visible">
            <table className="w-full">
              <Table.Head className="bg-primary-400/70 text-white">
                <Table.HeadCell>Role Name</Table.HeadCell>
                <Table.HeadCell>Permissions Count</Table.HeadCell>
                <Table.HeadCell>Created At</Table.HeadCell>
                <Table.HeadCell>Actions</Table.HeadCell>
              </Table.Head>
              <Table.Body className="divide-y divide-primary-400/20">
                {roles?.items?.map((role, index) => (
                  <Table.Row
                    key={role.name}
                    className="bg-primary-600/30 hover:bg-primary-500/40 text-white transition-colors"
                  >
                    <Table.Cell>
                      <div className="flex items-center gap-3">
                        <div className="p-2 bg-accent-500/20 rounded-full">
                          <TbUserCog className="text-accent-400" size={16} />
                        </div>
                        <div>
                          <div className="font-semibold text-white">
                            {role.name}
                          </div>
                          <div className="text-sm text-white/60 capitalize">
                            {role.company_id ? "Custom Role" : "System Role"}
                          </div>
                        </div>
                      </div>
                    </Table.Cell>
                    <Table.Cell>
                      <span className="bg-accent-500/20 text-accent-400 px-2 py-1 rounded-full text-sm border border-accent-500/30">
                        {role.permissions?.length || 0} permissions
                      </span>
                    </Table.Cell>
                    <Table.Cell className="text-white/80">
                      {role.created_at.seconds
                        ? formatTime(role.created_at.seconds)
                        : "N/A"}
                    </Table.Cell>
                    <Table.Cell>
                      <ContextMenu
                        placement="left"
                        trigger={<BsThreeDotsVertical />}
                        items={[
                          {
                            icon: <TbEdit />,
                            label: "Edit Role",
                            onClick: () => {
                              setSelectedRole(role);
                              setEditRoleModal(true);
                            },
                          },
                          {
                            icon: <TbTrash />,
                            label: "Delete Role",
                            onClick: () => handleDeleteRole(role.name),
                          },
                        ]}
                      />
                    </Table.Cell>
                  </Table.Row>
                ))}
              </Table.Body>
            </table>
            <button
              disabled={!hasPermission(Permissions.ROLE_WRITE_PERMISSION)}
              className={cn(
                "w-full cursor-pointer group hover:bg-primary-500/40 py-4 group:transition-all duration-300",
                !hasPermission(Permissions.ROLE_WRITE_PERMISSION) &&
                  "cursor-not-allowed opacity-50"
              )}
              onClick={() => {
                setNewRoleModal(true);
              }}
            >
              <BsFillPlusCircleFill
                size="30"
                className="mx-auto text-white/60 group-hover:text-accent-400"
              />
            </button>

            {(!roles?.items || roles.items.length === 0) && (
              <div className="text-center py-8 text-white/70">
                <TbUserCog size={48} className="mx-auto mb-4 opacity-50" />
                <p>No roles found. Create your first role to get started.</p>
              </div>
            )}
          </div>

          {roles && roles.items && roles.total_items > 0 && (
            <div className="mt-4">
              <Paginator
                page={roles.page ?? 0}
                per_page={roles.per_page ?? 0}
                total_items={roles.total_items ?? 0}
                total_pages={roles.total_pages ?? 0}
                onPageChange={(page) => {
                  setRoleFilter({ ...roleFilter, page: page });
                }}
              />
            </div>
          )}
        </div>
      </section>

      {/* People Section */}
      <section className="pb-10">
        <div className="container mx-auto">
          <h2 className="text-2xl font-semibold mb-5">People</h2>
          <div className="bg-primary-500/50 backdrop-blur-sm rounded-lg border border-primary-400/30 overflow-visible">
            <table className="w-full">
              <Table.Head className="bg-primary-400/70 text-white">
                <Table.HeadCell></Table.HeadCell>
                <Table.HeadCell>№</Table.HeadCell>
                <Table.HeadCell>Name</Table.HeadCell>
                <Table.HeadCell>Email</Table.HeadCell>
                <Table.HeadCell>Joined</Table.HeadCell>
                <Table.HeadCell>Role</Table.HeadCell>
                <Table.HeadCell></Table.HeadCell>
              </Table.Head>
              <Table.Body className="divide-y divide-primary-400/20">
                {assignees?.items?.map((assignee, index) => (
                  <Table.Row
                    key={assignee.id}
                    className="bg-primary-600/30 hover:bg-primary-500/40 text-white transition-colors"
                  >
                    <Table.Cell>
                      <div className="aspect-square w-[2rem]">
                        {assignee.avatar_url ? (
                          <img
                            src={assignee.avatar_url}
                            alt={`${assignee.first_name}'s avatar`}
                            className="rounded-full w-full h-full object-cover"
                          />
                        ) : assignee.avatar_img ? (
                          <img
                            src={`data:image/jpeg;base64,${assignee.avatar_img}`}
                            alt={`${assignee.first_name}'s avatar`}
                            className="rounded-full w-full h-full object-cover"
                          />
                        ) : (
                          <div className="aspect-square w-[2rem] bg-accent-500/20 rounded-full flex items-center justify-center">
                            <span className="text-accent-400 text-sm font-medium">
                              {assignee.first_name?.[0]}
                              {assignee.last_name?.[0]}
                            </span>
                          </div>
                        )}
                      </div>
                    </Table.Cell>
                    <Table.Cell className="text-white/80">
                      {index + 1}
                    </Table.Cell>
                    <Table.Cell className="text-white">
                      {assignee.first_name} {assignee.last_name}
                    </Table.Cell>
                    <Table.Cell className="text-white/80">
                      {assignee.email}
                    </Table.Cell>
                    <Table.Cell className="text-white/80">
                      {formatTime(assignee.created_at.seconds)}
                    </Table.Cell>
                    <Table.Cell>
                      <span className="bg-accent-500/20 text-accent-400 px-2 py-1 rounded-full text-sm border border-accent-500/30">
                        {assignees?.items &&
                        assignee.participants &&
                        assignee.participants.length > 0
                          ? assignee.participants?.find(
                              (participant) =>
                                participant.company_id ===
                                metaCache.metadata.selectedCompany?.id
                            )?.role
                          : "N/A"}
                      </span>
                    </Table.Cell>
                    <Table.Cell>
                      <ContextMenu
                        placement="left"
                        trigger={<BsThreeDotsVertical />}
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
                                usersRefetch();
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
              className={cn(
                "w-full cursor-pointer group hover:bg-primary-500/40 py-4 group:transition-all duration-300",
                !hasPermission(Permissions.USER_INVITE_PERMISSION) &&
                  "cursor-not-allowed opacity-50"
              )}
              disabled={!hasPermission(Permissions.USER_INVITE_PERMISSION)}
              onClick={() => {
                setAddUserModal(true);
              }}
            >
              <BsFillPlusCircleFill
                size="30"
                className="mx-auto text-white/60 group-hover:text-accent-400"
              />
            </button>

            {(!assignees?.items || assignees.items.length === 0) && (
              <div className="text-center py-8 text-white/70">
                <TbUserCog size={48} className="mx-auto mb-4 opacity-50" />
                <p>
                  No people found. Add your first team member to get started.
                </p>
              </div>
            )}
          </div>

          {users && users.items && users.total_items > 0 && (
            <div className="mt-4">
              <Paginator
                page={users.page ?? 0}
                per_page={users.per_page ?? 0}
                total_items={users.total_items ?? 0}
                total_pages={users.total_pages ?? 0}
                onPageChange={(page) => {
                  setUserFilter({ ...userFilter, page: page });
                }}
              />
            </div>
          )}
        </div>
      </section>
    </div>
  );
};

export default CompanyOverviewPage;
