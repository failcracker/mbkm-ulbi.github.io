# MBKM ULBI Backend

Backend API untuk sistem MBKM (Merdeka Belajar Kampus Merdeka) Universitas Logistik dan Bisnis Internasional.

## Features

- **Authentication & Authorization**: JWT-based authentication with role-based access control
- **User Management**: Support for multiple user roles (mahasiswa, dosen, prodi, cdc, mitra, superadmin)
- **Job Management**: CRUD operations for internship job postings
- **Application Management**: Handle student applications for internships
- **Report Management**: Manage internship reports and approvals
- **Evaluation System**: Grade and evaluate student performance
- **Monthly Logs**: Track student activities during internship
- **Academic Integration**: Grade conversion and academic credit management
- **Dashboard**: Overview statistics and analytics
- **File Upload**: Cloudinary integration for file storage

## Tech Stack

- **Framework**: Gin (Go)
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT
- **File Storage**: Cloudinary
- **Password Hashing**: bcrypt

## Setup

1. **Clone the repository**
```bash
git clone <repository-url>
cd go-backend
```

2. **Install dependencies**
```bash
go mod download
```

3. **Setup environment variables**
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Setup PostgreSQL database**
```sql
CREATE DATABASE mbkm_ulbi;
```

5. **Run the application**
```bash
go run main.go
```

## Environment Variables

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_SSLMODE`: SSL mode (default: disable)
- `JWT_SECRET`: JWT secret key
- `JWT_EXPIRE_HOURS`: JWT expiration time in hours
- `CLOUDINARY_CLOUD_NAME`: Cloudinary cloud name
- `CLOUDINARY_API_KEY`: Cloudinary API key
- `CLOUDINARY_API_SECRET`: Cloudinary API secret
- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin mode (debug/release)

## API Endpoints

### Authentication
- `POST /api/v1/login` - User login
- `POST /api/v1/register` - User registration

### Protected Routes (require JWT token)

#### Profile
- `GET /api/v1/profile` - Get user profile
- `PUT /api/v1/profile` - Update user profile

#### Users
- `GET /api/v1/users/lecturer` - Get lecturers
- `GET /api/v1/users/student` - Get students

#### Roles
- `GET /api/v1/roles` - Get all roles
- `POST /api/v1/roles/assign` - Assign role to user

#### Companies
- `GET /api/v1/companies` - Get companies

#### Jobs
- `GET /api/v1/jobs` - Get jobs
- `GET /api/v1/jobs/:id` - Get job by ID
- `POST /api/v1/jobs` - Create job
- `PUT /api/v1/jobs/:id` - Update job
- `DELETE /api/v1/jobs/:id` - Delete job
- `POST /api/v1/jobs/:id/approve` - Approve job
- `POST /api/v1/jobs/:id/reject` - Reject job
- `GET /api/v1/jobs/:id/list` - Get job candidates

#### Apply Jobs
- `GET /api/v1/apply-jobs` - Get apply jobs
- `GET /api/v1/apply-jobs/:id` - Get apply job by ID
- `GET /api/v1/apply-jobs/user/:user_id` - Get apply jobs by user
- `GET /api/v1/apply-jobs/user/:user_id/last` - Get last apply job by user
- `POST /api/v1/apply-jobs` - Create apply job
- `POST /api/v1/apply-jobs/:id/approve` - Approve apply job
- `POST /api/v1/apply-jobs/:id/reject` - Reject apply job
- `POST /api/v1/apply-jobs/:id/activate` - Activate apply job
- `POST /api/v1/apply-jobs/:id/done` - Complete apply job
- `POST /api/v1/apply-jobs/:id/set-lecturer` - Set lecturer

#### Monthly Logs
- `GET /api/v1/apply-jobs/monthly-logs` - Get monthly logs
- `GET /api/v1/apply-jobs/monthly-logs/:id` - Get monthly log by ID
- `POST /api/v1/apply-jobs/monthly-logs` - Create monthly log
- `PUT /api/v1/apply-jobs/monthly-logs/update` - Update monthly log
- `POST /api/v1/apply-jobs/monthly-logs/:id/approve` - Approve monthly log
- `POST /api/v1/apply-jobs/monthly-logs/:id/revision` - Request revision

#### Reports
- `GET /api/v1/reports` - Get reports
- `GET /api/v1/reports/:id` - Get report by ID
- `POST /api/v1/reports` - Create report
- `POST /api/v1/reports/:id/check` - Check report

#### Evaluations
- `GET /api/v1/evaluations` - Get evaluations
- `GET /api/v1/evaluations/:id` - Get evaluation by ID
- `POST /api/v1/evaluations` - Create evaluation

#### Academic
- `GET /api/v1/program-studi` - Get program studi
- `GET /api/v1/mata-kuliah` - Get mata kuliah
- `GET /api/v1/konversi-nilai` - Get konversi nilai
- `GET /api/v1/konversi-nilai/:id` - Get konversi nilai by ID
- `POST /api/v1/konversi-nilai` - Create konversi nilai

#### Settings
- `GET /api/v1/settings/bobot-nilai` - Get bobot nilai
- `POST /api/v1/settings/bobot-nilai` - Update bobot nilai

#### Dashboard
- `GET /api/v1/dashboard/overview` - Get dashboard overview

## Database Schema

The application uses the following main entities:

- **Users**: Store user information for all roles
- **Roles**: Define user roles and permissions
- **Companies**: Company information
- **Jobs**: Job/internship postings
- **ApplyJobs**: Student applications for jobs
- **Reports**: Internship reports
- **Evaluations**: Performance evaluations
- **MonthlyLogs**: Monthly activity logs
- **Files**: File metadata for uploads
- **Academic entities**: Program studi, mata kuliah, konversi nilai

## Development

1. **Run in development mode**
```bash
GIN_MODE=debug go run main.go
```

2. **Build for production**
```bash
go build -o mbkm-backend main.go
```

3. **Run tests**
```bash
go test ./...
```

## Deployment

1. **Build the application**
```bash
go build -o mbkm-backend main.go
```

2. **Set environment variables**
```bash
export GIN_MODE=release
export PORT=8080
# Set other environment variables
```

3. **Run the application**
```bash
./mbkm-backend
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.