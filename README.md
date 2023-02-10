# protoc-gen-backstage
> Warning API is very subject to change before v1.0.0

This is a protoc plugin to allow you to generate backstage API kinds from your proto definitions


## Usage 

### Owner Field

The `spec.owner`  field in backstage is required, so we have provided a service option extension to be used in your proto definitions

### System Field

The `system` field in backstage is optional but is provided by this package so that it can populated from a proto service
