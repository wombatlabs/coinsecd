#!/bin/bash
rm -rf /tmp/coinsecd-temp

coinsecd --simnet --appdir=/tmp/coinsecd-temp --profile=6061 &
SECPAD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $SECPAD_PID

wait $SECPAD_PID
SECPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $SECPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SECPAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
