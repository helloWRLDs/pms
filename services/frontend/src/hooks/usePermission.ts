import { Permission } from "../lib/permission";
import { useAuthStore } from "../store/authStore";
import useMetaCache from "../store/useMetaCache";

export const usePermission = () => {
  const { auth } = useAuthStore();
  const metaCache = useMetaCache();

  const hasPermission = (requiredPermission: Permission): boolean => {
    if (!auth?.user?.permissions) return false;
    if (!metaCache.metadata.selectedCompany?.id) return false;

    const companyId = metaCache.metadata.selectedCompany.id;
    const companyPermissions = auth.user.permissions[companyId];

    if (!companyPermissions) return false;

    let permissionsArray: string[];

    if (Array.isArray(companyPermissions)) {
      permissionsArray = companyPermissions.filter(
        (p): p is string => typeof p === "string"
      );
    } else if (
      typeof companyPermissions === "object" &&
      "values" in companyPermissions &&
      Array.isArray(companyPermissions.values)
    ) {
      permissionsArray = companyPermissions.values.filter(
        (p): p is string => typeof p === "string"
      );
    } else {
      return false;
    }

    return permissionsArray.includes(requiredPermission);
  };

  return { hasPermission };
};
