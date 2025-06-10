import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { BsFillPlusCircleFill, BsThreeDotsVertical } from "react-icons/bs";
import { TbTrash, TbUserCog } from "react-icons/tb";
import { Modal } from "../../../components/ui/Modal";
import Table from "../../../components/ui/Table";
import { ContextMenu } from "../../../components/ui/ContextMenu";
import Paginator from "../../../components/ui/Paginator";
import AddParticipantForm from "../../../components/forms/AddParticipantForm";
import { Profile } from "../../../components/profile/Profile";
import companyAPI from "../../../api/company";
import authAPI from "../../../api/authAPI";
import useMetaCache from "../../../store/useMetaCache";
import { usePermission } from "../../../hooks/usePermission";
import { useAssigneeList } from "../../../hooks/useData";
import { Permissions } from "../../../lib/permission";
import { UserFilter, User } from "../../../lib/user/user";
import { formatTime } from "../../../lib/utils/time";
import { infoToast } from "../../../lib/utils/toast";
import { useAuthStore } from "../../../store/authStore";

interface PeopleSectionProps {
  companyId: string;
}

const PeopleSection = ({ companyId }: PeopleSectionProps) => {
  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const [addUserModal, setAddUserModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [showProfileModal, setShowProfileModal] = useState(false);

  const [userFilter, setUserFilter] = useState<UserFilter>({
    page: 1,
    per_page: 10,
    company_id: companyId,
  });

  const { data: users, refetch: usersRefetch } = useQuery({
    queryKey: [
      "users",
      userFilter.page,
      userFilter.per_page,
      userFilter.company_id,
    ],
    queryFn: () => authAPI.listUsers(userFilter),
    enabled: !!companyId && hasPermission(Permissions.COMPANY_READ_PERMISSION),
  });

  const { assignees } = useAssigneeList(companyId);

  const handleViewProfile = (user: User) => {
    setSelectedUser(user);
    setShowProfileModal(true);
  };

  const handleRemoveUser = async (userId: string) => {
    if (!hasPermission(Permissions.COMPANY_DELETE_PERMISSION)) return;

    try {
      await companyAPI.removeParticipant(companyId, userId);
      infoToast("User removed successfully");
      await usersRefetch();
    } catch (error) {
      console.error("Failed to remove user:", error);
    }
  };

  if (!hasPermission(Permissions.COMPANY_READ_PERMISSION)) {
    return null;
  }

  return (
    <section className="pb-10">
      <Modal
        title="Add user"
        visible={addUserModal}
        onClose={() => setAddUserModal(false)}
        className="bg-secondary-100"
      >
        <AddParticipantForm
          onFinish={async (userID, role) => {
            try {
              await companyAPI.addParticipant(companyId, userID, role);
              infoToast("User added successfully");
            } catch (error) {
              console.error("Failed to add user:", error);
            } finally {
              await usersRefetch();
              setAddUserModal(false);
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
              <Table.HeadCell>Actions</Table.HeadCell>
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
                  <Table.Cell className="text-white/80">{index + 1}</Table.Cell>
                  <Table.Cell className="text-white">
                    <button
                      onClick={() => handleViewProfile(assignee)}
                      className="hover:text-accent-400 transition-colors"
                    >
                      {assignee.first_name} {assignee.last_name}
                    </button>
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
                    {hasPermission(Permissions.COMPANY_DELETE_PERMISSION) &&
                      assignee.id !==
                        useAuthStore.getState().auth?.user?.id && (
                        <ContextMenu
                          placement="left"
                          trigger={<BsThreeDotsVertical />}
                          items={[
                            {
                              icon: <TbTrash />,
                              label: "Remove from Company",
                              onClick: () => handleRemoveUser(assignee.id),
                            },
                          ]}
                        />
                      )}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </table>

          {hasPermission(Permissions.COMPANY_INVITE_PERMISSION) && (
            <button
              className="w-full cursor-pointer group hover:bg-primary-500/40 py-4 group:transition-all duration-300"
              onClick={() => setAddUserModal(true)}
            >
              <BsFillPlusCircleFill
                size="30"
                className="mx-auto text-white/60 group-hover:text-accent-400"
              />
            </button>
          )}

          {(!assignees?.items || assignees.items.length === 0) && (
            <div className="text-center py-8 text-white/70">
              <TbUserCog size={48} className="mx-auto mb-4 opacity-50" />
              <p>No people found. Add your first team member to get started.</p>
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
  );
};

export default PeopleSection;
