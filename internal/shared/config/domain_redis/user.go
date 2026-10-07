package domainredis

type UserDraft struct {
	ID       int    `redis:"id"`
	Name     string `redis:"Name"`
	Username string `redis:"username"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
	Prodi    string `redis:"prodi"`
}
