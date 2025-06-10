import { Route, Routes } from "react-router-dom";
import TestPage from "./TestPage";
import HomePage from "./HomePage";
import AgileDashboard from "./sprints/AgileDashboard";
import LoginPage from "./auth/LoginPage";
import RegisterPage from "./auth/RegisterPage";
import ProfilePage from "./auth/ProfilePage";
import CompaniesPage from "./company/CompaniesPage";
import BacklogPage from "./backlog/BacklogPage";
import { useAuthStore } from "../store/authStore";
import { useEffect } from "react";
import CompanyOverviewPage from "./company/CompanyOverview";
import TestPage1 from "./TestPage1";
import SprintsPage from "./sprints/SprintsPage";
import DocumentsPage from "./analytics/DocumentsPage";
import DocumentPage from "./analytics/DocumentPage";
import AnalyticsPage from "./analytics/AnalyticsPage";
import OAuthCallback from "../components/auth/OAuthCallback";
import useMetaCache from "../store/useMetaCache";
import CalendarPage from "./calendar/CalendarPage";
import Breadcrumb from "../components/ui/Breadcrumb";
import { useBreadcrumb } from "../hooks/useBreadcrumb";
import SideBar from "../components/ui/SideBar";
import { useSideBar } from "../hooks/useSideBar";

const Main = () => {
  const { auth, isAuthenticated } = useAuthStore();
  const metaCache = useMetaCache();
  const { breadcrumbItems } = useBreadcrumb();
  const { sideBarSections } = useSideBar();

  useEffect(() => {
    const currentUserId = auth?.user?.id;
    if (currentUserId && metaCache.metadata.currentUserId !== currentUserId) {
      metaCache.clearCache();
      metaCache.setCurrentUser(currentUserId);
    }
  }, [auth?.user?.id, metaCache]);

  return (
    <>
      <SideBar
        logo={{
          href: "/",
          label: "Taskflow",
        }}
      >
        {sideBarSections.map((section) => (
          <div key={section.id} className="space-y-1">
            {section.title && section.items.length > 0 && (
              <div className="px-3 py-2">
                <p className="text-xs font-semibold text-gray-400 uppercase tracking-wider">
                  {section.title}
                </p>
              </div>
            )}
            {section.items.map((item) => (
              <SideBar.Element
                key={item.id}
                icon={item.icon}
                onClick={item.onClick}
                isActive={item.isActive}
                level={item.level}
              >
                {item.label}
              </SideBar.Element>
            ))}
            {section.id !== "auth-section" && <div className="h-2" />}
          </div>
        ))}
      </SideBar>

      <div className="transition-all bg-gradient-to-br from-primary-700 to-primary-600 duration-300 ml-0 lg:ml-72">
        {isAuthenticated() && breadcrumbItems.length > 2 && (
          <div className="px-4 py-3 hidden lg:block">
            <Breadcrumb
              items={breadcrumbItems}
              className="overflow-x-auto"
              maxItems={10}
            />
          </div>
        )}
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/test" element={<TestPage />} />
          <Route path="/test1" element={<TestPage1 />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />
          <Route path="/companies" element={<CompaniesPage />} />
          <Route path="/projects" element={<CompanyOverviewPage />} />
          <Route path="/backlog" element={<BacklogPage />} />
          <Route path="/sprints" element={<SprintsPage />} />
          <Route path="/sprints/:sprintID" element={<AgileDashboard />} />
          <Route path="/documents" element={<DocumentsPage />} />
          <Route path="/documents/:documentID" element={<DocumentPage />} />
          <Route path="/analytics" element={<AnalyticsPage />} />
          <Route path="/calendar" element={<CalendarPage />} />
          <Route path="/auth/google/callback" element={<OAuthCallback />} />
          <Route path="/auth/callback" element={<OAuthCallback />} />
        </Routes>
      </div>
    </>
  );
};

export default Main;
