@echo off

rem Check if .env file exists
if not exist .env (
    echo .env file not found!
    exit /b 1
)

rem Read and set each line from the .env file
for /f "usebackq tokens=* delims=" %%i in (".env") do (
    rem Ignore lines starting with # (comments)
    echo %%i | findstr /r "^#" >nul
    if errorlevel 1 (
        set "%%i"
    )
)

goose postgres %NMOS_DB_URL% -dir "nmosdb/migrations" down

echo Postgres downed