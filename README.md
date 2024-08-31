# JWT Generator

This Go application generates unsigned JWTs based on command-line arguments. It allows you to pass various claims as 
arguments and handles both single values and lists for claims.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/jwt-generator.git
   
2. **Navigate to the project directory:**

    ```bash
    cd jwt-generator
   
3. **Build the application:**

   ```bash
   go build -o jwt-generator

This will create an executable named jwt-generator in your project directory.

## Usage

The application accepts command-line arguments to specify claims for the JWT. Claims can be single values or lists. 
The general syntax is:

```bash
jwt-generator --key value --key -list value1,value2,...
```

### Arguments
--key value: Single key-value pair where key is the claim name and value is its value.
--key -list value1,value2,...: For claims that should be treated as lists, use the -list flag followed by comma-separated values.

### Example
Generate a JWT with a single value for companyTag and a list for roles:
```bash
jwt-generator --companyTag BAGE --purchaseOrganisation 3109 --roles -list admin,manager
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.txt) file for details.
