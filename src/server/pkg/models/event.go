package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"title"`
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
	Color     string `json:"color"`
}

func (e *Event) CreateEvent() *Event {
	log.Println("Inserting", e.Name, "...")
	statement, err := db.Prepare(`
	INSERT INTO event(name, start, end, color, userid) VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Println("Insertion: Failed to prepare statement")
		log.Println(err)
		return e
	}
	defer statement.Close()
	_, err = statement.Exec(e.Name, e.StartTime, e.EndTime, e.Color, GetUserId())
	if err != nil {
		log.Println("Insertion: Failed to execute statement")
		log.Println(err)
		return e
	}
	return e
}

func GetAllEvent() []Event {
	Events := make([]Event, 0)
	log.Println("Getting all events...")
	rows, err := db.Query(`
	SELECT id, name, start, end, color
	FROM event
	WHERE userid=?
	ORDER BY id;
	`, GetUserId())
	if err != nil {
		log.Println("Query: Failed to execute query")
		log.Println(err)
		return nil
	}
	defer rows.Close()

	//ignore := 0
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.StartTime, &e.EndTime, &e.Color); err != nil {
			log.Println("Query: Failed to read query")
			log.Println(err)
			return nil
		}
		Events = append(Events, e)
	}
	log.Println("Queried", len(Events), "Events")
	return Events
}

func GetEventByID(id int) Event {
	var getEvent Event
	log.Println("Getting event", id, "...")
	rows, err := db.Query(`
	SELECT *
	FROM event
	WHERE id=? and userid=?;
	`, id, GetUserId())
	if err != nil {
		log.Println("Query: Failed to execute query")
		log.Println(err)
		return getEvent
	}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&getEvent.ID, &getEvent.Name, &getEvent.StartTime, &getEvent.EndTime, &getEvent.Color, nil); err != nil {
		log.Println("Query: Failed to read query")
		log.Println(err)
		return getEvent
	}
	return getEvent
}

func DeleteEvent(id int) Event {
	var event Event
	log.Println("Deleting event", id, "...")
	statement, err := db.Prepare(`
	DELETE FROM event WHERE id=?;
	`)
	if err != nil {
		log.Println("Deletion: failed to prepare statement")
		log.Println(err)
		return event
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		log.Println("Deletion: failed to execute statement")
		log.Println(err)
		return event
	}
	event.ID = id
	return event
}

func (e *Event) UpdateEvent() *Event {
	log.Println("Updating event", e.ID, "...")
	statement, err := db.Prepare(`
	UPDATE event
	SET name = ?, start = ?, end = ?, color = ?
	WHERE id=?;
	`)
	if err != nil {
		log.Println("Update: Failed to prepare statement")
		log.Println(err)
		return e
	}
	defer statement.Close()
	_, err = statement.Exec(e.Name, e.StartTime, e.EndTime, e.Color, e.ID)
	if err != nil {
		log.Println("Update: Failed to execute statement")
		log.Println(err)
		return e
	}
	return e
}
