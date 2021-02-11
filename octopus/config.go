package octopus

type Config struct {
	Servers []Server `yaml:"octopus.servers"`
}
type Channels struct {
	Dev  string `yaml:"dev"`
	Prod string `yaml:"prod"`
}
type Envs struct {
	Test string `yaml:"test"`
}
type Server struct {
	Name      string   `yaml:"name"`
	URL      string   `yaml:"url"`
	Key      string   `yaml:"key"`
	Channels Channels `yaml:"channels"`
	Envs Envs `yaml:"env"`
}