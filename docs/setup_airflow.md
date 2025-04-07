# Setting Up Apache Airflow with Python 3.11 Virtual Environment

This document provides detailed steps to set up Apache Airflow using a Python 3.11 virtual environment.

## Prerequisites

1. Install Python 3.11 on your system. You can use Homebrew:
   ```bash
   brew install python@3.11
   ```
2. Ensure `pip` is up-to-date:
   ```bash
   python3.11 -m pip install --upgrade pip
   ```

## Steps to Set Up Airflow

### 1. Create a Virtual Environment

Create a Python 3.11 virtual environment for Airflow:
```bash
python3.11 -m venv airflow-env
source airflow-env/bin/activate
```

### 2. Install Apache Airflow

Install Apache Airflow in the virtual environment:
```bash
pip install apache-airflow
```

### 3. Set the `AIRFLOW_HOME` Environment Variable

Set the `AIRFLOW_HOME` environment variable to a directory of your choice (e.g., `~/airflow`):
```bash
export AIRFLOW_HOME=~/airflow
```

To make this change permanent, add the above line to your shell configuration file (e.g., `~/.zshrc` or `~/.bashrc`):
```bash
echo 'export AIRFLOW_HOME=~/airflow' >> ~/.zshrc
source ~/.zshrc
```

### 4. Initialize the Airflow Database

Initialize the Airflow metadata database:
```bash
airflow db init
```

### 5. Start the Airflow Services

Start the Airflow webserver and scheduler in separate terminals:

1. **Webserver**:
   ```bash
   airflow webserver
   ```
2. **Scheduler**:
   ```bash
   airflow scheduler
   ```

### 6. Access the Airflow Web Interface

Once all services are running, access the Airflow web interface at:
```
http://localhost:8080
```

Use the default credentials:
- Username: `airflow`
- Password: `airflow`

### 7. Place Your DAGs

Ensure your DAG files (e.g., `etl_pipeline_dag.py`) are placed in the `dags` folder inside the `AIRFLOW_HOME` directory:
```bash
mkdir -p $AIRFLOW_HOME/dags
cp /path/to/etl_pipeline_dag.py $AIRFLOW_HOME/dags/
```

### Additional Notes

- You can customize the Airflow configuration in the `airflow.cfg` file located in the `AIRFLOW_HOME` directory.
- For more information, refer to the [official Airflow documentation](https://airflow.apache.org/docs/).
