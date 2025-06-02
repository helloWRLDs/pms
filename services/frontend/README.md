# Frontend Service

The Frontend Service is the user interface component of the PMS (Project Management System), built with React, TypeScript, and modern web technologies. It provides an intuitive and responsive interface for managing projects, tasks, and team collaboration.

## Features

- **Project Management**

  - Project creation and configuration
  - Task management with drag-and-drop
  - Sprint planning and tracking
  - Team member assignment
  - Document sharing and collaboration

- **User Interface**

  - Modern, design
  - Real-time updates
  - Interactive dashboards
  - Rich text editing
  - Drag-and-drop functionality

- **Authentication & Authorization**

  - Secure login/logout
  - Session management
  - OAuth2 integration

- **Data Visualization**
  - Project progress charts
  - Team performance metrics
  - Task completion statistics
  - Custom reports and analytics

## Tech Stack

- **Core**

  - React 18
  - TypeScript
  - Vite
  - React Router v7
  - TanStack Query (React Query)

- **UI Components**

  - Tailwind CSS
  - Headless UI
  - Lucide React Icons
  - React Hot Toast
  - TipTap Editor

- **State Management**

  - Zustand
  - React Query

- **Development Tools**
  - ESLint
  - TypeScript
  - Vite
  - PostCSS
  - Tailwind CSS

## Prerequisites

- Node.js 18 or higher
- npm 9 or higher
- Modern web browser

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```
3. Create a `.env` file in the root directory:
   ```env
   VITE_API_URL=http://localhost:50051
   VITE_AUTH_URL=http://localhost:50052
   VITE_NOTIFIER_URL=http://localhost:50054
   ```
4. Start the development server:
   ```bash
   npm run dev
   ```

## Development

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run lint` - Run ESLint
- `npm run preview` - Preview production build

### Project Structure

```
src/
├── components/     # Reusable UI components
├── features/       # Feature-specific components
├── hooks/         # Custom React hooks
├── layouts/       # Page layouts
├── pages/         # Page components
├── services/      # API services
├── store/         # State management
├── types/         # TypeScript types
└── utils/         # Utility functions
```

### Component Guidelines

1. Use functional components with TypeScript
2. Implement proper prop typing
3. Follow atomic design principles
4. Use Tailwind CSS for styling
5. Implement responsive design
6. Add proper error handling
7. Include loading states

### State Management

- Use Zustand for global state
- Use React Query for server state
- Implement proper caching strategies
- Handle loading and error states

## Building for Production

1. Build the application:
   ```bash
   npm run build
   ```
2. Preview the build:
   ```bash
   npm run preview
   ```
3. Deploy the `dist` directory to your web server
