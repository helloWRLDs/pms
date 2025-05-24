import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { User, UserOptional } from "../lib/user/user";
import { Task, TaskOptional } from "../lib/task/task";

ChartJS.register(ArcElement, Tooltip, Legend);

type UserStats = Record<string, { done: number; archived: number }>;

// Dummy Users
const users: UserOptional[] = [
  { id: "u1", name: "Alice" },
  { id: "u2", name: "Bob" },
  { id: "u3", name: "Charlie" },
];

// Dummy Tasks
const tasks: TaskOptional[] = [
  { id: "t1", title: "Task 1", status: "done", assignee_id: "u1" },
  { id: "t2", title: "Task 2", status: "archived", assignee_id: "u1" },
  { id: "t3", title: "Task 3", status: "done", assignee_id: "u2" },
  { id: "t4", title: "Task 4", status: "archived", assignee_id: "u2" },
  { id: "t5", title: "Task 5", status: "done", assignee_id: "u2" },
  { id: "t6", title: "Task 6", status: "done", assignee_id: "u3" },
  { id: "t7", title: "Task 7", status: "done", assignee_id: "u3" },
];

// Group tasks by user and status
const groupTasksByUserStatus = (tasks: TaskOptional[]) => {
  const userStats: UserStats = {};

  tasks.forEach((task) => {
    if (!task.assignee_id) return;

    if (!userStats[task.assignee_id]) {
      userStats[task.assignee_id] = { done: 0, archived: 0 };
    }

    if (task.status === "done") userStats[task.assignee_id].done += 1;
    if (task.status === "archived") userStats[task.assignee_id].archived += 1;
  });

  return userStats;
};

const TestPage1 = () => {
  const stats = groupTasksByUserStatus(tasks);

  return (
    <div className="p-8 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8">
      {users.map((user) => {
        const data = stats[user.id] || { done: 0, archived: 0 };

        const chartData = {
          labels: ["Done", "Archived"],
          datasets: [
            {
              data: [data.done, data.archived],
              backgroundColor: ["#4ade80", "#f87171"],
              borderWidth: 1,
            },
          ],
        };

        return (
          <div key={user.id} className="bg-neutral-900 p-4 rounded shadow">
            <h3 className="text-lg font-semibold text-center text-white mb-2">
              {user.name}
            </h3>
            <Pie data={chartData} />
          </div>
        );
      })}
    </div>
  );
};

export default TestPage1;
