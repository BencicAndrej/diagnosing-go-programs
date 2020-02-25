#! /bin/bash

echo "Watching for file changes..."
fswatch -o . | xargs -n1 -I{} ./scripts/reloadActiveChromeTab.scptd
