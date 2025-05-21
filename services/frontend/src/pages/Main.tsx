import { GoOrganization, GoRepo } from "react-icons/go";
import { GiSprint } from "react-icons/gi";
import SideBar from "../components/ui/SideBar";
import { Route, Routes, useNavigate } from "react-router-dom";
import TestPage from "./TestPage";
import HomePage from "./HomePage";
import AgileDashboard from "./AgileDashboard";
import LoginPage from "./LoginPage";
import RegisterPage from "./RegisterPage";
import ProfilePage from "./ProfilePage";
import CompaniesPage from "./CompaniesPage";
import BacklogPage from "./BacklogPage";
import {
  MdOutlineDashboard,
  MdOutlineHome,
  MdOutlineLogin,
  MdTaskAlt,
} from "react-icons/md";
import { PiUserCirclePlusLight } from "react-icons/pi";
import { IoIosLogOut } from "react-icons/io";
import { useAuthStore } from "../store/authStore";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { SideBarItem } from "../lib/ui/sidebar";
import { useEffect, useState } from "react";
import CompanyOverviewPage from "./CompanyOverview";
import TestPage1 from "./TestPage1";
import { IoAnalyticsOutline, IoDocumentOutline } from "react-icons/io5";
import SprintsPage from "./SprintsPage";

const Main = () => {
  const navigate = useNavigate();
  const { isAuthenticated, clearAuth } = useAuthStore();
  const { selectCompany } = useCompanyStore();
  const [sideBarItems, setSideBarItems] = useState<Record<string, SideBarItem>>(
    {}
  );

  useEffect(() => {
    const isLoggedIn = isAuthenticated();

    setSideBarItems({
      Home: {
        isEnabled: true,
        label: "Home",
        icon: MdOutlineHome,
        onClick: () => navigate("/"),
      },
      Companies: {
        isEnabled: isLoggedIn,
        label: "Companies",
        icon: GoOrganization,
        onClick: () => navigate("/companies"),
      },
      Projects: {
        isEnabled: isLoggedIn && !!selectCompany,
        label: "Projects",
        icon: GoRepo,
        onClick: () => navigate("/projects"),
      },
      Tasks: {
        isEnabled: isLoggedIn,
        label: "Tasks",
        icon: MdTaskAlt,
        onClick: () => navigate("/backlog"),
      },
      Sprints: {
        isEnabled: isLoggedIn,
        label: "Sprints",
        icon: GiSprint,
        onClick: () => navigate("/sprints"),
      },
      Analytics: {
        isEnabled: isLoggedIn,
        label: "Analytics",
        icon: IoAnalyticsOutline,
      },
      Documents: {
        isEnabled: isLoggedIn,
        label: "Documents",
        icon: IoDocumentOutline,
      },
      AgileDashboard: {
        isEnabled: isLoggedIn,
        label: "Agile Dashboard",
        icon: MdOutlineDashboard,
        onClick: () => navigate("/agile-dashboard"),
      },
      Login: {
        className: "absolute bottom-15",
        isEnabled: !isLoggedIn,
        label: "Log in",
        icon: MdOutlineLogin,
        onClick: () => navigate("/login"),
      },
      Register: {
        className: "absolute bottom-5",
        isEnabled: !isLoggedIn,
        label: "Register",
        icon: PiUserCirclePlusLight,
        onClick: () => navigate("/register"),
      },
      Logout: {
        className: "absolute bottom-5",
        isEnabled: isLoggedIn,
        label: "Log out",
        icon: IoIosLogOut,
        onClick: () => {
          clearAuth();
          navigate("/login");
        },
      },
    });
  }, [isAuthenticated(), selectCompany]);

  return (
    <>
      <SideBar
        logo={{
          href: "/",
          imgSrc: "https://flowbite.com/docs/images/logo.svg",
          label: "Taskflow",
        }}
      >
        {Object.values(sideBarItems)
          .filter((item) => item.isEnabled)
          .map((barItem, i) => (
            <SideBar.Element
              key={i}
              className={`text-accent-500 hover:bg-accent-500 hover:text-black transition duration-300 ease-in-out ${
                barItem.className ?? ""
              }`}
              onClick={barItem.onClick}
            >
              <barItem.icon
                size="23"
                className="group-hover:text-black transition ease-in-out duration-300"
              />
              <span className="group-hover:text-black transition duration-300 ease-in-out">
                {barItem.label}
              </span>
              {barItem.badge}
            </SideBar.Element>
          ))}
      </SideBar>
      <div className="md:ml-64">
        <Routes>
          <Route path="/test1" element={<TestPage1 />} />
          <Route path="/test" element={<TestPage />} />
          <Route path="/" element={<HomePage />} />
          <Route path="/agile-dashboard" element={<AgileDashboard />} />

          {/* authorization */}
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />

          <Route path="/companies" element={<CompaniesPage />} />
          <Route path="/projects" element={<CompanyOverviewPage />} />
          <Route path="/backlog" element={<BacklogPage />} />
          <Route path="/sprints" element={<SprintsPage />} />
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
