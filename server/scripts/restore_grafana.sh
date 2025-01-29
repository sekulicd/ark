#!/bin/bash

# Check required environment variables
if [ -z "$GRAFANA_URL" ]; then
    echo "Error: GRAFANA_URL environment variable is not set"
    echo "Example: export GRAFANA_URL=http://your-grafana-url"
    exit 1
fi

if [ -z "$GRAFANA_API_KEY" ]; then
    echo "Error: GRAFANA_API_KEY environment variable is not set"
    echo "Example: export GRAFANA_API_KEY=your-api-key"
    exit 1
fi

# Check for required tools
command -v jq >/dev/null 2>&1 || { echo "Error: jq is required but not installed. Aborting." >&2; exit 1; }
command -v curl >/dev/null 2>&1 || { echo "Error: curl is required but not installed. Aborting." >&2; exit 1; }

# Function to handle errors
handle_error() {
    echo "Error: $1"
    exit 1
}

# Check if backup path is provided
if [ -z "$1" ]; then
    handle_error "Please provide the backup directory path"
fi

BACKUP_PATH="$1"

# Check if backup directory exists
if [ ! -d "$BACKUP_PATH" ]; then
    handle_error "Backup directory does not exist: $BACKUP_PATH"
fi

echo "Starting Grafana dashboard restore..."
echo "Using Grafana URL: $GRAFANA_URL"

# Process each JSON file in the backup directory
for file in "$BACKUP_PATH"/*.json; do
    if [ -f "$file" ]; then
        dashboard_name=$(basename "$file" .json)
        echo "Importing: $dashboard_name"

        # Import dashboard
        response=$(curl -s -X POST \
            -H "Authorization: Bearer $GRAFANA_API_KEY" \
            -H "Content-Type: application/json" \
            -d "@$file" \
            "$GRAFANA_URL/api/dashboards/db")

        # Check if import was successful
        if echo "$response" | jq -e '.status == "success"' > /dev/null; then
            echo "Successfully imported: $dashboard_name"
        else
            echo "Failed to import: $dashboard_name"
            echo "Error: $(echo "$response" | jq -r '.message // "Unknown error"')"
        fi
    fi
done

echo "Restore completed!"