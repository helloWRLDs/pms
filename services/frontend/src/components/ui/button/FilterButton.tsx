const FilterButton = ({
  // label,
  value,
  options,
  onChange,
}: {
  label: string;
  value: string | undefined;
  options: { label: string; value: string }[];
  onChange: (value: string) => void;
}) => (
  <select
    value={value ?? ""}
    onChange={(e) => onChange(e.target.value)}
    className="px-4 py-2 rounded-lg text-neutral-100 bg-secondary-200 border border-secondary-100 
                 hover:border-accent-500 transition-colors cursor-pointer focus:outline-none focus:border-accent-500"
  >
    {options.map((option) => (
      <option key={option.value} value={option.value}>
        {option.label}
      </option>
    ))}
  </select>
);

export default FilterButton;
