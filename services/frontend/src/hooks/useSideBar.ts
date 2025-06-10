import { useState, useMemo, createElement } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import {
  MdOutlineHome,
  MdTaskAlt,
  MdOutlineCalendarToday,
  MdOutlineChat,
  MdOutlineLogin,
} from "react-icons/md";
import { PiUserCirclePlusLight } from "react-icons/pi";
import { IoIosLogOut } from "react-icons/io";
import { IoAnalyticsOutline, IoDocumentOutline } from "react-icons/io5";
import { GoOrganization, GoRepo } from "react-icons/go";
import { GiSprint } from "react-icons/gi";
import { RiProfileLine } from "react-icons/ri";
import { useAuthStore } from "../store/authStore";
import useMetaCache from "../store/useMetaCache";

export interface SideBarItem {
  id: string;
  icon: React.ReactNode;
  label: string;
  onClick: () => void;
  isActive?: boolean;
  level?: number;
}

export interface SideBarSection {
  id: string;
  title?: string;
  items: SideBarItem[];
}

export const useSideBar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();
  const { auth, isAuthenticated, clearAuth } = useAuthStore();
  const metaCache = useMetaCache();

  const handleLogout = () => {
    clearAuth();
    navigate("/login");
  };

  const mainNavigationItems: SideBarItem[] = useMemo(() => {
    const isLoggedIn = isAuthenticated();
    const currentPath = location.pathname;

    console.log(isLoggedIn);

    const items: SideBarItem[] = [
      {
        id: "home",
        icon: createElement(MdOutlineHome),
        label: "Home",
        onClick: () => navigate("/"),
        isActive: currentPath === "/",
      },
    ];

    if (isLoggedIn) {
      items.push(
        {
          id: "companies",
          icon: createElement(GoOrganization),
          label: "Companies",
          onClick: () => navigate("/companies"),
          isActive: currentPath === "/companies",
        },
        {
          id: "projects",
          icon: createElement(GoRepo),
          label: "Projects",
          onClick: () => navigate("/projects"),
          isActive: currentPath === "/projects",
        }
      );

      const selectedProject = metaCache.metadata.selectedProject;
      if (selectedProject) {
        items.push(
          {
            id: "backlog",
            icon: createElement(MdTaskAlt),
            label: "Tasks",
            onClick: () => navigate("/backlog"),
            isActive: currentPath === "/backlog",
          },
          {
            id: "sprints",
            icon: createElement(GiSprint),
            label: "Sprints",
            onClick: () => navigate("/sprints"),
            isActive: currentPath.startsWith("/sprints"),
          },
          {
            id: "documents",
            icon: createElement(IoDocumentOutline),
            label: "Documents",
            onClick: () => navigate("/documents"),
            isActive: currentPath.startsWith("/documents"),
          },
          {
            id: "analytics",
            icon: createElement(IoAnalyticsOutline),
            label: "Analytics",
            onClick: () => navigate("/analytics"),
            isActive: currentPath === "/analytics",
          }
        );
      }

      items.push({
        id: "calendar",
        icon: createElement(MdOutlineCalendarToday),
        label: "Calendar",
        onClick: () => navigate("/calendar"),
        isActive: currentPath === "/calendar",
      });
    }

    return items;
  }, [
    isAuthenticated,
    metaCache.metadata.selectedProject,
    navigate,
    location.pathname,
  ]);

  const profileItems: SideBarItem[] = useMemo(() => {
    const isLoggedIn = isAuthenticated();
    const currentPath = location.pathname;

    if (!isLoggedIn || !auth?.user?.id) return [];

    return [
      {
        id: `profile-${auth.user.id}`,
        icon: createElement(RiProfileLine),
        label: "Profile",
        onClick: () => navigate("/profile"),
        isActive: currentPath === "/profile",
      },
    ];
  }, [isAuthenticated, auth?.user?.id, navigate, location.pathname]);

  const authItems: SideBarItem[] = useMemo(() => {
    const isLoggedIn = isAuthenticated();
    const currentPath = location.pathname;

    if (isLoggedIn) {
      const userEmail = auth?.user?.email || "Unknown User";
      const truncatedEmail =
        userEmail.length > 20 ? `${userEmail.substring(0, 17)}...` : userEmail;

      return [
        {
          id: `logout-${auth?.user?.id}`,
          icon: createElement(IoIosLogOut),
          label: `Log out (${truncatedEmail})`,
          onClick: handleLogout,
        },
      ];
    }

    return [
      {
        id: "login",
        icon: createElement(MdOutlineLogin),
        label: "Log in",
        onClick: () => navigate("/login"),
        isActive: currentPath === "/login",
      },
      {
        id: "register",
        icon: createElement(PiUserCirclePlusLight),
        label: "Register",
        onClick: () => navigate("/register"),
        isActive: currentPath === "/register",
      },
    ];
  }, [isAuthenticated, handleLogout, navigate, auth?.user, location.pathname]);

  const sideBarSections: SideBarSection[] = useMemo(
    () => [
      {
        id: "main-navigation",
        title: "Navigation",
        items: mainNavigationItems,
      },
      {
        id: "profile-section",
        title: "Profile",
        items: profileItems,
      },
      {
        id: "auth-section",
        title: "Account",
        items: authItems,
      },
    ],
    [mainNavigationItems, profileItems, authItems]
  );

  return {
    isOpen,
    setIsOpen,
    sideBarSections,
    mainNavigationItems,
    profileItems,
    authItems,
    toggleSidebar: () => setIsOpen(!isOpen),
    closeSidebar: () => setIsOpen(false),
    openSidebar: () => setIsOpen(true),
  };
};

export default useSideBar;
