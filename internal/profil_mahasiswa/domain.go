package profilmahasiswa

// SELECT id, name, username, email, password, prodi FROM mahasiswa.profil;
//PostgreSQL
type Profil struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Prodi    string `json:"prodi"`
	Email    string `json:"email"`
}

//Redis
type UserDraft struct {
	ID       int    `redis:"id"`
	Name     string `redis:"Name"`
	Username string `redis:"username"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
	Prodi    string `redis:"prodi"`
}

// GetOTP
type DataOTP struct {
	KeyUser string `json:"keyuser"`
	OTP     string `json:"otp"`
}
