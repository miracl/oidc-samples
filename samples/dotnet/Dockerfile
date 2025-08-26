FROM mcr.microsoft.com/dotnet/framework/sdk:4.8-windowsservercore-ltsc2022 AS build

WORKDIR /app

ENV HOST="localhost"
ENV PORT="8000"
EXPOSE ${PORT}

COPY . .

RUN nuget restore
RUN msbuild /p:Configuration=Release -r:False

FROM mcr.microsoft.com/dotnet/framework/aspnet:4.8-windowsservercore-ltsc2022
RUN powershell -Command "New-Item -Path 'C:\app' -Type Directory; New-Website -Name 'app' -PhysicalPath 'C:\app' -Port 8000 -Force"

COPY --from=build /app/. /app
