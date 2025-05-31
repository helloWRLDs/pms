import { FC } from "react";
import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { TaskStats } from "../../lib/stats/stats";

// Register Chart.js components
ChartJS.register(ArcElement, Tooltip, Legend);

type Props = {
  stats: TaskStats;
  userName: string;
  className?: string;
};

const UserTaskPieChart: FC<Props> = ({ stats, userName, className }) => {
  // Show only task distribution - not points
  const chartData = {
    labels: ["Done Tasks", "In Progress", "To Do Tasks"],
    datasets: [
      {
        data: [
          stats.done_tasks || 0,
          stats.in_progress_tasks || 0,
          stats.to_do_tasks || 0,
        ],
        backgroundColor: [
          "#10b981", // Done - green
          "#3b82f6", // In Progress - blue
          "#facc15", // To Do - yellow
        ],
        borderColor: ["#059669", "#2563eb", "#eab308"],
        borderWidth: 2,
      },
    ],
  };

  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: "bottom" as const,
        labels: {
          color: "#ffffff",
          padding: 20,
          usePointStyle: true,
        },
      },
      tooltip: {
        callbacks: {
          label: function (context: any) {
            const label = context.label;
            const value = context.parsed;
            const total = stats.total_tasks;
            const percentage =
              total > 0 ? ((value / total) * 100).toFixed(1) : "0";

            return `${label}: ${value} (${percentage}%)`;
          },
        },
      },
    },
  };

  const hasData = stats.total_tasks > 0;

  return (
    <div className={`p-6 rounded-lg shadow-lg bg-secondary-200 ${className}`}>
      <h3 className="text-lg font-semibold text-center text-white mb-4">
        {userName}
      </h3>

      {hasData ? (
        <div className="h-64 relative">
          <Pie data={chartData} options={options} />
        </div>
      ) : (
        <div className="h-64 flex items-center justify-center text-neutral-400">
          <div className="text-center">
            <div className="text-4xl mb-2">📊</div>
            <p>No task data available</p>
          </div>
        </div>
      )}

      {/* Stats Summary */}
      <div className="mt-4 pt-4 border-t border-secondary-100">
        <div className="grid grid-cols-2 gap-2 text-sm">
          <div className="text-neutral-300">
            <span className="font-medium">Total Tasks:</span>{" "}
            {stats.total_tasks}
          </div>
          <div className="text-purple-400">
            <span className="font-medium">Total Points:</span>{" "}
            {stats.total_points}
          </div>
          <div className="text-green-400">
            <span className="font-medium">Completed:</span>{" "}
            {stats.done_tasks || 0}
          </div>
          <div className="text-blue-400">
            <span className="font-medium">In Progress:</span>{" "}
            {stats.in_progress_tasks || 0}
          </div>
        </div>

        {/* Completion Rate */}
        {stats.total_tasks > 0 && (
          <div className="mt-3 pt-3 border-t border-secondary-100">
            <div className="flex items-center justify-between text-sm">
              <span className="text-neutral-300">Completion Rate:</span>
              <span className="text-accent-400 font-bold">
                {((stats.done_tasks / stats.total_tasks) * 100).toFixed(1)}%
              </span>
            </div>
            <div className="mt-2 w-full bg-secondary-100 rounded-full h-2">
              <div
                className="bg-accent-500 h-2 rounded-full transition-all duration-300"
                style={{
                  width: `${(stats.done_tasks / stats.total_tasks) * 100}%`,
                }}
              ></div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default UserTaskPieChart;
