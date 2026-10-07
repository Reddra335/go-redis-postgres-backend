package profilmahasiswa

// SELECT id, name, username, email, password, prodi FROM mahasiswa.profil;

type Profil struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Prodi    string `json:"prodi"`
	Email    string `json:"email"`
}
