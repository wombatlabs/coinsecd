#!/bin/bash
rm -rf /tmp/coinsecd-temp

coinsecd --devnet --appdir=/tmp/coinsecd-temp --profile=6061 &
COINSECD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $COINSECD_PID

wait $COINSECD_PID
COINSECD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $COINSECD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $COINSECD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
