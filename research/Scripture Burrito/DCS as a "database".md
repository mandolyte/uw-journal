# DCS as a "database"

If you think of DCS as a database, then the data stored in the database has a limited number of data types along with keys and content which are used to query (filter, join, etc.) and to update.

Our data types are:
- USFM
- Markdown
- TSV
- YAML

Our keys are:
- Organization
- Language Id
- Resource Type Id
- Book Id (depending on resource type)
- Filepath

Since DCS isn't a real database, much of the filtering and joining must be in client code. This means that keeping the set of data types used to a minimum is important to the tooling. 