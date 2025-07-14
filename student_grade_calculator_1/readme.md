# Student Grade Calculator (Go)

This is a console-based Go application that allows students to calculate their average grade based on different subjects.

## Features

- Prompts student for name and number of subjects
- Accepts subject names and grades with input validation
- Calculates and displays:
  - Subject names and corresponding grades
  - Average grade
- Includes unit tests for grade average calculation logic

## Prerequisites

- Go 1.18 or higher

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/student_grade_calculator_1.git
   cd student_grade_calculator_1
   ```

2. Initialize Go module (if not already initialized):

   ```bash
   go mod init student_grade_calculator_1
   ```

3. Run the application:

   ```bash
   go run main.go grades.go
   ```

## Example Run

Below is an example of how the program behaves when run from the terminal:

<img width="499" height="366" alt="Screenshot_20250714_154425" src="https://github.com/user-attachments/assets/941cd0e9-7609-4113-ae7c-8cb94699d6de" />


## Running Tests

Unit tests are provided for the grade average calculation logic:

```bash
go test
```

Example output of running tests:

<img width="756" height="253" alt="Screenshot_20250714_154841" src="https://github.com/user-attachments/assets/6fafa9c2-7ed4-431f-919d-cabb9ea773d1" />


## Project Structure

```
student_grade_calculator_1/
├── main.go            // Main application logic
├── grades.go          // Grade processing logic (average calculation)
├── grades_test.go     // Unit tests for grades.go
└── go.mod             // Go module definition
```

## License

This project is licensed under the MIT License.
