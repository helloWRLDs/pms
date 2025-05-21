import { useEffect } from "react";
import { useCacheStore } from "../store/cacheStore";
import { Company } from "../lib/company/company";
import { Project } from "../lib/project/project";
import { Sprint } from "../lib/sprint/sprint";

export const useCacheLoader = ({
  companyList,
  projectList,
  sprintList,
}: {
  companyList?: { items: Company[] };
  projectList?: { items: Project[] };
  sprintList?: { items: Sprint[] };
}) => {
  const { setCompanies, setProjects, setSprints } = useCacheStore();

  useEffect(() => {
    if (companyList?.items?.length) {
      const companyMap = Object.fromEntries(
        companyList.items.map((c) => [c.id, c])
      );
      setCompanies(companyMap);
    }
  }, [companyList]);

  useEffect(() => {
    if (projectList?.items?.length) {
      const projectMap = Object.fromEntries(
        projectList.items.map((p) => [p.id, p])
      );
      setProjects(projectMap);
    }
  }, [projectList]);

  useEffect(() => {
    if (sprintList?.items?.length) {
      const sprintMap = Object.fromEntries(
        sprintList.items.map((s) => [s.id, s])
      );
      setSprints(sprintMap);
    }
  }, [sprintList]);
};
