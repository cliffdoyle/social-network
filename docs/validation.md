 ## User Input Validation System in Registration Flow

This document explains the custom validation logic implemented across the validator, models, and service packages — focusing on how it's used to validate user input during registration.

## validator Package
The validator package defines a small utility framework to assist with validating user input throughout the application.

## Validator Struct

```go
type Validator struct {
    Errors map[string]string
}
```

Holds a map of field-specific validation errors.

* Only records the first error per field.

### Key Methods

* **New() *Validator**

  * Initializes a new Validator with an empty Errors map.

* **Valid() bool**

  * Returns true if no errors are recorded (i.e., the input is valid).

* **AddError(key, message string)**

  * Adds a validation error message for a specific field, only if one doesn’t already exist.

* **Check(ok bool, key, message string)**

  * Helper to conditionally add an error — only if the `ok` condition fails.

### Generic Validators

*  **Matches(string, *regexp.Regexp)**

  * Checks if a string matches a given regex (used for validating email format).

* **PermittedValue\[T comparable]\(value T, permittedValues ...T)**

  * Checks if a value exists in a list of allowed values (generic).

* **Unique\[T comparable]\(values \[]T)**

  * Checks if all values in a slice are unique (generic).

## models Package
This package contains the validation logic for specific fields related to the user registration form.

 **ValidateEmail(v \*validator.Validator, email string)**

* Ensures the email is not empty.
* Ensures it matches a valid regex pattern for emails.

 **ValidatePasswordPlaintext(v \*validator.Validator, password string)**
Ensures the password is:

* Not empty.
* At least 8 characters.
* No more than 72 characters (bcrypt limit).

 **ValidateUser(v \*validator.Validator, user \*User)**
Validates fields like:

* FirstName (must be provided)
* LastName (max 500 chars)
* Email (using ValidateEmail)
* Password (using ValidatePasswordPlaintext)

Panics if the password hash is missing — considered a critical logic error.

## service Package
This is where validation is applied in real user registration flow.

 **UserService.Register()**

* Parses DateOfBirth:

  * If invalid, an error is added to the Validator.

* Builds the User model from request input.

* Hashes the password:

  * If hashing fails (e.g., internal bcrypt error), returns a server error.

* Runs validation:

  * Calls models.ValidateUser() to perform field-level checks using the Validator.

* Checks for duplicate email:

  * This only runs if the first round of validation passes.
  * Uses repository lookup to determine if the email already exists.

* If any errors are present in the Validator, the function returns:

```go
return nil, v, nil
```

So the handler can return the validation feedback to the client.

 Example Flow During Registration

```go
v := validator.New()

models.ValidateUser(v, user)

if !v.Valid() {
    // Return v.Errors to client as structured validation errors
}
```

The result is a clean, reusable, and centralized validation mechanism that keeps business logic decoupled from other areas of our code.

 ### Why This Design is Great
 **Reusable**: You can use the same validator in multiple places (registration, profile updates, etc.).

 **Clean**: Field-specific validation is in the model layer — not scattered in handlers.

 **Safe**: Prevents nil maps and gracefully handles unexpected cases.

 **Extensible**: You can easily add new helpers like MinLength, MaxLength, IsURL, etc.
