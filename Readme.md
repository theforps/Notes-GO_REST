## 📝 Notes API Documentation

This API allows you to perform CRUD operations on notes.

### Base URL

```
http://localhost:5050
```

---

### 📌 Endpoints

---

### 1. Get All Notes

* **URL:** `/get-notes`
* **Method:** `GET`
* **Description:** Retrieves a list of all notes.
* **Response:**

```json
{
  "notes": [
    {
      "id": 1,
      "title": "Example Title",
      "description": "Example Description"
    }
  ]
}
```

---

### 2. Get Note by ID

* **URL:** `/get-note/:id`
* **Method:** `GET`
* **Description:** Retrieves a single note by its ID.
* **Path Parameters:**

    * `id` (integer) — ID of the note.
* **Response:**

```json
{
  "note": {
    "id": 1,
    "title": "Example Title",
    "description": "Example Description"
  }
}
```

---

### 3. Create Note

* **URL:** `/add-note`
* **Method:** `POST`
* **Description:** Creates a new note.
* **Request Body:**

```json
{
  "title": "New Note Title",
  "description": "New Note Description"
}
```

* **Response:**

```json
"Success"
```

---

### 4. Update Note

* **URL:** `/update-note/:id`
* **Method:** `PUT`
* **Description:** Updates an existing note by ID.
* **Path Parameters:**

    * `id` (integer) — ID of the note to update.
* **Request Body:**

```json
{
  "title": "Updated Note Title",
  "description": "Updated Note Description"
}
```

* **Response:**

```json
"Success"
```

---

### 5. Delete Note

* **URL:** `/delete-note/:id`
* **Method:** `DELETE`
* **Description:** Deletes a note by ID.
* **Path Parameters:**

    * `id` (integer) — ID of the note to delete.
* **Response:**

```json
"Success"
```

---

### ⚠️ Error Handling

* Returns `400 Bad Request` if:

    * The `id` parameter is not a valid integer.
    * The request body is not properly formatted JSON.
