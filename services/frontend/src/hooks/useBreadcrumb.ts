import { useMemo, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { BreadcrumbItem } from "../components/ui/Breadcrumb";
import useMetaCache from "../store/useMetaCache";
import { useSprintList } from "./useSprintList";
import { Sprint } from "../lib/sprint/sprint";

export const useBreadcrumb = () => {
  const navigate = useNavigate();
  const metaCache = useMetaCache();
  const [currentSprint, setCurrentSprint] = useState<Sprint | null>(null);

  const selectedCompany = metaCache.metadata.selectedCompany;
  const selectedProject = metaCache.metadata.selectedProject;

  // Get sprint data for sprint pages
  const { sprints } = useSprintList(selectedProject?.id ?? "");

  // Extract sprint ID from URL when on sprint page
  useEffect(() => {
    const currentPath = window.location.pathname;
    if (currentPath.includes("/sprints/")) {
      const sprintId = currentPath.split("sprints/")[1];
      if (sprintId && sprints?.items) {
        const sprint = sprints.items.find((s) => s.id === sprintId);
        setCurrentSprint(sprint || null);
      }
    } else {
      setCurrentSprint(null);
    }
  }, [sprints, window.location.pathname]);

  const breadcrumbItems: BreadcrumbItem[] = useMemo(() => {
    const items: BreadcrumbItem[] = [];

    // Always include Home as first item
    items.push({
      id: "home",
      label: "Home",
      onClick: () => navigate("/"),
      isClickable: true,
      isActive: window.location.pathname === "/",
    });

    // Add Companies level
    items.push({
      id: "companies",
      label: "Companies",
      onClick: () => navigate("/companies"),
      isClickable: true,
      isActive: window.location.pathname === "/companies",
    });

    // Add selected company if available
    if (selectedCompany) {
      items.push({
        id: `company-${selectedCompany.id}`,
        label: selectedCompany.name,
        onClick: () => {
          navigate("/projects");
        },
        isClickable: true,
        isActive: false,
      });

      // Add Projects level when in company context
      items.push({
        id: "projects",
        label: "Projects",
        onClick: () => navigate("/projects"),
        isClickable: true,
        isActive: window.location.pathname === "/projects",
      });

      // Add selected project if available
      if (selectedProject) {
        items.push({
          id: `project-${selectedProject.id}`,
          label: selectedProject.title,
          onClick: () => {
            navigate("/backlog");
          },
          isClickable: true,
          isActive: false,
        });

        // Add current page context based on route
        const currentPath = window.location.pathname;
        if (currentPath.includes("/backlog")) {
          items.push({
            id: "backlog",
            label: "Tasks",
            isClickable: false,
            isActive: true,
          });
        } else if (currentPath.includes("/sprints")) {
          items.push({
            id: "sprints",
            label: "Sprints",
            onClick: () => navigate("/sprints"),
            isClickable: true,
            isActive: currentPath === "/sprints",
          });

          // Add current sprint if viewing specific sprint
          if (
            currentSprint &&
            currentPath.includes(`/sprints/${currentSprint.id}`)
          ) {
            items.push({
              id: `sprint-${currentSprint.id}`,
              label: currentSprint.title,
              isClickable: false,
              isActive: true,
            });
          }
        } else if (currentPath.includes("/documents")) {
          items.push({
            id: "documents",
            label: "Documents",
            isClickable: false,
            isActive: true,
          });
        } else if (currentPath.includes("/analytics")) {
          items.push({
            id: "analytics",
            label: "Analytics",
            isClickable: false,
            isActive: true,
          });
        }
      }
    }

    return items;
  }, [
    selectedCompany,
    selectedProject,
    currentSprint,
    navigate,
    typeof window !== "undefined" ? window.location.pathname : "",
  ]);

  return {
    breadcrumbItems,
    currentLevel: breadcrumbItems.length,
    hasCompany: !!selectedCompany,
    hasProject: !!selectedProject,
  };
};
