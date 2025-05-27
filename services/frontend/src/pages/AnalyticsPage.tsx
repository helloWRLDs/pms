import { useQueries, useQuery } from "@tanstack/react-query";
import { useCompanyStore } from "../store/selectedCompanyStore";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import authAPI from "../api/auth";
import { useEffect, useState } from "react";
import { ListItems } from "../lib/utils/list";
import { Task } from "../lib/task/task";
import { taskAPI } from "../api/taskAPI";
import UserTaskPieChart from "../components/charts/UserTaskPieChart";
import UserRadarChart from "../components/charts/UserRadarChart";

ChartJS.register(ArcElement, Tooltip, Legend);

const AnalyticsPage = () => {
  const [userTasks, setUserTasks] = useState<Record<string, ListItems<Task>>>();

  const { selectedCompany } = useCompanyStore();
  const { data: users, isLoading: isUsersLoading } = useQuery({
    queryKey: ["users", selectedCompany?.id],
    queryFn: () =>
      authAPI.listUsers({
        page: 1,
        per_page: 50,
        company_id: selectedCompany?.id ?? "",
      }),
    enabled: !!selectedCompany?.id,
  });

  const collectTasksByUser = async (userID: string) => {
    try {
      const tasks = await taskAPI.list({
        page: 1,
        per_page: 10000,
        assignee_id: userID,
      });
      setUserTasks((prev) => ({
        ...prev,
        [userID]: tasks,
      }));
    } catch (e) {
      return;
    }
  };

  useEffect(() => {
    if (!isUsersLoading && users && users.items && users.items.length > 0) {
      users.items.forEach((user) => {
        if (!userTasks || !userTasks[user.id]) {
          collectTasksByUser(user.id);
        }
      });
    }
  }, [users]);

  useEffect(() => {
    console.log(JSON.stringify(userTasks));
  }, [userTasks]);
  return (
    <div className="w-full px-5 py-10 bg-primary-600 min-h-lvh text-neutral-100">
      <section className="mb-5">
        <div className="container mx-auto">
          <h1 className="text-3xl font-bold">
            <span className="text-accent-500">
              {selectedCompany?.name} Company
            </span>{" "}
            Charts
          </h1>
        </div>
      </section>
      <section>
        <div className="container mx-auto min-h-[200px]">
          <h2 className="text-2xl font-bold">Task Statuses Pie Charts</h2>
          <div className="container flex flex-row items-center flex-wrap justify-center">
            {!isUsersLoading &&
              users?.items.map((user) => {
                if (!userTasks) return null;

                const tasks = userTasks[user.id];
                if (!tasks || !tasks.items || tasks.items.length === 0)
                  return null;

                return (
                  <UserTaskPieChart
                    key={user.id}
                    userName={user.name}
                    tasks={tasks}
                    className="w-[30%]"
                  />
                );
              })}
          </div>
        </div>
      </section>

      <section>
        <div className="container mx-auto min-h-[200px]">
          <h2 className="text-2xl font-bold">Tasks Pie Charts</h2>
          <div className="container flex flex-row items-center flex-wrap ">
            {!isUsersLoading &&
              users &&
              users.items &&
              users.items.length > 0 &&
              userTasks && (
                <UserRadarChart
                  userTasks={userTasks}
                  userNames={users.items.reduce((acc, user) => {
                    acc[user.id] = user.name;
                    return acc;
                  }, {} as Record<string, string>)}
                />
              )}
          </div>
        </div>
      </section>
    </div>
  );
};

export default AnalyticsPage;
