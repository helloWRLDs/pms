import { FC, useState } from "react";
import { User, UserOptional } from "../../lib/user/user";
import { useMutation } from "@tanstack/react-query";
import authAPI from "../../api/authAPI";
import { toast } from "react-hot-toast";
import { Button } from "../ui/Button";

interface ProfileProps {
  user: User;
  isEditable?: boolean;
  onClose?: () => void;
  onUpdate?: () => void;
  setAvatarModal?: (show: boolean) => void;
  className?: string;
  variant?: "page" | "modal";
}

export const Profile: FC<ProfileProps> = ({
  user: initialUser,
  isEditable = true,
  onClose,
  onUpdate,
  setAvatarModal,
  className = "",
  variant = "page",
}) => {
  const [userProfile, setUserProfile] = useState<UserOptional>(initialUser);

  const updateProfileMutation = useMutation({
    mutationFn: (updatedProfile: UserOptional) => {
      // Only send fields that can be updated, excluding email and system fields
      const updateData: Partial<User> = {
        first_name: updatedProfile.first_name,
        last_name: updatedProfile.last_name,
        phone: updatedProfile.phone,
        bio: updatedProfile.bio,
        avatar_img: updatedProfile.avatar_img,
        avatar_url: updatedProfile.avatar_url,
      };
      return authAPI.updateUser(initialUser.id, updateData);
    },
    onSuccess: () => {
      toast.success("Profile updated successfully!");
      onUpdate?.();
    },
    onError: (error) => {
      toast.error("Failed to update profile. Please try again.");
      console.error("Profile update error:", error);
    },
  });

  const handleSaveChanges = async () => {
    // Validation for required fields
    if (!userProfile?.first_name?.trim()) {
      toast.error("First name is required");
      return;
    }
    if (!userProfile?.last_name?.trim()) {
      toast.error("Last name is required");
      return;
    }

    updateProfileMutation.mutate(userProfile);
  };

  return (
    <div className={`${className} ${variant === "modal" ? "p-6" : ""}`}>
      <div className="grid grid-cols-12 gap-8">
        <div
          className={
            variant === "modal" ? "col-span-12" : "col-span-12 md:col-span-3"
          }
        >
          <div className="bg-primary-500/30 backdrop-blur-sm p-6 rounded-xl border border-primary-400/30">
            <div className="flex flex-col items-center">
              <div className="relative">
                <div className="w-32 h-32 rounded-full bg-accent-500/20 flex items-center justify-center overflow-hidden">
                  {userProfile?.avatar_img ? (
                    <img
                      src={`data:image/jpeg;base64,${userProfile.avatar_img}`}
                      alt="Profile"
                      className="w-full h-full object-cover"
                    />
                  ) : userProfile.avatar_url ? (
                    <img
                      src={userProfile.avatar_url}
                      alt="Profile"
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <span className="text-4xl text-accent-500">
                      {userProfile.first_name?.[0]?.toUpperCase() || "U"}
                    </span>
                  )}
                </div>
                {isEditable && (
                  <button
                    className="absolute bottom-0 right-0 bg-accent-500 text-white rounded-full p-2 hover:bg-accent-600 transition-colors duration-200 cursor-pointer"
                    onClick={() => setAvatarModal?.(true)}
                  >
                    <svg
                      className="w-5 h-5"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"
                      />
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"
                      />
                    </svg>
                  </button>
                )}
              </div>
              <h3 className="mt-4 text-xl font-medium text-white">
                {`${userProfile?.first_name} ${userProfile?.last_name}`}
              </h3>
              <p className="text-white/60">{userProfile?.email}</p>
            </div>
          </div>
        </div>

        <div
          className={
            variant === "modal" ? "col-span-12" : "col-span-12 md:col-span-9"
          }
        >
          <div className="bg-primary-500/30 backdrop-blur-sm p-6 rounded-xl border border-primary-400/30">
            <h2 className="text-xl font-semibold text-white mb-6">
              General Information
            </h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  First Name <span className="text-red-400">*</span>
                </label>
                <input
                  type="text"
                  value={userProfile?.first_name}
                  onChange={(e) =>
                    setUserProfile({
                      ...userProfile,
                      first_name: e.target.value,
                    })
                  }
                  disabled={!isEditable}
                  required
                  className={`w-full bg-primary-400/20 border rounded-lg py-2 px-3 text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-accent-500 disabled:opacity-50 disabled:cursor-not-allowed ${
                    !userProfile?.first_name?.trim()
                      ? "border-red-400/50"
                      : "border-primary-400/30"
                  }`}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Last Name <span className="text-red-400">*</span>
                </label>
                <input
                  type="text"
                  value={userProfile?.last_name}
                  onChange={(e) =>
                    setUserProfile({
                      ...userProfile,
                      last_name: e.target.value,
                    })
                  }
                  disabled={!isEditable}
                  required
                  className={`w-full bg-primary-400/20 border rounded-lg py-2 px-3 text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-accent-500 disabled:opacity-50 disabled:cursor-not-allowed ${
                    !userProfile?.last_name?.trim()
                      ? "border-red-400/50"
                      : "border-primary-400/30"
                  }`}
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Email{" "}
                  <span className="text-white/50 text-xs">(Read-only)</span>
                </label>
                <input
                  type="email"
                  value={userProfile?.email}
                  disabled={true}
                  className="w-full bg-primary-400/10 border border-primary-400/20 rounded-lg py-2 px-3 text-white/60 placeholder-white/30 cursor-not-allowed opacity-60"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Phone
                </label>
                <input
                  type="tel"
                  value={userProfile.phone}
                  placeholder="Add phone number"
                  onChange={(e) =>
                    setUserProfile({
                      ...userProfile,
                      phone: e.target.value,
                    })
                  }
                  disabled={!isEditable}
                  className="w-full bg-primary-400/20 border border-primary-400/30 rounded-lg py-2 px-3 text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-accent-500 disabled:opacity-50 disabled:cursor-not-allowed"
                />
              </div>
            </div>

            <div className="mt-8">
              <h2 className="text-xl font-semibold text-white mb-6">
                Additional Information
              </h2>
              <div className="space-y-6">
                <div>
                  <label className="block text-sm font-medium text-white/90 mb-2">
                    Bio
                  </label>
                  <textarea
                    value={userProfile.bio}
                    className="w-full bg-primary-400/20 border border-primary-400/30 rounded-lg py-2 px-3 text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-accent-500 disabled:opacity-50 disabled:cursor-not-allowed"
                    rows={4}
                    placeholder="Write a short bio about yourself..."
                    onChange={(e) => {
                      setUserProfile({
                        ...userProfile,
                        bio: e.target.value,
                      });
                    }}
                    disabled={!isEditable}
                  ></textarea>
                </div>
              </div>
            </div>

            {isEditable && (
              <div className="mt-8 flex justify-end space-x-4">
                {onClose && (
                  <Button variant="outline" onClick={onClose} className="px-6">
                    Cancel
                  </Button>
                )}
                <Button
                  variant="default"
                  onClick={handleSaveChanges}
                  className="px-6"
                  disabled={updateProfileMutation.isPending}
                >
                  {updateProfileMutation.isPending ? (
                    <div className="flex items-center gap-2">
                      <div className="w-4 h-4 border-t-2 border-b-2 border-white rounded-full animate-spin"></div>
                      Saving...
                    </div>
                  ) : (
                    "Save Changes"
                  )}
                </Button>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
