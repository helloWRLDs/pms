import { FC, useState } from "react";
import Dashboard from "../components/dashboard/Dashboard";
import SidePanel from "../components/dashboard/SidePanel";
import { RxDashboard } from "react-icons/rx";
import { Icon } from "../components/ui/Icon";

const DashboardPage: FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <div className="flex justify-between bg-primary-500">
        {/* SidePanel (Fixed Width) */}
        <SidePanel
          className={`transition-all duration-300 ease-in-out bg-primary-500 shadow-lg text-neutral-100 ${
            isOpen ? "w-1/5 opacity-100" : "w-0 opacity-0 overflow-hidden"
          }`}
        >
          {isOpen && (
            <div className="p-5">
              <div id="backlogs" className="mb-10">
                <h2 className="text-lg font-semibold mb-4">Project Backlogs</h2>
                <ul className="space-y-2">
                  {[
                    { name: "frontend", icon: "browser", color: "green" },
                    { name: "devops", icon: "docker", color: "blue" },
                    { name: "backend", icon: "server", color: "red" },
                    { name: "mobile", icon: "apple", color: "white" },
                  ].map((item) => (
                    <li
                      key={item.name}
                      className="border-b border-neutral-400 pb-2 flex items-center hover:bg-primary-400 py-2 px-2 cursor-pointer rounded-md"
                    >
                      <Icon
                        name={item.icon}
                        size={20}
                        color={item.color}
                        className="mr-2"
                      />
                      <span>{item.name}</span>
                    </li>
                  ))}
                </ul>
              </div>

              <div id="reports" className="mb-10">
                <h2 className="text-lg font-semibold mb-4">Reports</h2>
                <ul className="space-y-4">
                  {[
                    {
                      title: "Sprint 1",
                      duration: "1 Jan - 31 Jan",
                      progress: 80,
                      icon: "browser",
                      color: "green",
                    },
                    {
                      title: "Sprint 1",
                      duration: "1 Jan - 31 Jan",
                      progress: 60,
                      icon: "docker",
                      color: "blue",
                    },
                    {
                      title: "Sprint 2",
                      duration: "1 Feb - 28 Feb",
                      progress: 45,
                      icon: "server",
                      color: "red",
                    },
                  ].map((sprint) => (
                    <li
                      key={sprint.title}
                      className="border border-neutral-400 p-3 rounded-md hover:bg-primary-400 cursor-pointer"
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
                          className="h-2 rounded-full bg-primary-300"
                          style={{ width: `${sprint.progress}%` }}
                        ></div>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          )}
        </SidePanel>

        {/* Dashboard (Fills Remaining Space) */}
        <div className="w-full flex flex-col p-4 bg-primary-600 text-neutral-100">
          <button
            className="mb-4 p-2 bg-primary-300 hover:bg-primary-200 transition rounded-md self-start flex items-center gap-2 cursor-pointer"
            onClick={() => setIsOpen(!isOpen)}
          >
            <RxDashboard size={24} />
            <span>{isOpen ? "Hide Sidebar" : "Show Sidebar"}</span>
          </button>
          <h1 className="text-3xl font-bold mb-4">Dashboard</h1>
          <Dashboard className="w-full" />
        </div>
      </div>
    </>
  );
};

export default DashboardPage;
