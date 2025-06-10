import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import {
  Document,
  DocumentCreation,
  DocumentFilter,
} from "../../lib/document/document";
import { Button } from "../../components/ui/Button";
import { Modal } from "../../components/ui/Modal";
import NewDocumentForm from "../../components/forms/NewDocumentForm";
import Table from "../../components/ui/Table";
import { useNavigate } from "react-router-dom";
import Paginator from "../../components/ui/Paginator";
import { SlOptionsVertical } from "react-icons/sl";
import { ContextMenu } from "../../components/ui/ContextMenu";
import {
  BsDownload,
  BsFillPlusCircleFill,
  BsFileEarmarkText,
  BsFolder,
} from "react-icons/bs";
import { MdOpenInNew } from "react-icons/md";
import { formatTime } from "../../lib/utils/time";
import { infoToast, errorToast } from "../../lib/utils/toast";
import useMetaCache from "../../store/useMetaCache";
import { useProjectList } from "../../hooks/useData";
import analyticsAPI from "../../api/analyticsAPI";

const DocumentsPage = () => {
  const metaCache = useMetaCache();

  const { getProjectName } = useProjectList(
    metaCache.metadata.selectedProject?.company_id ?? ""
  );

  const [filter, setFilter] = useState<DocumentFilter>({
    page: 1,
    per_page: 10,
    project_id: metaCache.metadata.selectedProject?.id,
    title: "",
  });
  const [newDocumentModal, setNewDocumentModal] = useState(false);

  const navigate = useNavigate();

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
      filter.project_id,
    ],
    queryFn: () => analyticsAPI.list(filter),
    enabled: !!metaCache.metadata.selectedProject,
  });

  useEffect(() => {
    if (!metaCache.metadata.selectedProject) {
      navigate("/projects");
    }
  }, [metaCache.metadata.selectedProject]);

  const handleDownloadDocument = async (doc: Document) => {
    try {
      const docPDF = await analyticsAPI.download(doc.id);
      const blob = new Blob([docPDF.body]);
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = docPDF.title || "document.pdf";
      document.body.appendChild(a);
      a.click();
      a.remove();
      window.URL.revokeObjectURL(url);
      infoToast("Document downloaded successfully");
    } catch (e) {
      console.error(e);
      errorToast("Failed to download document");
    }
  };

  const handleCreateDocument = async (newDoc: DocumentCreation) => {
    try {
      await analyticsAPI.create(newDoc);
      infoToast("Document created successfully");
      refetch();
      setNewDocumentModal(false);
    } catch (e) {
      console.error("failed to create doc: ", e);
      errorToast("Failed to create document");
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary-700 to-primary-600 text-neutral-100">
      <div className="px-6 py-8">
        <div className="max-w-7xl mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
            <div>
              <h1 className="text-3xl font-bold flex items-center gap-3">
                <BsFileEarmarkText className="text-accent-500" />
                <span className="text-accent-500">
                  {metaCache.metadata.selectedProject?.title}
                </span>
                <span>Documents</span>
              </h1>
              <p className="text-neutral-300 mt-2">
                Manage and organize your project documents
              </p>
            </div>

            <div className="flex items-center gap-4">
              <Button
                onClick={() => setNewDocumentModal(true)}
                className="flex items-center gap-2 bg-accent-500 text-primary-700 hover:bg-accent-400"
              >
                <BsFillPlusCircleFill />
                New Document
              </Button>
            </div>
          </div>
        </div>
      </div>

      <div className="px-6 pb-8">
        <div className="max-w-7xl mx-auto">
          <div className="bg-secondary-200 rounded-lg overflow-hidden shadow-lg">
            <div className="overflow-x-auto">
              <table className="w-full border-collapse">
                <Table.Head className="bg-primary-400 text-neutral-100 sticky top-0 z-10">
                  <tr>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[8%] border-r border-primary-300">
                      #
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[40%] border-r border-primary-300">
                      Title
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[25%] border-r border-primary-300">
                      Project
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[20%] border-r border-primary-300">
                      Created
                    </Table.HeadCell>
                    <Table.HeadCell className="px-6 py-4 text-left font-semibold text-sm uppercase tracking-wide w-[7%]">
                      Actions
                    </Table.HeadCell>
                  </tr>
                </Table.Head>
                <Table.Body className="divide-y divide-secondary-100">
                  {isDocumentsLoading ? (
                    Array(10)
                      .fill(0)
                      .map((_, index) => (
                        <Table.Row
                          key={index}
                          className="bg-secondary-200 hover:bg-secondary-100 transition-all duration-200 animate-pulse"
                        >
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-8"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-3/4"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-32"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                            <div className="h-4 bg-secondary-100 rounded w-24"></div>
                          </Table.Cell>
                          <Table.Cell className="px-6 py-4">
                            <div className="h-8 bg-secondary-100 rounded w-8"></div>
                          </Table.Cell>
                        </Table.Row>
                      ))
                  ) : !documents?.items || documents.items.length === 0 ? (
                    <Table.Row className="bg-secondary-200">
                      <Table.Cell
                        className="px-6 py-12 text-center text-neutral-400 italic text-lg"
                        {...({ colSpan: 5 } as any)}
                      >
                        <div className="flex flex-col items-center gap-3">
                          <BsFileEarmarkText className="text-4xl text-neutral-500" />
                          <span>No documents found</span>
                          <span className="text-sm text-neutral-500">
                            Create your first document to get started
                          </span>
                        </div>
                      </Table.Cell>
                    </Table.Row>
                  ) : (
                    documents.items.map((doc, index) => (
                      <Table.Row
                        key={doc.id}
                        className="bg-secondary-200 hover:bg-secondary-100 transition-all duration-200 cursor-pointer group border-l-4 border-l-transparent hover:border-l-accent-500"
                        onClick={() => navigate(`/documents/${doc.id}`)}
                      >
                        <Table.Cell className="px-6 py-4 font-mono text-neutral-400 text-sm border-r border-secondary-100 group-hover:text-accent-400 transition-colors">
                          {(filter.page - 1) * filter.per_page + index + 1}
                        </Table.Cell>
                        <Table.Cell className="px-6 py-4 font-medium border-r border-secondary-100 group-hover:text-accent-400 transition-colors">
                          <div className="flex items-center gap-3">
                            <BsFileEarmarkText
                              className="text-accent-500 flex-shrink-0"
                              size={20}
                            />
                            <div
                              className="truncate max-w-xs"
                              title={doc.title}
                            >
                              {doc.title}
                            </div>
                          </div>
                        </Table.Cell>
                        <Table.Cell className="px-6 py-4 border-r border-secondary-100">
                          <div className="flex items-center gap-2">
                            <BsFolder
                              className="text-accent-500 flex-shrink-0"
                              size={16}
                            />
                            <span className="truncate">
                              {getProjectName(doc.project_id)}
                            </span>
                          </div>
                        </Table.Cell>
                        <Table.Cell className="px-6 py-4 border-r border-secondary-100 text-sm text-neutral-300">
                          {doc.created_at
                            ? formatTime(doc.created_at.seconds)
                            : "Unknown"}
                        </Table.Cell>
                        <Table.Cell className="px-6 py-4">
                          <ContextMenu
                            // placement="right"
                            trigger={<SlOptionsVertical size={16} />}
                            items={[
                              {
                                icon: <MdOpenInNew />,
                                label: "Open Document",
                                onClick: (e) => {
                                  e.stopPropagation();
                                  navigate(`/documents/${doc.id}`);
                                },
                              },
                              {
                                icon: <BsDownload />,
                                label: "Download PDF",
                                onClick: (e) => {
                                  e.stopPropagation();
                                  handleDownloadDocument(doc);
                                },
                              },
                            ]}
                          />
                        </Table.Cell>
                      </Table.Row>
                    ))
                  )}
                </Table.Body>
              </table>
            </div>
          </div>

          {documents && documents.items && documents.items.length > 0 && (
            <div className="flex justify-center mt-6">
              <Paginator
                page={Number(documents.page)}
                per_page={Number(documents.per_page)}
                total_items={Number(documents.total_items)}
                total_pages={Number(documents.total_pages)}
                onPageChange={(page) =>
                  setFilter({ ...filter, page: Number(page) })
                }
              />
            </div>
          )}
        </div>
      </div>

      <Modal
        visible={newDocumentModal}
        title="Create New Document"
        onClose={() => setNewDocumentModal(false)}
        className="max-w-3xl mx-auto bg-secondary-300"
      >
        <NewDocumentForm onSubmit={handleCreateDocument} />
      </Modal>
    </div>
  );
};

export default DocumentsPage;
