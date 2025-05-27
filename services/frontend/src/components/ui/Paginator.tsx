type PaginationProps = {
  page: number;
  per_page: number;
  total_pages: number;
  total_items: number;
  onPageChange?: (page: number, per_page: number) => void;
  className?: string;
};

const Paginator = (props: PaginationProps) => {
  type PanelItem = {
    value: string;
    page?: number;
  };
  const panel: PanelItem[] = [];
  const start = Math.max(2, props.page - 2);
  const end = Math.min(props.total_pages - 1, props.page + 2);

  if (props.page > 1) {
    panel.push({ value: "<", page: props.page - 1 });
  }
  panel.push({ value: "1", page: 1 });

  if (start > 2) {
    panel.push({ value: "..." });
  }

  for (let i = start; i <= end; i++) {
    panel.push({ value: `${i}`, page: i });
  }

  if (end < props.total_pages - 1) {
    panel.push({ value: "..." });
  }

  if (props.total_pages > 1) {
    panel.push({ value: `${props.total_pages}`, page: props.total_pages });
  }
  if (props.page < props.total_pages) {
    panel.push({ value: ">", page: props.page + 1 });
  }
  return (
    <div
      className={`flex justify-center gap-1 border-t border-secondary-200 shadow-xl px-2 py-1 rounded-md ${props.className}`}
    >
      {panel.map((item, i) => (
        <div
          key={i}
          className={`px-2 py-1 text-center m-1 cursor-pointer rounded-md select-none   ${
            item.page === props.page ? "bg-accent-500 text-primary-500" : ""
          } ${item.page && "hover:bg-accent-500"}`}
          onClick={() =>
            item.page &&
            props.onPageChange &&
            props.onPageChange(item.page, props.per_page)
          }
        >
          {item.value}
        </div>
      ))}
    </div>
  );
};

export default Paginator;
