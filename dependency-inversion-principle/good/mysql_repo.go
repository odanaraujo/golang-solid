package good

type MySQLRepository struct {
	//connection *sql.DB
}

func (repo *MySQLRepository) InsertUser(user_id string, name string) *User {
	user := User{
		ID:   user_id,
		Name: name,
	}
	inMemoUsers[user_id] = user
	return &user
}

func (repo *MySQLRepository) GetUserByID(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}
