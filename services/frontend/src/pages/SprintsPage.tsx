import { useEffect, useState } from "react";
import { usePageSettings } from "../hooks/usePageSettings";
import { Modal } from "../components/ui/Modal";
import NewSprintForm from "../components/forms/NewSprintForm";
import sprintAPI from "../api/sprintAPI";
import { useQuery } from "@tanstack/react-query";
import { SprintFilter } from "../lib/sprint/sprint";
import { useProjectStore } from "../store/selectedProjectStore";
import Table from "../components/ui/Table";
import { formatTime } from "../lib/utils/time";
import Paginator from "../components/ui/Paginator";
import { useSprintStore } from "../store/selectedSprintStore";
import { useNavigate } from "react-router-dom";
import { useCacheLoader } from "../hooks/useCacheLoader";
import { useCacheStore } from "../store/cacheStore";
import Input from "../components/ui/Input";
import { BsFillPlusCircleFill } from "react-icons/bs";

const SprintsPage = () => {
  usePageSettings({ title: "Sprints", requireAuth: true });

  const navigate = useNavigate();
  const { selectSprint } = useSprintStore();
  const { project: selectedProject } = useProjectStore();
  const [newSprintModal, setNewSprintModal] = useState(false);
  const [filter, setFilter] = useState<SprintFilter>({
    page: 1,
    per_page: 10,
    project_id: selectedProject?.id,
    title: "",
  });
  const {
    data: sprintList,
    isLoading: isSprintListLoading,
    refetch: refetchSprints,
  } = useQuery({
    queryKey: [
      "sprints",
      filter.page,
      filter.per_page,
      filter.project_id,
      filter.title,
    ],
    queryFn: () => sprintAPI.list(filter),
  });

  useCacheLoader({ sprintList: sprintList });

  useEffect(() => {
    console.log(sprintList);
  }, [sprintList]);

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-primary-600 text-neutral-100">
      <section id="sprint-modals">
        <Modal
          title="Create Sprint"
          visible={newSprintModal}
          onClose={() => setNewSprintModal(false)}
          className="bg-secondary-100"
        >
          <NewSprintForm
            onFinish={async (creation) => {
              console.log(creation);
              await sprintAPI.create(creation);
              setNewSprintModal(false);
              refetchSprints();
            }}
          />
        </Modal>
      </section>

      <section id="companies-header">
        <div className="container mx-auto flex justify-between items-center mb-5">
          <h2 className="font-bold text-3xl">
            <span className="text-accent-500">{selectedProject?.title}</span>{" "}
            Sprints
          </h2>
          <div className="flex gap-4 items-baseline">
            <Input>
              <Input.Element
                type="text"
                label="Title"
                value={filter.title}
                onInput={(e) => {
                  setFilter({ ...filter, title: e.currentTarget.value });
                }}
              />
            </Input>
          </div>
        </div>
      </section>

      <section id="sprint-table">
        <div className="overflow-x-auto container mx-auto shadow-xl">
          <div>
            <div className="h-[75lvh] w-full">
              <Table className="rounded-lg">
                <Table.Head>
                  <Table.Row className="text-neutral-100 bg-primary-400">
                    <Table.HeadCell>№</Table.HeadCell>
                    <Table.HeadCell>Title</Table.HeadCell>
                    <Table.HeadCell>Description</Table.HeadCell>
                    <Table.HeadCell>Period</Table.HeadCell>
                    <Table.HeadCell>Created</Table.HeadCell>
                    <Table.HeadCell>Tasks</Table.HeadCell>
                  </Table.Row>
                </Table.Head>

                {isSprintListLoading ? (
                  <p>Sprints loading...</p>
                ) : !sprintList ||
                  !sprintList.items ||
                  sprintList.items.length === 0 ? (
                  <Table.Body></Table.Body>
                ) : (
                  <Table.Body>
                    {sprintList.items.map((sprint, i) => (
                      <Table.Row
                        key={i}
                        onClick={() => {
                          selectSprint(sprint);
                          navigate("/agile-dashboard");
                        }}
                        className="cursor-pointer bg-secondary-200 text-neutral-100 hover:bg-secondary-100 py-10"
                      >
                        <Table.Cell>{i + 1}</Table.Cell>
                        <Table.Cell>{sprint.title}</Table.Cell>
                        <Table.Cell>{sprint.description}</Table.Cell>
                        <Table.Cell>
                          {formatTime(sprint.start_date.seconds)} -{" "}
                          {formatTime(sprint.end_date.seconds)}
                        </Table.Cell>
                        <Table.Cell>
                          {formatTime(sprint.created_at.seconds)}
                        </Table.Cell>
                        <Table.Cell>{sprint.tasks?.length ?? 0}</Table.Cell>
                      </Table.Row>
                    ))}
                  </Table.Body>
                )}
              </Table>
              <button
                className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
                onClick={() => {
                  setNewSprintModal(true);
                }}
              >
                <BsFillPlusCircleFill
                  size="30"
                  className="mx-auto text-neutral-300 group-hover:text-accent-300 "
                />
              </button>
            </div>
            {sprintList && sprintList.items && (
              <Paginator
                page={sprintList.page}
                per_page={sprintList.per_page}
                total_items={sprintList.total_items}
                total_pages={sprintList.total_pages}
                onPageChange={(page) => {
                  setFilter({ ...filter, page: page });
                }}
              />
            )}
          </div>
        </div>
      </section>
    </div>
  );
};
export default SprintsPage;
