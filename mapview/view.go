package mapview

type MapView struct {
	IncludedPaths []string
	ExcludedPaths []string
	RequiredPaths []string
	TypeHandlers map[typ][]func(interface{}) (interface{}, err)
	PathHandlers map[string][]func(interface{}) (interface{}, err)
	// AllowedTypes []interface
}

func NewMapView() MapView {
	return MapView{}
}

func (v *MapView) IncludePaths (paths []string) MapView {
	v.IncludedPaths = paths
	return v
}

// TODO: Add some chainable methods for easier config/creation