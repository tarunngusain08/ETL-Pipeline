# Setting Up and Running MinIO on macOS

## 1. Install MinIO

### Using Homebrew
MinIO can be installed easily using Homebrew:
```sh
brew install minio/stable/minio
```

### Manual Installation
Alternatively, you can download the binary manually:
```sh
curl -O https://dl.min.io/server/minio/release/darwin-amd64/minio
chmod +x minio
sudo mv minio /usr/local/bin/
```

## 2. Create a Directory for MinIO Storage
MinIO requires a directory to store data. Create one:
```sh
mkdir -p ~/minio/data
```

## 3. Start MinIO Server
Run the MinIO server pointing to the storage directory:
```sh
minio server ~/minio/data --console-address ":9001"
```

### Running MinIO with Custom Credentials
You can set custom credentials using environment variables:
```sh
export MINIO_ROOT_USER=minioadmin
export MINIO_ROOT_PASSWORD=minioadminpassword
minio server ~/minio/data --console-address ":9001"
```

## 4. Access MinIO Web Console
MinIO provides a web-based management UI accessible at:
```
http://localhost:9001
```
Use the credentials defined above to log in.

## 5. Install MinIO Client (mc)
The MinIO Client (mc) is used to interact with the MinIO server.
Install it via Homebrew:
```sh
brew install minio/stable/mc
```

Alternatively, install manually:
```sh
curl -O https://dl.min.io/client/mc/release/darwin-amd64/mc
chmod +x mc
sudo mv mc /usr/local/bin/
```

## 6. Configure MinIO Client
To connect the MinIO client to the server:
```sh
mc alias set localminio http://localhost:9000 minioadmin minioadminpassword
```

## 7. Create a Bucket
Buckets in MinIO are like folders in object storage. Create one:
```sh
mc mb localminio/my-bucket
```

## 8. Upload and Download Files
### Upload a file:
```sh
mc cp myfile.txt localminio/my-bucket/
```

### Download a file:
```sh
mc cp localminio/my-bucket/myfile.txt ./
```

## 9. Run MinIO as a Background Service
To keep MinIO running in the background, use:
```sh
nohup minio server ~/minio/data --console-address ":9001" &
```

For a more persistent setup, consider using `launchctl` to manage MinIO as a service.

## 10. Stopping MinIO
If running in the foreground, use `Ctrl + C`.
If running in the background, find the process and kill it:
```sh
ps aux | grep minio
kill <PID>
```

## Conclusion
MinIO is now set up and running on your macOS machine. You can store and retrieve objects using the MinIO client or API. For more advanced configurations, refer to the [MinIO documentation](https://min.io/docs/).

