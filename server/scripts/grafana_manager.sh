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

# Configuration
BACKUP_DIR="grafana_backup"

# Function to show usage
show_usage() {
    echo "Usage:"
    echo "  $0 backup              - Create a new backup"
    echo "  $0 restore [path]      - Restore from specified backup path"
    echo "  $0 list               - List available backups"
    exit 1
}

# Function to list backups
list_backups() {
    echo "Available backups:"
    if [ -d "$BACKUP_DIR" ]; then
        ls -lt "$BACKUP_DIR"
    else
        echo "No backups found"
    fi
}

# Check command line arguments
if [ $# -lt 1 ]; then
    show_usage
fi

case "$1" in
    "backup")
        ./backup_grafana.sh
        ;;
    "restore")
        if [ -z "$2" ]; then
            echo "Error: Please specify backup path"
            show_usage
        fi
        ./restore_grafana.sh "$2"
        ;;
    "list")
        list_backups
        ;;
    *)
        show_usage
        ;;
esac