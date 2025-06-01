import { useEffect, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { TreeNode } from "../components/ui/TreeView";
import useMetaCache from "../store/useMetaCache";
import { useAuthStore } from "../store/authStore";
import { MdOutlineHome, MdOutlineDashboard, MdTaskAlt } from "react-icons/md";
import { GoOrganization, GoRepo } from "react-icons/go";
import { GiSprint } from "react-icons/gi";
import { IoAnalyticsOutline, IoDocumentOutline } from "react-icons/io5";
import { useProjectList, useSprintList } from "./useSprintList";

export const useNavigationTree = () => {
  const navigate = useNavigate();
  const metaCache = useMetaCache();
  const { isAuthenticated, auth } = useAuthStore();

  const selectedCompany = metaCache.metadata.selectedCompany;
  const selectedProject = metaCache.metadata.selectedProject;
  const currentUserId = auth?.user?.id;

  // Fetch projects for the selected company
  const { projects } = useProjectList(selectedCompany?.id ?? "");

  // Fetch sprints for the selected project
  const { sprints } = useSprintList(selectedProject?.id ?? "");

  // Use useMemo with proper dependencies to ensure tree updates when user changes
  const navigationTree = useMemo(() => {
    const isLoggedIn = isAuthenticated();

    console.log("🌳 Rebuilding navigation tree for user:", currentUserId, {
      isLoggedIn,
      selectedCompany: selectedCompany?.name,
      selectedProject: selectedProject?.title,
      projectsCount: projects?.items?.length ?? 0,
      sprintsCount: sprints?.items?.length ?? 0,
    });

    const mainNodes: TreeNode[] = [
      {
        id: "home",
        label: "Home",
        icon: MdOutlineHome,
        onClick: () => navigate("/"),
        isActive: window.location.pathname === "/",
      },
    ];

    if (isLoggedIn && currentUserId) {
      // Companies node - directly navigable to companies page
      const companiesNode: TreeNode = {
        id: `companies-${currentUserId}`,
        label: "Companies",
        icon: GoOrganization,
        onClick: () => navigate("/companies"),
        isActive: window.location.pathname === "/companies",
        children: [],
      };

      // If a company is selected, show it in the tree
      if (selectedCompany) {
        const companyChildren: TreeNode[] = [
          {
            id: `projects-overview-${currentUserId}`,
            label: "Projects Overview",
            icon: GoRepo,
            onClick: () => navigate("/projects"),
            isActive: window.location.pathname === "/projects",
          },
        ];

        // Add individual projects to the company node
        if (projects?.items && projects.items.length > 0) {
          projects.items.forEach((project) => {
            const projectNode: TreeNode = {
              id: `project-${project.id}-${currentUserId}`,
              label: project.title,
              icon: GoRepo,
              isActive: selectedProject?.id === project.id,
              onClick: () => {
                metaCache.setSelectedProject(project);
                navigate("/projects");
              },
              children: [],
            };

            // If this is the selected project, add its sub-items
            if (selectedProject?.id === project.id) {
              projectNode.children = [
                {
                  id: `tasks-${currentUserId}`,
                  label: "Tasks",
                  icon: MdTaskAlt,
                  onClick: () => navigate("/backlog"),
                  isActive: window.location.pathname === "/backlog",
                },
                {
                  id: `sprints-${currentUserId}`,
                  label: "Sprints",
                  icon: GiSprint,
                  onClick: () => navigate("/sprints"),
                  isActive: window.location.pathname === "/sprints",
                  children: [],
                },
                {
                  id: `documents-${currentUserId}`,
                  label: "Documents",
                  icon: IoDocumentOutline,
                  onClick: () => navigate("/documents"),
                  isActive: window.location.pathname.startsWith("/documents"),
                },
              ];

              // Add individual sprints directly to the sprints node
              if (sprints?.items && sprints.items.length > 0) {
                const sprintNodes: TreeNode[] = sprints.items.map((sprint) => ({
                  id: `sprint-${sprint.id}-${currentUserId}`,
                  label: sprint.title,
                  icon: MdOutlineDashboard,
                  onClick: () => navigate(`/sprints/${sprint.id}`),
                  isActive: window.location.pathname.includes(
                    `/sprints/${sprint.id}`
                  ),
                }));

                // Find sprints node and add individual sprints as children
                const sprintsNode = projectNode.children?.find(
                  (child) => child.id === `sprints-${currentUserId}`
                );
                if (sprintsNode && sprintsNode.children) {
                  sprintsNode.children.push(...sprintNodes);
                }
              }
            }

            companyChildren.push(projectNode);
          });
        }

        // Add analytics at company level (after projects)
        companyChildren.push({
          id: `analytics-${currentUserId}`,
          label: "Analytics",
          icon: IoAnalyticsOutline,
          onClick: () => navigate("/analytics"),
          isActive: window.location.pathname === "/analytics",
        });

        const selectedCompanyNode: TreeNode = {
          id: `company-${selectedCompany.id}-${currentUserId}`,
          label: selectedCompany.name,
          icon: GoOrganization,
          isActive: true,
          children: companyChildren,
        };

        companiesNode.children = [selectedCompanyNode];
      }

      mainNodes.push(companiesNode);
    }

    return mainNodes;
  }, [
    isAuthenticated,
    currentUserId,
    selectedCompany,
    selectedProject,
    projects,
    sprints,
    navigate,
    metaCache,
    // Add window.location.pathname as dependency to update active states
    typeof window !== "undefined" ? window.location.pathname : "",
  ]);

  return navigationTree;
};

export default useNavigationTree;
