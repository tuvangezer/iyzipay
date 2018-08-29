package iyzipay

import "strconv"

type pkiBuilder struct {
	rString string
}

func (p *pkiBuilder) append(k, v string) *pkiBuilder {
	if v != "" {
		p.appendKeyValue(k, v)
	}
	return p
}

func (p *pkiBuilder) appendKeyValue(k, v string) *pkiBuilder {
	if v != "" {
		p.rString += k + "=" + v + ","
	}
	return p
}

// TODO: A need for checking if v exists might need to be implemented.
func (p *pkiBuilder) appendPrice(k string, v string) *pkiBuilder {
	if v != "" {
		p.appendKeyValue(k, sanitizePrice(v))
	}
	return p
}

func (p *pkiBuilder) appendArray(k string, arr []string) *pkiBuilder {
	if arr != nil {
		val := ""
		for _, v := range arr {
			val += v + ", "
		}
		p.appendKeyValueArray(k, val)
	}
	return p
}
func (p *pkiBuilder) appendIntArray(k string, arr []int) *pkiBuilder {
	if arr != nil {
		val := ""
		for _, v := range arr {
			val += strconv.Itoa(v) + ", "
		}
		p.appendKeyValueArray(k, val)
	}
	return p
}
func (p *pkiBuilder) appendKeyValueArray(k, v string) *pkiBuilder {
	if v != "" {
		v = v[:len(v)-2]
		p.rString += k + "=[" + v + "],"
	}
	return p
}
func (p *pkiBuilder) removeTrailingComma() *pkiBuilder {
	p.rString = p.rString[:len(p.rString)-1]
	return p
}

func (p *pkiBuilder) appendPrefix() *pkiBuilder {
	p.rString = "[" + p.rString + "]"
	return p
}

func (p *pkiBuilder) GetReqString() string {
	p.removeTrailingComma()
	p.appendPrefix()
	return p.rString
}
func newResourcePKI(locale, convId string) *pkiBuilder {
	p := pkiBuilder{}
	p.append("locale",locale)
	p.append("conversationId", convId)
	return &p
}
