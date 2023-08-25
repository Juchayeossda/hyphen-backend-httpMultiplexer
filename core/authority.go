package core

type Authority struct {
	authorities map[string]string
}

type AuthoritySet struct {
	Name          string
	DestAuthority string
}

func NewAuthority(authoritySets []AuthoritySet) *Authority {
	authority := &Authority{make(map[string]string)}

	for _, authoritySet := range authoritySets {
		authority.authorities[authoritySet.Name] = authoritySet.DestAuthority
	}

	return authority
}
func (a *Authority) Get(name string) string {
	return a.authorities[name]
}
