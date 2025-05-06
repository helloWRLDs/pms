import { FC, useEffect, useState } from "react";
import authAPI from "../api/auth";
import useAuth from "../hooks/useAuth";
import { usePageSettings } from "../hooks/usePageSettings";
import { User } from "../lib/user/new";

const ProfilePage: FC = () => {
  usePageSettings({ title: "Profile", requireAuth: true });
  const [userProfile, setUserProfile] = useState<User>({
    id: "",
    email: "",
    name: "",
    bio: "",
    avatar_img: "",
    created_at: 0,
    updated_at: 0,
    phone: "",
  });
  const { access_token, user, isAuthenticated } = useAuth();

  useEffect(() => {
    if (user) {
      authAPI(access_token)
        .getUser(user.id)
        .then((res) => {
          setUserProfile(res);
        });
    }
  }, []);

  return (
    <div className="bg-gray-100 min-h-screen py-8 px-4">
      <div className="relative mx-auto bg-white min-h-screen">
        <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center mb-8">
            <h1 className="text-3xl font-bold text-gray-900">
              Profile Settings
            </h1>
            <button className="text-gray-400 hover:text-gray-500 !rounded-button">
              <i className="fas fa-times text-xl"></i>
            </button>
          </div>

          <div className="grid grid-cols-12 gap-8">
            {/* Left Sidebar */}
            <div className="col-span-12 md:col-span-3">
              <div className="bg-white p-6 rounded-lg shadow">
                <div className="flex flex-col items-center">
                  <div className="relative">
                    <img
                      src={`data:image/jpeg;base64,${userProfile?.avatar_img}`}
                      alt="Profile"
                      className="w-32 h-32 rounded-full object-cover"
                    />
                    <button className="absolute bottom-0 right-0 bg-[rgb(41,43,41)] text-white rounded-full p-2 hover:bg-[rgb(31,33,31)] !rounded-button">
                      <i className="fas fa-camera"></i>
                    </button>
                  </div>
                  <h3 className="mt-4 text-xl font-medium text-gray-900">
                    {userProfile?.name}
                  </h3>
                  <p className="text-gray-500">{"profileData.role"}</p>
                </div>

                <div className="mt-6 space-y-4">
                  <button className="w-full flex items-center px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 !rounded-button">
                    <i className="fas fa-user-circle mr-3"></i>
                    General Info
                  </button>
                  <button className="w-full flex items-center px-4 py-2 text-gray-700 rounded-md hover:bg-gray-100 !rounded-button">
                    <i className="fas fa-shield-alt mr-3"></i>
                    Security
                  </button>
                  <button className="w-full flex items-center px-4 py-2 text-gray-700 rounded-md hover:bg-gray-100 !rounded-button">
                    <i className="fas fa-bell mr-3"></i>
                    Notifications
                  </button>
                  <button className="w-full flex items-center px-4 py-2 text-gray-700 rounded-md hover:bg-gray-100 !rounded-button">
                    <i className="fas fa-palette mr-3"></i>
                    Preferences
                  </button>
                </div>
              </div>
            </div>

            {/* Main Content */}
            <div className="col-span-12 md:col-span-9">
              <div className="bg-white p-6 rounded-lg shadow">
                <h2 className="text-xl font-semibold text-gray-900 mb-6">
                  General Information
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Name
                    </label>
                    <input
                      type="text"
                      value={userProfile?.name}
                      onChange={(e) =>
                        setUserProfile({
                          ...userProfile,
                          name: e.target.value,
                        })
                      }
                      className="w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Email
                    </label>
                    <input
                      type="email"
                      value={userProfile?.email}
                      onChange={(e) =>
                        setUserProfile({
                          ...userProfile,
                          email: e.target.value,
                        })
                      }
                      className="w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
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
                      className="w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm"
                    />
                  </div>
                </div>

                <div className="mt-8">
                  <h2 className="text-xl font-semibold text-gray-900 mb-6">
                    Additional Information
                  </h2>
                  <div className="space-y-6">
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Bio
                      </label>
                      <textarea
                        value={userProfile.bio}
                        className="w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm"
                        rows={4}
                        placeholder="Write a short bio about yourself..."
                        onChange={(e) => {
                          setUserProfile({
                            ...userProfile,
                            bio: e.target.value,
                          });
                        }}
                      ></textarea>
                    </div>
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Time Zone
                      </label>
                      <div className="relative">
                        <select className="w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-[rgb(41,43,41)] focus:border-[rgb(41,43,41)] sm:text-sm appearance-none">
                          <option>Pacific Time (PT)</option>
                          <option>Mountain Time (MT)</option>
                          <option>Central Time (CT)</option>
                          <option>Eastern Time (ET)</option>
                        </select>
                        <i className="fas fa-chevron-down absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400"></i>
                      </div>
                    </div>
                  </div>
                </div>

                <div className="mt-8 flex justify-end space-x-4">
                  <button
                    // onClick={() => setShowProfileSettings(false)}
                    className="px-6 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[rgb(41,43,41)] !rounded-button"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={() => {
                      // Handle save profile
                      // setShowProfileSettings(false);
                    }}
                    className="px-6 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[rgb(41,43,41)] hover:bg-[rgb(31,33,31)] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[rgb(41,43,41)] !rounded-button"
                  >
                    Save Changes
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;
