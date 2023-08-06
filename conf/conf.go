package conf

import (
	"fmt"
)

func SettingPostgresDsn() (dsn string) {
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		"127.0.0.1", // host
		"giltaeho",  // user
		"",          // password
		"giltaeho",  // dbname
		"5432",      // port
	)
	return dsn
}
func FrontDomain() (domain string) {
	return "http://localhost:3000"
}

func RunBackDomain() (domain string) {
	return "127.0.0.1:8000"
}
func BuildBackendDomain() (domain string) {
	return "http://127.0.0.1:8000/v1"
}
