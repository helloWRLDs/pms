type TableProps = React.TableHTMLAttributes<HTMLTableElement> & {};
const Table = ({ className, ...props }: TableProps) => {
  return (
    <table className={`w-full border-gray-500 ${className}`} {...props}>
      {props.children}
    </table>
  );
};

type TableHeadProps = React.HTMLAttributes<HTMLTableSectionElement> & {};
const TableHead = ({ className, ...props }: TableHeadProps) => {
  return (
    <thead className={`bg-gray-100 text-left ${className}`} {...props}>
      {props.children}
    </thead>
  );
};

type TableHeadCellProps = React.HTMLAttributes<HTMLTableHeaderCellElement> & {};
const TableHeadCell = ({ className, ...props }: TableHeadCellProps) => {
  return (
    <th className={`px-4 py-3 font-semibold ${className}`} {...props}>
      {props.children}
    </th>
  );
};

type TableBodyProps = React.HTMLAttributes<HTMLTableSectionElement> & {};
const TableBody = ({ className, ...props }: TableBodyProps) => {
  return (
    <tbody
      className={`bg-white divide-y divide-gray-200 ${className}`}
      {...props}
    >
      {props.children}
    </tbody>
  );
};

type TableRowProps = React.HTMLAttributes<HTMLTableRowElement> & {};
const TableRow = ({ className, ...props }: TableRowProps) => {
  return (
    <tr
      className={`hover:bg-accent-50 transition duration-150 ${className}`}
      {...props}
    >
      {props.children}
    </tr>
  );
};

type TableBodyCellProps = React.HTMLAttributes<HTMLTableCellElement> & {};
const TableBodyCell = ({ className, ...props }: TableBodyCellProps) => {
  return (
    <td className={`px-4 py-2 ${className}`} {...props}>
      {props.children}
    </td>
  );
};

Table.Head = TableHead;
Table.HeadCell = TableHeadCell;
Table.Body = TableBody;
Table.Row = TableRow;
Table.Cell = TableBodyCell;

export default Table;
