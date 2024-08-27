package app

type ClientAPPConfig struct {
}

type ClientAPP struct {
	cfg *ClientAPPConfig
}

func NewClientAPP(cfg *ClientAPPConfig) *ClientAPP {
	return &ClientAPP{
		cfg: cfg,
	}
}

func (app *ClientAPP) Build() error {

	return nil
}

func (app *ClientAPP) Run() error {

	return nil
}
