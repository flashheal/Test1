# Online Record Label

This project is an online record label application built using Go for the backend and SvelteKit for the frontend. It aims to provide a platform for artists to showcase their music and for users to discover new talent.

## Project Structure

```
online-record-label
├── backend
│   ├── main.go          # Entry point for the Go backend application
│   ├── go.mod           # Module dependencies for the Go application
│   └── go.sum           # Checksums for module dependencies
├── frontend
│   ├── src
│   │   ├── routes
│   │   │   └── index.svelte  # Main page of the frontend application
│   │   ├── lib
│   │   │   └── styles.css     # CSS styles for the frontend application
│   │   └── app.html           # Main HTML template for the Svelte application
│   ├── svelte.config.js        # Configuration for Svelte
│   ├── tailwind.config.js      # Configuration for Tailwind CSS
│   ├── package.json            # npm configuration file
│   ├── postcss.config.js       # Configuration for PostCSS
│   └── vite.config.js          # Configuration for Vite
└── README.md                   # Documentation for the project
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Node.js (version 14 or later)
- npm (Node package manager)

### Backend Setup

1. Navigate to the `backend` directory:
   ```
   cd backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the Go application:
   ```
   go run main.go
   ```

### Frontend Setup

1. Navigate to the `frontend` directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Start the development server:
   ```
   npm run dev
   ```

### Features

- Artist profiles
- Music discovery
- User-friendly interface
- Responsive design using Tailwind CSS

### Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

### License

This project is licensed under the MIT License. See the LICENSE file for more details.