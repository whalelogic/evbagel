# EvBagel - Go Web Application

A modern web application demonstrating the integration of Go's `html/template` package with Templ components, using Fiber for HTTP routing and Tailwind CSS for styling.

## Features

- **Go html/template**: Traditional HTML templates for page structure
- **Templ Components**: Type-safe, reusable components for header, footer, and buttons
- **Fiber v2**: Fast and efficient web framework
- **Tailwind CSS**: Modern utility-first CSS framework
- **3 Pages**: Home, About, and Contact with full routing

## Project Structure

```
.
├── components/          # Templ components
│   ├── header.templ    # Navigation header component
│   ├── footer.templ    # Footer component
│   └── button.templ    # Button component
├── handlers/           # HTTP request handlers
│   └── handlers.go     # Page handlers for all routes
├── templates/          # HTML templates
│   ├── home.html       # Home page template
│   ├── about.html      # About page template
│   └── contact.html    # Contact page template
├── static/             # Static assets (CSS, JS)
├── main.go             # Application entry point
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

## Prerequisites

- Go 1.24 or higher
- Templ CLI tool

## Installation

1. Clone the repository:
```bash
git clone https://github.com/whalelogic/evbagel.git
cd evbagel
```

2. Install dependencies:
```bash
go mod download
```

3. Install Templ CLI:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Building

Generate Go code from Templ files:
```bash
templ generate
```

Build the application:
```bash
go build -o evbagel .
```

Or use the combined command:
```bash
templ generate && go build -o evbagel .
```

## Running

Start the server:
```bash
./evbagel
```

The application will be available at `http://localhost:3000`

## Pages

- **Home** (`/`) - Welcome page with features overview
- **About** (`/about`) - Information about the project and technology stack
- **Contact** (`/contact`) - Contact form and information

## Development

The project follows a hybrid approach:

1. **HTML Templates** are used for the overall page structure and content
2. **Templ Components** are used for reusable UI elements (header, footer, buttons)
3. Components are rendered to strings and injected into templates
4. Tailwind CSS is loaded via CDN for styling

This approach combines the simplicity of Go's standard `html/template` package with the type safety and reusability of Templ components.

## Technologies

- **Backend**: Go 1.24+, Fiber v2
- **Components**: Templ
- **Styling**: Tailwind CSS
- **Templating**: Go html/template

## License

See LICENSE file for details.
