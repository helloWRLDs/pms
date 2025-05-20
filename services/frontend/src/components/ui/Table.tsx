type TableProps = React.TableHTMLAttributes<HTMLTableElement> & {};
const Table = (props: TableProps) => {
  return (
    <table className="w-full border-gray-500" {...props}>
      {props.children}
    </table>
  );
};

type TableHeadProps = React.HTMLAttributes<HTMLTableSectionElement> & {};
const TableHead = (props: TableHeadProps) => {
  return (
    <thead className="bg-gray-100 text-left" {...props}>
      {props.children}
    </thead>
  );
};

type TableHeadCellProps = React.HTMLAttributes<HTMLTableHeaderCellElement> & {};
const TableHeadCell = (props: TableHeadCellProps) => {
  return (
    <th className="px-4 py-3 border-b font-semibold" {...props}>
      {props.children}
    </th>
  );
};

type TableBodyProps = React.HTMLAttributes<HTMLTableSectionElement> & {};
const TableBody = (props: TableBodyProps) => {
  return (
    <tbody className="bg-white divide-y divide-gray-200" {...props}>
      {props.children}
    </tbody>
  );
};

type TableRowProps = React.HTMLAttributes<HTMLTableRowElement> & {};
const TableRow = (props: TableRowProps) => {
  return (
    <tr className="hover:bg-accent-50 transition duration-150" {...props}>
      {props.children}
    </tr>
  );
};

type TableBodyCellProps = React.HTMLAttributes<HTMLTableCellElement> & {};
const TableBodyCell = (props: TableBodyCellProps) => {
  return (
    <td className="px-4 py-2 whitespace-nowrap text-gray-800" {...props}>
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
