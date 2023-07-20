package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env        *Env
	Postgresql *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	postgresqlDB, err := NewPostgresDatabase(app.Env)
	if err != nil {
		panic(err)
	}
	app.Postgresql = postgresqlDB
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgresql)
}
