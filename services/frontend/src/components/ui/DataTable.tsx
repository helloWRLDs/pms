type Options = {
  heads: { label?: string; key: string }[];
  data: any[] | undefined;
  borders?: "none" | "";
  className?: string;
  textSize?: string;
  actions?: {
    label?: string;
    fns: { icon: JSX.Element; onClick: (item: any) => void }[];
  };
  onRowClick?: (item: any) => void;
};

const DataTable = (opts: Options) => {
  return (
    <div
      className={`overflow-x-auto rounded-lg shadow ${opts.className ?? ""}`}
    >
      <table
        className={`w-full border-gray-500 text-${
          opts.textSize ?? "sm"
        } table-auto`}
      >
        <thead className="bg-gray-100 text-left">
          <tr className="px-6 py-3 text-xs font-medium uppercase tracking-wider">
            {opts.heads[0].label &&
              opts.heads.map((head, i) => (
                <th key={i} className="px-4 py-3 border-b font-semibold">
                  {head.label}
                </th>
              ))}
            {opts.actions?.label && (
              <th className="px-4 py-3 border-b font-semibold">
                {opts.actions.label}
              </th>
            )}
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {opts.data &&
            opts.data.map((item, i) => (
              <tr
                key={i}
                className={`hover:bg-accent-50 transition duration-150 ${
                  opts.onRowClick && "cursor-pointer hover:bg-accent-500"
                }`}
              >
                {opts.heads.map((head, j) => (
                  <td
                    key={j}
                    className={`px-4 py-2 whitespace-nowrap text-gray-800 `}
                    onClick={() => {
                      opts.onRowClick && opts.onRowClick(item);
                    }}
                  >
                    {item[head.key]}
                  </td>
                ))}
                {opts.actions && (
                  <td className="px-4 py-2 whitespace-nowrap">
                    <div className="flex space-x-2 justify-center mx-auto">
                      {opts.actions.fns.map((fn, k) => (
                        <button
                          key={k}
                          onClick={() => fn.onClick(item)}
                          className="text-gray-600 hover:text-accent-600 transition cursor-pointer"
                        >
                          {fn.icon}
                        </button>
                      ))}
                    </div>
                  </td>
                )}
              </tr>
            ))}
        </tbody>
      </table>
    </div>
  );
};

export default DataTable;
