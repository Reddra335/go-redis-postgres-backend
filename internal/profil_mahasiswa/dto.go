package profilmahasiswa

type ReqMahasiswa struct {
	Name     string `validate:"required,min=1,max=50" json:"name"`
	Username string `validate:"required,min=1,max=50" json:"username"`
	Password string `validate:"required,min=8,max=255" json:"password"`
	Prodi    string `validate:"required,min=1,max=20" json:"prodi"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
}

type ResponseMahasiswa struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Prodi    string `json:"prodi"`
	KeyUser  string
}
