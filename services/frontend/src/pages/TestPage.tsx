import { useState } from "react";
import DropDownList from "../components/ui/DropDown";

const TestPage = () => {
  const [dd, setDD] = useState(false);
  return (
    <div className="px-8 py-5 bg-primary-500">
      <p className="text-white border border-white py-2 relative">test</p>
      <button
        className="mt-16 px-4 py-2 border border-black rounded-md text-white bg-orange-500"
        onClick={() => {
          setDD(!dd);
        }}
      >
        Open
      </button>

      <p className="mt-7">dsadasfgad</p>
      <DropDownList visible={dd} className="rounded-md">
        <DropDownList.Element className="px-3 py-2 bg-white text-black hover:bg-gray-600">
          bob 21
        </DropDownList.Element>
        <DropDownList.Element className="px-3 py-2 bg-white text-black hover:bg-gray-600">
          john 32
        </DropDownList.Element>
      </DropDownList>
    </div>
  );
};

export default TestPage;
