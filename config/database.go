package config

var Database *DatabaseConfig

type DatabaseConfig struct {
	MasterDB DBConfig
	SlaveDB  DBConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string
}

func initDatabase() {
	Database = &DatabaseConfig{
		MasterDB: DBConfig{
			DBHost:     v.GetString("MASTER_DB_HOST"),
			DBPort:     v.GetInt("MASTER_DB_PORT"),
			DBDatabase: v.GetString("MASTER_DB_DATABASE"),
			DBUsername: v.GetString("MASTER_DB_USERNAME"),
			DBPassword: v.GetString("MASTER_DB_PASSWORD"),
		},
		SlaveDB: DBConfig{
			DBHost:     v.GetString("SLAVE_DB_HOST"),
			DBPort:     v.GetInt("SLAVE_DB_PORT"),
			DBDatabase: v.GetString("SLAVE_DB_DATABASE"),
			DBUsername: v.GetString("SLAVE_DB_USERNAME"),
			DBPassword: v.GetString("SLAVE_DB_PASSWORD"),
		},
	}
}
