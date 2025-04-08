$OutputDir = "dist"

Write-Host "🧼 Cleaning and creating $OutputDir directory..."
Remove-Item -Recurse -Force $OutputDir -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path $OutputDir | Out-Null

Write-Host "🚀 Building pong for all major OS targets..."

$env:CGO_ENABLED = "0"

# Windows
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "$OutputDir\pong.exe" pong.go

# Linux
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "$OutputDir\pong-linux" pong.go

# macOS (Intel)
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o "$OutputDir\pong-mac" pong.go

# macOS (ARM)
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o "$OutputDir\pong-mac-arm" pong.go

# Raspberry Pi / ARMv7
$env:GOOS = "linux"
$env:GOARCH = "arm"
$env:GOARM = "7"
go build -o "$OutputDir\pong-armv7" pong.go

Write-Host "✅ Done! All pong builds are living their best life in '$OutputDir\' 💼"
