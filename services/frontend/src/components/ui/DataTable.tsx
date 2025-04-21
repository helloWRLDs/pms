import { FC, ReactNode } from "react";

type Column<T> = {
  header: string;
  accessor: keyof T | ((row: T) => ReactNode);
  className?: string;
};

interface DataTableProps<T> {
  columns: Column<T>[];
  data: T[];
  onRowClick?: (row: T) => void;
  renderActions?: (row: T) => ReactNode;

  // Pagination support
  page?: number;
  perPage?: number;
  total?: number;
  onPageChange?: (newPage: number) => void;
}

export const DataTable = <T extends { id: string | number }>({
  columns,
  data,
  onRowClick,
  renderActions,
  page,
  perPage,
  total,
  onPageChange,
}: DataTableProps<T>) => {
  const totalPages = total && perPage ? Math.ceil(total / perPage) : 1;

  return (
    <div className="space-y-4">
      <div className="overflow-x-auto shadow-2xl rounded-md">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              {columns.map((col, i) => (
                <th
                  key={i}
                  className={`px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider ${
                    col.className ?? ""
                  }`}
                >
                  {col.header}
                </th>
              ))}
              {renderActions && (
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              )}
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {data.map((row) => (
              <tr
                key={row.id}
                onClick={onRowClick ? () => onRowClick(row) : undefined}
                className={`${
                  onRowClick ? "cursor-pointer hover:bg-gray-50" : ""
                }`}
              >
                {columns.map((col, i) => {
                  const value =
                    typeof col.accessor === "function"
                      ? col.accessor(row)
                      : (row[col.accessor] as ReactNode); // 🧠 Cast as ReactNode

                  return (
                    <td
                      key={i}
                      className="px-6 py-4 whitespace-nowrap text-sm text-gray-700"
                    >
                      {value}
                    </td>
                  );
                })}
                {renderActions && (
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    {renderActions(row)}
                  </td>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      {totalPages > 1 && onPageChange && (
        <div className="flex justify-end items-center gap-2">
          <button
            className="px-3 py-1 border rounded disabled:opacity-40"
            onClick={() => onPageChange((page || 1) - 1)}
            disabled={page === 1}
          >
            Previous
          </button>
          <span className="text-sm text-gray-600">
            Page {page} of {totalPages}
          </span>
          <button
            className="px-3 py-1 border rounded disabled:opacity-40"
            onClick={() => onPageChange((page || 1) + 1)}
            disabled={page === totalPages}
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
};
