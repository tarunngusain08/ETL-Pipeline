import psycopg2
from psycopg2.extras import execute_values
from datetime import datetime, timedelta
import random

def insert_transactions():
    # Database connection parameters
    conn = psycopg2.connect(
        dbname="postgres",
        user="etl_user",
        password="",
        host="localhost",
        port="5432"
    )
    cursor = conn.cursor()

    # Generate 10,000 rows of data
    rows = []
    numRows = 10**6
    for i in range(numRows):
        amount = round(random.uniform(1.00, 1000.00), 2)
        transaction_date = datetime.now() - timedelta(days=random.randint(0, 365))
        description = f"Transaction {i + 1}"
        rows.append((amount, transaction_date, description))

    # Insert data using execute_values for efficiency
    insert_query = """
        INSERT INTO transactions (amount, transaction_date, description)
        VALUES %s
    """
    execute_values(cursor, insert_query, rows)

    # Commit the transaction and close the connection
    conn.commit()
    cursor.close()
    conn.close()
    print(numRows, "rows inserted successfully.")

if __name__ == "__main__":
    insert_transactions()
