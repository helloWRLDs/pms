import { useState, useEffect } from "react";
import { CompanyCreation } from "../../lib/company/company";
import { generateCodeName } from "../../lib/company/codename";
import { BiChevronDown } from "react-icons/bi";

type NewCompanyFormProps = React.HTMLAttributes<HTMLFormElement> & {
  onFinish: (data: CompanyCreation) => void;
};

const NewCompanyForm = ({ onFinish, ...props }: NewCompanyFormProps) => {
  const NULL_COMPANY: CompanyCreation = {
    name: "",
    codename: "",
  };
  const [newCompany, setNewCompany] = useState<CompanyCreation>(NULL_COMPANY);
  const [recommendedCodeNames, setRecommendedCodeNames] = useState<string[]>(
    []
  );
  const [showDropdown, setShowDropdown] = useState(false);

  // Generate recommendations when company name changes
  useEffect(() => {
    if (newCompany.name.trim().length > 2) {
      const recommendations = [];
      const name = newCompany.name.trim();

      // Generate multiple variations using different approaches
      recommendations.push(generateCodeName(name)); // Primary recommendation

      // Create additional variations
      const words = name.split(/\s+/);
      if (words.length > 1) {
        // Initials version
        const initials = words
          .map((word) => word.charAt(0).toUpperCase())
          .join("");
        recommendations.push(initials.toLowerCase());

        // First + last word
        if (words.length >= 2) {
          const firstLast = words[0] + words[words.length - 1];
          recommendations.push(generateCodeName(firstLast));
        }

        // Shortened version - first 3 chars of first word + first char of others
        const shortened =
          words[0].substring(0, 3) +
          words
            .slice(1)
            .map((w) => w.charAt(0))
            .join("");
        recommendations.push(shortened.toLowerCase());
      } else {
        // Single word variations
        const singleWord = words[0];
        if (singleWord.length > 4) {
          // First 4 characters
          recommendations.push(singleWord.substring(0, 4).toLowerCase());
          // First 3 + last character
          recommendations.push(
            (singleWord.substring(0, 3) + singleWord.slice(-1)).toLowerCase()
          );
        }
      }

      // Remove duplicates and filter out empty values
      const uniqueRecommendations = [...new Set(recommendations)].filter(
        (name) => name && name.length > 0
      );
      setRecommendedCodeNames(uniqueRecommendations);

      // Auto-select first recommendation if no codename is set
      if (!newCompany.codename && uniqueRecommendations.length > 0) {
        setNewCompany((prev) => ({
          ...prev,
          codename: uniqueRecommendations[0],
        }));
      }
    } else {
      setRecommendedCodeNames([]);
    }
  }, [newCompany.name]);

  const handleCodenameSelect = (codename: string) => {
    setNewCompany({ ...newCompany, codename });
    setShowDropdown(false);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (newCompany.name.trim() && newCompany.codename.trim()) {
      onFinish(newCompany);
      setNewCompany(NULL_COMPANY);
      setRecommendedCodeNames([]);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mx-auto space-y-6" {...props}>
      {/* Company Name Field */}
      <div className="relative z-0">
        <input
          type="text"
          value={newCompany.name}
          id="new-company-name"
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          required={true}
          onChange={(e) => {
            setNewCompany({ ...newCompany, name: e.target.value });
          }}
        />
        <label
          htmlFor="new-company-name"
          className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
        >
          Company Name
        </label>
      </div>

      {/* Code Name Field with Dropdown */}
      <div className="relative z-10">
        <div className="relative">
          <input
            type="text"
            value={newCompany.codename}
            id="new-company-codename"
            className="block py-2.5 px-0 pr-8 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
            placeholder=" "
            required={true}
            onChange={(e) =>
              setNewCompany({ ...newCompany, codename: e.target.value })
            }
            onFocus={() => setShowDropdown(recommendedCodeNames.length > 0)}
          />

          {/* Dropdown trigger button */}
          {recommendedCodeNames.length > 0 && (
            <button
              type="button"
              onClick={() => setShowDropdown(!showDropdown)}
              className="absolute right-0 top-3 text-gray-400 hover:text-accent-500 transition-colors"
            >
              <BiChevronDown
                className={`transition-transform ${
                  showDropdown ? "rotate-180" : ""
                }`}
              />
            </button>
          )}

          <label
            htmlFor="new-company-codename"
            className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto"
          >
            Code Name
          </label>
        </div>

        {/* Dropdown Menu */}
        {showDropdown && recommendedCodeNames.length > 0 && (
          <div className="absolute top-full left-0 right-0 mt-1 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-md shadow-lg z-50">
            <div className="py-1">
              {/* Recommended Options */}
              <div className="px-3 py-2 text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wide border-b border-gray-200 dark:border-gray-700">
                Recommended
              </div>
              {recommendedCodeNames.map((codename, index) => (
                <button
                  key={index}
                  type="button"
                  onClick={() => handleCodenameSelect(codename)}
                  className={`w-full text-left px-3 py-2 text-sm hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors ${
                    newCompany.codename === codename
                      ? "bg-accent-50 dark:bg-accent-900 text-accent-600 dark:text-accent-400"
                      : "text-gray-900 dark:text-gray-100"
                  }`}
                >
                  {codename}
                  {index === 0 && (
                    <span className="ml-2 text-xs text-accent-500 font-medium">
                      Recommended
                    </span>
                  )}
                </button>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Submit Button */}
      <div className="mx-auto w-fit mt-8">
        <input
          type="submit"
          value="Create Company"
          className="cursor-pointer px-6 py-2 border border-black bg-accent-500 text-black hover:bg-accent-300 active:bg-accent-600 transition-colors font-semibold rounded-md mx-auto"
        />
      </div>
    </form>
  );
};

export default NewCompanyForm;
