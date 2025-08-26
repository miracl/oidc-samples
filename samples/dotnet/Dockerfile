# FROM mcr.microsoft.com/dotnet/framework/sdk:4.8 AS build
# WORKDIR /app

# ENV HOST="localhost"
# ENV PORT="8000"
# EXPOSE ${PORT}

# # copy csproj and restore as distinct layers
# COPY *.sln ./
# COPY *.csproj ./
# COPY *.config ./
# RUN nuget restore

# # copy everything else and build app
# COPY . .
# WORKDIR /app
# RUN msbuild /p:Configuration=Release -r:False

# FROM mcr.microsoft.com/dotnet/framework/aspnet:4.8 AS runtime
# RUN New-Item -Path 'C:\app' -Type Directory; New-Website -Name 'app' -PhysicalPath 'C:\app' -Port 8000 -Force

# COPY --from=build /app/. /app
#########################################################################
# Stage 1: Build the application
# Use a specific Windows Server 2022-compatible SDK image
FROM mcr.microsoft.com/dotnet/framework/sdk:4.8-windowsservercore-ltsc2022 AS build

# Set the working directory
WORKDIR /app


ENV HOST="localhost"
ENV PORT="8000"
EXPOSE ${PORT}

# copy csproj and restore as distinct layers
# COPY *.sln ./
# COPY *.csproj ./
# COPY *.config ./

# Copy the project files
COPY . .

# Restore NuGet packages
RUN nuget restore

# Build the application and publish it
# RUN msbuild /p:Configuration=Release /p:DeployOnBuild=true /p:PublishProfile=FileSystem /p:publishUrl=c:\publish
RUN msbuild /p:Configuration=Release -r:False

# Stage 2: Create the runtime image
# Use a compatible Windows Server 2022 ASP.NET runtime image
FROM mcr.microsoft.com/dotnet/framework/aspnet:4.8-windowsservercore-ltsc2022
# RUN New-Item -Path 'C:\app' -Type Directory; New-Website -Name 'app' -PhysicalPath 'C:\app' -Port 8000 -Force
RUN powershell -Command "New-Item -Path 'C:\app' -Type Directory; New-Website -Name 'app' -PhysicalPath 'C:\app' -Port 8000 -Force"

# # Set the working directory for the final image
# WORKDIR /inetpub/wwwroot

# # Copy the published application from the build stage
# COPY --from=build /publish/ .

COPY --from=build /app/. /app
