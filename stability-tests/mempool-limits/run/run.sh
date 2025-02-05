#!/bin/bash

APPDIR=/tmp/coinsecd-temp
COINSECD_RPC_PORT=29587

rm -rf "${APPDIR}"

coinsecd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${COINSECD_RPC_PORT}" --profile=6061 &
COINSECD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${COINSECD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $COINSECD_PID

wait $COINSECD_PID
COINSECD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Coinsecd exit code: $COINSECD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $COINSECD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
