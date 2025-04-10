{{ define "schema" }}
        CREATE TABLE IF NOT EXISTS books (
                id INTEGER PRIMARY KEY,
                title TEXT NOT NULL,
                author TEXT NOT NULL
        )
{{ end }}

{{ define "insert_book" }}
        INSERT INTO books (title, author) VALUES ({{ .Title }}, {{ .Author }}) RETURNING id;
{{ end }}

{{ define "insert_books" }}
        INSERT INTO books (title, author) VALUES 
        {{ range $i, $c := . }}
                {{ if $i }}, {{ end }}
                ({{ $c.Title }}, {{ $c.Author }})
        {{ end }}
        RETURNING id;
{{ end }}

{{ define "get_book" }}
        SELECT 
                id              {{ Scan "ID" }}
                , title         {{ Scan "Title" }}
                , author        {{ Scan "Author" }}
        FROM books
        WHERE id = {{ . }};
{{ end }}