# Database schema migrations

This folder will contain ordered PostgreSQL schema changes for **Phase 4**.
Migrations describe how the schema evolves; database access code belongs in
[storage](../internals/storage/).

Planned contents:

- Numbered SQL files for repository, scan, finding, and rule-metadata tables.
- Primary keys, foreign keys, constraints, and indexes justified by actual queries.
- Migration instructions for a fresh database and an existing installation.
- Rollback files where supported and safe, following the migration tool selected
  during Phase 4.

For example, a tool that supports paired migrations might use
`000001_initial_schema.up.sql` and `000001_initial_schema.down.sql`. This naming
is illustrative; no migration tool or schema has been chosen yet.

Keep credentials, database dumps, and real scan data out of this folder. Test
migrations against a disposable database before applying them to retained data.

**Current phase:** planned only. Do not design the full database before the
parser, rules, and scan result structures are understood.
