package factory

// 1. 简单工厂模式
// IRuleConfigParser IRuleConfigPraser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser json config parser
type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
