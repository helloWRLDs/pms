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
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";
import { UserTaskStats } from "../../api/analyticsAPI";

ChartJS.register(
  RadialLinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend
);

type Props = {
  userStats: UserTaskStats[];
  className?: string;
};

const UserRadarChart: FC<Props> = ({ userStats, className }) => {
  const chartData = {
    labels: getTaskStatuses.map((status) =>
      capitalize(status.toLowerCase().replace(/_/g, " "))
    ),
    datasets: userStats.map((stat, index) => ({
      label: `${stat.first_name} ${stat.last_name}`,
      data: getTaskStatuses.map((status) => stat.tasks_by_status[status] || 0),
      backgroundColor: `hsla(${
        (index * 360) / userStats.length
      }, 70%, 50%, 0.2)`,
      borderColor: `hsla(${(index * 360) / userStats.length}, 70%, 50%, 1)`,
      borderWidth: 2,
    })),
  };

  const options = {
    scales: {
      r: {
        angleLines: {
          color: "rgba(255, 255, 255, 0.1)",
        },
        grid: {
          color: "rgba(255, 255, 255, 0.1)",
        },
        pointLabels: {
          color: "rgba(255, 255, 255, 0.7)",
          font: {
            size: 12,
          },
        },
        ticks: {
          color: "rgba(255, 255, 255, 0.7)",
          backdropColor: "transparent",
        },
      },
    },
    plugins: {
      legend: {
        position: "top" as const,
        labels: {
          color: "rgba(255, 255, 255, 0.7)",
          font: {
            size: 12,
          },
        },
      },
    },
  };

  return (
    <div className={className}>
      <Radar data={chartData} options={options} />
    </div>
  );
};

export default UserRadarChart;
