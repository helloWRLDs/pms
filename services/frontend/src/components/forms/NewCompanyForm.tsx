import { useState } from "react";
import { CompanyCreation } from "../../lib/company/company";

type NewCompanyFormProps = React.HTMLAttributes<HTMLFormElement> & {
  onFinish: (data: CompanyCreation) => void;
};

const NewCompanyForm = ({ onFinish, ...props }: NewCompanyFormProps) => {
  const NULL_COMPANY: CompanyCreation = {
    name: "",
    codename: "",
  };
  const [newCompany, setNewCompany] = useState<CompanyCreation>(NULL_COMPANY);
  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        onFinish(newCompany);
        setNewCompany(NULL_COMPANY);
      }}
      className="mx-auto"
      {...props}
    >
      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={newCompany.name}
          id="new-company-name"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewCompany({ ...newCompany, name: e.target.value })
          }
        />
        <label
          htmlFor="new-company-name"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Company Name
        </label>
      </div>

      <div className="relative z-0 mb-4">
        <input
          type="text"
          value={newCompany.codename}
          id="new-company-codename"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) =>
            setNewCompany({ ...newCompany, codename: e.target.value })
          }
        />
        <label
          htmlFor="new-company-codename"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Code Name
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

export default NewCompanyForm;
