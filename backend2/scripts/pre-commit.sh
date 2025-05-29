# scripts/pre-commit
#!/bin/bash

echo "🔍 Running golangci-lint..."
if ! golangci-lint run; then
  echo "❌ Lint failed. Aborting commit."
  exit 1
fi
echo "✅ Lint passed. Proceeding."