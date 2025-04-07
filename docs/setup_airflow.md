# Setting Up Apache Airflow with Docker

This document provides detailed steps to set up Apache Airflow using Docker.

## Prerequisites

1. Install Docker and Docker Compose on your system.
2. Ensure you have sufficient permissions to run Docker commands.

## Steps to Set Up Airflow

### 1. Pull the Airflow Docker Image

Run the following command to pull the latest Airflow image:

```bash
docker pull apache/airflow:latest
```

### 2. Create the Airflow Directory Structure

Create the necessary directories for Airflow:

```bash
mkdir -p ~/airflow
cd ~/airflow
mkdir -p ./dags ./logs ./plugins ./config
```

### 3. Download the Docker Compose File

Download the official Airflow `docker-compose.yaml` file:

```bash
curl -LfO 'https://airflow.apache.org/docs/apache-airflow/2.7.3/docker-compose.yaml'
```

### 4. Set the Airflow User ID

Set the Airflow user ID in an `.env` file:

```bash
echo -e "AIRFLOW_UID=$(id -u)" > .env
```

### 5. Initialize the Airflow Environment

Run the following command to initialize the Airflow environment:

```bash
docker compose up airflow-init
```

> **Note:** If you encounter warnings about the `version` attribute being obsolete, you can ignore them or remove the `version` attribute from the `docker-compose.yaml` file.

### 6. Start the Airflow Services

Start the Airflow services in detached mode:

```bash
docker compose up -d
```

### 7. Troubleshooting

If the `airflow-postgres-1` container is unhealthy, follow these steps:

1. Check the logs of the `airflow-postgres-1` container:

   ```bash
   docker logs airflow-postgres-1
   ```

2. Ensure that the PostgreSQL service is running correctly. If necessary, restart the container:

   ```bash
   docker restart airflow-postgres-1
   ```

3. Verify that the required ports (e.g., 5432 for PostgreSQL) are not being used by other services.

4. If the issue persists, try removing the container and starting it again:

   ```bash
   docker compose down
   docker compose up -d
   ```

### 8. Access the Airflow Web Interface

Once all services are running, access the Airflow web interface at:

```
http://localhost:8080
```

Use the default credentials:
- Username: `airflow`
- Password: `airflow`

## Additional Notes

- You can customize the `docker-compose.yaml` file to suit your requirements.
- For more information, refer to the [official Airflow documentation](https://airflow.apache.org/docs/).
