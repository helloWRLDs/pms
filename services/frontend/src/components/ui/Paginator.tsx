type Options = {
  page: number;
  per_page: number;
  total_pages: number;
  total_items: number;
  onPageChange?: (page: number, per_page: number) => void;
};

const Paginator = (opts: Options) => {
  type PanelItem = {
    value: string;
    page?: number;
  };
  const panel: PanelItem[] = [];
  const start = Math.max(2, opts.page - 2);
  const end = Math.min(opts.total_pages - 1, opts.page + 2);

  if (opts.page > 1) {
    panel.push({ value: "<", page: opts.page - 1 });
  }
  panel.push({ value: "1", page: 1 });

  if (start > 2) {
    panel.push({ value: "..." });
  }

  for (let i = start; i <= end; i++) {
    panel.push({ value: `${i}`, page: i });
  }

  if (end < opts.total_pages - 1) {
    panel.push({ value: "..." });
  }

  if (opts.total_pages > 1) {
    panel.push({ value: `${opts.total_pages}`, page: opts.total_pages });
  }
  if (opts.page < opts.total_pages) {
    panel.push({ value: ">", page: opts.page + 1 });
  }
  return (
    <div className="flex justify-center gap-1 border-t border-gray-200 shadow-xl px-2 py-1 rounded-md">
      {panel.map((item, i) => (
        <div
          key={i}
          className={`px-2 py-1 text-center m-1 cursor-pointer rounded-md select-none   ${
            item.page === opts.page ? "bg-accent-500" : ""
          } ${item.page && "hover:bg-accent-500"}`}
          onClick={() =>
            item.page &&
            opts.onPageChange &&
            opts.onPageChange(item.page, opts.per_page)
          }
        >
          {item.value}
        </div>
      ))}
    </div>
  );
};

export default Paginator;
