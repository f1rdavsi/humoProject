# Finance Report

This project is a finance management application that allows users to track their expenses, manage accounts, and generate financial reports. It provides a comprehensive set of features, including CRUD operations for users, accounts, and categories, transaction handling, comments on transactions, transaction history, date-based statistics, exporting statistics in various formats, user avatars, account avatars, category avatars, and the ability to hide accounts.

## Features

1. **User Management**:
   - Create, read, update, and delete users (CRUD operations).
   - User authentication and authorization using JWT (JSON Web Tokens).
   - User avatars for personalized profiles.

2. **Account Management**:
   - Create, read, update, and delete accounts (CRUD operations).
   - Hide accounts to exclude them from financial reports.
   - Account avatars for visual identification.

3. **Category Management**:
   - Create, read, update, and delete categories (CRUD operations).
   - Category avatars for visual identification.

4. **Transactions**:
   - Record income (top-up) transactions.
   - Record expense transactions.
   - Track transaction history.
   - Add comments to transactions.

5. **Statistics**:
   - Generate date-based statistics for financial analysis.
   - Export statistics in XLSX and CSV formats for further processing.

## Technology Stack

The project is developed using the following technologies and frameworks:

- **Golang**: The programming language used for backend development.
- **gin-gonic**: A lightweight web framework for building RESTful APIs in Golang.
- **gORM**: An ORM (Object-Relational Mapping) library for Golang, used for database operations.
- **PostgreSQL**: The chosen database system for data storage.
- **lumberjack.v2**: A rolling logger for Golang, used for efficient log management.
- **spf13/viper**: A library for configuration management.
- **spf13/cast**: A utility library for type casting.
- **joho/godotenv**: A library for reading environment variables from a .env file.
- **bcrypt**: A library for password hashing and encryption.
- **golang-jwt/jwt/v5**: A library for JSON Web Tokens (JWT) authentication and authorization.
- **excelize/v2**: A library for working with Microsoft Excel files (XLSX format).

## Installation

To install and run the project locally, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/icoder-new/finance-report.git
   ```

2. Install the required dependencies:

   ```shell
   go mod download
   ```

3. Set up the PostgreSQL database and configure the connection details in the `.env` file.

4. Run the application:

   ```shell
   go run main.go
   ```

   The application will start running on the specified port (default is 8080).

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/icoder-new/reporter/blob/main/LICENSE) file for details.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please submit an issue or create a pull request.