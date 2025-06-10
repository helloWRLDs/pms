import { useState } from "react";
import { Role } from "../../lib/roles";
import { Button } from "../ui/Button";
import { Permission, Permissions } from "../../lib/permission";
import useMetaCache from "../../store/useMetaCache";

const RoleForm = ({
  initialRole,
  onSubmit,
  onCancel,
  isEditing = false,
}: {
  initialRole?: Role;
  onSubmit: (role: Role) => void;
  onCancel: () => void;
  isEditing?: boolean;
}) => {
  const [roleName, setRoleName] = useState(initialRole?.name || "");
  const [selectedPermissions, setSelectedPermissions] = useState<Permission[]>(
    initialRole?.permissions || []
  );

  const availablePermissions = Object.values(Permissions);
  const metaCache = useMetaCache();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      name: roleName,
      permissions: selectedPermissions,
      company_id: metaCache.metadata.selectedCompany?.id,
      created_at: initialRole?.created_at || {
        seconds: Math.floor(Date.now() / 1000),
        nanos: 0,
      },
    });
  };

  const togglePermission = (permission: Permission) => {
    setSelectedPermissions((prev) =>
      prev.includes(permission)
        ? prev.filter((p) => p !== permission)
        : [...prev, permission]
    );
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div>
        <label className="block text-sm font-medium mb-2">Role Name</label>
        <input
          type="text"
          value={roleName}
          onChange={(e) => setRoleName(e.target.value)}
          placeholder="Enter role name"
          className="w-full px-3 py-2 bg-secondary-200 border border-secondary-100 rounded-md text-white placeholder-neutral-400"
          required
          disabled={isEditing}
        />
      </div>

      <div>
        <label className="block text-sm font-medium mb-4">Permissions</label>
        <div className="grid grid-cols-2 gap-3 max-h-60 overflow-y-auto">
          {availablePermissions.map((permission) => (
            <label
              key={permission}
              className="flex items-center space-x-3 cursor-pointer"
            >
              <input
                type="checkbox"
                checked={selectedPermissions.includes(permission)}
                onChange={() => togglePermission(permission)}
                className="w-4 h-4 text-accent-500 bg-secondary-200 border-secondary-100 rounded focus:ring-accent-500"
              />
              <span className="text-sm">
                {permission.replace(/[_:]/g, " ").toLowerCase()}
              </span>
            </label>
          ))}
        </div>
      </div>

      <div className="flex gap-3 pt-4">
        <Button
          type="submit"
          className="bg-accent-500 hover:bg-accent-600 text-white px-6 py-2 rounded-md"
        >
          {isEditing ? "Update Role" : "Create Role"}
        </Button>
        <Button
          type="button"
          onClick={onCancel}
          className="bg-secondary-300 hover:bg-secondary-400 text-white px-6 py-2 rounded-md"
        >
          Cancel
        </Button>
      </div>
    </form>
  );
};

export default RoleForm;
