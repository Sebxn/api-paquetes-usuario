package config

import (
	"fmt"
)

const (
	// Datos de conexión a la base de datos PostgreSQL
	DBHost     = "34.72.243.84" // Cambia a la dirección IP o el nombre de host de tu servidor PostgreSQL
	DBPort     = 5432           // Cambia al puerto de tu servidor PostgreSQL si es diferente
	DBUser     = "postgres"     // Cambia a tu nombre de usuario de PostgreSQL
	DBPassword = "123456"       // Cambia a tu contraseña de PostgreSQL
	DBName     = "bdtest"       // Cambia al nombre de tu base de datos PostgreSQL
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}
