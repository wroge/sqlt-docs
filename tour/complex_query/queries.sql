{{ define "schema" }}
        CREATE TABLE IF NOT EXISTS books (
                id INTEGER PRIMARY KEY,
                title TEXT NOT NULL,
                added_at DATE NOT NULL
        );

        CREATE TABLE IF NOT EXISTS authors (
                id INTEGER PRIMARY KEY,
                name TEXT UNIQUE
        );

        CREATE TABLE IF NOT EXISTS book_authors (
                book_id INTEGER REFERENCES books(id),
                author_id INTEGER REFERENCES authors(id),
                PRIMARY KEY (book_id, author_id)
        );
{{ end }}

{{ define "upsert_authors" }}
        INSERT INTO authors (name) VALUES 
        {{ range $i, $a := .Authors }}
                {{ if $i }}, {{ end }}
                ({{ $a }})
        {{ end }}
        ON CONFLICT (name) DO UPDATE SET 
                id = authors.id,
                name = EXCLUDED.name
        RETURNING id;
{{ end }}

{{ define "insert_book" }}
        INSERT INTO books (title, added_at) VALUES ({{ .Title }}, {{ now }}) RETURNING id;
{{ end }}

{{ define "link_book_authors" }}
        {{ $book_id := Context "insert_book" }}
        INSERT INTO book_authors (book_id, author_id) VALUES
        {{ range $i, $a := Context "upsert_authors" }}
                {{ if $i }}, {{ end }}
                ({{ $book_id }}, {{ $a }})
        {{ end }}
        ON CONFLICT DO NOTHING;
{{ end }}

{{ define "query_books" }}
        SELECT 
                books.id                        {{ Scan "ID" }}
                , books.title                   {{ Scan "Title" }}
                , GROUP_CONCAT(authors.name)    {{ ScanStringSlice "Authors" "," }}
                , books.added_at                {{ Scan "AddedAt" }}
        FROM books
        LEFT JOIN book_authors ON books.id = book_authors.book_id
        LEFT JOIN authors ON authors.id = book_authors.author_id
        WHERE 1=1
        {{ with .Title }}
                AND books.title = {{ . }}
        {{ end }}
        {{ with .Author }}
                AND books.id IN (
                        SELECT ba.book_id
                        FROM book_authors ba
                        JOIN authors a ON a.id = ba.author_id
                        WHERE a.name = {{ . }}
                )
        {{ end }}
        {{ if not .AddedBefore.IsZero }}
                AND books.added_at < {{ .AddedBefore }}
        {{ end }}
        GROUP BY books.id, books.title, books.added_at;
{{ end }}