package main

import (
	"bytes"
	"context"
	"database/sql"
	"log"

	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/xitongsys/parquet-go-source/buffer"
	"github.com/xitongsys/parquet-go/writer"
)

type Config struct {
	PostgresDSN     string
	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioBucketName string
}

func GetConfig() Config {
	return Config{
		PostgresDSN:     "postgres://etl_user:password@localhost:5432/postgres?sslmode=disable",
		MinioEndpoint:   "localhost:9000",
		MinioAccessKey:  "minio",
		MinioSecretKey:  "password",
		MinioBucketName: "transactions-bucket",
	}
}

type Transaction struct {
	ID              int64   `parquet:"name=id, type=INT64"`
	Amount          float64 `parquet:"name=amount, type=DOUBLE"`
	TransactionDate string  `parquet:"name=transaction_date, type=BYTE_ARRAY, convertedtype=UTF8"`
	Description     string  `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
}

func extractData(db *sql.DB, query string) ([]Transaction, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.Amount, &t.TransactionDate, &t.Description); err != nil {
			return nil, err
		}
		results = append(results, t)
	}
	return results, nil
}

func transformData(data []map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func transformToParquet(data []Transaction) ([]byte, error) {
	buf := new(bytes.Buffer)
	fw := buffer.NewBufferFile()
	pw, err := writer.NewParquetWriter(fw, new(Transaction), 1)
	if err != nil {
		return nil, err
	}

	for _, record := range data {
		if err := pw.Write(record); err != nil {
			return nil, err
		}
	}

	if err := pw.WriteStop(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func loadToMinio(config Config, data []byte, objectName string) error {
	minioClient, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, config.MinioBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, config.MinioBucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	_, err = minioClient.PutObject(ctx, config.MinioBucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	return err
}

func main() {
	config := GetConfig()
	db, err := sql.Open("postgres", config.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	query := "SELECT id, amount, transaction_date, description FROM transactions"
	data, err := extractData(db, query)
	if err != nil {
		log.Fatalf("Failed to extract data: %v", err)
	}

	parquetData, err := transformToParquet(data)
	if err != nil {
		log.Fatalf("Failed to transform data to Parquet: %v", err)
	}

	// transformedData, err := transformData(data)
	// if err != nil {
	// 	log.Fatalf("Failed to transform data: %v", err)
	// }

	objectName := "data.parquet"
	err = loadToMinio(config, parquetData, objectName)
	if err != nil {
		log.Fatalf("Failed to load data to MinIO: %v", err)
	}

	log.Println("ETL process completed successfully!")
}
