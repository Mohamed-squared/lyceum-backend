# Lyceum - Architecture Overview

## 1. Guiding Principles

This document outlines the high-level architecture of the Lyceum Learning Platform. Our architecture is designed to be scalable, maintainable, and secure, following a decoupled, service-oriented approach.

- **Decoupled Services:** The frontend and backend are independent applications. They communicate exclusively through a versioned REST API. This allows for independent development, deployment, and scaling.
- **Single Source of Truth:**
    - The backend is the source of truth for all business logic and data.
    - The database schema is the source of truth for data structure.
    - This documentation is the source of truth for architectural decisions.
- **Asynchronous Operations:** Long-running, resource-intensive tasks (e.g., AI-powered content generation) are handled as asynchronous background jobs to ensure a responsive user experience.

## 2. Core Components

| Component | Technology | Hosted On | Role & Responsibilities |
| :--- | :--- | :--- | :--- |
| **Frontend** | Next.js (React) | Vercel | Responsible for all user-facing interfaces, client-side state management, and interaction. It consumes the backend API. |
| **Backend** | Go (Golang) | Railway.com | The core engine. Responsible for business logic, serving the REST API, managing background jobs, and interfacing with the database and third-party AI services. |
| **Database** | PostgreSQL | Supabase | The primary datastore for all persistent data, including users, courses, progress, etc. |
| **Authentication** | Supabase Auth | Supabase | Manages user identity, sign-up, login, and session tokens (JWTs). |
| **Storage** | Supabase Storage | Supabase | Stores user-uploaded files and AI-generated assets (e.g., PDFs, images, video thumbnails). |
| **AI Services** | Gemini, AssemblyAI | Google, AssemblyAI | External services consumed by the backend for tasks like content generation and transcription. API keys are managed securely as environment variables. |

## 3. Communication Flow

1.  **User Authentication:** The frontend communicates directly with Supabase Auth to log a user in. Supabase returns a JWT.
2.  **Authenticated API Requests:** For all subsequent requests to our backend, the frontend includes the Supabase JWT in the `Authorization` header (`Bearer <token>`).
3.  **Backend Verification:** The backend API receives the request, uses a middleware to validate the JWT with Supabase, and identifies the user.
4.  **Data Flow:** The backend processes the request, interacts with the PostgreSQL database, and returns data to the frontend in JSON format.
