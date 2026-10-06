package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel: Menyediakan ID, CreatedAt, UpdatedAt, dan Soft Delete untuk semua tabel
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// USer dan role
type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
}

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"role"`
}

// Data Master Akademik
type Prodi struct {
	BaseModel
	Kode string `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama string `gorm:"type:varchar(100);not null" json:"nama"`
}

type TahunAkademik struct {
	BaseModel
	Kode     string `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama     string `gorm:"type:varchar(100);not null" json:"nama"`
	IsActive bool   `gorm:"default:false" json:"is_active"`
}

type Ruangan struct {
	BaseModel
	Kode      string `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama      string `gorm:"type:varchar(100);not null" json:"nama"`
	Kapasitas int    `gorm:"not null" json:"kapasitas"`
}

// Entitas utama Dosen & mahasiswa
type Dosen struct {
	BaseModel
	UserID  uint   `gorm:"uniqueIndex" json:"user_id"` //1 user = 1 dosen
	User    User   `gorm:"foreignKey:UserID"`
	NIDN    string `gorm:"type:varchar(20);uniqueIndex;not null" json:"nidn"`
	Nama    string `gorm:"type:varchar(150);not null" json:"nama"`
	ProdiID uint   `json:"prodi_id"`
	Prodi   Prodi  `gorm:"foreignKey:ProdiID"`

	// Tambahan opsional: Agar mudah menarik data mahasiswa yang dibimbingnya (Has Many)
	MahasiswaBimbingan []Mahasiswa `gorm:"foreignKey:DosenPAID" json:"mahasiswa_bimbingan,omitempty"`
}

type Mahasiswa struct {
	BaseModel
	UserID  uint   `gorm:"uniqueIndex" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	NPM     string `gorm:"type:varchar(20);uniqueIndex;not null" json:"npm"`
	Nama    string `gorm:"type:varchar(150);not null" json:"nama"`
	ProdiID uint   `json:"prodi_id"`
	Prodi   Prodi  `gorm:"foreignKey:ProdiID"`
	Kelas   string `gorm:"type:varchar(20)" json:"kelas"`

	//Tambahan untuk dosen Pembimbing akademik
	DosenPAID *uint  `json:"dosen_pa_id"`
	DosenPA   *Dosen `gorm:"foreignKey:DosenPAID" json:"dosen_pa,omitempty"`
}

// Perkuliahan dan Jadwal
type MataKuliah struct {
	BaseModel
	Kode     string `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama     string `gorm:"type:varchar(150);not null" json:"nama"`
	SKS      int    `gorm:"not null" json:"sks"`
	Semester int    `gorm:"not null" json:"semester"`
	ProdiID  uint   `json:"prodi_id"`
	Prodi    Prodi  `gorm:"foreignKey:ProdiID"`
}

type Kelas struct {
	BaseModel
	Nama            string        `gorm:"type:varchar(50);not null" json:"nama"`
	TahunAkademikID uint          `json:"tahun_akademik_id"`
	TahunAkademik   TahunAkademik `gorm:"foreignKey:TahunAkademikID"`
	ProdiID         uint          `json:"prodi_id"`
	Prodi           Prodi         `gorm:"foreignKey:ProdiID"`
}

type Jadwal struct {
	BaseModel
	TahunAkademikID uint          `json:"tahun_akademik_id"`
	TahunAkademik   TahunAkademik `gorm:"foreignKey:TahunAkademikID" json:"tahun_akademik"`
	MataKuliahID    uint          `json:"mata_kuliah_id"`
	MataKuliah      MataKuliah    `gorm:"foreignKey:MataKuliahID"`
	KelasID         uint          `json:"kelas_id"`
	Kelas           Kelas         `gorm:"foreignKey:KelasID"`
	DosenID         uint          `json:"dosen_id"`
	Dosen           Dosen         `gorm:"foreignKey:DosenID"`
	RuanganID       uint          `json:"ruangan_id"`
	Ruangan         Ruangan       `gorm:"foreignKey:RuanganID"`
	Hari            string        `gorm:"type:varchar(20);not null" json:"hari"`
	JamMulai        string        `gorm:"type:varchar(5);not null" json:"jam_mulai"`
	JamSelesai      string        `gorm:"type:varchar(5);not null" json:"jam_selesai"`
}

