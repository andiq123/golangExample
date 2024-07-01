package models

import "api/db"

type Registration struct {
	ID       int64
	Event_id int64
	User_id  int64
}

func (r *Registration) Save() error {
	query := "INSERT INTO registration (event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Event_id, r.User_id)
	if err != nil {
		return nil
	}

	return nil
}
