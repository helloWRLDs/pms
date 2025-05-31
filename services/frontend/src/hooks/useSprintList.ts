import { useQuery } from "@tanstack/react-query";
import sprintAPI from "../api/sprintAPI";
import authAPI from "../api/authAPI";
import projectAPI from "../api/projectsAPI";

const useProjectList = (companyID: string) => {
  const {
    data: projects,
    isLoading: isLoadingProjects,
    error: errorProjects,
    refetch: refetchProjects,
  } = useQuery({
    queryKey: ["projects", companyID],
    queryFn: () =>
      projectAPI.list({
        company_id: companyID,
        page: 1,
        per_page: 1000,
      }),
    enabled: !!companyID,
    staleTime: 5 * 60 * 1000,
    gcTime: 10 * 60 * 1000,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
  });

  const getProjectName = (projectID: string | undefined) => {
    if (!projectID) return "none";
    return (
      projects?.items?.find((project) => project.id === projectID)?.title ??
      "none"
    );
  };

  return {
    projects,
    isLoadingProjects,
    errorProjects,
    refetchProjects,
    getProjectName,
  };
};

const useSprintList = (projectID: string) => {
  const {
    data: sprints,
    isLoading: isLoadingSprints,
    error: errorSprints,
    refetch: refetchSprints,
  } = useQuery({
    queryKey: ["sprints", projectID],
    queryFn: () =>
      sprintAPI.list({
        project_id: projectID,
        page: 1,
        per_page: 1000,
      }),
    enabled: !!projectID,
    staleTime: 5 * 60 * 1000,
    gcTime: 10 * 60 * 1000,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
  });

  const getSprintName = (sprintID: string | undefined) => {
    if (!sprintID) return "none";
    return (
      sprints?.items?.find((sprint) => sprint.id === sprintID)?.title ?? "none"
    );
  };

  return {
    sprints,
    isLoadingSprints,
    errorSprints,
    getSprintName,
    refetchSprints,
  };
};

const useAssigneeList = (companyID: string) => {
  const {
    data: assignees,
    isLoading: isLoadingAssignees,
    error: errorAssignees,
    refetch: refetchAssignees,
  } = useQuery({
    queryKey: ["assignees", companyID],
    queryFn: () =>
      authAPI.listUsers({
        company_id: companyID,
        page: 1,
        per_page: 1000,
      }),
    enabled: !!companyID,
  });

  const getAssigneeName = (assigneeID: string | undefined | null) => {
    if (!assigneeID) return "none";
    const assignee = assignees?.items?.find(
      (assignee) => assignee.id === assigneeID
    );
    if (!assignee) return "none";
    return `${assignee.first_name} ${assignee.last_name}`;
  };

  return {
    assignees,
    isLoadingAssignees,
    errorAssignees,
    refetchAssignees,
    getAssigneeName,
  };
};

export { useSprintList, useAssigneeList, useProjectList };
