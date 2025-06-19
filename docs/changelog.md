# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Task B-1.2: Implement "Get My Courses" API Endpoint
- Added `GetCoursesByUserID` to the store to fetch all courses a user is enrolled in.
- Added `GetUserCoursesHandler` to the API to expose this functionality.
- Registered the `GET /api/v1/courses` endpoint in the main router, protected by auth.
- Updated `api_contract.md` with the new endpoint definition.
