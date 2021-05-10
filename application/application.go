package application

func Load(config Config) (*Application, error) {
	return &Application{}, nil
}

type Application struct {
}

func (a *Application) Run() <-chan error {
	return nil
}

func (a *Application) Shutdown() {
}
