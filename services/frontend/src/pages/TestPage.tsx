import { useQuery } from "@tanstack/react-query";
import authAPI from "../api/authAPI";
import { RoleFilter } from "../lib/roles";
import { useEffect, useState } from "react";

const TestPage = () => {
  const [filter, setFilter] = useState<RoleFilter>({
    companyID: "dee3b9c8-b6a4-4106-9304-525b3da7dc30",
    withDefault: true,
    page: 1,
    perPage: 10,
  });
  const {
    data: roles,
    isLoading: isLoadingRoles,
    error: errorRoles,
  } = useQuery({
    queryKey: [
      "roles",
      filter.companyID,
      filter.withDefault,
      filter.page,
      filter.perPage,
    ],
    queryFn: () => authAPI.listRoles(filter),
  });

  useEffect(() => {
    if (roles) {
      console.log(JSON.stringify(roles, null, 2));
    }
  }, [roles]);

  return <div className="px-8 py-5 bg-primary-500">test</div>;
};

export default TestPage;
