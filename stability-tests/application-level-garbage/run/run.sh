#!/bin/bash
rm -rf /tmp/coinsecd-temp

coinsecd --devnet --appdir=/tmp/coinsecd-temp --profile=6061 --loglevel=debug &
COINSECD_PID=$!
COINSECD_KILLED=0
function killCoinsecdIfNotKilled() {
    if [ $COINSECD_KILLED -eq 0 ]; then
      kill $COINSECD_PID
    fi
}
trap "killCoinsecdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $COINSECD_PID

wait $COINSECD_PID
COINSECD_KILLED=1
COINSECD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $COINSECD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $COINSECD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
