import { FC } from "react";
import { Pie } from "react-chartjs-2";
import { ListItems } from "../../lib/utils/list";
import { Task } from "../../lib/task/task";
import { getTaskStatuses } from "../../lib/task/status";
import { capitalize } from "../../lib/utils/string";

type Props = {
  userName: string;
  tasks: ListItems<Task>;
  className?: string;
};

const UserTaskPieChart: FC<Props> = ({ userName, tasks, className }) => {
  const countStatuses = (status: string): number => {
    if (!tasks || !tasks.items || tasks.items.length === 0) {
      return 0;
    }
    return tasks.items.filter((task) => task.status === status).length;
  };

  const chartData = {
    labels: getTaskStatuses.map((status) =>
      capitalize(status.toLowerCase().replace(/_/g, " "))
    ),
    datasets: [
      {
        data: getTaskStatuses.map((status) => countStatuses(status)),
        backgroundColor: [
          "#22c55e", // CREATED
          "#3b82f6", // IN_PROGRESS
          "#facc15", // PENDING
          "#10b981", // DONE
          "#ef4444", // ARCHIVED
        ],
        borderWidth: 2,
      },
    ],
  };

  return (
    <div className={` p-4 rounded shadow ${className}`}>
      <h3 className="text-lg font-semibold text-center text-white mb-2">
        {userName}
      </h3>
      <Pie data={chartData} />
    </div>
  );
};

export default UserTaskPieChart;
