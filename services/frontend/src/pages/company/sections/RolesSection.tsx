import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { BsFillPlusCircleFill, BsThreeDotsVertical } from "react-icons/bs";
import { TbTrash, TbEdit, TbUserCog } from "react-icons/tb";
import { Modal } from "../../../components/ui/Modal";
import Table from "../../../components/ui/Table";
import { Button } from "../../../components/ui/Button";
import { ContextMenu } from "../../../components/ui/ContextMenu";
import Paginator from "../../../components/ui/Paginator";
import RoleForm from "../../../components/forms/RoleForm";
import authAPI from "../../../api/authAPI";
import { usePermission } from "../../../hooks/usePermission";
import { Permissions } from "../../../lib/permission";
import { Role, RoleFilter } from "../../../lib/roles";
import { formatTime } from "../../../lib/utils/time";
import { infoToast } from "../../../lib/utils/toast";

interface RolesSectionProps {
  companyId: string;
}

const RolesSection = ({ companyId }: RolesSectionProps) => {
  const { hasPermission } = usePermission();
  const [newRoleModal, setNewRoleModal] = useState(false);
  const [editRoleModal, setEditRoleModal] = useState(false);
  const [selectedRole, setSelectedRole] = useState<Role | null>(null);

  const [roleFilter, setRoleFilter] = useState<RoleFilter>({
    company_id: companyId,
    page: 1,
    per_page: 10,
  });

  const { data: roles, refetch: refetchRoles } = useQuery({
    queryKey: [
      "roles",
      roleFilter.company_id,
      roleFilter.page,
      roleFilter.per_page,
    ],
    queryFn: () => authAPI.listRoles(roleFilter),
    enabled: !!companyId && hasPermission(Permissions.ROLE_READ_PERMISSION),
  });

  const handleCreateRole = async (roleData: Omit<Role, "created_at">) => {
    try {
      await authAPI.createRole({
        ...roleData,
        company_id: companyId,
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

  // Don't render if user doesn't have permission to read roles
  if (!hasPermission(Permissions.ROLE_READ_PERMISSION)) {
    return null;
  }

  return (
    <section className="mb-10">
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
              {roles?.items?.map((role) => (
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
                    {(hasPermission(Permissions.ROLE_WRITE_PERMISSION) ||
                      hasPermission(Permissions.ROLE_DELETE_PERMISSION)) && (
                      <ContextMenu
                        placement="left"
                        trigger={<BsThreeDotsVertical />}
                        items={[
                          ...(hasPermission(Permissions.ROLE_WRITE_PERMISSION)
                            ? [
                                {
                                  icon: <TbEdit />,
                                  label: "Edit Role",
                                  onClick: () => {
                                    setSelectedRole(role);
                                    setEditRoleModal(true);
                                  },
                                },
                              ]
                            : []),
                          ...(hasPermission(Permissions.ROLE_DELETE_PERMISSION)
                            ? [
                                {
                                  icon: <TbTrash />,
                                  label: "Delete Role",
                                  onClick: () => handleDeleteRole(role.name),
                                },
                              ]
                            : []),
                        ]}
                      />
                    )}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </table>

          {hasPermission(Permissions.ROLE_WRITE_PERMISSION) && (
            <button
              className="w-full cursor-pointer group hover:bg-primary-500/40 py-4 group:transition-all duration-300"
              onClick={() => setNewRoleModal(true)}
            >
              <BsFillPlusCircleFill
                size="30"
                className="mx-auto text-white/60 group-hover:text-accent-400"
              />
            </button>
          )}

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
  );
};

export default RolesSection;
