import { FC, useState } from "react";
import {
  MdNavigateBefore,
  MdNavigateNext,
  MdEvent,
  MdToday,
  MdAdd,
} from "react-icons/md";
import { Modal } from "../ui/Modal";
import { DAY_ABBR, DAY_NAMES, MONTH_NAMES } from "../../lib/utils/time";
import {
  CalendarEventType,
  eventTypeIcon,
  eventTypeStyle,
} from "../../lib/calendar/calendar";
import { CalendarEvent } from "../../lib/calendar/event";
import { useCalendarData } from "../../hooks/useCalendarData";

interface CalendarProps {
  month: number; // 0-11 (JavaScript month format)
  year: number;
  onMonthChange?: (month: number, year: number) => void;
  events?: CalendarEvent[];
  onAddEvent?: (event: Omit<CalendarEvent, "id">) => void;
}

const HARDCODED_EVENTS: CalendarEvent[] = [
  {
    id: "1",
    date: "2025-06-01",
    body: "Sprint Planning Meeting",
    company_id: "company1",
    type: "meeting",
    time: "09:00",
  },
  {
    id: "2",
    date: "2025-06-01",
    body: "Project Deadline",
    company_id: "company1",
    type: "deadline",
    time: "17:00",
  },
  {
    id: "3",
    date: "2025-06-01",
    body: "Team Retrospective",
    company_id: "company1",
    type: "review",
    time: "14:00",
  },
  {
    id: "4",
    date: "2025-06-01",
    body: "Client Presentation",
    company_id: "company1",
    type: "presentation",
    time: "10:30",
  },
  {
    id: "5",
    date: "2025-06-01",
    body: "Code Review Session",
    company_id: "company1",
    type: "review",
    time: "15:00",
  },
  {
    id: "6",
    date: "2025-06-01",
    body: "Monthly Review",
    company_id: "company1",
    type: "review",
    time: "11:00",
  },
  {
    id: "7",
    date: "2025-06-03",
    body: "Holiday Break",
    company_id: "company1",
    type: "holiday",
  },
  {
    id: "8",
    date: "2025-06-05",
    body: "Year End Review",
    company_id: "company1",
    type: "review",
    time: "16:00",
  },
  {
    id: "9",
    date: "2025-06-08",
    body: "Team Standup",
    company_id: "company1",
    type: "meeting",
    time: "11:30",
  },
];

