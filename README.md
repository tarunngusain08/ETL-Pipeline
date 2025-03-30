# ETL-Pipeline

This project implements an ETL (Extract, Transform, Load) pipeline that extracts data from a PostgreSQL database, transforms it into Parquet format, and loads it into MinIO object storage.

<img width="1512" alt="Screenshot 2025-03-30 at 5 31 52 PM" src="https://github.com/user-attachments/assets/d8761439-c786-4bc3-8bc4-ccd3748d47eb" />
<img width="686" alt="Screenshot 2025-03-30 at 5 35 15 PM" src="https://github.com/user-attachments/assets/d4930d51-8aef-419f-9f00-768107ba58bb" />


## Prerequisites

1. **PostgreSQL**: Ensure PostgreSQL is installed and running.
2. **MinIO**: Install and run MinIO for object storage.
3. **Go**: Install Go to run the ETL pipeline.
4. **Python**: Install Python for data insertion scripts.
5. **Dependencies**: Install required libraries for Go and Python.

---

## Setup Instructions

### 1. PostgreSQL Setup
1. Start PostgreSQL:
   ```bash
   sudo service postgresql start
   ```
2. Create a database and user:
   ```sql
   CREATE DATABASE postgres;
   CREATE USER etl_user WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE postgres TO etl_user;
   ```
3. Create the `transactions` table:
   ```sql
   CREATE TABLE transactions (
       id SERIAL PRIMARY KEY,
       amount NUMERIC(10, 2) NOT NULL,
       transaction_date TIMESTAMP NOT NULL,
       description TEXT
   );
   ```

### 2. MinIO Setup
1. Start MinIO:
   ```bash
   minio server /data
   ```
2. Access the MinIO web interface at `http://localhost:9000` and log in using the default credentials:
   - Access Key: `minio`
   - Secret Key: `password`
3. Create a bucket named `transactions-bucket`.

---

## Running the ETL Pipeline

### 1. Insert Sample Data
Run the Python script to insert 1 million rows into the `transactions` table:
```bash
python3 insert_transactions.py
```

### 2. Run the ETL Pipeline
Run the Go program to extract data from PostgreSQL, transform it into Parquet format, and load it into MinIO:
```bash
go run main.go
```

---

## High-Level Overview of the Code

### 1. **`insert_transactions.py`**
- Inserts 1 million rows of sample data into the `transactions` table in PostgreSQL.
- Uses the `psycopg2` library for database interaction.

### 2. **`main.go`**
- **Extract**: Queries the `transactions` table in PostgreSQL to fetch data.
- **Transform**: Converts the data into Parquet format using the `parquet-go` library.
- **Load**: Uploads the Parquet file to MinIO using the `minio-go` library.

### 3. **`commit_changes.sh`**
- Automates the process of committing and pushing changes to the Git repository.

---

## Notes
- Ensure the PostgreSQL and MinIO services are running
