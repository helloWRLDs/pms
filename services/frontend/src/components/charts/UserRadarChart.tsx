import { FC } from "react";
import { Radar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  RadialLinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend,
} from "chart.js";
import { ListItems } from "../../lib/utils/list";
import { Task } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";

ChartJS.register(
  RadialLinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend
);

type Props = {
  userTasks: Record<string, ListItems<Task>>;
  userNames: Record<string, string>;
};

const UserRadarChart: FC<Props> = ({ userTasks, userNames }) => {
  const labels = getTaskStatuses.map((status) =>
    capitalize(status.toLowerCase().replace(/_/g, " "))
  );
  if (
    Object.values(userTasks).length === 0 ||
    !Object.values(userTasks)[0].items
  ) {
    return;
  }

  const datasets = Object.entries(userTasks).map(
    ([userId, taskList], index) => {
      const colorBase = 100 + index * 30;
      return {
        label: userNames[userId],
        data: getTaskStatuses.map(
          (status) =>
            taskList.items.filter((task) => task.status === status).length
        ),
        backgroundColor: `rgba(${colorBase}, 100, 255, 0.2)`,
        borderColor: `rgba(${colorBase}, 100, 255, 1)`,
        borderWidth: 1,
      };
    }
  );

  const chartData = {
    labels,
    datasets,
  };

  return (
    <div className="bg-neutral-900 p-6 rounded shadow mt-6">
      <h2 className="text-white text-xl font-semibold mb-4 text-center">
        Task Status Comparison by User
      </h2>
      <Radar data={chartData} />
    </div>
  );
};

export default UserRadarChart;
