package site

type accessMapping struct {
	AccessName string
	Group      string
}

func newAccessMapping(accessName, group string) accessMapping {
	return accessMapping{
		AccessName: accessName,
		Group:      group,
	}
}

type accessMap map[string]accessMapping
