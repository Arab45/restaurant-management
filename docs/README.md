# API Docs

This folder contains a minimal OpenAPI spec and an HTML viewer.

To view the docs:

- Serve the `docs` folder over HTTP (e.g., `python -m http.server 8000` or any static file server).
- Open `http://localhost:8000/swagger.html` in your browser.

Notes:
- `swagger.json` is a sample OpenAPI 3.0 file; expand it to include other endpoints (menus, orders, tables, users, etc.).
- Alternatively, integrate with `swaggo/swag` to auto-generate docs from code if you prefer annotations in Go.
