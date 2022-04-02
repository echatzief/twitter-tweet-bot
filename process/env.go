package process

type EnviromentVariables struct {
	CONSUMER_KEY        string
	CONSUMER_SECRET     string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
}

func Variables() EnviromentVariables {
	return EnviromentVariables{
		CONSUMER_KEY:        "....",
		CONSUMER_SECRET:     "....",
		ACCESS_TOKEN:        "....",
		ACCESS_TOKEN_SECRET: "....",
	}
}
