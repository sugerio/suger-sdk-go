#!/bin/bash

# Script to regenerate OpenAPI specification from [marketplace-service] and update Go SDK
# This script automates the process of:
# 1. Copy OpenAPI spec from [marketplace-service]
# 2. Converting Swagger 2.0 to OpenAPI 3.0 format

set -e  # Exit on any error

# Configuration
GO_SERVER_PATH="../marketplace-service"
GO_SDK_PATH="../suger-sdk-go"
SWAGGER_SOURCE="$GO_SERVER_PATH/service/marketplace_service/api/files/swagger.yaml"
OPENAPI_TARGET="$GO_SDK_PATH/api/openapi.yaml"

echo "🚀 Starting OpenAPI specification update process..."

# Step 1: Check if openapi-generator is available for conversion
if ! command -v openapi-generator &> /dev/null; then
    echo "❌ Error: 'openapi-generator' command not found."
    echo "Install it with: npm install @openapitools/openapi-generator-cli -g"
    echo "Or: brew install openapi-generator"
    exit 1
fi

# Step 2: Check if the Swagger file exists
if [ ! -f "$SWAGGER_SOURCE" ]; then
    echo "❌ Error: Swagger specification was not generated at $SWAGGER_SOURCE"
    exit 1
fi

# Step 3: Convert Swagger 2.0 to OpenAPI 3.0
echo "🔄 Converting Swagger 2.0 to OpenAPI 3.0 format..."
openapi-generator generate \
  -i "$SWAGGER_SOURCE" \
  -g openapi-yaml \
  -o "/tmp/openapi-conversion" \
  --skip-validate-spec

# The converted file will be in the output directory
CONVERTED_FILE="/tmp/openapi-conversion/openapi/openapi.yaml"
if [ ! -f "$CONVERTED_FILE" ]; then
    echo "❌ Error: OpenAPI 3.0 conversion failed"
    exit 1
fi

echo "✅ Successfully converted to OpenAPI 3.0 format"

# Step 4: Copy the converted OpenAPI spec to Go SDK project
echo "📋 Copying converted OpenAPI specification to Go SDK project..."
cp "$CONVERTED_FILE" "$OPENAPI_TARGET"

echo "✅ OpenAPI 3.0 specification copied to $OPENAPI_TARGET"

# Step 5: Cleanup
echo "🧹 Cleaning up temporary files..."
rm -rf "/tmp/openapi-conversion"

echo ""
echo "🎉 OpenAPI specification update completed!"
echo ""
echo "📋 Summary of changes:"
echo "  - Converted Swagger 2.0 to OpenAPI 3.0 format"
echo "  - Updated: $OPENAPI_TARGET"
echo ""
echo "📝 Next steps:"
echo "  Regenerated Go SDK code"