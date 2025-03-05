package dto

type Mentor struct {
	Email        string `json:"email" validate:"required,email"`
	Nama         string `json:"nama" validate:"required"`
	Spesialisasi string `json:"spesialisasi" validate:"required"`
	ProfilMentor string `json:"profil_mentor" validate:"required"`
}

type MentorsDetails struct {
	ProfilMentor    string `json:"profil_mentor" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	Deskripsi       string `json:"deskripsi" validate:"required"`
	Spesialisasi    string `json:"spesialisasi" validate:"required"`
	Pendidikan      string `json:"pendidikan" validate:"required"`
	PengalamanKerja string `json:"pengalaman_kerja" validate:"required"`
	Pencapaian      string `json:"pencapaian" validate:"required"`
	Keahlian        string `json:"keahlian" validate:"required"`
	TopikAjar       string `json:"topik_ajar" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
}
