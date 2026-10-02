#!/bin/bash

###############################################################################
# Common reporting functions for Harbormaster certification workflows
###############################################################################

post_result() {

    local TIER="$1"
    local STEP="$2"
    local STATUS="$3"
    local SCORE="$4"
    local MESSAGE="$5"
    local DETAIL="$6"
    local TERMINATE="${7:-false}"

    local RESPONSE_FILE="${RUNNER_TEMP}/response.txt"
    local MAX=5000

    if [ ${#DETAIL} -gt $MAX ]; then
        DETAIL="${DETAIL:0:$MAX}"
    fi

    echo "Reporting result to Harbormaster..."
    echo "URL: ${RESULT_URL}/${ACTION_TO_CALL}"
    echo "Tier: ${TIER}"
    echo "Step: ${STEP}"
    echo "Status: ${STATUS}"
    echo "Message: ${MESSAGE}"
    echo "Detail: ${DETAIL}"

    local HTTP_STATUS
    local CURL_EXIT

    set +e

    # Do not use `curl ... || echo 000` — on connect failure curl already
    # writes http_code 000, and appending makes 000000. Never let reporting
    # fail the job (workflows run with bash -e).
    HTTP_STATUS=$(
        curl \
            --silent \
            --show-error \
            --write-out "%{http_code}" \
            --output "${RESPONSE_FILE}" \
            --connect-timeout 5 \
            --max-time 10 \
            -X POST "${RESULT_URL}/${ACTION_TO_CALL}" \
            --data-urlencode "certificationIdentifier=${CERTIFICATION_IDENTIFIER}" \
            --data-urlencode "tier=${TIER}" \
            --data-urlencode "step=${STEP}" \
            --data-urlencode "status=${STATUS}" \
            --data-urlencode "score=${SCORE}" \
            --data-urlencode "message=${MESSAGE}" \
            --data-urlencode "detail=${DETAIL}" \
            --data-urlencode "application=${APP_NAME}" \
            --data-urlencode "terminate=${TERMINATE}"
    )

    CURL_EXIT=$?

    set -e

    echo "curl exit code: ${CURL_EXIT}"
    echo "Harbormaster returned HTTP ${HTTP_STATUS}"

    if [ -f "${RESPONSE_FILE}" ]; then
        echo "Response:"
        cat "${RESPONSE_FILE}"
        rm -f "${RESPONSE_FILE}"
    fi

    if [ "$CURL_EXIT" -ne 0 ]; then
        echo "WARNING: Unable to report result to Harbormaster."
        echo "curl exit code: ${CURL_EXIT}"
        return 0
    fi

    if [ "${HTTP_STATUS}" != "200" ]; then
        echo "WARNING: Harbormaster returned HTTP ${HTTP_STATUS}."
        return 0
    fi

    echo "Result successfully reported to Harbormaster."
}

post_pass() {

    post_result \
        "$1" \
        "$2" \
        "PASSED" \
        "$3" \
        "$4" \
        "$5"
}

post_fail() {

    post_result \
        "$1" \
        "$2" \
        "FAILED" \
        "$3" \
        "$4" \
        "$5"
}

post_warning() {

    post_result \
        "$1" \
        "$2" \
        "WARNING" \
        "$3" \
        "$4" \
        "$5"
}

post_terminate_prebuild() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "BUILD" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}

post_terminate_build() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "BUILD" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}

post_terminate_runtime() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "RUNTIME" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}

post_terminate_delivery() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "DELIVERY" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}

post_terminate_cloud() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "CLOUD" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}

post_terminate_publish() {

    local EXIT_WORKFLOW="${1:-true}"

    post_result \
        "PUBLISH" \
        "" \
        "" \
        0 \
        "" \
        "" \
        "true"

    if [ "$EXIT_WORKFLOW" = "true" ]; then
        exit 1
    fi
}
