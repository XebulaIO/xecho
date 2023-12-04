package xecho

var (
	scopeMap = map[Scope]any{}
)

type (
	Scope         string
	Scopes        []Scope
	scopeRegistry struct{}
)

var (
	r = &scopeRegistry{}
)

func (r *scopeRegistry) Register(scope ...Stringer) *scopeRegistry {
	for _, s := range scope {
		scopeMap[Scope(s.String())] = struct{}{}
	}
	return r
}

func (s Scope) IsValid() bool {
	_, ok := scopeMap[s]
	return ok
}

func (s Scope) String() string {
	return string(s)
}

func ScopesFrom(s []string) Scopes {
	scopes := make(Scopes, len(s))
	for i, scope := range s {
		scopes[i] = Scope(scope)
	}

	return scopes
}

func (s Scopes) Validate() bool {
	for _, scope := range s {
		if !scope.IsValid() {
			return false
		}
	}

	return true
}

func ScopeRegistry() *scopeRegistry {
	return r
}
