
<h1>EvBagel 🥯

![image](https://img.icons8.com/?size=100&id=u0SDdwgQXX72&format=png&color=000000)


### This application works fast! 

#### Feel free to fork this repo and use it as your starting point. 

## Start here

```bash
# Install dependencies
go mod download

# Install Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Build and run
make run
```

The application will be available at `http://localhost:3000`

## Features

- 🚀 **Fiber v2** - Fast and efficient web framework
- 📝 **Go html/template** - Traditional HTML templating
- 🎨 **Templ Components** - Type-safe, reusable UI components
- 💅 **Tailwind CSS** - Modern utility-first styling
- 🔧 **Makefile** - Easy build and run commands

## Pages

- `/` - Home page with features overview
- `/about` - About page with technology stack information
- `/contact` - Contact page with form

## Documentation

See [README_PROJECT.md](README_PROJECT.md) for detailed documentation.

## Project Structure

```
├── components/       # Templ components (header, footer, button)
├── handlers/        # HTTP request handlers
├── templates/       # HTML templates
├── static/          # Static assets
└── main.go         # Application entry point
```

## Building

```bash
make build    # Build the application
make run      # Build and run
make clean    # Clean build artifacts
make help     # Show all commands
```

## License

See LICENSE file for details.
