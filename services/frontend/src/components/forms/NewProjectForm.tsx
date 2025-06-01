import { useEffect, useState } from "react";
import { ProjectCreation } from "../../lib/project/project";
import { useCompanyStore } from "../../store/selectedCompanyStore";
import useMetaCache from "../../store/useMetaCache";

type NewProjectFormProps = React.HTMLAttributes<HTMLFormElement> & {
  onSubmit?: (data: ProjectCreation) => void;
  onFinish: (data: ProjectCreation) => void;
};

const NewProjectForm = ({
  onFinish,
  className,
  ...props
}: NewProjectFormProps) => {
  const metaCache = useMetaCache();
  const NULL_PROJECT: ProjectCreation = {
    title: "",
    description: "",
    code_name: "",
    company_id: metaCache.metadata.selectedCompany?.id ?? "",
  };
  const [newProject, setNewProject] = useState<ProjectCreation>(NULL_PROJECT);

  useEffect(() => {
    setNewProject({
      ...newProject,
      company_id: metaCache.metadata.selectedCompany?.id ?? "",
    });
  }, [metaCache.metadata.selectedCompany?.id]);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        onFinish(newProject);
        setNewProject(NULL_PROJECT);
      }}
      className="mx-auto"
      {...props}
    >
      <div className="relative z-0 mb-4 mt-8">
        <input
          type="text"
          value={metaCache.metadata.selectedCompany?.name}
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          disabled
        />
        <label
          htmlFor="new-project-description"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Company
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={newProject.title}
          id="new-project-title"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewProject({ ...newProject, title: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-project-title"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Title
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={newProject.code_name}
          id="new-project-codename"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewProject({ ...newProject, code_name: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-project-codename"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Codename
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <textarea
          value={newProject.description}
          id="new-project-description"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewProject({ ...newProject, description: e.currentTarget.value })
          }
        />
        <label
          htmlFor="new-project-description"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Description
        </label>
      </div>

      <div className="mx-auto w-fit mt-8">
        <input
          type="submit"
          value="Create"
          className="cursor-pointer px-4 py-2 border border-black bg-accent-500 text-black hover:bg-accent-300 active:bg-accent-600 transition-colors font-semibold rounded-md mx-auto"
        />
      </div>
    </form>
  );
};

export default NewProjectForm;
