# Lyceum - API Contract v1

This document defines the contract for the REST API served by the `lyceum-backend`. All communication between the frontend and backend MUST adhere to this contract.

## 1. General Principles

- **Format:** All data is sent and received as JSON (`Content-Type: application/json`).
- **Authentication:** All endpoints, unless specified otherwise (e.g., status checks), require a valid Supabase JWT to be passed in the `Authorization` header as a Bearer token.
- **Root URL:** The base URL for the API is `https://<backend-service-url>.up.railway.app`.

## 2. Versioning

The API is versioned in the URL path. The initial version is `v1`.
**Example:** `https://.../api/v1/courses`

## 3. Standard Response Formats

### Success Response

- **Status Code:** `200 OK` for `GET`, `200 OK` or `204 No Content` for `PUT`/`DELETE`, `201 Created` for `POST`.
- **Body:**
```json
{
  "success": true,
  "data": { ... } // The requested or created resource
}
```
For list results, `data` will be an array. For single results, it will be an object. For `DELETE` or actions without a return value, `data` may be omitted.

### Error Response

- **Status Code:** `4xx` for client-side errors, `5xx` for server-side errors.
- **Body:**
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE_SLUG",
    "message": "A human-readable description of the error."
  }
}
```

## 4. Endpoints

*This section will be populated as endpoints are developed.*

### Authentication
- **`POST /api/v1/users/register`**: Endpoint to finalize user registration in our DB after Supabase sign-up.

### Courses
- **`POST /api/v1/courses`**: Create a new course.
- **`GET /api/v1/courses`**: Get a list of courses.
- **`GET /api/v1/courses/{courseId}`**: Get details for a specific course.
