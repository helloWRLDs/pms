import { useState } from "react";
import { usePageSettings } from "../../hooks/usePageSettings";
import useMetaCache from "../../store/useMetaCache";
import { Calendar } from "../../components/calendar/Calendar";

const CalendarPage = () => {
  usePageSettings({ title: "Calendar", requireAuth: true });
  const metaCache = useMetaCache();

  // Initialize with current month and year
  const today = new Date();
  const [currentMonth, setCurrentMonth] = useState(today.getMonth());
  const [currentYear, setCurrentYear] = useState(today.getFullYear());

  const handleMonthChange = (month: number, year: number) => {
    setCurrentMonth(month);
    setCurrentYear(year);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-primary-700 to-primary-600 py-8 px-4">
      <div className="max-w-7xl mx-auto">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-white mb-2">Calendar</h1>
          <p className="text-white/70">
            {metaCache.metadata.selectedCompany?.name
              ? `Events for ${metaCache.metadata.selectedCompany.name}`
              : "Company calendar events"}
          </p>
        </div>

        <Calendar
          month={currentMonth}
          year={currentYear}
          onMonthChange={handleMonthChange}
        />
      </div>
    </div>
  );
};

export default CalendarPage;
