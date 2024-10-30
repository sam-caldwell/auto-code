package arguments

// SourceCount - Return the number of data sources
func (p *ParameterDefinition) SourceCount() int {
	return p.sources.Count()
}
