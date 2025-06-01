import { FC, useState } from "react";
import { usePageSettings } from "../../hooks/usePageSettings";
import { useAuthStore } from "../../store/authStore";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import authAPI from "../../api/authAPI";
import { Profile } from "../../components/profile/Profile";
import { Modal } from "../../components/ui/Modal";
import { ChangeAvatarForm } from "../../components/forms/ChangeAvatarForm";

const ProfilePage: FC = () => {
  usePageSettings({ title: "Profile", requireAuth: true });
  const { auth, isAuthenticated } = useAuthStore();
  const navigate = useNavigate();
  const [showAvatarModal, setShowAvatarModal] = useState(false);

  const {
    data: user,
    isLoading,
    refetch,
  } = useQuery({
    queryKey: ["user", auth?.user.id ?? ""],
    queryFn: () => authAPI.getUser(auth?.user?.id ?? ""),
    enabled: !!auth?.user?.id && isAuthenticated(),
  });

  const handleAvatarSuccess = () => {
    setShowAvatarModal(false);
    console.log("Avatar updated");
    refetch();
  };

  const handleAvatarCancel = () => {
    setShowAvatarModal(false);
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-primary-700 to-primary-600 py-8 px-4">
        <div className="flex items-center justify-center min-h-[60vh]">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-accent-500"></div>
        </div>
      </div>
    );
  }

  if (!user) {
    return null;
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary-700 to-primary-600 py-8 px-4">
      <div className="max-w-7xl mx-auto bg-primary-500/30 backdrop-blur-lg rounded-xl shadow-xl overflow-hidden">
        <div className="p-6 sm:p-8">
          <Profile
            user={user}
            setAvatarModal={setShowAvatarModal}
            variant="page"
            onClose={() => navigate(-1)}
          />
        </div>
      </div>

      {/* Avatar Change Modal */}
      <Modal
        visible={showAvatarModal}
        onClose={handleAvatarCancel}
        title="Change Avatar"
        size="md"
      >
        <ChangeAvatarForm
          user={user}
          onSuccess={handleAvatarSuccess}
          onCancel={handleAvatarCancel}
        />
      </Modal>
    </div>
  );
};

export default ProfilePage;
