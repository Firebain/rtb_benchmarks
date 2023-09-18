FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["RpsTest/RpsTest.csproj", "RpsTest/"]
RUN dotnet restore "RpsTest/RpsTest.csproj"
COPY . .
WORKDIR "/src/RpsTest"
RUN dotnet build "RpsTest.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "RpsTest.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "RpsTest.dll"]
