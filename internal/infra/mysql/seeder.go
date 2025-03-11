package mysql

import (
	"innovaspace/internal/domain/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed hash password")
	}
	return string(hashedPassword)
}

func SeedMentors(db *gorm.DB) {
	mentors := []entity.Mentor{
		{
			Id:         uuid.New(),
			Email:      "dimas.arya@email.com",
			Username:   "dimasarya",
			Password:   hashPassword("DimasMakanSate123"),
			Nama:       "Dimas Arya",
			Deskripsi:  "Seorang Digital Marketing Specialist berpengalaman dengan keahlian dalam mengembangkan strategi pemasaran yang kreatif dan berbasis data. Berfokus pada SEO, SEM, manajemen media sosial, dan kampanye iklan digital, Dimas telah membantu berbagai brand meningkatkan visibilitas online dan mencapai target pasar yang tepat. Dengan pendekatan yang analitis dan inovatif, ia selalu berkomitmen untuk menghasilkan hasil terbalk dalam setiap kampanye pemasaran digital.",
			Pendidikan: "S1 Ilmu Komunikasi - Universitas Indonesia (2014-2018)",
			Preferensi: "Teknologi & Startup",
			PengalamanKerja: datatypes.JSON([]byte(`[
				"Digital Marketing Specialist - Arya Digital Solutions (2020 - sekarang)", 
				"SEO Specialist - Bright Agency (2018 - 2020)"
			]`)),
			Pencapaian: datatypes.JSON([]byte(`[
				"Meningkatkan traffic organik sebesar 150% dalam 6 bulan untuk klien e-commerce", 
				"Memimpin kampanye iklan yang menghasilkan ROI hingga 300%"
			]`)),
			Keahlian: datatypes.JSON([]byte(`[
				"Strategi Pemasaran Digital", "SEO & SEM", "Manajemen Media Sosial",
				"Google Ads & Facebook Ads", "Email Marketing", "Content Marketing",
				"Analisis Data & Google Analytics"
			]`)),
			TopikAjar: datatypes.JSON([]byte(`[
				"Strategi Pemasaran Digital", "Cara Efektif Mengembangkan Brand",
				"Optimasi SEO & SEM Untuk Bisnis", "Meningkatkan Engagement di Media Sosial"
			]`)),
			Spesialisasi: "Social Media Specialist",
			ProfilMentor: "https://rshdseakqgwspflewctn.supabase.co/storage/v1/object/sign/innovaspace-userprofile/mentor-profile/innovaspace-mentor2.jpg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJpbm5vdmFzcGFjZS11c2VycHJvZmlsZS9tZW50b3ItcHJvZmlsZS9pbm5vdmFzcGFjZS1tZW50b3IyLmpwZyIsImlhdCI6MTc0MTE0MTk0OCwiZXhwIjoyMDU2NTAxOTQ4fQ.BbunvPy3zoj3H0nXIOuGVQi3MWDJKVmJkxEdJejkWoc",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:         uuid.New(),
			Email:      "johndoe@email.com",
			Username:   "johndoe",
			Password:   hashPassword("JohnDoe123"),
			Nama:       "John Doe",
			Deskripsi:  "Seorang Digital Marketing Specialist berpengalaman dengan keahlian dalam mengembangkan strategi pemasaran yang kreatif dan berbasis data. Berfokus pada SEO, SEM, manajemen media sosial, dan kampanye iklan digital, Dimas telah membantu berbagai brand meningkatkan visibilitas online dan mencapai target pasar yang tepat. Dengan pendekatan yang analitis dan inovatif, ia selalu berkomitmen untuk menghasilkan hasil terbalk dalam setiap kampanye pemasaran digital.",
			Pendidikan: "S1 Peternakan - Universitas Brawijaya (2016-2020)",
			Preferensi: "Perdagangan",
			PengalamanKerja: datatypes.JSON([]byte(`[
				"Digital Marketing Specialist - Arya Digital Solutions (2020 - sekarang)", 
				"SEO Specialist - Bright Agency (2018 - 2020)"
			]`)),
			Pencapaian: datatypes.JSON([]byte(`[
				"Meningkatkan traffic organik sebesar 150% dalam 6 bulan untuk klien e-commerce", 
				"Memimpin kampanye iklan yang menghasilkan ROI hingga 300%"
			]`)),
			Keahlian: datatypes.JSON([]byte(`[
				"Strategi Pemasaran Digital", "SEO & SEM", "Manajemen Media Sosial",
				"Google Ads & Facebook Ads", "Email Marketing", "Content Marketing",
				"Analisis Data & Google Analytics"
			]`)),
			TopikAjar: datatypes.JSON([]byte(`[
				"Strategi Pemasaran Digital", "Cara Efektif Mengembangkan Brand",
				"Optimasi SEO & SEM Untuk Bisnis", "Meningkatkan Engagement di Media Sosial"
			]`)),
			ProfilMentor: "https://rshdseakqgwspflewctn.supabase.co/storage/v1/object/sign/innovaspace-userprofile/mentor-profile/innovaspace-mentor1.jpg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJpbm5vdmFzcGFjZS11c2VycHJvZmlsZS9tZW50b3ItcHJvZmlsZS9pbm5vdmFzcGFjZS1tZW50b3IxLmpwZyIsImlhdCI6MTc0MTE0MTc5MCwiZXhwIjoyMDU2NTAxNzkwfQ.nAOFzQ9QlnDcvWN6ixIW9TDS_cGx93EbU-fEf3bQ4ak",
			Spesialisasi: "Social Media Specialist",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
	}

	for _, mentor := range mentors {
		var existing entity.Mentor
		result := db.Where("nama = ?", mentor.Nama).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&mentor)
		}
	}

}

