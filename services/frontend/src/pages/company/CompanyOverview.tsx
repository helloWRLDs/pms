import { useQuery } from "@tanstack/react-query";
import companyAPI from "../../api/company";
import { useEffect } from "react";
import { capitalize } from "../../lib/utils/string";
import { useNavigate } from "react-router-dom";
import { usePageSettings } from "../../hooks/usePageSettings";
import { Layouts } from "../../lib/layout/layout";
import useMetaCache from "../../store/useMetaCache";
import { usePermission } from "../../hooks/usePermission";
import { Permissions } from "../../lib/permission";

// Section Components
import { ProjectsSection, RolesSection, PeopleSection } from "./sections";

const CompanyOverviewPage = () => {
  usePageSettings({
    title: "Dashboard",
    requireAuth: true,
    layout: Layouts.Companies,
  });

  const { hasPermission } = usePermission();
  const metaCache = useMetaCache();
  const navigate = useNavigate();

  const { data: company } = useQuery({
    queryKey: ["company", metaCache.metadata.selectedCompany?.id],
    queryFn: () => companyAPI.get(metaCache.metadata.selectedCompany?.id ?? ""),
    enabled: !!metaCache.metadata.selectedCompany?.id,
  });

  useEffect(() => {
    if (!metaCache.metadata.selectedCompany?.id) {
      navigate("/companies");
    }
  }, [metaCache.metadata.selectedCompany?.id, navigate]);

  if (!company?.id) {
    return (
      <div className="min-h-[100lvh] px-5 py-10 bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100 flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-semibold mb-4">Loading Company...</h2>
          <p className="text-neutral-200">
            Please wait while we load your company data.
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-[100lvh] px-5 py-10 bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100">
      {/* Header Section */}
      <section>
        <div className="container mx-auto mb-6 flex items-center justify-between">
          <h1 className="text-3xl font-bold">
            <span className="text-accent-500">
              {capitalize(company?.name ?? "")}
            </span>{" "}
            Company Dashboard
          </h1>
        </div>
      </section>

      {hasPermission(Permissions.PROJECT_READ_PERMISSION) && (
        <ProjectsSection companyId={company.id} />
      )}

      {hasPermission(Permissions.ROLE_READ_PERMISSION) && (
        <RolesSection companyId={company.id} />
      )}

      {hasPermission(Permissions.COMPANY_READ_PERMISSION) && (
        <PeopleSection companyId={company.id} />
      )}

      {!hasPermission(Permissions.PROJECT_READ_PERMISSION) &&
        !hasPermission(Permissions.ROLE_READ_PERMISSION) &&
        !hasPermission(Permissions.COMPANY_READ_PERMISSION) && (
          <section className="pb-10">
            <div className="container mx-auto text-center py-20">
              <div className="bg-primary-500/50 backdrop-blur-sm rounded-lg border border-primary-400/30 p-12">
                <h2 className="text-2xl font-semibold mb-4 text-white/80">
                  Access Restricted
                </h2>
                <p className="text-neutral-200">
                  You don't have permission to view any sections of this company
                  dashboard. Please contact your administrator for access.
                </p>
              </div>
            </div>
          </section>
        )}
    </div>
  );
};

export default CompanyOverviewPage;
