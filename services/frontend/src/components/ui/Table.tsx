import React from "react";
import { twMerge } from "tailwind-merge";

export type TableColumn<T> = {
  header: React.ReactNode;
  accessor: keyof T | ((item: T, index?: number) => React.ReactNode);
  className?: string;
};

type TableProps<T> = {
  data?: T[];
  columns: TableColumn<T>[];
  isLoading?: boolean;
  loadingRows?: number;
  onRowClick?: (item: T) => void;
  emptyMessage?: string;
  className?: string;
  rowClassName?: string | ((item: T) => string);
  stickyHeader?: boolean;
};

const LoadingCell = () => (
  <div className="animate-pulse">
    <div className="h-4 bg-secondary-100 rounded w-3/4"></div>
  </div>
);

const LoadingRow = <T,>({ columns }: { columns: TableColumn<T>[] }) => (
  <TableRow>
    {columns.map((_, index) => (
      <TableCell key={index}>
        <LoadingCell />
      </TableCell>
    ))}
  </TableRow>
);

function Table<T>({
  data,
  columns,
  isLoading = false,
  loadingRows = 5,
  onRowClick,
  emptyMessage = "No data available",
  className,
  rowClassName,
  stickyHeader = false,
}: TableProps<T>) {
  const renderCell = (item: T, column: TableColumn<T>, index: number) => {
    if (typeof column.accessor === "function") {
      return column.accessor(item, index);
    }
    return item[column.accessor] as React.ReactNode;
  };

  return (
    <div className="overflow-x-auto">
      <table className={twMerge("w-full border-collapse", className)}>
        <thead
          className={twMerge(
            "bg-primary-400 text-neutral-100",
            stickyHeader && "sticky top-0"
          )}
        >
          <tr>
            {columns.map((column, index) => (
              <th
                key={index}
                className={twMerge(
                  "px-4 py-3 text-left font-semibold",
                  column.className
                )}
              >
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="divide-y divide-secondary-100">
          {isLoading ? (
            Array(loadingRows)
              .fill(0)
              .map((_, index) => <LoadingRow key={index} columns={columns} />)
          ) : !data || data.length === 0 ? (
            <tr>
              <td
                colSpan={columns.length}
                className="px-4 py-8 text-center text-neutral-400"
              >
                {emptyMessage}
              </td>
            </tr>
          ) : (
            data.map((item, rowIndex) => (
              <tr
                key={rowIndex}
                onClick={() => onRowClick?.(item)}
                className={twMerge(
                  "bg-secondary-200 text-neutral-100 transition-colors duration-200",
                  "hover:bg-secondary-100 cursor-pointer",
                  typeof rowClassName === "function"
                    ? rowClassName(item)
                    : rowClassName
                )}
              >
                {columns.map((column, colIndex) => (
                  <td
                    key={colIndex}
                    className={twMerge("px-4 py-3", column.className)}
                  >
                    {renderCell(item, column, rowIndex)}
                  </td>
                ))}
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
}

// Legacy support for compound components
const TableHead: React.FC<React.HTMLAttributes<HTMLTableSectionElement>> = ({
  className,
  ...props
}) => (
  <thead
    className={twMerge("bg-primary-400 text-neutral-100", className)}
    {...props}
  />
);

const TableHeadCell: React.FC<
  React.HTMLAttributes<HTMLTableHeaderCellElement>
> = ({ className, ...props }) => (
  <th
    className={twMerge("px-4 py-3 text-left font-semibold", className)}
    {...props}
  />
);

const TableBody: React.FC<React.HTMLAttributes<HTMLTableSectionElement>> = ({
  className,
  ...props
}) => (
  <tbody
    className={twMerge("divide-y divide-secondary-100", className)}
    {...props}
  />
);

const TableRow: React.FC<React.HTMLAttributes<HTMLTableRowElement>> = ({
  className,
  ...props
}) => (
  <tr
    className={twMerge(
      "bg-secondary-200 hover:bg-secondary-100 transition-colors duration-200",
      className
    )}
    {...props}
  />
);

const TableCell: React.FC<React.HTMLAttributes<HTMLTableCellElement>> = ({
  className,
  ...props
}) => <td className={twMerge("px-4 py-3", className)} {...props} />;

// Attach legacy compound components
Table.Head = TableHead;
Table.HeadCell = TableHeadCell;
Table.Body = TableBody;
Table.Row = TableRow;
Table.Cell = TableCell;

export default Table;
