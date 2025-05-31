import { useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { TreeNode } from "../components/ui/TreeView";
import useMetaCache from "../store/useMetaCache";
import { useAuthStore } from "../store/authStore";
import { MdOutlineHome, MdOutlineDashboard, MdTaskAlt } from "react-icons/md";
import { GoOrganization, GoRepo } from "react-icons/go";
import { GiSprint } from "react-icons/gi";
import { IoAnalyticsOutline, IoDocumentOutline } from "react-icons/io5";

export const useNavigationTree = () => {
  const navigate = useNavigate();
  const metaCache = useMetaCache();
  const { isAuthenticated } = useAuthStore();

  const selectedCompany = metaCache.metadata.selectedCompany;
  const selectedProject = metaCache.metadata.selectedProject;

  const navigationTree = useMemo(() => {
    const isLoggedIn = isAuthenticated();

    const mainNodes: TreeNode[] = [
      {
        id: "home",
        label: "Home",
        icon: MdOutlineHome,
        onClick: () => navigate("/"),
        isActive: window.location.pathname === "/",
      },
    ];

    if (isLoggedIn) {
      const companiesNode: TreeNode = {
        id: "companies",
        label: "Companies",
        icon: GoOrganization,
        isActive: window.location.pathname === "/companies",
        onClick: () => navigate("/companies"),
      };

      const allCompaniesNode: TreeNode = {
        id: "all-companies",
        label: "All Companies",
        icon: GoOrganization,
        onClick: () => navigate("/companies"),
        isActive: window.location.pathname === "/companies",
      };

      if (selectedCompany) {
        const companyNode: TreeNode = {
          id: `company-${selectedCompany.id}`,
          label: selectedCompany.name,
          icon: GoRepo,
          isActive: true,
          children: [
            {
              id: "all-projects",
              label: "All Projects",
              icon: GoRepo,
              onClick: () => navigate("/projects"),
              isActive: window.location.pathname === "/projects",
            },
          ],
        };

        if (selectedProject) {
          const projectNode: TreeNode = {
            id: `project-${selectedProject.id}`,
            label: selectedProject.title,
            icon: GoRepo,
            isActive: true,
            children: [
              {
                id: "tasks",
                label: "Tasks",
                icon: MdTaskAlt,
                onClick: () => navigate("/backlog"),
                isActive: window.location.pathname === "/backlog",
              },
              {
                id: "sprints",
                label: "Sprints",
                icon: GiSprint,
                isActive:
                  window.location.pathname.includes("/sprints") ||
                  window.location.pathname === "/agile-dashboard",
                children: [
                  {
                    id: "sprint-list",
                    label: "Sprint List",
                    icon: GiSprint,
                    onClick: () => navigate("/sprints"),
                    isActive: window.location.pathname === "/sprints",
                  },
                  {
                    id: "agile-dashboard",
                    label: "Agile Dashboard",
                    icon: MdOutlineDashboard,
                    onClick: () => navigate("/agile-dashboard"),
                    isActive: window.location.pathname === "/agile-dashboard",
                  },
                ],
              },
              {
                id: "analytics",
                label: "Analytics",
                icon: IoAnalyticsOutline,
                onClick: () => navigate("/analytics"),
                isActive: window.location.pathname === "/analytics",
              },
              {
                id: "documents",
                label: "Documents",
                icon: IoDocumentOutline,
                onClick: () => navigate("/documents"),
                isActive: window.location.pathname.startsWith("/documents"),
              },
            ],
          };

          if (companyNode.children) {
            companyNode.children = companyNode.children.concat(projectNode);
          }
        }

        companiesNode.children = [allCompaniesNode, companyNode];
      } else {
        companiesNode.children = [allCompaniesNode];
      }

      mainNodes.push(companiesNode);
    }

    return mainNodes;
  }, [isAuthenticated, selectedCompany, selectedProject, navigate]);

  return navigationTree;
};

export default useNavigationTree;
