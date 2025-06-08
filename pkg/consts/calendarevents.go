package consts

type CalendarEventType string

const (
	CalendarEventTypeMeeting      CalendarEventType = "meeting"
	CalendarEventTypeDeadline     CalendarEventType = "deadline"
	CalendarEventTypeReview       CalendarEventType = "review"
	CalendarEventTypeHoliday      CalendarEventType = "holiday"
	CalendarEventTypePresentation CalendarEventType = "presentation"
	CalendarEventTypeOther        CalendarEventType = "other"
)
