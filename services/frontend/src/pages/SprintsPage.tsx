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
    refetch,
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
    <div>
      <section id="sprint-modals">
        <Modal
          title="Create Sprint"
          visible={newSprintModal}
          onClose={() => setNewSprintModal(false)}
          className="bg-white"
        >
          <NewSprintForm
            onSubmit={(creation) => {
              sprintAPI.create(creation);
              setNewSprintModal(false);
              refetch();
            }}
          />
        </Modal>
      </section>

      <section id="sprint-options">
        <button
          className="px-4 py-2 bg-primary-400 text-accent-500 cursor-pointer rounded-md hover:bg-accent-500 hover:text-primary-500"
          onClick={() => {
            setNewSprintModal(true);
          }}
        >
          Create Sprint
        </button>
      </section>

      <section id="sprint-table">
        {isSprintListLoading ? (
          <p>Sprints loading...</p>
        ) : !sprintList ||
          !sprintList.items ||
          sprintList.items.length === 0 ? (
          <p>No sprints found</p>
        ) : (
          <div>
            <Table>
              <Table.Head>
                <Table.Row>
                  <Table.HeadCell>№</Table.HeadCell>
                  <Table.HeadCell>Title</Table.HeadCell>
                  <Table.HeadCell>Description</Table.HeadCell>
                  <Table.HeadCell>Period</Table.HeadCell>
                  <Table.HeadCell>Created</Table.HeadCell>
                  <Table.HeadCell>Tasks</Table.HeadCell>
                </Table.Row>
              </Table.Head>

              <Table.Body>
                {sprintList.items.map((sprint, i) => (
                  <Table.Row
                    key={i}
                    onClick={() => {
                      selectSprint(sprint);
                      navigate("/agile-dashboard");
                    }}
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
            </Table>
            <Paginator
              page={sprintList.page}
              per_page={sprintList.per_page}
              total_items={sprintList.total_items}
              total_pages={sprintList.total_pages}
              onPageChange={(page) => {
                setFilter({ ...filter, page: page });
              }}
            />
          </div>
        )}
      </section>
    </div>
  );
};
export default SprintsPage;
