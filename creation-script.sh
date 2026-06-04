#!/bin/bash

create_netns() {
# This is for creating and setting up the Network Namespaces

    set -e

    # 1) Create earth-base, orb-sat and m-rover Namespaces
    ip netns add earth-base
    ip netns add orb-sat
    ip netns add m-rover

    # 2) Turn on the loopback interface
    ip netns exec earth-base ip link set lo up
    ip netns exec orb-sat ip link set lo up
    ip netns exec m-rover ip link set lo up

    # 3) Create and connect Virtual Ethernet between earth-base and orb-sat
    ip link add veth-earth type veth peer name veth-sat-earth
    ip link set veth-earth netns earth-base
    ip link set veth-sat-earth netns orb-sat

    # 4) Between orb-sat and m-rover
    ip link add veth-sat-rover type veth peer name veth-rover
    ip link set veth-sat-rover netns orb-sat
    ip link set veth-rover netns m-rover

    # 5) Turn on veth
    ip netns exec earth-base ip link set veth-earth up
    ip netns exec orb-sat ip link set veth-sat up
    ip netns exec m-rover ip link set veth-rover up

    # 6) Assign IP Addresses
    ip netns exec earth-base ip addr add 10.0.1.1/24 dev veth-earth
    ip netns exec orb-sat ip addr add 10.0.1.2/24 dev veth-sat
    ip netns exec m-rover ip addr add 10.0.1.3/24 dev veth-rover

    set +e
}

if create_netns; then
    echo "✅ Namespaces created."
else
    echo "❌ Error: Something went wrong."
fi
