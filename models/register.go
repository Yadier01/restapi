package models

import "github.com/Yadier01/rest-api/db"

type Register struct {
	ID      int64
	EventID int64
	UserID  int64
}

func (r Register) Save() error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(r.EventID, r.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r Register) Delete() error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	_, err := db.DB.Exec(query, r.EventID, r.UserID)
	if err != nil {
		return err
	}
	return nil
}
