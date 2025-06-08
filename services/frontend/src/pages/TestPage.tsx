import { SlOptions } from "react-icons/sl";
import { ContextMenu } from "../components/ui/ContextMenu";
import { BsPencil, BsTrash } from "react-icons/bs";
import useMetaCache from "../store/useMetaCache";

const TestPage = () => {
  const metaCache = useMetaCache();
  return (
    <div className="px-8 py-5 bg-primary-500">
      <ContextMenu
        placement="right"
        trigger={
          <button className="text-accent-500">
            <SlOptions />
          </button>
        }
        items={[
          {
            label: "Edit",
            onClick: () => console.log("Edit"),
            icon: <BsPencil />,
          },
          {
            label: "Delete",
            onClick: () => console.log("Delete"),
            icon: <BsTrash />,
          },
        ]}
      />
      <div className="flex flex-row">
        <div className="bg-secondary-100 text-accent-500 px-6 py-2 rounded-r-3xl z-10">
          <a href="/companies">Companies</a>
        </div>
        <div className="bg-accent-500 px-6 py-2 rounded-r-3xl z-9 -ml-4">
          {metaCache.metadata.selectedCompany?.name}
        </div>
        <div className="bg-secondary-100 text-accent-500 px-6 py-2 rounded-r-3xl z-8 -ml-4">
          Projects
        </div>
        <div className="bg-accent-500 px-6 py-2 rounded-r-3xl z-7 -ml-4">
          {metaCache.metadata.selectedProject?.title}
        </div>
        <div className="bg-accent-100 px-6 py-2 rounded-r-3xl z-6 -ml-4">
          Sprints
        </div>
      </div>
    </div>
  );
};

export default TestPage;
