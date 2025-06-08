import { useMemo } from "react";

export interface CalendarDay {
  date: number;
  month: "prev" | "current" | "next";
  fullDate: string;
  isToday: boolean;
}

export const useCalendarData = (month: number, year: number): CalendarDay[] => {
  return useMemo(() => {
    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);
    const firstDayOfWeek = (firstDay.getDay() + 6) % 7;
    const daysInMonth = lastDay.getDate();

    console.log("Calendar Debug:", {
      year,
      month,
      firstDay,
      lastDay,
      firstDayOfWeek,
      daysInMonth,
    });

    const prevMonth = month === 0 ? 11 : month - 1;
    const prevYear = month === 0 ? year - 1 : year;
    const daysInPrevMonth = new Date(prevYear, prevMonth + 1, 0).getDate();

    const calendarDays: CalendarDay[] = [];

    // Add days from previous month
    for (let i = firstDayOfWeek - 1; i >= 0; i--) {
      const date = daysInPrevMonth - i;
      const fullDate = `${prevYear}-${String(prevMonth + 1).padStart(
        2,
        "0"
      )}-${String(date).padStart(2, "0")}`;
      calendarDays.push({
        date,
        month: "prev",
        fullDate,
        isToday: false,
      });
    }

    // Add days from current month
    const today = new Date();
    for (let date = 1; date <= daysInMonth; date++) {
      const fullDate = `${year}-${String(month + 1).padStart(2, "0")}-${String(
        date
      ).padStart(2, "0")}`;
      const isToday =
        today.getFullYear() === year &&
        today.getMonth() === month &&
        today.getDate() === date;
      calendarDays.push({
        date,
        month: "current",
        fullDate,
        isToday,
      });
    }

    const nextMonth = month === 11 ? 0 : month + 1;
    const nextYear = month === 11 ? year + 1 : year;
    const remainingCells = 42 - calendarDays.length;
    for (let date = 1; date <= remainingCells; date++) {
      const fullDate = `${nextYear}-${String(nextMonth + 1).padStart(
        2,
        "0"
      )}-${String(date).padStart(2, "0")}`;
      calendarDays.push({
        date,
        month: "next",
        fullDate,
        isToday: false,
      });
    }

    console.log(
      "Calendar Days:",
      calendarDays.length,
      calendarDays.slice(0, 10)
    );
    return calendarDays;
  }, [month, year]);
};
