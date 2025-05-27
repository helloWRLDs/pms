import { useState } from "react";
import { SprintCreation } from "../../lib/sprint/sprint";
import { useProjectStore } from "../../store/selectedProjectStore";
import Input from "../ui/Input";
import { errorToast, infoToast } from "../../lib/utils/toast";

type NewSprintFormProps = {
  className?: string;
  onFinish: (data: SprintCreation) => void;
};

const NewSprintForm = ({
  onFinish,
  className,
  ...props
}: NewSprintFormProps) => {
  const { project: selectedProject } = useProjectStore();

  const NULL_SPRINT: SprintCreation = {
    title: "",
    description: "",
    start_date: {
      seconds: new Date().getTime() / 1000,
    },
    end_date: {
      seconds: new Date().getTime() / 1000,
    },
    project_id: selectedProject?.id ?? "",
    tasks: [],
  };

  const validateSprint = (creation: SprintCreation): boolean => {
    if (creation.title.trim().length == 0) {
      infoToast("Please define sprint title");
      return false;
    }
    if (creation.description.trim().length === 0) {
      errorToast("Please define sprint description");
      return false;
    }
    return true;
  };

  const [newSprint, setNewSprint] = useState<SprintCreation>(NULL_SPRINT);
  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        if (!validateSprint(newSprint)) {
          return;
        }
        const fixedSprint = {
          ...newSprint,
          start_date: {
            seconds: Math.floor(newSprint.start_date.seconds),
          },
          end_date: {
            seconds: Math.floor(newSprint.end_date.seconds),
          },
        };
        onFinish(fixedSprint);
        setNewSprint(NULL_SPRINT);
      }}
      className="text-black text-lg"
      {...props}
    >
      <Input>
        <Input.Element
          type="text"
          required
          value={newSprint.title}
          label="Title"
          onChange={(e) => {
            setNewSprint({ ...newSprint, title: e.currentTarget.value });
          }}
        />
      </Input>

      <Input>
        <Input.Element
          label="Description"
          type="textarea"
          value={newSprint.description}
          required
          onChange={(e) => {
            setNewSprint({ ...newSprint, description: e.currentTarget.value });
          }}
        />
      </Input>

      <Input>
        <Input.Element
          type="date"
          label="Start Date"
          value={
            new Date(newSprint.start_date.seconds * 1000)
              .toISOString()
              .split("T")[0]
          }
          onChange={(e) => {
            setNewSprint({
              ...newSprint,
              start_date: {
                seconds: Math.floor(
                  new Date(e.currentTarget.value).getTime() / 1000
                ),
              },
            });
          }}
        />
      </Input>

      <Input>
        <Input.Element
          type="date"
          label="End Date"
          value={
            new Date(newSprint.end_date.seconds * 1000)
              .toISOString()
              .split("T")[0]
          }
          onChange={(e) => {
            setNewSprint({
              ...newSprint,
              end_date: {
                seconds: Math.floor(
                  new Date(e.currentTarget.value).getTime() / 1000
                ),
              },
            });
          }}
        />
      </Input>

      <div className="mx-auto w-fit">
        <input
          type="submit"
          value="Create"
          className="px-4 py-2 bg-primary-500 text-accent-400 rounded-md cursor-pointer hover:bg-accent-500 hover:text-primary-500"
        />
      </div>
    </form>
  );
};

export default NewSprintForm;
