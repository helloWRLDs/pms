import { useState } from "react";
import Input from "../ui/Input";
import { DocumentCreation } from "../../lib/document/document";
import { useCacheStore } from "../../store/cacheStore";
import { Button } from "../ui/Button";
import { useProjectStore } from "../../store/selectedProjectStore";
import { errorToast } from "../../lib/utils/toast";

type NewDocumentProps = {
  onSubmit: (newDoc: DocumentCreation) => Promise<void>;
  className?: string;
};

const NewDocumentForm = ({
  className,
  onSubmit,
  ...props
}: NewDocumentProps) => {
  const { project } = useProjectStore();
  const NULL_DOCUMENT: DocumentCreation = {
    project_id: project?.id ?? "",
    sprint_id: "",
    title: "",
  };
  const { sprints } = useCacheStore();
  const [newDocument, setNewDocument] =
    useState<DocumentCreation>(NULL_DOCUMENT);
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
          options={[
            { label: "None", value: "" },
            ...(sprints
              ? Object.values(sprints).map((sprint) => ({
                  label: sprint.title,
                  value: sprint.id,
                }))
              : []),
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
