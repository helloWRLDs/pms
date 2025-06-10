import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { usePageSettings } from "../../hooks/usePageSettings";
import { useQuery } from "@tanstack/react-query";
import analyticsAPI from "../../api/analyticsAPI";
import useMetaCache from "../../store/useMetaCache";
import UserTaskPieChart from "../../components/charts/UserTaskPieChart";
import { useSprintList } from "../../hooks/useData";
import { IoAnalyticsOutline } from "react-icons/io5";
import { BsPeople, BsTrophy } from "react-icons/bs";
import { BiChevronDown, BiChevronRight } from "react-icons/bi";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { safeNumber, safePercentage } from "../../lib/utils/num";
import { safeDisplay } from "../../lib/utils/num";

ChartJS.register(ArcElement, Tooltip, Legend);

const AnalyticsPage = () => {
  usePageSettings({ requireAuth: true, title: "Analytics" });

  const metaCache = useMetaCache();
  const navigate = useNavigate();
  const [expandedUsers, setExpandedUsers] = useState<Set<string>>(new Set());

  const { getSprintName } = useSprintList(
    metaCache.metadata.selectedProject?.id ?? ""
  );

  const { data: userStats, isLoading: isLoadingUserStats } = useQuery({
    queryKey: ["analytics", "stats", metaCache.metadata.selectedCompany?.id],
    queryFn: () =>
      analyticsAPI.getStats(metaCache.metadata.selectedCompany?.id ?? ""),
    enabled: !!metaCache.metadata.selectedCompany,
  });

  useEffect(() => {
    if (!metaCache.metadata.selectedCompany) {
      navigate("/companies");
    }
  }, [metaCache.metadata.selectedCompany, navigate]);

  const toggleUserExpanded = (userId: string) => {
    setExpandedUsers((prev) => {
      const newSet = new Set(prev);
      if (newSet.has(userId)) {
        newSet.delete(userId);
      } else {
        newSet.add(userId);
      }
      return newSet;
    });
  };

  const expandAll = () => {
    if (userStats) {
      setExpandedUsers(new Set(userStats.map((user) => user.user_id)));
    }
  };

  const collapseAll = () => {
    setExpandedUsers(new Set());
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100">
      <div className="px-6 py-8">
        <div className="max-w-7xl mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
            <div>
              <h1 className="text-3xl font-bold flex items-center gap-3">
                <IoAnalyticsOutline className="text-accent-500" />
                <span className="text-accent-500">
                  {metaCache.metadata.selectedCompany?.name ||
                    "Unknown Company"}
                </span>
                <span>Analytics</span>
              </h1>
              <p className="text-neutral-300 mt-2">
                Track team performance and task completion metrics
              </p>
            </div>

            <div className="flex items-center gap-4">
              <div className="flex items-center gap-2 text-sm text-neutral-400">
                <BsPeople className="text-accent-500" />
                <span>
                  {safeDisplay(userStats?.length)} team member
                  {userStats?.length !== 1 ? "s" : ""}
                </span>
              </div>

              {userStats && userStats.length > 0 && (
                <div className="flex items-center gap-2">
                  <button
                    onClick={expandAll}
                    className="px-3 py-1 text-xs bg-accent-500 text-primary-700 rounded hover:bg-accent-400 transition-colors"
                  >
                    Expand All
                  </button>
                  <button
                    onClick={collapseAll}
                    className="px-3 py-1 text-xs bg-secondary-100 text-neutral-300 rounded hover:bg-secondary-50 transition-colors"
                  >
                    Collapse All
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>

      <div className="px-6 pb-8">
        <div className="max-w-7xl mx-auto">
          {isLoadingUserStats && (
            <div className="space-y-4">
              {Array(3)
                .fill(0)
                .map((_, index) => (
                  <div
                    key={index}
                    className="bg-secondary-200 rounded-lg p-6 animate-pulse"
                  >
                    <div className="flex items-center gap-4">
                      <div className="w-12 h-12 bg-secondary-100 rounded-full"></div>
                      <div className="flex-1">
                        <div className="h-6 bg-secondary-100 rounded w-48 mb-2"></div>
                        <div className="h-4 bg-secondary-100 rounded w-32"></div>
                      </div>
                      <div className="flex gap-4">
                        <div className="text-center">
                          <div className="h-8 w-12 bg-secondary-100 rounded mb-1"></div>
                          <div className="h-3 w-16 bg-secondary-100 rounded"></div>
                        </div>
                        <div className="text-center">
                          <div className="h-8 w-12 bg-secondary-100 rounded mb-1"></div>
                          <div className="h-3 w-16 bg-secondary-100 rounded"></div>
                        </div>
                      </div>
                    </div>
                  </div>
                ))}
            </div>
          )}

          {!isLoadingUserStats && (!userStats || userStats.length === 0) && (
            <div className="bg-secondary-200 rounded-lg shadow-lg">
              <div className="px-6 py-12 text-center">
                <div className="flex flex-col items-center gap-4">
                  <IoAnalyticsOutline className="text-6xl text-neutral-500" />
                  <div>
                    <h3 className="text-xl font-semibold text-neutral-200 mb-2">
                      No Analytics Data Available
                    </h3>
                    <p className="text-neutral-400">
                      Task analytics will appear here once team members start
                      working on tasks
                    </p>
                  </div>
                </div>
              </div>
            </div>
          )}

          {!isLoadingUserStats && userStats && userStats.length > 0 && (
            <div className="space-y-4">
              {userStats.map((userStat) => {
                const isExpanded = expandedUsers.has(userStat.user_id);
                const overallStats = userStat.stats.overall;

                return (
                  <div
                    key={userStat.user_id}
                    className="bg-secondary-200 rounded-lg shadow-lg overflow-hidden"
                  >
                    <button
                      onClick={() => toggleUserExpanded(userStat.user_id)}
                      className="w-full bg-primary-400 px-6 py-4 border-b border-primary-300 hover:bg-primary-300 transition-colors"
                    >
                      <div className="flex items-center justify-between">
                        <div className="flex items-center gap-3">
                          <div className="flex items-center gap-2">
                            {isExpanded ? (
                              <BiChevronDown className="text-accent-500 text-xl" />
                            ) : (
                              <BiChevronRight className="text-accent-500 text-xl" />
                            )}
                            <div className="w-10 h-10 bg-accent-500 rounded-full flex items-center justify-center">
                              <span className="text-primary-700 font-bold text-lg">
                                {(userStat.first_name || "?").charAt(0)}
                                {(userStat.last_name || "?").charAt(0)}
                              </span>
                            </div>
                          </div>
                          <div className="text-left">
                            <h2 className="text-xl font-bold text-neutral-100">
                              {userStat.first_name || "Unknown"}{" "}
                              {userStat.last_name || "User"}
                            </h2>
                            <p className="text-neutral-300 text-sm">
                              {Object.keys(userStat.stats || {}).length} context
                              {Object.keys(userStat.stats || {}).length !== 1
                                ? "s"
                                : ""}{" "}
                              available
                            </p>
                          </div>
                        </div>

                        {overallStats && (
                          <div className="flex items-center gap-6 text-sm">
                            <div className="text-center">
                              <div className="text-accent-400 font-bold text-xl">
                                {safeDisplay(overallStats.total_tasks)}
                              </div>
                              <div className="text-neutral-400 text-xs">
                                Total Tasks
                              </div>
                            </div>
                            <div className="text-center">
                              <div className="text-green-400 font-bold text-xl">
                                {safeDisplay(overallStats.done_tasks)}
                              </div>
                              <div className="text-neutral-400 text-xs">
                                Completed
                              </div>
                            </div>
                            <div className="text-center">
                              <div className="text-purple-400 font-bold text-xl">
                                {safeDisplay(overallStats.total_points)}
                              </div>
                              <div className="text-neutral-400 text-xs">
                                Points
                              </div>
                            </div>
                            {safeNumber(overallStats.total_tasks) > 0 && (
                              <div className="text-center">
                                <div className="text-accent-500 font-bold text-xl">
                                  {safePercentage(
                                    overallStats.done_tasks,
                                    overallStats.total_tasks
                                  )}
                                  %
                                </div>
                                <div className="text-neutral-400 text-xs">
                                  Complete
                                </div>
                              </div>
                            )}
                          </div>
                        )}
                      </div>
                    </button>

                    {isExpanded && (
                      <div className="p-6 border-t border-secondary-100">
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                          {Object.entries(userStat.stats || {}).map(
                            ([key, value]) => {
                              const statsData = value || {
                                total_tasks: 0,
                                done_tasks: 0,
                                in_progress_tasks: 0,
                                to_do_tasks: 0,
                                total_points: 0,
                              };

                              const sanitizedStats = {
                                total_tasks: safeNumber(statsData.total_tasks),
                                done_tasks: safeNumber(statsData.done_tasks),
                                in_progress_tasks: safeNumber(
                                  statsData.in_progress_tasks
                                ),
                                to_do_tasks: safeNumber(statsData.to_do_tasks),
                                total_points: safeNumber(
                                  statsData.total_points
                                ),
                              };

                              return (
                                <div key={key} className="relative">
                                  {key === "overall" && (
                                    <div className="absolute top-2 right-2 z-10">
                                      <BsTrophy className="text-accent-500 text-lg" />
                                    </div>
                                  )}

                                  <div className="mb-4">
                                    <h3 className="text-lg font-semibold text-neutral-200 text-center">
                                      {key === "overall"
                                        ? "Overall Performance"
                                        : getSprintName(key) ||
                                          "Unknown Sprint"}
                                    </h3>
                                    {key !== "overall" && (
                                      <p className="text-neutral-400 text-sm text-center">
                                        Sprint Statistics
                                      </p>
                                    )}
                                  </div>

                                  <UserTaskPieChart
                                    stats={sanitizedStats}
                                    userName={`${
                                      userStat.first_name || "Unknown"
                                    } ${userStat.last_name || "User"}`}
                                    className="h-full"
                                  />
                                </div>
                              );
                            }
                          )}
                        </div>
                      </div>
                    )}
                  </div>
                );
              })}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default AnalyticsPage;