func SeedKelas(db *gorm.DB) {
	kelass := []entity.Kelas{
		{
			Id:               "COURSEB0001",
			Nama:             "Introduction to Business Intelligence",
			Deskripsi:        "Kursus dasar tentang Business Intelligence (BI) yang mengajarkan cara menganalisis data bisnis untuk pengambilan keputusan yang lebih baik.",
			Kategori:         "Business",
			JumlahMateri:     3,
			CoverCourse:      "https://rshdseakqgwspflewctn.supabase.co/storage/v1/object/sign/innovaspace-userprofile/cover-course/innovaspace-course1.jpg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJpbm5vdmFzcGFjZS11c2VycHJvZmlsZS9jb3Zlci1jb3Vyc2UvaW5ub3Zhc3BhY2UtY291cnNlMS5qcGciLCJpYXQiOjE3NDE3MDQwNTAsImV4cCI6MjA1NzA2NDA1MH0.wM3eU3oHBUY06P6BtX6RNdVaFyG6OF2anejTQoLxAs8",
			TingkatKesulitan: "Beginner",
			Durasi:           5,
			CreatedDate:      time.Now(),
			ModifiedDate:     time.Now(),
		},
		{
			Id:               "COURSEB0002",
			Nama:             "Advanced Data Analytics",
			Deskripsi:        "Kursus lanjutan tentang Data Analytics untuk bisnis, dengan fokus pada teknik analisis data yang mendalam dan aplikasinya dalam meningkatkan performa bisnis.",
			Kategori:         "Business",
			JumlahMateri:     3,
			CoverCourse:      "https://rshdseakqgwspflewctn.supabase.co/storage/v1/object/sign/innovaspace-userprofile/cover-course/innovaspace-course2.jpg?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cmwiOiJpbm5vdmFzcGFjZS11c2VycHJvZmlsZS9jb3Zlci1jb3Vyc2UvaW5ub3Zhc3BhY2UtY291cnNlMi5qcGciLCJpYXQiOjE3NDE3MDQyNTgsImV4cCI6MjA1NzA2NDI1OH0.IXdAQmlB1f-yXwHZaoQWGkIWqkI6qMEAk923pKjMC5I",
			TingkatKesulitan: "Advanced",
			Durasi:           10,
			CreatedDate:      time.Now(),
			ModifiedDate:     time.Now(),
		},
	}

	for _, kelas := range kelass {
		var existing entity.Kelas
		result := db.Where("nama = ?", kelas.Nama).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&kelas)
		}
	}
}

func SeedMateri(db *gorm.DB) {
	materis := []entity.Materi{
		{
			Id:           uuid.New(),
			KelasId:      "COURSEB0001",
			IsFree:       true,
			Judul:        "Intro to Business Intelligence",
			Deskripsi:    "",
			PathFile:     "https://youtu.be/dn97ux9exbY?si=s9ZHhRSr1MJ37A4v",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:           uuid.New(),
			KelasId:      "COURSEB0001",
			IsFree:       true,
			Judul:        "Introduction to Business Intelligence",
			Deskripsi:    "Pengenalan dasar tentang Business Intelligence.",
			PathFile:     "",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:      uuid.New(),
			KelasId: "COURSEB0001",
			IsFree:  true,
			Judul:   "Study Case",
			Deskripsi: `1. Lotte.com: BI Increases Company Revenue
Lotte.com is the leading internet shopping mall in Korea with 13 million customers.
Challenge: With more than 1 million site visitors daily, company executives wanted to understand why customers abandon shopping carts.`,
			PathFile:     "",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:           uuid.New(),
			KelasId:      "COURSEB0002",
			IsFree:       true,
			Judul:        "Google Advanced Data Analytics Certification",
			Deskripsi:    "",
			PathFile:     "https://youtu.be/gTf2GLEhKgA?si=MP8BD7ICHjlkXTsC",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:           uuid.New(),
			KelasId:      "COURSEB0002",
			IsFree:       true,
			Judul:        "Introduction to Advanced Data Analytics",
			Deskripsi:    "Pengenalan tentang analisis data lanjutan.",
			PathFile:     "",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
		{
			Id:      uuid.New(),
			KelasId: "COURSEB0002",
			IsFree:  true,
			Judul:   "Study case",
			Deskripsi: `GE Renewable Energy: Harnessing AI for Wind Turbine Optimization
			The Challenge
GE Renewable Energy sought to improve the performance and reliability of its wind turbines, aiming to reduce downtime and maintenance costs while maximizing energy output.`,
			PathFile:     "",
			CreatedDate:  time.Now(),
			ModifiedDate: time.Now(),
		},
	}

	for _, materi := range materis {
		var existing entity.Materi
		result := db.Where("kelas_id = ? AND path_file = ?", materi.KelasId, materi.PathFile).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&materi)
		}
	}
}
