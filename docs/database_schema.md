# Lyceum - Database Schema (PostgreSQL)

This document describes the PostgreSQL database schema. The database is hosted on Supabase.

## 1. Notes on Supabase

- Supabase manages its own `auth` schema, which includes the `auth.users` table. Our `users` table will have a one-to-one relationship with `auth.users`.
- Timestamps should use `TIMESTAMPTZ` (timestamp with time zone) to ensure consistency.
- UUIDs are used for primary keys to prevent enumeration attacks and simplify distributed data management.

## 2. Core Tables

### `users` Table
Stores public profile information and application-specific data for each user.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `uuid` | `PRIMARY KEY`, `DEFAULT uuid_generate_v4()` | Primary key. |
| `auth_id` | `uuid` | `FOREIGN KEY (auth.users.id)`, `UNIQUE`, `NOT NULL` | Links to the Supabase user. |
| `username` | `varchar(50)` | `UNIQUE`, `NOT NULL` | The user's unique handle. |
| `display_name` | `varchar(100)`| `NOT NULL` | The user's public display name. |
| `avatar_url` | `text` | | URL to the user's profile picture. |
| `credits` | `integer` | `DEFAULT 0`, `NOT NULL` | Gamification credits. |
| `created_at` | `timestamptz` | `DEFAULT now()`, `NOT NULL` | Timestamp of creation. |
| `updated_at` | `timestamptz` | `DEFAULT now()`, `NOT NULL` | Timestamp of last update. |


### `courses` Table
Stores the definition and metadata for each course.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `uuid` | `PRIMARY KEY`, `DEFAULT uuid_generate_v4()` | Primary key. |
| `creator_id`| `uuid` | `FOREIGN KEY (users.id)`, `NOT NULL` | The user who created the course. |
| `title` | `varchar(255)` | `NOT NULL` | The full title of the course. |
| `description`| `text` | | A detailed description of the course. |
| `visibility` | `varchar(20)` | `CHECK (visibility IN ('public', 'private'))`, `NOT NULL` | Course visibility status. |
| `created_at` | `timestamptz` | `DEFAULT now()`, `NOT NULL` | Timestamp of creation. |
| `updated_at` | `timestamptz` | `DEFAULT now()`, `NOT NULL` | Timestamp of last update. |


### `enrollments` Table
A join table to track which users are enrolled in which courses.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `user_id` | `uuid` | `PRIMARY KEY`, `FOREIGN KEY (users.id)` | Part of the composite primary key. |
| `course_id` | `uuid` | `PRIMARY KEY`, `FOREIGN KEY (courses.id)`| Part of the composite primary key. |
| `role` | `varchar(20)` | `CHECK (role IN ('student', 'instructor'))`, `NOT NULL` | User's role in the course. |
| `enrolled_at` | `timestamptz` | `DEFAULT now()`, `NOT NULL` | Timestamp of enrollment. |
