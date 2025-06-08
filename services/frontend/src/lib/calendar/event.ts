import { CalendarEventType } from "./calendar";

export type CalendarEvent = {
  id: string;
  date: string; // YYYY-MM-DD format
  body: string;
  company_id: string;
  type?: CalendarEventType;
  time?: string; // HH:MM format
};
