import { useState } from "react";
import Paginator from "../components/ui/Paginator";

const TestPage = () => {
  const [page, setPage] = useState(1);
  const perPage = 10;

  return (
    <div className="px-8 py-5 bg-primary-500">
      <Paginator
        page={page}
        per_page={perPage}
        total_pages={10}
        total_items={100}
        onPageChange={setPage}
      />
    </div>
  );
};

export default TestPage;
