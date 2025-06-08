import DropDownable from "../components/ui/DropDownable";
import { usePermission } from "../hooks/usePermission";
import { Permissions } from "../lib/permission";

const TestPage1 = () => {
  const { hasPermission } = usePermission();
  console.log(
    "Has user:read permission:",
    hasPermission(Permissions.USER_READ_PERMISSION)
  );
  console.log(
    "Has project:write permission:",
    hasPermission(Permissions.PROJECT_WRITE_PERMISSION)
  );

  return (
    <div>
      <DropDownable
        options={[
          {
            label: "test1",
            isActive: true,
            onClick: () => {
              console.log("test1");
            },
          },
          {
            label: "test2",
            isActive: false,
            onClick: () => {
              console.log("test2");
            },
          },
        ]}
      >
        <span>test</span>
      </DropDownable>
    </div>
  );
};

export default TestPage1;
