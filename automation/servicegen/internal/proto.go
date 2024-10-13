package internal

type Service struct {
	Name string   `yaml:"name"`
	Rpcs []string `yaml:"rpcs"`
}
type Template struct {
	Pkg     string  `yaml:"package"`
	Service Service `yaml:"service"`
}

func (p *Template) validate() {
	if p.Pkg == "" {
		panic("package name is required")
	}
	if p.Service.Name == "" {
		panic("service name is required")
	}
	if !isSnakeCase(p.Service.Name) {
		panic("service name should be in camelcase")
	}
	if len(p.Service.Rpcs) == 0 {
		panic("at least one rpc is required")
	}
}

func (p *Template) GenProtoFile() string {
	content := `syntax = "proto3";` + "\n"
	content += "package " + p.Pkg + ";" + "\n"
	content += "option go_package = " + "\"" + p.Pkg + "/proto" + "\"" + ";" + "\n"
	content += p.Service.genServiceStr()
	return content
}

func (s *Service) genServiceStr() string {
	var msgs []string
	content := "service " + snakeToPascal(s.Name) + "{" + "\n"
	for _, rpc := range s.Rpcs {
		cap := capitalize(rpc)
		content += "\t" + "rpc " + cap + " (" + cap + "Request" + ")" + " returns (" + cap + "Response" + ")" + "{}" + "\n"
		msgs = append(msgs, cap)
	}
	content += "}\n"
	for _, msg := range msgs {
		content += boilerMessage(msg, "Request")
		content += boilerMessage(msg, "Response")
	}
	return content
}

func boilerMessage(str, suffix string) string {
	return "message " + capitalize(str) + suffix + " {" + "\n" + "}" + "\n"
}
