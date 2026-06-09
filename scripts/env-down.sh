#!/bin/bash

env_down() {
# This deletes and clean the Namespaces

    set -e

    # 1) Delete earth-base
    ip netns del earth-base

    # 2) Delete orb-sat
    ip netns del orb-sat

    # 3) Delete m-rover
    ip netns del m-rover

    set +e
}

if env_down; then
    echo "✅ Namespeces deleted: everything inside cleaned up."
else
    echo "❌ Error: Something went wrong."
fi
