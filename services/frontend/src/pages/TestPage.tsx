import { useState } from "react";
import { UserOptional } from "../lib/user/user";
import Paginator from "../components/ui/Paginator";
import { FaAddressBook } from "react-icons/fa";
import { ListItems } from "../lib/utils/list";

interface Options {
  heads: { label?: string; key: string }[];
  data: any[];
  borders?: "none" | "";
  className?: string;
  textSize?: string;
  actions?: {
    label?: string;
    fns: { icon: JSX.Element; onClick: (item: any) => void }[];
  };
  onRowClick?: (item: any) => void;
}

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
          <tr>
            {opts.heads[0].label &&
              opts.heads.map((head, i) => (
                <th
                  key={i}
                  className="px-4 py-3 border-b font-semibold text-gray-700"
                >
                  {head.label}
                </th>
              ))}
            {opts.actions?.label && (
              <th className="px-4 py-3 border-b font-semibold text-gray-700">
                {opts.actions.label}
              </th>
            )}
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {opts.data.map((item, i) => (
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

const userList: ListItems<UserOptional> = {
  total_items: 10,
  total_pages: 100,
  page: 1,
  per_page: 10,
  items: [
    {
      id: "1",
      name: "Bob",
    },
    {
      id: "2",
      name: "Alice",
    },
    {
      id: "3",
      name: "John",
    },
  ],
};

const TestPage = () => {
  const [list, setList] = useState<ListItems<UserOptional>>(userList);

  return (
    <>
      <div className="w-1/2 mx-auto">
        <DataTable
          data={list.items}
          heads={[
            { label: "ID", key: "id" },
            { label: "Name", key: "name" },
          ]}
          actions={{
            label: "Actions",
            fns: [
              {
                icon: <FaAddressBook size={20} />,
                onClick: (item) => console.log(item),
              },
            ],
          }}
          onRowClick={(item) => console.log(`clicked at ${item.id}`)}
          className="text-3xl mb-3"
        />
      </div>

      <div className="mx-auto w-fit">
        <Paginator
          page={list.page}
          per_page={list.per_page}
          total_items={list.total_items}
          total_pages={list.total_pages}
          onPageChange={(page, per_page) => {
            setList({ ...userList, page: page, per_page: per_page });
            console.log(`fetching data: page=${page} & per_page=${per_page}`);
          }}
        />
      </div>
    </>
  );
};

export default TestPage;
