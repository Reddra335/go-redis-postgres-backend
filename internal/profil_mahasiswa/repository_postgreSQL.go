package profilmahasiswa

import (
	"context"
	"newProjek/internal/shared/helper/helperdb"
)

// Implementasi
type mhsRepoMhsImpln struct {
	exec helperdb.Executor
}

func NewRepoMhsImpln(exec helperdb.Executor) MhsRepo {

	return &mhsRepoMhsImpln{
		exec: exec,
	}

}

func (m *mhsRepoMhsImpln) TambahMhs(ctx context.Context, d Profil) (Profil, error) {
	query := "INSERT INTO mahasiswa.profil (name, username, email, password, prodi) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := m.exec.QueryRowContext(ctx, query, d.Name, d.Username, d.Email, d.Password, d.Prodi).Scan(&id)

	if err != nil {

		return Profil{}, err

	}
	d.ID = id

	resId := int(id)
	d.ID = resId

	return d, nil
}
func (m *mhsRepoMhsImpln) GetById(ctx context.Context, id int) (Profil, error) {

	query := "SELECT id, name, username, email, password, prodi FROM mahasiswa.profil WHERE id=$1"

	row := m.exec.QueryRowContext(ctx, query, id)

	var mhs Profil

	err := row.Scan(&mhs.ID, &mhs.Name, &mhs.Username, &mhs.Email, &mhs.Password, &mhs.Prodi)

	if err != nil {

		return Profil{}, err

	}

	return mhs, nil
}
func (m *mhsRepoMhsImpln) GetByEmailUsername(ctx context.Context, email, username string) (Profil, error) {

	return Profil{}, nil
}
