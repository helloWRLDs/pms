import { FiChevronLeft, FiChevronRight } from "react-icons/fi";

type PaginationProps = {
  page: number;
  per_page: number;
  total_pages: number;
  total_items: number;
  onPageChange?: (page: number, per_page: number) => void;
  className?: string;
};

const PageButton = ({
  children,
  isActive,
  onClick,
  disabled,
  ariaLabel,
}: {
  children: React.ReactNode;
  isActive?: boolean;
  onClick?: () => void;
  disabled?: boolean;
  ariaLabel?: string;
}) => (
  <button
    onClick={onClick}
    disabled={disabled}
    aria-label={ariaLabel}
    aria-current={isActive ? "page" : undefined}
    className={`
      relative inline-flex items-center justify-center min-w-[2.5rem] h-10 px-3
      text-sm font-medium rounded-md transition-all duration-200
      focus:outline-none focus:ring-2 focus:ring-accent-500 focus:ring-offset-2 focus:ring-offset-primary-600
      disabled:cursor-not-allowed disabled:opacity-50
      ${
        isActive
          ? "bg-accent-500 text-primary-700 hover:bg-accent-400"
          : "text-neutral-300 hover:bg-secondary-100 hover:text-accent-400"
      }
    `}
  >
    {children}
  </button>
);

const Ellipsis = () => (
  <span className="px-2 py-1 text-neutral-400">
    <span className="tracking-wider">•••</span>
  </span>
);

const Paginator = ({
  page,
  per_page,
  total_pages,
  total_items,
  onPageChange,
  className = "",
}: PaginationProps) => {
  // Don't render pagination if there's only one page
  if (total_pages <= 1) return null;

  const handlePageChange = (newPage: number) => {
    if (onPageChange && newPage >= 1 && newPage <= total_pages) {
      onPageChange(newPage, per_page);
    }
  };

  // Calculate visible page numbers
  const getVisiblePages = () => {
    const delta = 2; // Number of pages to show on each side of current page
    const range: (number | "ellipsis")[] = [];

    for (let i = 1; i <= total_pages; i++) {
      if (
        i === 1 || // First page
        i === total_pages || // Last page
        (i >= page - delta && i <= page + delta) // Pages around current page
      ) {
        range.push(i);
      } else if (
        (i === 2 && page - delta > 2) || // First ellipsis
        (i === total_pages - 1 && page + delta < total_pages - 1) // Last ellipsis
      ) {
        range.push("ellipsis");
      }
    }

    return range;
  };

  const visiblePages = getVisiblePages();

  return (
    <nav
      role="navigation"
      aria-label="Pagination Navigation"
      className={`flex flex-col items-center gap-4 ${className}`}
    >
      <div className="flex items-center gap-2 bg-secondary-200/50 rounded-lg p-2">
        {/* Previous Button */}
        <PageButton
          onClick={() => handlePageChange(page - 1)}
          disabled={page === 1}
          ariaLabel="Go to previous page"
        >
          <FiChevronLeft className="w-5 h-5" />
        </PageButton>

        {/* Page Numbers */}
        <div className="flex items-center gap-1">
          {visiblePages.map((pageNum, index) =>
            pageNum === "ellipsis" ? (
              <Ellipsis key={`ellipsis-${index}`} />
            ) : (
              <PageButton
                key={pageNum}
                isActive={pageNum === page}
                onClick={() => handlePageChange(pageNum)}
                ariaLabel={`Go to page ${pageNum}`}
              >
                {pageNum}
              </PageButton>
            )
          )}
        </div>

        {/* Next Button */}
        <PageButton
          onClick={() => handlePageChange(page + 1)}
          disabled={page === total_pages}
          ariaLabel="Go to next page"
        >
          <FiChevronRight className="w-5 h-5" />
        </PageButton>
      </div>

      {/* Page Info */}
      <div className="text-sm text-neutral-400">
        Showing page {page} of {total_pages} ({total_items} items)
      </div>
    </nav>
  );
};

export default Paginator;