// KRS - Kartu Rencana Studi
type KRS struct {
	BaseModel
	MahasiswaID     uint          `gorm:"uniqueIndex:idx_mhs_ta" json:"mahasiswa_id"`
	Mahasiswa       Mahasiswa     `gorm:"foreignKey:MahasiswaID"`
	TahunAkademikID uint          `gorm:"uniqueIndex:idx_mhs_ta" json:"tahun_akademik_id"`
	TahunAkademik   TahunAkademik `gorm:"foreignKey:TahunAkademikID"`
	Status          string        `gorm:"type:varchar(20);default:'draft'" json:"status"` //draft, diajukan, disetujui, ditolak
	Catatan         string        `gorm:"type:text" json:"catatan"`
	TotalSKS        int           `gorm:"default:0" json:"total_sks"`
	KRSDetail       []KRSDetail   `gorm:"foreignKey:KRSID" json:"krs_details,omitempty"`
}

type KRSDetail struct {
	BaseModel
	KRSID      uint    `gorm:"uniqueIndex:idx_krs_jadwal" json:"krs_id"`
	KRS        KRS     `gorm:"foreignKey:KRSID"`
	JadwalID   uint    `gorm:"uniqueIndex:idx_krs_jadwal" json:"jadwal_id"`
	Jadwal     Jadwal  `gorm:"foreignKey:JadwalID"`
	NilaiTugas float64 `json:"nilai_tugas"`
	NilaiUTS   float64 `json:"nilai_uts"`
	NilaiUAS   float64 `json:"nilai_uas"`
	NilaiAkhir float64 `json:"nilai_akhir"`
	NilaiHuruf string  `gorm:"type:varchar(2)" json:"nilai_huruf"`
	BobotNilai float64 `json:"bobot_nilai"`
}

// Presensi
type Presensi struct {
	BaseModel
	JadwalID    uint             `json:"jadwal_id"`
	Jadwal      Jadwal           `gorm:"foreignKey:JadwalID"`
	Tanggal     time.Time        `gorm:"type:date;not null" json:"tanggal"`
	PertemuanKe int              `gorm:"not null" json:"pertemuan_ke"`
	Materi      string           `gorm:"type:text" json:"materi"`
	Status      string           `gorm:"type:varchar(20);default:'open'" json:"status"` //open, closed
	Details     []PresensiDetail `gorm:"foreignKey:PresensiID" json:"details,omitempty"`
}

type PresensiDetail struct {
	BaseModel
	PresensiID  uint      `gorm:"uniqueIndex:idx_presensi_mhs" json:"presensi_id"`
	MahasiswaID uint      `gorm:"uniqueIndex:idx_presensi_mhs" json:"mahasiswa_id"`
	Mahasiswa   Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	Status      string    `gorm:"type:varchar(20);not null" json:"status"`
}

// Nilai
type Nilai struct {
	BaseModel
	KRSDetailID uint      `gorm:"uniqueIndex;not null" json:"krs_detail_id"` // 1 KRS Detail = 1 Nilai Akhir
	KRSDetail   KRSDetail `gorm:"foreignKey:KRSDetailID"`
	Tugas       float64   `gorm:"default:0" json:"tugas"`
	UTS         float64   `gorm:"default:0" json:"uts"`
	UAS         float64   `gorm:"default:0" json:"uas"`
	Kehadiran   float64   `gorm:"default:0" json:"kehadiran"`
	NilaiAkhir  float64   `gorm:"default:0" json:"nilai_akhir"`
	NilaiHuruf  string    `gorm:"type:varchar(2)" json:"nilai_huruf"`
	Bobot       float64   `gorm:"default:0" json:"bobot"`
	IsLocked    bool      `gorm:"default:false" json:"is_locked"`
}

type Pengumuman struct {
	BaseModel
	Judul     string `gorm:"type:varchar(255);not null" json:"judul"`
	Isi       string `gorm:"type:text;not null" json:"isi"`
	PenulisID uint   `json:"penulis_id"`
	Penulis   User   `gorm:"foreignKey:PenulisID"`
}

// Model untuk mencatat sesi kelas (Dosen mengajar)
type Pertemuan struct {
	BaseModel
	JadwalID    uint   `json:"jadwal_id"`
	Jadwal      Jadwal `gorm:"foreignKey:JadwalID"`
	PertemuanKe int    `json:"pertemuan_ke"`             // Contoh: Pertemuan 1, 2, 3...
	Tanggal     string `gorm:"type:date" json:"tanggal"` // Format: YYYY-MM-DD
	Materi      string `gorm:"type:text" json:"materi"`  // Materi yang diajarkan hari itu
}

// Model untuk mencatat kehadiran tiap mahasiswa di sesi tersebut
type Absensi struct {
	BaseModel
	PertemuanID uint      `json:"pertemuan_id"`
	Pertemuan   Pertemuan `gorm:"foreignKey:PertemuanID"`
	MahasiswaID uint      `json:"mahasiswa_id"`
	Mahasiswa   Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	Status      string    `gorm:"type:varchar(10)" json:"status"` // Hadir, Izin, Sakit, Alpa
}
