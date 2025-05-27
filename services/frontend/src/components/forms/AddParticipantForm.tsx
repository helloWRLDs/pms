import { useState } from "react";
import { UserFilter } from "../../lib/user/user";
import authAPI from "../../api/auth";
import Input from "../ui/Input";
import { Button } from "../ui/Button";

type AddParticipantFormProps = React.HTMLAttributes<HTMLDivElement> & {
  onFinish: (userID: string) => void;
};

const AddParticipantForm = ({
  className,
  onFinish,
  ...props
}: AddParticipantFormProps) => {
  const [filter, setFilter] = useState<UserFilter>({
    page: 1,
    per_page: 1,
    company_id: "",
  });

  return (
    <div className={`${className}`} {...props}>
      <div className="flex flex-row gap-2 items-center w-full">
        <Input className="w-full">
          <Input.Element
            label="Enter email"
            type="text"
            value={filter.email}
            onChange={(e) => {
              setFilter({ ...filter, email: e.currentTarget.value });
            }}
          />
        </Input>
        <Button
          onClick={async () => {
            const res = await authAPI.listUsers(filter);
            if (res && res.items && res.total_items !== 0) {
              onFinish(res.items[0].id);
            }
          }}
        >
          Add
        </Button>
      </div>
    </div>
  );
};

export default AddParticipantForm;
