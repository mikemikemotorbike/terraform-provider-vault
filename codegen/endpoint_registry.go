package codegen

// endpointRegistry is a registry of all the endpoints we'd
// like to have generated, along with the type of template
// we should use.
var endpointRegistry = map[string]templateType{
	"/transform/role/{name}": templateTypeResource,
	// TODO this will eventually list all endpoints and data sources we want to generate.
}
