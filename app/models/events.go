package models

import (
	"log"
	"time"
)

type Event struct {
	ID         int
	Content    string
	Location   string
    StartTime time.Time
    EndTime   time.Time
    HostID     int
    GroupID     int
	CreatedAt time.Time
}

func (u *User) CreateEvent(content string,
                           location string,
                           start_time time.Time,
                           end_time time.Time,
                           group_id int) (err error) {
    cmd := `INSERT INTO events (
                content,
                location,
                start_time,
                end_time,
                host_id,
                group_id,
                created_at)
            VALUES ($1, $2, $3, $4, $5, $6)`
    _, err = Db.Exec(cmd, content, location, start_time, end_time,
                     u.ID, group_id, time.Now())
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

// func GetEvent(id int) (event Event, err error) {
//     cmd := `SELECT id, content, location, user_id,
//                    start_time, end_time, created_at
//             FROM events
//             WHERE id = $1`
//     event = Event{}
//     err = Db.QueryRow(cmd, id).Scan(
//         &event.ID,
//         &event.Content,
//         &event.Location,
//         &event.StartTime,
//         &event.EndTime,
//         &event.CreatedAt,
//     )
//
//     return event, err
// }

// func GetEvents() (events []Event, err error) {
//     cmd := `SELECT id, content, location, user_id,
//                    start_time, end_time, created_at
//             FROM events`
// 	rows, err := Db.Query(cmd)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	for rows.Next() {
// 		var event Event
// 		err = rows.Scan(
// 			&event.ID,
//             &event.Content,
// 			&event.Location,
//             &event.HostID,
//             &event.StartTime,
//             &event.EndTime,
// 			&event.CreatedAt)
//         if err != nil {
//             log.Fatalln(err)
//         }
//         events = append(events, event)
// 	}
//     rows.Close()
//
// 	return events, err
// }

func (u *User) GetEventsByUser() (events []Event, err error) {
    cmd := `SELECT id, content, location, start_time, end_time,
                   host_id, group_id, created_at
            FROM events
            WHERE user_id = $1`
    rows, err := Db.Query(cmd, u.ID)
    if err != nil {
        log.Fatalln(err)
    }
    for rows.Next() {
        var event Event
        err = rows.Scan(
            &event.ID,
            &event.Content,
            &event.Location,
            &event.StartTime,
            &event.EndTime,
            &event.HostID,
            &event.GroupID,
            &event.CreatedAt)
        if err != nil {
            log.Fatalln(err)
        }
        events = append(events, event)
    }
    rows.Close()

    return events, err
}
// func (e *Event) UpdateEvent() error {
//     cmd := `UPDATE events SET content = $1, location = $2, user_id = $3,
//                    start_time = $4, end_time = $5
//             WHERE id = $6`
//     _, err = Db.Exec(cmd, e.Content, e.Location, e.HostID,
//                      e.StartTime, e.EndTime, e.ID)
//     if err != nil {
//         log.Fatalln(err)
//     }
//     return err
// }
//
// func (e *Event) DeleteEvent() error {
//     cmd := `DELETE FROM events
//             WHERE id = $1`
//     _, err = Db.Exec(cmd, e.ID)
//     if err != nil {
//         log.Fatalln(err)
//     }
//     return err
// }
