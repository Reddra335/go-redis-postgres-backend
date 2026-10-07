package profilmahasiswa

import (
	"context"
	"newProjek/internal/helperdb"
)

// Interface untuk repo

type MhsRepo interface {
	TambahMhs(ctx context.Context, domain Profil, exec helperdb.Executor) (Profil, error)
	GetById(ctx context.Context, id int, exec helperdb.Executor) (Profil, error)
}

// Implementasi
type MhsRepoMhsImpln struct {
}

func NewRepoMhsImpln() MhsRepo {

	return &MhsRepoMhsImpln{}

}

func (m *MhsRepoMhsImpln) TambahMhs(ctx context.Context, domain Profil, exec helperdb.Executor) (Profil, error) {
	query := "INSERT INTO mahasiswa.profil (name, username, email, password, prodi) VALUES ('?', '?', '?', '?', '?')"

	data, err := exec.ExecContext(ctx, query, domain.Name, domain.Username, domain.Email, domain.Password, domain.Prodi)

	if err != nil {

		return Profil{}, err

	}

	id, err := data.LastInsertId()

	if err != nil {
		return Profil{}, err
	}

	resId := int(id)
	domain.ID = resId

	return domain, nil
}
func (m *MhsRepoMhsImpln) GetById(ctx context.Context, id int, exec helperdb.Executor) (Profil, error) {

	query := "SELECT id, name, username, email, password, prodi FROM mahasiswa.profil WHERE id=?"

	row := exec.QueryRowContext(ctx, query, id)

	var mhs Profil

	err := row.Scan(&mhs.ID, &mhs.Name, &mhs.Username, &mhs.Email, &mhs.Prodi)

	if err != nil {

		return Profil{}, err

	}

	return mhs, nil
}
