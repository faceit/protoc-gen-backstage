# protoc-gen-backstage
> Warning API is very subject to change before v1.0.0

This is a protoc plugin to allow you to generate backstage API kinds from your proto definitions


## Usage 

### Owner Field

The `spec.owner`  field in backstage is required, so we have provided a service option extension to be used in your proto definitions

### System Field

The `system` field in backstage is optional but is provided by this package so that it can populated from a proto service


## Plugin Options

The plugin takes options in teh form of a comma separated string the options are in the following order

1. The repo root url e.g. if you kept your protos in  a repository called "protos" and a directory "team-1" https://github.com/organisation/protos/blob/master/team-1 - This field lets us link back to the protos from backstage
2. How you name your catalog-info.yaml files - controls the filenames that are output by the plugin
3. A regex string to capture parts of the protos fullname. You could use this to truncate the fullname to service name
