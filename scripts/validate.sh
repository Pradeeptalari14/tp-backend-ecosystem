#!/usr/bin/env bash
set -euo pipefail

echo "======================================================================"
echo "  Validating Backend Ecosystem Multi-Stack Repository (v2.0.0)"
echo "======================================================================"

# 1. Check file existence
echo "==> Step 1: Checking project files..."
test -f main.py
test -f docker-compose.yml
test -f nginx.conf
test -f docs/backend_ecosystem_flow.png
test -f projects/ecommerce-order-engine/main.py
test -f projects/notification-event-hub/src/index.js
test -f projects/saas-billing-dispatcher/main.go
echo "    All required project files verified."

# 2. Syntax check Python
echo "==> Step 2: Validating Python syntax..."
python3 -m py_compile main.py
python3 -m py_compile projects/ecommerce-order-engine/main.py
python3 -m py_compile projects/ecommerce-order-engine/worker.py
echo "    Python code syntax clean."

# 3. Syntax check Node.js
if command -v node >/dev/null 2>&1; then
  echo "==> Step 3: Validating Node.js syntax..."
  node --check projects/notification-event-hub/src/index.js
  node --check projects/notification-event-hub/src/worker.js
  echo "    Node.js syntax clean."
fi

# 4. Check simulation run
echo "==> Step 4: Running lifecycle verification..."
python3 main.py

echo "======================================================================"
echo "  ALL BACKEND ECOSYSTEM VERIFICATION CHECKS PASSED SUCCESSFULLY!"
echo "======================================================================"
