type CommonProps = {
  label: string;
  className?: string;
};

type InputElementProps = React.InputHTMLAttributes<HTMLInputElement> &
  CommonProps & {
    type: "text" | "password" | "date" | "email" | "number" | "password";
  };

type TextareaElementProps = React.TextareaHTMLAttributes<HTMLTextAreaElement> &
  CommonProps & { type: "textarea" };

type SelectElementProps = React.SelectHTMLAttributes<HTMLSelectElement> &
  CommonProps & {
    type: "select";
    options: { label: string; value: string | number }[];
  };

type InputElementsProps =
  | InputElementProps
  | TextareaElementProps
  | SelectElementProps;

const InputElement = ({
  type,
  label,
  className,
  ...props
}: InputElementsProps) => {
  if (type === "textarea") {
    const textareaProps =
      props as React.TextareaHTMLAttributes<HTMLTextAreaElement>;
    return (
      <>
        <textarea
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          placeholder=" "
          {...textareaProps}
        />
        <label className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto">
          {label}
        </label>
      </>
    );
  }

  if (type === "select") {
    const { options, ...selectProps } = props as SelectElementProps;
    return (
      <>
        <select
          className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
          {...selectProps}
        >
          {/* <option value="" disabled hidden></option> */}
          {options.map((opt) => (
            <option
              key={opt.value}
              value={opt.value}
              className="bg-secondary-200 dark:text-white text-black cursor-pointer hover:bg-secondary-100 "
            >
              {opt.label}
            </option>
          ))}
        </select>
        <label className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto">
          {label}
        </label>
      </>
    );
  }

  const inputProps = props as React.InputHTMLAttributes<HTMLInputElement>;
  return (
    <>
      <input
        type={type}
        className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-accent-500 focus:outline-none focus:ring-0 focus:border-accent-600 peer"
        placeholder=" "
        {...inputProps}
      />
      <label className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:start-0 peer-focus:text-accent-600 peer-focus:dark:text-accent-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6 rtl:peer-focus:translate-x-1/4 rtl:peer-focus:left-auto">
        {label}
      </label>
    </>
  );
};

type InputProps = React.HTMLAttributes<HTMLDivElement> & {};

const Input = ({ ...props }: InputProps) => {
  return <div className="relative z-0 mb-4 mt-8">{props.children}</div>;
};

Input.Element = InputElement;

export default Input;
