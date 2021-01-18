package utils

import "github.com/MathisBurger/apache2-automatisation/config"

func CheckCORS(ip string) bool {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}
	return cfg.AllowedOrigins == ip
}
