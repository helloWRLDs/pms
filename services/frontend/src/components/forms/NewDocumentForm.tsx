import { useState } from "react";
import Input from "../ui/Input";
import { DocumentCreation } from "../../lib/document/document";
import { errorToast } from "../../lib/utils/toast";
import useMetaCache from "../../store/useMetaCache";
import { useSprintList } from "../../hooks/useData";

type NewDocumentProps = {
  onSubmit: (newDoc: DocumentCreation) => Promise<void>;
  className?: string;
};

const NewDocumentForm = ({
  className,
  onSubmit,
  ...props
}: NewDocumentProps) => {
  const metaCache = useMetaCache();
  const NULL_DOCUMENT: DocumentCreation = {
    project_id: metaCache.metadata.selectedProject?.id ?? "",
    sprint_id: "",
    title: "",
  };
  const [newDocument, setNewDocument] =
    useState<DocumentCreation>(NULL_DOCUMENT);

  // Use the direct sprint list hook instead of cache collections
  const { sprints } = useSprintList(
    metaCache.metadata.selectedProject?.id ?? ""
  );

  return (
    <form
      className={className}
      {...props}
      onSubmit={(e) => {
        e.preventDefault();
        if (!newDocument.title) {
          errorToast("Define title");
          return;
        }
        if (!newDocument.project_id) {
          errorToast("Failed to resolve project");
          return;
        }
        onSubmit(newDocument);
      }}
    >
      <Input>
        <Input.Element
          type="text"
          value={newDocument.title}
          label="Title"
          onChange={(e) => {
            setNewDocument({ ...newDocument, title: e.currentTarget.value });
          }}
        />
      </Input>

      <Input>
        <Input.Element
          type="select"
          label="Sprint"
          value={newDocument.sprint_id}
          onChange={(e) => {
            setNewDocument({
              ...newDocument,
              sprint_id: e.currentTarget.value,
            });
          }}
          options={[
            { label: "Unassigned", value: "" },
            ...(sprints?.items.map((sprint) => ({
              label: sprint.title ?? `Sprint ${sprint.id}`,
              value: sprint.id,
            })) ?? []),
          ]}
        />
      </Input>

      <input
        type="submit"
        value="Create"
        className="px-4 py-2 cursor-pointer rounded-md outline-1 outline-accent-500 bg-secondary-100 text-accent-500 hover:text-secondary-100  hover:bg-accent-500"
      />
    </form>
  );
};

export default NewDocumentForm;
