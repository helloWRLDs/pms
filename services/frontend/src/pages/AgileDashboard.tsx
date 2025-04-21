import { FC, useEffect, useState } from "react";
import Dashboard from "../components/dashboard/Dashboard";
import { Icon } from "../components/ui/Icon";
import SidePanel from "../components/ui/SidePanel";
import { PageSettings } from "./page";
import { usePageSettings } from "../hooks/usePageSettings";

const AgileDashboard: FC = () => {
  usePageSettings({ requireAuth: false, title: "Agile Dashboard" });

  return (
    <>
      <div className="flex justify-between bg-primary-200">
        <SidePanel
          isOpen={true}
          openContent={
            <div className="p-5">
              <h2 className="text-lg font-semibold mb-4">Reports</h2>
              <ul className="space-y-4">
                {[
                  {
                    id: 1,
                    title: "Sprint 1",
                    duration: "1 Jan - 31 Jan",
                    progress: 80,
                    icon: "browser",
                    color: "green",
                  },
                  {
                    id: 2,
                    title: "Sprint 1",
                    duration: "1 Jan - 31 Jan",
                    progress: 60,
                    icon: "docker",
                    color: "blue",
                  },
                  {
                    id: 3,
                    title: "Sprint 2",
                    duration: "1 Feb - 28 Feb",
                    progress: 45,
                    icon: "server",
                    color: "red",
                  },
                ].map((sprint) => (
                  <li
                    key={sprint.id}
                    className="h-24 border border-neutral-400 p-3 rounded-md hover:bg-primary-400 cursor-pointer"
                  >
                    <div className="flex items-center mb-2">
                      <Icon
                        name={sprint.icon}
                        size={20}
                        color={sprint.color}
                        className="mr-3"
                      />
                      <div>
                        <h3 className="text-md font-semibold">
                          {sprint.title}
                        </h3>
                        <p className="text-sm text-neutral-300">
                          {sprint.duration}
                        </p>
                      </div>
                    </div>
                    {/* ✅ Progress Bar */}
                    <div className="w-full bg-neutral-700 h-2 rounded-full">
                      <div
                        className="h-2 rounded-full bg-accent-700"
                        style={{ width: `${sprint.progress}%` }}
                      ></div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          }
          closedContent={
            <div className="p-5">
              <h2 className="text-lg font-semibold mb-4">Reports</h2>
              <ul className="space-y-4">
                {[
                  {
                    id: 1,
                    // title: "Sprint 1",
                    icon: "browser",
                    color: "green",
                  },
                  {
                    id: 2,
                    // title: "Sprint 1",
                    icon: "docker",
                    color: "blue",
                  },
                  {
                    // title: "Sprint 2",
                    id: 3,
                    icon: "server",
                    color: "red",
                  },
                ].map((sprint) => (
                  <li
                    key={sprint.id}
                    className="h-24 border border-neutral-400 p-3 rounded-md hover:bg-primary-400 cursor-pointer"
                  >
                    <div className="flex items-center mb-2">
                      <Icon
                        name={sprint.icon}
                        size={20}
                        color={sprint.color}
                        className="mr-3"
                      />
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          }
        />
        <div className="w-full flex flex-col p-4 bg-primary-600 text-neutral-100">
          <h1 className="text-3xl font-bold mb-4">Dashboard</h1>
          <Dashboard className="w-full" />
        </div>
      </div>
    </>
  );
};

export default AgileDashboard;
