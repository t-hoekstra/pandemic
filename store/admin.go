package store

type Users struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessLevel int    `json:"accesslevel"`
	Balance     int    `json:"balance"`
	Manufacture int    `json:"manufacture"`
}

func (us *UserStore) ListAll(offset int64, limit int64) ([]*Users, error) {
	// Query
	var (
		// users       []Users
		id          string
		username    string
		accesslevel int
		balance     int
		manufacture int
	)
	q := `
	SELECT id, username, accesslevel, balance, manufacture
	FROM user
	LIMIT ? OFFSET ?;`

	rows, err := us.db.Query(q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*Users, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &username, &accesslevel, &balance, &manufacture)
		if err != nil {
			return nil, err
		}
		results = append(results, &Users{id, username, accesslevel, balance, manufacture})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

func (us *UserStore) MakeUserAdmin(userId string) error {

	q := `
	UPDATE
	user
	SET accesslevel = 150
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return err

}

func (us *UserStore) DeleteUser(userId string) error {
	//TODO make sure its save delete for foreign keys
	q := `
	DELETE
	FROM user
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return err

}
