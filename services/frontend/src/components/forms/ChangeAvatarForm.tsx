import { FC, useState, useRef } from "react";
import { useMutation } from "@tanstack/react-query";
import { toast } from "react-hot-toast";
import authAPI from "../../api/authAPI";
import { User } from "../../lib/user/user";
import {
  fileToBase64,
  validateImageFile,
  validateUrl,
} from "../../lib/utils/base64";
import { Button } from "../ui/Button";

interface ChangeAvatarFormProps {
  user: User;
  onSuccess: () => void;
  onCancel: () => void;
}

export const ChangeAvatarForm: FC<ChangeAvatarFormProps> = ({
  user,
  onSuccess,
  onCancel,
}) => {
  const [avatarUrl, setAvatarUrl] = useState("");
  const [isUploading, setIsUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Avatar update mutation
  const updateAvatarMutation = useMutation({
    mutationFn: async ({
      avatarImg,
      avatarUrl,
    }: {
      avatarImg?: string;
      avatarUrl?: string;
    }) => {
      const updateData: { avatar_img?: string; avatar_url?: string } = {};
      if (avatarImg) {
        updateData.avatar_img = avatarImg;
        updateData.avatar_url = ""; // Clear URL when using image
      } else if (avatarUrl) {
        updateData.avatar_url = avatarUrl;
        updateData.avatar_img = ""; // Clear image when using URL
      }

      return authAPI.updateUser(user.id, updateData);
    },
    onSuccess: () => {
      toast.success("Avatar updated successfully!");
      setAvatarUrl("");
      onSuccess();
    },
    onError: (error) => {
      toast.error("Failed to update avatar. Please try again.");
      console.error("Avatar update error:", error);
    },
    onSettled: () => {
      setIsUploading(false);
    },
  });

  // Handle file upload
  const handleFileUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (!file) return;

    const validation = validateImageFile(file);
    if (!validation.isValid) {
      toast.error(validation.error!);
      return;
    }

    try {
      setIsUploading(true);
      const base64 = await fileToBase64(file);
      updateAvatarMutation.mutate({ avatarImg: base64, avatarUrl: "" });
    } catch (error) {
      toast.error("Failed to process image");
      setIsUploading(false);
    }
  };

  // Handle URL submit
  const handleUrlSubmit = () => {
    const validation = validateUrl(avatarUrl);
    if (!validation.isValid) {
      toast.error(validation.error!);
      return;
    }

    setIsUploading(true);
    updateAvatarMutation.mutate({ avatarUrl: avatarUrl.trim(), avatarImg: "" });
  };

  // Remove avatar
  const handleRemoveAvatar = () => {
    setIsUploading(true);
    updateAvatarMutation.mutate({ avatarImg: "", avatarUrl: "" });
  };

  return (
    <div className="space-y-6">
      {/* Current Avatar Preview */}
      <div className="flex justify-center">
        <div className="w-24 h-24 rounded-full bg-accent-500/20 flex items-center justify-center overflow-hidden">
          {user.avatar_img ? (
            <img
              src={`data:image/jpeg;base64,${user.avatar_img}`}
              alt="Current Avatar"
              className="w-full h-full object-cover"
            />
          ) : user.avatar_url ? (
            <img
              src={user.avatar_url}
              alt="Current Avatar"
              className="w-full h-full object-cover"
            />
          ) : (
            <span className="text-2xl text-accent-500">
              {user.first_name?.[0]?.toUpperCase() || "U"}
            </span>
          )}
        </div>
      </div>

      <div className="space-y-4">
        {/* File Upload Option */}
        <div>
          <label className="block text-sm font-medium text-white/90 mb-2">
            Upload Image
          </label>
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            onChange={handleFileUpload}
            className="hidden"
            disabled={isUploading}
          />
          <button
            onClick={() => fileInputRef.current?.click()}
            disabled={isUploading}
            className="w-full bg-primary-500/50 hover:bg-primary-500/70 border border-primary-400/30 rounded-lg py-3 px-4 text-white transition-colors duration-200 flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
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
                d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
              />
            </svg>
            Choose Image File
          </button>
          <p className="text-xs text-white/60 mt-1">
            Maximum size: 5MB. Supported formats: JPG, PNG, GIF
          </p>
        </div>

        {/* Divider */}
        <div className="flex items-center gap-4">
          <div className="flex-1 h-px bg-primary-400/30"></div>
          <span className="text-white/60 text-sm">OR</span>
          <div className="flex-1 h-px bg-primary-400/30"></div>
        </div>

        {/* URL Input Option */}
        <div>
          <label className="block text-sm font-medium text-white/90 mb-2">
            Image URL
          </label>
          <div className="flex gap-2">
            <input
              type="url"
              value={avatarUrl}
              onChange={(e) => setAvatarUrl(e.target.value)}
              placeholder="https://example.com/avatar.jpg"
              disabled={isUploading}
              className="flex-1 bg-primary-400/20 border border-primary-400/30 rounded-lg py-2 px-3 text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-accent-500 disabled:opacity-50"
            />
            <Button
              onClick={handleUrlSubmit}
              disabled={isUploading || !avatarUrl.trim()}
              variant="default"
              className="px-4"
            >
              Set
            </Button>
          </div>
        </div>

        {/* Remove Avatar Option */}
        {(user.avatar_img || user.avatar_url) && (
          <>
            <div className="flex items-center gap-4">
              <div className="flex-1 h-px bg-primary-400/30"></div>
              <span className="text-white/60 text-sm">OR</span>
              <div className="flex-1 h-px bg-primary-400/30"></div>
            </div>

            <Button
              onClick={handleRemoveAvatar}
              disabled={isUploading}
              variant="outline"
              className="w-full bg-red-500/20 hover:bg-red-500/30 border-red-400/30 text-red-400 hover:text-red-300"
            >
              Remove Current Avatar
            </Button>
          </>
        )}
      </div>

      {/* Loading State */}
      {isUploading && (
        <div className="flex items-center justify-center gap-2 text-white/70">
          <div className="w-4 h-4 border-t-2 border-b-2 border-accent-500 rounded-full animate-spin"></div>
          <span>Updating avatar...</span>
        </div>
      )}

      {/* Action Buttons */}
      <div className="flex justify-end gap-3 pt-4">
        <Button variant="outline" onClick={onCancel} disabled={isUploading}>
          Cancel
        </Button>
      </div>
    </div>
  );
};
