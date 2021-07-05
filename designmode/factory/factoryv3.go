package factory

// 抽象工厂
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser json system config parser
type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

// jsonConfigParserFactory
type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
