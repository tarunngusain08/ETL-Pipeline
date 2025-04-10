from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime, timedelta
import subprocess

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}

def run_etl():
    """Function to execute the ETL process."""
    subprocess.run(
        ["go", "run", "/Users/radhakrishna/GolandProjects/ETL-Pipeline/main.go"],
        check=True
    )

with DAG(
    'etl_pipeline',
    default_args=default_args,
    description='ETL pipeline DAG for migrating data from PostgreSQL to MinIO in Parquet format',
    schedule_interval=timedelta(days=1),
    start_date=datetime(2023, 1, 1),
    catchup=False,
) as dag:

    run_etl_task = PythonOperator(
        task_id='run_etl_pipeline',
        python_callable=run_etl,
    )

    run_etl_task
