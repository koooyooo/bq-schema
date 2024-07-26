## Usage
### Settings
Define following environment variables in `.env` file or set variables directly. Renaming `.env.template` to `.env` and define your original settings could be useful.
- `GOOGLE_APPLICATION_CREDENTIALS`
- `GOOGLE_CLOUD_PROJECT`
- `BIGQUERY_DATASET`

> [optional]  
> you could define exclude tables by `::` separated regular expression patterns
> - `EXCLUDE_TABLE_PATTERNS`

### Run
```bash
$ go run bq-schema
```

then, it outputs `schemas.puml` in `target/` directory.

> I strongly recommend not defining relations in `schemas.puml` directly.
> Instead, I recommend creating another `.puml` file and include `schemas.puml` in it. 
> Then defining relations in it. 
> This is useful in case of re-generating `schemas.puml` because it won't override and delete your added relations defined in another `.puml` file.

e.g. `target/my-custom.puml`
```plantuml
@startuml
!include schemas.puml

# YOUR RELATION DEFINITIONS

@enduml
```

## Trouble Shoot
### PlantUML Env in Mac 
If you cannot view PlantUML model. May be you need to install Java and related Tools.

```bash
# Install Java
- $ brew install Java
- $ echo 'export PATH="/opt/homebrew/opt/openjdk/bin:$PATH"' >> ~/.zshrc
- $ export CPPFLAGS="-I/opt/homebrew/opt/openjdk/include"

# Install graphviz
- $ brew install graphviz
```

## See
- [PlantUML-ER grammer](https://plantuml.com/ie-diagram)
