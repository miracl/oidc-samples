FROM mcr.microsoft.com/dotnet/framework/sdk:4.8 AS build
WORKDIR /app

ENV HOST="localhost"
ENV PORT="8000"
EXPOSE ${PORT} 

# copy csproj and restore as distinct layers
COPY *.sln ./
COPY *.csproj ./
COPY *.config ./
RUN nuget restore

# copy everything else and build app
COPY . .
WORKDIR /app
RUN msbuild /p:Configuration=Release -r:False

FROM mcr.microsoft.com/dotnet/framework/aspnet:4.8 AS runtime
RUN New-Item -Path 'C:\app' -Type Directory; New-Website -Name 'app' -PhysicalPath 'C:\app' -Port 8000 -Force

COPY --from=build /app/. /app
