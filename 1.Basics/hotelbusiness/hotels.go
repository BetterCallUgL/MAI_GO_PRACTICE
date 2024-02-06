//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

type Event struct {
	x    int
	flag int
}

func ComputeLoad(guests []Guest) []Load {
	var events []Event
	var result []Load
	for _, human := range guests {
		events = append(events, Event{human.CheckInDate, 1})
		events = append(events, Event{human.CheckOutDate, -1})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].x < events[j].x || events[i].x == events[j].x && events[i].flag > events[j].flag
	})

	cnt := 0
	if len(events) > 0 {
		prev_date := events[0].x

		for _, date := range events {

			if date.x != prev_date {
				if len(result) == 0 || len(result) > 0 && result[len(result)-1].GuestCount != cnt {
					result = append(result, Load{prev_date, cnt})
				}
			}

			prev_date = date.x
			cnt += date.flag
		}

		if len(result) == 0 || len(result) > 0 && result[len(result)-1].GuestCount != cnt {
			result = append(result, Load{prev_date, cnt})
		}
	}

	return result
}
