# job-sync
## HOW TO USE

# 1. Create Database
# 2. Fill the env file
# 3. Migrate database by installing or using goose
`goose postgres "user=postgres dbname=job_sync sslmode=disable" status`
# 3. Start Minio Server

# notes: for migrating data
`goose create file_name sql`