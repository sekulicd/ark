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

# Configuration
BACKUP_DIR="grafana_backup"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_PATH="${BACKUP_DIR}/${DATE}"

# Create backup directory
mkdir -p "$BACKUP_PATH"

# Function to handle errors
handle_error() {
    echo "Error: $1"
    exit 1
}

# Function to sanitize filename
sanitize_filename() {
    echo "$1" | sed 's/[^a-zA-Z0-9._-]/_/g'
}

echo "Starting Grafana dashboard backup..."
echo "Using Grafana URL: $GRAFANA_URL"

# Get all dashboards
dashboards=$(curl -s -H "Authorization: Bearer $GRAFANA_API_KEY" "$GRAFANA_URL/api/search" || handle_error "Failed to get dashboard list")

# Check if dashboards were retrieved successfully
if [ -z "$dashboards" ]; then
    handle_error "No dashboards found or API call failed"
fi

# Export each dashboard
echo "$dashboards" | jq -r '.[] | select(.type == "dash-db") | [.uid, .title] | @tsv' | while IFS=$'\t' read -r uid title; do
    echo "Exporting: $title"

    # Get dashboard JSON
    dashboard_json=$(curl -s -H "Authorization: Bearer $GRAFANA_API_KEY" "$GRAFANA_URL/api/dashboards/uid/$uid")

    # Generate safe filename
    safe_filename=$(sanitize_filename "$title")

    # Clean and prepare JSON for import
    echo "$dashboard_json" | jq '.dashboard | del(.id) | del(.uid) | {dashboard: ., folderId: 0, overwrite: true}' > "$BACKUP_PATH/${safe_filename}.json"

    if [ $? -ne 0 ]; then
        echo "Warning: Failed to export dashboard: $title"
    fi
done

echo "Backup completed successfully!"
echo "Backup location: $BACKUP_PATH"