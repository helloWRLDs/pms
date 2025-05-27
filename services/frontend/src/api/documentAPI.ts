import {
  Document,
  DocumentCreation,
  DocumentFilter,
} from "../lib/document/document";
import { DocumentPDF } from "../lib/document/documentPDF";
import { buildQuery, ListItems } from "../lib/utils/list";
import { API } from "./api";

class DocumentAPI extends API {
  async create(creation: DocumentCreation): Promise<void> {
    try {
      await this.req.post(`${this.baseURL}/docs`, creation);
    } catch (e) {
      console.error("failed creating document", e);
      throw e;
    }
  }

  async download(docID: string): Promise<DocumentPDF> {
    try {
      const res = await this.req.get(`${this.baseURL}/docs/${docID}/download`, {
        responseType: "blob",
      });
      return {
        title: res.headers["X-Document-Title"],
        body: res.data,
        id: res.headers["X-Document-ID"],
      } as DocumentPDF;
    } catch (e) {
      console.error(e);
      throw e;
    }
  }

  async get(docID: string): Promise<Document> {
    try {
      const res = await this.req.get(`${this.baseURL}/docs/${docID}`);
      return res.data;
    } catch (e) {
      console.error("failed getting document", e);
    }
    return {} as Document;
  }

  async list(filter: DocumentFilter): Promise<ListItems<Document>> {
    try {
      const res = await this.req.get(
        buildQuery(`${this.baseURL}/docs/`, filter)
      );
      return res.data;
    } catch (e) {
      console.error("failed getting document", e);
    }
    return {} as ListItems<Document>;
  }

  async update(id: string, updatedDoc: Document): Promise<void> {
    try {
      await this.req.post(`${this.baseURL}/docs/${id}`, updatedDoc);
    } catch (e) {
      console.error("failed updating document", e);
    }
  }
}

const documentAPI = new DocumentAPI();

export default documentAPI;
