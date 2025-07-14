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

![Run Example](https://gist.github.com/user-attachments/assets/c3f50814-2a77-4b4d-b02b-a99d69804bf3)

## Running Tests

Unit tests are provided for the grade average calculation logic:

```bash
go test
```

Example output of running tests:

![Test Output](https://gist.github.com/user-attachments/assets/ed1b4546-7e4e-4214-8ec3-5ad816e971d8)

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
