# Catalog Next Initialization

This issue covers the initial setup of the package plus the Catalog Next (CN) initialization feature.

DoD: 

- [ ] Repo catalog-next-toolkit created
- [ ] Tag search demo
- [ ] Catalog initialization demo
- [ ] Unit Tests

## Overview
Create a headless component library that integrates CN with DCS.

The features of this CN component are:

1. For production users, restrict application access to production releases for our resources.
2. For developers, allow specification of the CN release to use for resources
3. For improved performance (and reduced network and server resources), provide automatic management of caching based on resource releases. A description of the caching features is below.

## Proposal - Details

In the design below there are two parameters that are used extensively but will not be provided by the application:
1. SERVER URL: 
2. API VERSION:
 
These values will be hard coded in the environment file (`.env`) and will be made available via Node.js process access APIs.
The default for the server will be `https://qa.door43.org`. When deployed, this needs to be provided by Netlify and should point to production.

The API Version for the catalog is `v5` and will be specific fully as `api/catalog/v5`. At present this the only version needed and should also be provided by Netlify.

### Tag Search

A function will allow an app to specify an owner, language, and resource and it will return a list of production tags. This list can be used to allow the user to select the release they want to use of the resource.

### Initialization

The initialization parameters are:

- OWNER: this is the owner of the resources. The default is 'unfoldingWord'. 
- LANGID: this provides the language to be used in identifying resources. Default will be `en`.
- RESOURCES: this is a comma delimited list of resource IDs. The complete list is:
'ta', 'tw', 'twl', 'tn', 'tq', 'obs', 'obs-tq', 'obs-tn', 'obs-sn', 'obs-sq'. *Issue: how to specify Original Language resources (Greek and Hebrew texts)*.
- TAG: this is the value for CN component which allows retrieval of the exact resource released desired.

Questions:
1. Should the Original Languages always use the highest prod release? (I think "yes")
2. Should the tag parameter be optional and, if not specified, default to use the highest tag? (I think "no", but see below for implications)
3. Should this API enforce the entanglement of language and resource? (I think "yes"). In other words:
	- if language is specified first, only show resources which exist with that language
	- if resource is specified first, only show languages which exist with that resource
	- this devolves to having two API searches, one for each way

*On the tag being optional:* I doubt there is a use case for a GL team to prefer the highest production release. Instead, they want to point to a release and stick with it until they intentionally want to switch. The upshot of this is that this component can never just go off and pre-load the resources. The app must always direct which resources are to be used.
This simplifies the API some at the cost of complicating the UI/X.


### API

The CN component will provide a set of functions to access the content of a resource.

Given a server, owner, language, resource type, and tag, then `initializeCatalog()`: 
	- will clear the "unzipped catalog"; this storage contains items already extracted from the zip. 
	- if the tagged release zip file is already present, then done. Otherwise, continue...
	- If not local, then it must fetch the "resource zip" from CN and store it. See Example #2
