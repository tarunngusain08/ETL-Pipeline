package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"bytes"
	"context"
	"log"
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

func extractData(db *sql.DB, query string) ([]map[string]interface{}, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	for rows.Next() {
		row := make([]interface{}, len(columns))
		rowPointers := make([]interface{}, len(columns))
		for i := range row {
			rowPointers[i] = &row[i]
		}
		if err := rows.Scan(rowPointers...); err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = row[i]
		}
		results = append(results, rowMap)
	}
	return results, nil
}

func transformData(data []map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
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
		ContentType: "application
