package factory

/*
工厂方法
当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，
做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，
让每个工厂类都不至于过于复杂
*/

// IRUleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

// CreateParser create yaml config parser
func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &yamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory json config parser
type jsonRuleConfigParserFactory struct {
}

// CreateParser create json config parser
func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

func NewRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
