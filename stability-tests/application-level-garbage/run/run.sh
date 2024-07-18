#!/bin/bash
rm -rf /tmp/coinsecd-temp

coinsecd --devnet --appdir=/tmp/coinsecd-temp --profile=6061 --loglevel=debug &
SECPAD_PID=$!
SECPAD_KILLED=0
function killCoinsecdIfNotKilled() {
    if [ $SECPAD_KILLED -eq 0 ]; then
      kill $SECPAD_PID
    fi
}
trap "killCoinsecdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SECPAD_PID

wait $SECPAD_PID
SECPAD_KILLED=1
SECPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $SECPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SECPAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
