package project

var (
	description string = "Microservice that is part of the Giant Swarm App Platform."
	gitSHA             = "n/a"
	name        string = "app-service"
	source      string = "https://github.com/giantswarm/app-service"
	version            = "0.1.0"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
