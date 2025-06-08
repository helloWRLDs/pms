export const EVENT_TYPES = {
  MEETING: "meeting",
  DEADLINE: "deadline",
  REVIEW: "review",
  HOLIDAY: "holiday",
  PRESENTATION: "presentation",
} as const;

export type CalendarEventType = (typeof EVENT_TYPES)[keyof typeof EVENT_TYPES];

export const eventTypeStyle = (type: CalendarEventType) => {
  switch (type) {
    case EVENT_TYPES.MEETING:
      return "bg-blue-500 border-blue-400 text-blue-50";
    case EVENT_TYPES.DEADLINE:
      return "bg-red-500 border-red-400 text-red-50";
    case EVENT_TYPES.REVIEW:
      return "bg-purple-500 border-purple-400 text-purple-50";
    case EVENT_TYPES.HOLIDAY:
      return "bg-green-500 border-green-400 text-green-50";
    case EVENT_TYPES.PRESENTATION:
      return "bg-orange-500 border-orange-400 text-orange-50";
    default:
      return "bg-indigo-500 border-indigo-400 text-indigo-50";
  }
};

export const eventTypeIcon = (type: CalendarEventType) => {
  switch (type) {
    case EVENT_TYPES.MEETING:
      return "👥";
    case EVENT_TYPES.DEADLINE:
      return "⏰";
    case EVENT_TYPES.REVIEW:
      return "📋";
    case EVENT_TYPES.HOLIDAY:
      return "🎉";
    case EVENT_TYPES.PRESENTATION:
      return "📊";
    default:
      return "📅";
  }
};
