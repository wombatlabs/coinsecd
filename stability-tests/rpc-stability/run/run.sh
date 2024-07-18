#!/bin/bash
rm -rf /tmp/coinsecd-temp

coinsecd --devnet --appdir=/tmp/coinsecd-temp --profile=6061 --loglevel=debug &
SECPAD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $SECPAD_PID

wait $SECPAD_PID
SECPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $SECPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SECPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
