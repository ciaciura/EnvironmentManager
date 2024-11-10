# Environment manager

App for managing access to environments in a organization,

### Samople vscode launch.json 
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch API",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/api/main.go",
            "cwd": "${workspaceFolder}/api",
            "env": {
                "DOCKER_ENV": "false"
            }
        }
    ]
}
```
### sample .env for docker
should be placed in main project directory
```
JWT_SECRET=your_secret_key
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=password
MONGODB_DATABASE=envmanager_dev
ADMIN_USERNAME=admin
ADMIN_PASSWORD=password
MONGODB_URI=mongodb://admin:password@mongodb_service:27017
APP_PORT=37564
Env=DEV
DOCKER_ENV=true
```
### sample .env-debug for local debuging
should be placed min /api directory
```
JWT_SECRET=your_secret_key
MONGODB_DATABASE=envmanager_dev
ADMIN_USERNAME=admin
ADMIN_PASSWORD=password
MONGODB_URI=mongodb://admin:password@localhost:27017
APP_PORT=37564
Env=DEV
DOCKER_ENV=false
```
