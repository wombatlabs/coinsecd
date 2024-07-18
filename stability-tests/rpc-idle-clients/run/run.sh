#!/bin/bash
rm -rf /tmp/coinsecd-temp

NUM_CLIENTS=128
coinsecd --devnet --appdir=/tmp/coinsecd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SECPAD_PID=$!
SECPAD_KILLED=0
function killCoinsecdIfNotKilled() {
  if [ $SECPAD_KILLED -eq 0 ]; then
    kill $SECPAD_PID
  fi
}
trap "killCoinsecdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SECPAD_PID

wait $SECPAD_PID
SECPAD_EXIT_CODE=$?
SECPAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $SECPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SECPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
