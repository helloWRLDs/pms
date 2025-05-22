import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { DocumentCreation, DocumentFilter } from "../lib/document/document";
import documentAPI from "../api/documentAPI";
import { Button } from "../components/ui/Button";
import { Modal } from "../components/ui/Modal";
import NewDocumentForm from "../components/forms/NewDocumentForm";
import Table from "../components/ui/Table";
import { useCacheStore } from "../store/cacheStore";
import { useNavigate } from "react-router-dom";
import Paginator from "../components/ui/Paginator";
import { SlOptionsVertical } from "react-icons/sl";
import { Menu, MenuButton, MenuItem, MenuItems } from "@headlessui/react";
import type { IconType } from "react-icons";
import { BsFillPlusCircleFill } from "react-icons/bs";

const DocumentsPage = () => {
  const [filter, setFilter] = useState<DocumentFilter>({
    page: 1,
    per_page: 10,
    company_id: "",
    title: "",
  });
  const [newDocumentModal, setNewDocumentModal] = useState(false);
  const { projects } = useCacheStore();
  const navigate = useNavigate();

  type MenuButtonItem = {
    label: string;
    onClick: (id: string) => void;
    icon: IconType;
  };

  const PROJECT_ACTIONS: MenuButtonItem[] = [
    {
      label: "Open in Editor",
      onClick: (docID: string) => navigate(`documents/${docID}`),
      icon: SlOptionsVertical,
    },
    {
      label: "Download",
      onClick: (docID: string) => console.log(`download`),
      icon: SlOptionsVertical,
    },
  ];

  const {
    data: documents,
    isLoading: isDocumentsLoading,
    refetch,
  } = useQuery({
    queryKey: [
      "documents",
      filter.page,
      filter.per_page,
      filter.title,
      filter.company_id,
    ],
    queryFn: () => documentAPI.list(filter),
    enabled: true,
  });

  return (
    <div className="w-full h-[100lvh] px-5 py-10 bg-primary-600 text-neutral-100">
      <section id="modals">
        <Modal
          visible={newDocumentModal}
          title="Create new document"
          onClose={() => {
            setNewDocumentModal(false);
          }}
        >
          <NewDocumentForm
            onSubmit={async (newDoc: DocumentCreation) => {
              console.log(newDoc);
              try {
                await documentAPI.create(newDoc);
              } catch (e) {
                console.error("failed to create doc: ", e);
              } finally {
                refetch();
                setNewDocumentModal(false);
              }
            }}
          />
        </Modal>
      </section>
      <section>
        <Button
          onClick={() => {
            setNewDocumentModal(true);
          }}
        >
          Create
        </Button>
      </section>

      <section>
        <div className="overflow-x-auto container mx-auto shadow-xl">
          <div>
            <div className="h-[75lvh] w-full">
              <Table className="rounded-lg">
                <Table.Head>
                  <Table.Row className="text-neutral-100 bg-primary-400">
                    <Table.HeadCell>№</Table.HeadCell>
                    <Table.HeadCell>Title</Table.HeadCell>
                    <Table.HeadCell>Project</Table.HeadCell>
                    <Table.HeadCell className="w-[3rem]"></Table.HeadCell>
                  </Table.Row>
                </Table.Head>
                {isDocumentsLoading ? (
                  <p>Documents are loading...</p>
                ) : !documents ||
                  !documents.items ||
                  documents.items.length === 0 ? (
                  <Table.Body></Table.Body>
                ) : (
                  <Table.Body>
                    {documents.items.map((item, i) => (
                      <Table.Row
                        className="cursor-pointer bg-secondary-200 text-neutral-100 hover:bg-secondary-100 p-5"
                        key={i}
                        // onClick={() => {
                        //   navigate(`/documents/${item.id}`);
                        // }}
                      >
                        <Table.Cell>{i + 1}</Table.Cell>
                        <Table.Cell>{item.title}</Table.Cell>
                        <Table.Cell>
                          {projects ? projects[item?.project_id].title : "None"}
                        </Table.Cell>
                        <Table.Cell>
                          <Menu>
                            <MenuButton
                              className={
                                "text-neutral-300 hover:text-neutral-100 cursor-pointer "
                              }
                            >
                              <SlOptionsVertical />
                            </MenuButton>
                            <MenuItems anchor="bottom end">
                              <MenuItem>
                                <button
                                  onClick={() => {
                                    navigate(`/documents/${item.id}`);
                                  }}
                                  className="block w-full px-2 text-left py-2 text-neutral-100 bg-secondary-200 hover:bg-secondary-100 transition-all duration-300 cursor-pointer"
                                >
                                  Open in Editor
                                </button>
                              </MenuItem>
                              <MenuItem>
                                <button className="block w-full px-2 text-left py-2 text-neutral-100 bg-secondary-200 hover:bg-secondary-100 transition-all duration-300 cursor-pointer">
                                  Download
                                </button>
                              </MenuItem>
                            </MenuItems>
                          </Menu>
                        </Table.Cell>
                      </Table.Row>
                    ))}
                  </Table.Body>
                )}
              </Table>
              <button
                className="w-full cursor-pointer group hover:bg-secondary-100 py-4 group:transition-all duration-300"
                onClick={() => {
                  setNewDocumentModal(true);
                }}
              >
                <BsFillPlusCircleFill
                  size="30"
                  className="mx-auto text-neutral-300 group-hover:text-accent-300 "
                />
              </button>
            </div>
            {documents && documents.items && (
              <Paginator
                page={documents.page ?? 0}
                per_page={documents.per_page ?? 0}
                total_items={documents.total_items ?? 0}
                total_pages={documents.total_pages ?? 0}
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

export default DocumentsPage;
