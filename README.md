# Setting Environment Variables

To configure the following environment variables, execute the following commands:

```bash
export ZINCSEARCH_IP=${zincsearch_ip}
export ENTRYPOINT_APIREST_ENABLED=${apirest_enabled}
export ENTRYPOINT_APIREST_PORT=${apirest_port}
export ZINCSEARCH_IP=${zincsearch_ip}
export ZINCSEARCH_PORT=${zincsearch_port}
export ZINCSEARCH_INDEX=${zincsearch_index_name}
export ZINC_FIRST_ADMIN_USER=${zincsearch_user}
export ZINC_FIRST_ADMIN_PASSWORD=${zincsearch_pass}
export BASIC_AUTH_USER=${basic_auth_user}
export BASIC_AUTH_PASS=${basic_auth_pass}
export EXTERNAL_USER=${external_auth_user}
export EXTERNAL_PASS=${external_auth_pass}
export JWT_AUTH_SECRET=${jwt_secret}

```

# Running the Project

To run the project, you need to have GoLang installed on your system. Once installed, follow these steps:

1. **Install GoLang:**
   If you haven't already, download and install GoLang from the [official website](https://golang.org/doc/install).

2. **Navigate to Project Directory:**
   Open a terminal and navigate to the directory where your project is located.

3. **Run the Application:**
   Execute the following command in your terminal:
   ```bash
   go run ./application
   ```