export const Calendar: FC<CalendarProps> = ({
  month,
  year,
  onMonthChange,
  events = HARDCODED_EVENTS,
  onAddEvent,
}) => {
  const [selectedDate, setSelectedDate] = useState<string | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [newEventTitle, setNewEventTitle] = useState("");
  const [newEventTime, setNewEventTime] = useState("");
  const [newEventType, setNewEventType] =
    useState<CalendarEvent["type"]>("meeting");

  const calendarData = useCalendarData(month, year);

  const getEventsForDate = (date: string) => {
    return events.filter((event) => event.date === date);
  };

  const handlePrevMonth = () => {
    if (onMonthChange) {
      const prevMonth = month === 0 ? 11 : month - 1;
      const prevYear = month === 0 ? year - 1 : year;
      onMonthChange(prevMonth, prevYear);
    }
  };

  const handleNextMonth = () => {
    if (onMonthChange) {
      const nextMonth = month === 11 ? 0 : month + 1;
      const nextYear = month === 11 ? year + 1 : year;
      onMonthChange(nextMonth, nextYear);
    }
  };

  const handleDayClick = (date: string, month: string) => {
    if (month === "current") {
      setSelectedDate(date);
      setIsModalOpen(true);
    }
  };

  const handleAddEvent = () => {
    if (newEventTitle.trim() && selectedDate && onAddEvent) {
      const newEvent: Omit<CalendarEvent, "id"> = {
        date: selectedDate,
        body: newEventTitle.trim(),
        company_id: "company1",
        type: newEventType,
        time: newEventTime || undefined,
      };
      onAddEvent(newEvent);
      setNewEventTitle("");
      setNewEventTime("");
      setNewEventType("meeting");
    }
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedDate(null);
    setNewEventTitle("");
    setNewEventTime("");
    setNewEventType("meeting");
  };

  const selectedDateEvents = selectedDate ? getEventsForDate(selectedDate) : [];
  const selectedDateObj = selectedDate ? new Date(selectedDate) : null;

  return (
    <>
      <div className="bg-primary-500/30 backdrop-blur-lg rounded-2xl shadow-2xl border border-primary-400/30 overflow-hidden w-full">
        {/* Calendar Header */}
        <div className="bg-primary-600/50 px-6 py-4 border-b border-primary-400/30">
          <div className="flex items-center justify-between">
            <button
              onClick={handlePrevMonth}
              className="group p-2 rounded-lg bg-white/10 hover:bg-white/20 text-white transition-all duration-200 hover:scale-105 border border-white/20"
            >
              <MdNavigateBefore
                size={20}
                className="group-hover:-translate-x-0.5 transition-transform"
              />
            </button>

            <div className="text-center">
              <h2 className="text-2xl font-bold text-white">
                {MONTH_NAMES[month]} {year}
              </h2>
            </div>

            <button
              onClick={handleNextMonth}
              className="group p-2 rounded-lg bg-white/10 hover:bg-white/20 text-white transition-all duration-200 hover:scale-105 border border-white/20"
            >
              <MdNavigateNext
                size={20}
                className="group-hover:translate-x-0.5 transition-transform"
              />
            </button>
          </div>
        </div>

        {/* Days of Week Header */}
        <div className="grid grid-cols-7 bg-primary-600/30 border-b border-primary-400/30">
          {DAY_ABBR.map((day) => (
            <div
              key={day}
              className="p-3 text-center text-sm font-semibold text-white/90 uppercase tracking-wide"
            >
              {day}
            </div>
          ))}
        </div>

        {/* Calendar Grid */}
        <div className="grid grid-cols-7 bg-primary-500/20 min-h-[600px]">
          {calendarData.map((day, index) => {
            const dayEvents = getEventsForDate(day.fullDate);
            const isWeekend = index % 7 === 5 || index % 7 === 6; // Saturday (5) or Sunday (6) in Monday-first week

            return (
              <div
                key={`${day.fullDate}-${index}`}
                onClick={() => handleDayClick(day.fullDate, day.month)}
                className={`
                  relative h-24 sm:h-32 p-2 border-b border-r border-primary-400/20 
                  cursor-pointer transition-all duration-200 group flex flex-col
                  ${
                    day.month === "current"
                      ? isWeekend
                        ? "bg-primary-600/20 hover:bg-primary-400/30"
                        : "bg-primary-500/20 hover:bg-primary-400/30"
                      : "bg-primary-600/10 text-white/50 hover:bg-primary-500/20"
                  }
                  ${
                    day.isToday
                      ? "bg-accent-500/30 border-accent-400/50 ring-1 ring-accent-400/50"
                      : ""
                  }
                  hover:shadow-lg hover:z-10 backdrop-blur-sm
                `}
              >
                {/* Date Number */}
                <div className="flex items-center justify-between mb-1 flex-shrink-0">
                  <span
                    className={`
                      inline-flex items-center justify-center w-6 h-6 sm:w-7 sm:h-7 rounded-full text-xs sm:text-sm font-semibold
                      ${
                        day.isToday
                          ? "bg-accent-600 text-white shadow-lg"
                          : day.month === "current"
                          ? "text-white group-hover:bg-white/20 group-hover:text-white"
                          : "text-white/50"
                      }
                      transition-all duration-200
                    `}
                  >
                    {day.date}
                  </span>

                  <div className="flex items-center space-x-1">
                    {day.isToday && (
                      <MdToday className="text-accent-400" size={12} />
                    )}

                    {dayEvents.length > 0 && (
                      <span className="text-xs bg-white/20 text-white px-1 py-0.5 rounded-full font-medium backdrop-blur-sm">
                        {dayEvents.length}
                      </span>
                    )}
                  </div>
                </div>

                {/* Events Preview */}
                <div className="space-y-1 flex-1 overflow-hidden">
                  {dayEvents.slice(0, 2).map((event) => (
                    <div
                      key={event.id}
                      className={`text-xs px-1.5 py-0.5 sm:px-2 sm:py-1 rounded truncate ${eventTypeStyle(
                        event.type as CalendarEventType
                      )} backdrop-blur-sm border border-white/20`}
                      title={`${event.time ? event.time + " - " : ""}${
                        event.body
                      }`}
                    >
                      {event.time && (
                        <span className="font-mono text-xs opacity-90">
                          {event.time}
                        </span>
                      )}{" "}
                      <span className="truncate">{event.body}</span>
                    </div>
                  ))}

                  {dayEvents.length > 2 && (
                    <div className="text-xs text-white/70 px-1.5 py-0.5 bg-white/10 rounded truncate backdrop-blur-sm">
                      +{dayEvents.length - 2} more
                    </div>
                  )}
                </div>

                {/* Add event button on hover */}
                {day.month === "current" && (
                  <div className="absolute bottom-1 right-1 opacity-0 group-hover:opacity-100 transition-opacity">
                    <div className="w-4 h-4 sm:w-5 sm:h-5 bg-accent-600 text-white rounded-full flex items-center justify-center text-xs shadow-lg backdrop-blur-sm">
                      <MdAdd size={10} />
                    </div>
                  </div>
                )}
              </div>
            );
          })}
        </div>
      </div>

      {/* Day Details Modal */}
      {isModalOpen && selectedDate && (
        <Modal
          visible={isModalOpen}
          onClose={closeModal}
          title={`${selectedDateObj && DAY_NAMES[selectedDateObj.getDay()]} - ${
            selectedDateObj &&
            selectedDateObj.toLocaleDateString("en-US", {
              month: "long",
              day: "numeric",
              year: "numeric",
            })
          }`}
          size="md"
        >
          {/* Existing Events */}
          <div className="mb-6">
            <h4 className="font-semibold text-white mb-3 flex items-center">
              <MdEvent className="mr-2" size={16} />
              Events ({selectedDateEvents.length})
            </h4>

            {selectedDateEvents.length === 0 ? (
              <p className="text-white/70 text-sm italic">
                No events scheduled for this day
              </p>
            ) : (
              <div className="space-y-2">
                {selectedDateEvents.map((event) => (
                  <div
                    key={event.id}
                    className="p-3 rounded-lg border-l-4 bg-white/10 backdrop-blur-sm border-white/20"
                  >
                    <div className="flex items-start justify-between">
                      <div>
                        <p className="font-medium text-white">{event.body}</p>
                        {event.time && (
                          <p className="text-sm text-white/70 font-mono">
                            {event.time}
                          </p>
                        )}
                        <p className="text-xs text-white/60 capitalize">
                          {event.type}
                        </p>
                      </div>
                      <span className="text-lg">
                        {eventTypeIcon(event.type as CalendarEventType)}
                      </span>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>

          {/* Add New Event Form */}
          <div className="border-t border-white/20 pt-6">
            <h4 className="font-semibold text-white mb-3 flex items-center">
              <MdAdd className="mr-2" size={16} />
              Add New Event
            </h4>

            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Event Title
                </label>
                <input
                  type="text"
                  value={newEventTitle}
                  onChange={(e) => setNewEventTitle(e.target.value)}
                  placeholder="Enter event title..."
                  className="w-full px-3 py-2 bg-white/10 border border-white/20 rounded-lg text-white placeholder-white/50 focus:ring-2 focus:ring-accent-500 focus:border-accent-500 backdrop-blur-sm"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Time (Optional)
                </label>
                <input
                  type="time"
                  value={newEventTime}
                  onChange={(e) => setNewEventTime(e.target.value)}
                  className="w-full px-3 py-2 bg-white/10 border border-white/20 rounded-lg text-white focus:ring-2 focus:ring-accent-500 focus:border-accent-500 backdrop-blur-sm"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-white/90 mb-2">
                  Event Type
                </label>
                <select
                  value={newEventType}
                  onChange={(e) =>
                    setNewEventType(e.target.value as CalendarEvent["type"])
                  }
                  className="w-full px-3 py-2 bg-white/10 border border-white/20 rounded-lg text-white focus:ring-2 focus:ring-accent-500 focus:border-accent-500 backdrop-blur-sm"
                >
                  <option value="meeting">👥 Meeting</option>
                  <option value="deadline">⏰ Deadline</option>
                  <option value="review">📋 Review</option>
                  <option value="presentation">📊 Presentation</option>
                  <option value="holiday">🎉 Holiday</option>
                </select>
              </div>

              <button
                onClick={handleAddEvent}
                disabled={!newEventTitle.trim()}
                className="w-full bg-accent-600 text-white py-2 px-4 rounded-lg hover:bg-accent-700 disabled:bg-white/20 disabled:cursor-not-allowed transition-colors font-medium"
              >
                Add Event
              </button>
            </div>
          </div>
        </Modal>
      )}
    </>
  );
};

export default Calendar;
