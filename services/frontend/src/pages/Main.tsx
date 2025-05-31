import { Route, Routes, useNavigate } from "react-router-dom";
import TestPage from "./TestPage";
import HomePage from "./HomePage";
import AgileDashboard from "./sprints/AgileDashboard";
import LoginPage from "./auth/LoginPage";
import RegisterPage from "./auth/RegisterPage";
import ProfilePage from "./auth/ProfilePage";
import CompaniesPage from "./company/CompaniesPage";
import BacklogPage from "./BacklogPage";
import { MdOutlineLogin } from "react-icons/md";
import { PiUserCirclePlusLight } from "react-icons/pi";
import { IoIosLogOut } from "react-icons/io";
import { useAuthStore } from "../store/authStore";
import { useCallback, useMemo } from "react";
import CompanyOverviewPage from "./company/CompanyOverview";
import TestPage1 from "./TestPage1";
import SprintsPage from "./sprints/SprintsPage";
import DocumentsPage from "./analytics/DocumentsPage";
import DocumentPage from "./analytics/DocumentPage";
import { RiProfileLine } from "react-icons/ri";
import TreeView, { TreeNode } from "../components/ui/TreeView";
import useNavigationTree from "../hooks/useNavigationTree";
import AnalyticsPage from "./analytics/AnalyticsPage";

const Main = () => {
  const navigate = useNavigate();
  const { isAuthenticated, clearAuth, auth } = useAuthStore();

  const handleLogout = useCallback(() => {
    clearAuth();
    navigate("/login");
  }, [clearAuth, navigate]);

  const navigationTree = useNavigationTree();

  const profileNodes: TreeNode[] = useMemo(() => {
    const isLoggedIn = isAuthenticated();

    if (!isLoggedIn || !auth?.user?.id) return [];

    return [
      {
        id: "profile",
        label: "Profile",
        icon: RiProfileLine,
        onClick: () => navigate("/profile"),
        isActive: window.location.pathname === "/profile",
      },
    ];
  }, [isAuthenticated, auth?.user?.id, navigate]);

  const authNodes: TreeNode[] = useMemo(() => {
    const isLoggedIn = isAuthenticated();

    if (isLoggedIn) {
      return [
        {
          id: "logout",
          label: "Log out",
          icon: IoIosLogOut,
          onClick: handleLogout,
          className: "text-red-400 hover:text-red-300",
        },
      ];
    }

    return [
      {
        id: "login",
        label: "Log in",
        icon: MdOutlineLogin,
        onClick: () => navigate("/login"),
        isActive: window.location.pathname === "/login",
      },
      {
        id: "register",
        label: "Register",
        icon: PiUserCirclePlusLight,
        onClick: () => navigate("/register"),
        isActive: window.location.pathname === "/register",
      },
    ];
  }, [isAuthenticated, handleLogout, navigate]);

  return (
    <>
      <div className="fixed top-0 left-0 w-72 h-screen bg-[#1a1a1a] border-r border-[#2a2a2a]">
        <div className="flex items-center gap-3 px-6 py-4 border-b border-[#2a2a2a] bg-[#1a1a1a]">
          <a href="/" className="flex items-center gap-3">
            <img
              src="https://flowbite.com/docs/images/logo.svg"
              className="h-8"
              alt="Logo"
            />
            <span className="text-xl font-semibold text-white">Taskflow</span>
          </a>
        </div>

        <div className="flex flex-col h-[calc(100vh-4rem)] bg-[#1a1a1a]">
          <div className="flex-grow overflow-y-auto">
            <TreeView nodes={navigationTree} />
          </div>

          <div className="border-t border-[#2a2a2a] bg-[#1a1a1a]">
            <TreeView nodes={profileNodes} />
            <TreeView nodes={authNodes} />
          </div>
        </div>
      </div>

      <div className="transition-all duration-300 md:ml-72">
        <Routes>
          <Route path="/test1" element={<TestPage1 />} />
          <Route path="/test" element={<TestPage />} />
          <Route path="/" element={<HomePage />} />
          <Route path="/sprints/:sprintID" element={<AgileDashboard />} />
          <Route path="/agile-dashboard" element={<AgileDashboard />} />

          {/* authorization */}
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />

          {/* <Route path="/analytics" element={<AnalyticsPage />} /> */}
          <Route path="/companies" element={<CompaniesPage />} />
          <Route path="/projects" element={<CompanyOverviewPage />} />
          <Route path="/backlog" element={<BacklogPage />} />
          <Route path="/sprints" element={<SprintsPage />} />
          <Route path="/analytics" element={<AnalyticsPage />} />
          <Route path="/documents/:documentID" element={<DocumentPage />} />
          <Route path="/documents" element={<DocumentsPage />} />
          <Route path="/projects/:projectID/sprints/:sprintID" />
          <Route path="/projects/:projectID/sprints/:sprintID/tasks/:taskID" />

          {/* admin */}
          <Route path="/users" />
        </Routes>
      </div>
    </>
  );
};

export default Main;
